// Package mailer sends transactional email.
//
// In production the Resend HTTP API is used. When no API key is configured the
// Log mailer simply records the would-be send at INFO level, which is the same
// behaviour we relied on during development.
package mailer

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

type Mailer interface {
	Send(ctx context.Context, msg Message) error
}

type Message struct {
	To      string
	Subject string
	Text    string
	HTML    string
}

type Config struct {
	ResendAPIKey string
	From         string
}

func New(cfg Config) Mailer {
	if cfg.ResendAPIKey != "" && cfg.From != "" {
		return &resendMailer{apiKey: cfg.ResendAPIKey, from: cfg.From, http: &http.Client{Timeout: 10 * time.Second}}
	}
	return &logMailer{from: cfg.From}
}

type logMailer struct{ from string }

func (m *logMailer) Send(_ context.Context, msg Message) error {
	log.Info().Str("from", m.from).Str("to", msg.To).Str("subject", msg.Subject).Msg("mailer: log-only send")
	return nil
}

type resendMailer struct {
	apiKey string
	from   string
	http   *http.Client
}

type resendPayload struct {
	From    string   `json:"from"`
	To      []string `json:"to"`
	Subject string   `json:"subject"`
	Text    string   `json:"text,omitempty"`
	HTML    string   `json:"html,omitempty"`
}

func (m *resendMailer) Send(ctx context.Context, msg Message) error {
	body, err := json.Marshal(resendPayload{
		From:    m.from,
		To:      []string{msg.To},
		Subject: msg.Subject,
		Text:    msg.Text,
		HTML:    msg.HTML,
	})
	if err != nil {
		return err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "https://api.resend.com/emails", bytes.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+m.apiKey)
	req.Header.Set("Content-Type", "application/json")
	resp, err := m.http.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		buf := make([]byte, 1024)
		n, _ := resp.Body.Read(buf)
		return fmt.Errorf("resend: %s: %s", resp.Status, string(buf[:n]))
	}
	return nil
}

var ErrNotConfigured = errors.New("mailer not configured")
