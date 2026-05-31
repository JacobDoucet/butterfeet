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
			HTML: fmt.Sprintf(
				`<p>Hi,</p><p>You reserved <strong>%s</strong> on the <strong>%s</strong> registry but haven't confirmed yet. We're holding it for you for about %s longer.</p>`+
					`<p>If you've completed your purchase, head back and click <strong>"I've bought this"</strong>. If you changed your mind, please release the reservation so another guest can grab it.</p>`+
					`<p><a href="%s">Open the registry</a></p>`,
				htmlEscape(item.Title), htmlEscape(reg.Title), formatDuration(expiresIn), link,
			),
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

func htmlEscape(s string) string {
	r := strings.NewReplacer("&", "&amp;", "<", "&lt;", ">", "&gt;", "\"", "&quot;")
	return r.Replace(s)
}
