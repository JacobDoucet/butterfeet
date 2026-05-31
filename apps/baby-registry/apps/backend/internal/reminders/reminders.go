// Package reminders runs a periodic loop that emails buyers when they have a
// Reserved hold that's older than the reminder threshold (so they don't forget
// to either confirm the purchase or release the hold).
package reminders

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/api"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/enum_reservation_status"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/permissions"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry"
	registryapi "github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_api"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_item"
	registryitemapi "github.com/butterfeetlabs/baby-registry/apps/backend/generated/registry_item_api"
	"github.com/butterfeetlabs/baby-registry/apps/backend/generated/reservation"
	reservationapi "github.com/butterfeetlabs/baby-registry/apps/backend/generated/reservation_api"
	"github.com/butterfeetlabs/baby-registry/apps/backend/internal/mailer"
)

type Config struct {
	Client     api.Client
	Mailer     mailer.Mailer
	AppBaseURL string
	// Threshold after which a Reserved hold should trigger a reminder.
	Threshold time.Duration
	// Interval between scans.
	Interval time.Duration
}

func (c *Config) defaults() {
	if c.Threshold <= 0 {
		c.Threshold = 4 * time.Hour
	}
	if c.Interval <= 0 {
		c.Interval = 15 * time.Minute
	}
}

// Run blocks until ctx is cancelled, periodically scanning for reservations
// that need a reminder and emailing the buyer.
func Run(ctx context.Context, cfg Config) {
	cfg.defaults()
	if cfg.Mailer == nil || cfg.Client == nil {
		log.Warn().Msg("reminders.Run skipped: mailer or client is nil")
		return
	}
	base := strings.TrimRight(cfg.AppBaseURL, "/")

	tick := time.NewTicker(cfg.Interval)
	defer tick.Stop()

	// Run once on startup so a recently-restarted server catches up.
	scan(ctx, cfg, base)
	for {
		select {
		case <-ctx.Done():
			return
		case <-tick.C:
			scan(ctx, cfg, base)
		}
	}
}

func scan(ctx context.Context, cfg Config, base string) {
	super := permissions.NewSuperActor()
	now := time.Now().UTC()
	cutoff := now.Add(-cfg.Threshold)
	status := enum_reservation_status.Reserved

	scanCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	sweepExpired(scanCtx, cfg, super, now)

	// Find Reserved holds that are still valid (not expired) and were created
	// more than `cfg.Threshold` ago.
	resvResult, _, err := cfg.Client.Reservation().Search(scanCtx, super, reservation.WhereClause{
		StatusEq:    &status,
		ExpiresAtGt: &now,
	}, reservationapi.QueryOptions{Limit: 500})
	if err != nil {
		log.Error().Err(err).Msg("reminders scan failed")
		return
	}

	for _, rsv := range resvResult.Data {
		// Skip rows we've already reminded.
		if !rsv.ReminderSentAt.IsZero() {
			continue
		}
		// Skip rows created within the threshold.
		if rsv.Created.At.After(cutoff) {
			continue
		}
		email := strings.TrimSpace(rsv.ContactEmail)
		if email == "" {
			continue
		}

		reg, _, err := cfg.Client.Registry().SelectById(scanCtx, super,
			registry.SelectByIdQuery{Id: rsv.RegistryId}, registryapi.NewProjection(true))
		if err != nil {
			log.Error().Err(err).Str("reservationId", rsv.Id).Msg("reminder registry lookup failed")
			continue
		}
		item, _, err := cfg.Client.RegistryItem().SelectById(scanCtx, super,
			registry_item.SelectByIdQuery{Id: rsv.ItemId}, registryitemapi.NewProjection(true))
		if err != nil {
			log.Error().Err(err).Str("reservationId", rsv.Id).Msg("reminder item lookup failed")
			continue
		}

		link := base + "/r/" + reg.Slug
		expiresIn := time.Until(rsv.ExpiresAt).Round(time.Minute)

		brand := mailer.Brand{AppName: "Stork Nest", AppBaseURL: base}
		sendCtx, sendCancel := context.WithTimeout(ctx, 10*time.Second)
		err = cfg.Mailer.Send(sendCtx, mailer.Message{
			To:      email,
			Subject: fmt.Sprintf("Still buying %q for the %s registry?", item.Title, reg.Title),
			Text: fmt.Sprintf(
				"Hi,\n\nYou reserved %q on the %s registry but haven't confirmed yet. "+
					"We're holding it for you for about %s longer.\n\n"+
					"If you've completed your purchase, head to the registry and click "+
					"\"I've bought this\". If you changed your mind, please release the "+
					"reservation so another guest can grab it.\n\n%s\n",
				item.Title, reg.Title, formatDuration(expiresIn), link,
			),
			HTML: brand.Render(mailer.Email{
				Preheader: "We're still holding your reserved gift.",
				Heading:   "Still buying " + item.Title + "?",
				Intro: "You reserved \"" + item.Title + "\" on the \"" + reg.Title + "\" registry but haven't confirmed yet. " +
					"We're holding it for you for about " + formatDuration(expiresIn) + " longer.",
				BodyHTML: `<p style="margin:16px 0 0 0;">If you've completed your purchase, open the registry and tap <strong>I've bought this</strong>. If you changed your mind, please release the reservation so another guest can grab it.</p>`,
				CTAText:  "Open the registry",
				CTAHref:  link,
			}),
		})
		sendCancel()
		if err != nil {
			log.Error().Err(err).Str("reservationId", rsv.Id).Str("email", email).Msg("reminder email send failed")
			continue
		}

		// Stamp the reservation so we don't send again.
		inner := rsv.Model
		inner.ReminderSentAt = time.Now().UTC()
		if _, _, err := cfg.Client.Reservation().Update(scanCtx, super, inner, reservation.NewProjection(true)); err != nil {
			log.Error().Err(err).Str("reservationId", rsv.Id).Msg("reminder stamp update failed")
		} else {
			log.Info().Str("reservationId", rsv.Id).Str("email", email).Msg("reservation reminder sent")
		}
	}
}

// sweepExpired deletes Reserved holds whose ExpiresAt has passed. Other
// statuses (Purchased, Received, Cancelled) are preserved for history.
func sweepExpired(ctx context.Context, cfg Config, actor permissions.Actor, now time.Time) {
	status := enum_reservation_status.Reserved
	expired, _, err := cfg.Client.Reservation().Search(ctx, actor, reservation.WhereClause{
		StatusEq:     &status,
		ExpiresAtLte: &now,
	}, reservationapi.QueryOptions{Limit: 500})
	if err != nil {
		log.Error().Err(err).Msg("reminders sweep search failed")
		return
	}
	for _, rsv := range expired.Data {
		if rsv.ExpiresAt.IsZero() {
			continue
		}
		if err := cfg.Client.Reservation().Delete(ctx, actor, rsv.Id); err != nil {
			log.Error().Err(err).Str("reservationId", rsv.Id).Msg("expired reservation delete failed")
			continue
		}
		log.Info().Str("reservationId", rsv.Id).Time("expiresAt", rsv.ExpiresAt).Msg("expired reservation removed")
	}
}

func formatDuration(d time.Duration) string {
	if d <= 0 {
		return "less than a minute"
	}
	h := int(d.Hours())
	m := int(d.Minutes()) % 60
	if h > 0 && m > 0 {
		return fmt.Sprintf("%dh %dm", h, m)
	}
	if h > 0 {
		return fmt.Sprintf("%dh", h)
	}
	return fmt.Sprintf("%dm", m)
}
