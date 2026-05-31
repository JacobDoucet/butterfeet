package mailer

import (
	"fmt"
	"html"
	"strings"
)

// Brand carries the per-environment values needed to render a branded email.
// Callers typically build one at startup and reuse it.
type Brand struct {
	AppName    string
	AppBaseURL string // absolute, no trailing slash
}

// Email is a small abstract description of a transactional message that the
// renderer turns into a brand-styled HTML body.
type Email struct {
	Preheader string // hidden inbox preview text
	Heading   string // hero heading (plain text; HTML-escaped)
	Intro     string // first paragraph (plain text; HTML-escaped, newlines preserved)
	BodyHTML  string // optional raw HTML inserted after the intro (caller is responsible for escaping)
	CTAText   string
	CTAHref   string
	Footnote  string // optional small grey paragraph below the CTA
}

const (
	brandPrimary    = "#7a9e7e"
	brandSecondary  = "#e8a87c"
	brandBg         = "#fbf7f2"
	brandSurface    = "#ffffff"
	brandText       = "#2d2a26"
	brandTextMuted  = "#6b665f"
	brandBorder     = "#ecdfce"
	brandFooterText = "#9b948b"
)

// Esc escapes a plain string for safe inclusion in HTML.
func Esc(s string) string { return html.EscapeString(s) }

// Render returns a full HTML document for the given email. The template is
// table-based and inline-styled so it survives the major email clients.
func (b Brand) Render(e Email) string {
	app := b.AppName
	if app == "" {
		app = "Stork Nest"
	}
	base := strings.TrimRight(b.AppBaseURL, "/")

	heading := Esc(e.Heading)
	intro := escapeWithBreaks(e.Intro)

	cta := ""
	if e.CTAText != "" && e.CTAHref != "" {
		cta = fmt.Sprintf(`
            <tr><td align="center" style="padding:24px 32px 8px 32px;">
              <a href="%s" style="background:%s;color:#ffffff;text-decoration:none;display:inline-block;padding:14px 28px;border-radius:12px;font-weight:600;font-family:Inter,Helvetica,Arial,sans-serif;font-size:15px;letter-spacing:0.01em;">%s</a>
            </td></tr>`, e.CTAHref, brandPrimary, Esc(e.CTAText))
	}

	footnote := ""
	if e.Footnote != "" {
		footnote = fmt.Sprintf(`
            <tr><td style="padding:12px 32px 0 32px;color:%s;font-size:12px;line-height:1.5;font-family:Inter,Helvetica,Arial,sans-serif;">%s</td></tr>`,
			brandTextMuted, escapeWithBreaks(e.Footnote))
	}

	body := ""
	if e.BodyHTML != "" {
		body = fmt.Sprintf(`
            <tr><td style="padding:0 32px;color:%s;font-size:15px;line-height:1.55;font-family:Inter,Helvetica,Arial,sans-serif;">%s</td></tr>`,
			brandText, e.BodyHTML)
	}

	preheader := ""
	if e.Preheader != "" {
		preheader = fmt.Sprintf(`<div style="display:none;max-height:0;overflow:hidden;opacity:0;color:transparent;visibility:hidden;font-size:1px;line-height:1px;">%s</div>`, Esc(e.Preheader))
	}

	return fmt.Sprintf(`<!DOCTYPE html>
<html><head><meta charset="utf-8"><meta name="viewport" content="width=device-width,initial-scale=1"><title>%s</title></head>
<body style="margin:0;padding:0;background:%s;font-family:Inter,Helvetica,Arial,sans-serif;color:%s;">
  %s
  <table role="presentation" width="100%%" cellpadding="0" cellspacing="0" border="0" style="background:%s;">
    <tr><td align="center" style="padding:32px 16px;">
      <table role="presentation" width="100%%" cellpadding="0" cellspacing="0" border="0" style="max-width:560px;width:100%%;">
				<tr><td align="center" style="padding-bottom:24px;">
					<a href="%s" style="text-decoration:none;color:%s;display:inline-block;">
						<div style="font-weight:700;letter-spacing:0.12em;text-transform:uppercase;font-size:14px;color:%s;">%s</div>
					</a>
				</td></tr>
        <tr><td style="background:%s;border:1px solid %s;border-radius:18px;overflow:hidden;">
          <table role="presentation" width="100%%" cellpadding="0" cellspacing="0" border="0">
            <tr><td style="padding:32px 32px 12px 32px;">
              <h1 style="margin:0;font-family:Inter,Helvetica,Arial,sans-serif;font-size:22px;line-height:1.25;font-weight:700;color:%s;letter-spacing:0.01em;">%s</h1>
            </td></tr>
            <tr><td style="padding:8px 32px 0 32px;color:%s;font-size:15px;line-height:1.55;">%s</td></tr>
            %s
            %s
            %s
            <tr><td style="padding:28px 32px 32px 32px;"></td></tr>
          </table>
        </td></tr>
        <tr><td align="center" style="padding:20px 16px 0 16px;color:%s;font-size:12px;line-height:1.5;">
          Sent by %s · <a href="%s" style="color:%s;text-decoration:underline;">%s</a>
        </td></tr>
      </table>
    </td></tr>
  </table>
</body></html>`,
		Esc(app), brandBg, brandText,
		preheader,
		brandBg,
		base, brandText,
		brandPrimary, Esc(app),
		brandSurface, brandBorder,
		brandText, heading,
		brandText, intro,
		body,
		cta,
		footnote,
		brandFooterText,
		Esc(app), base, brandPrimary, hostFromURL(base),
	)
}

func escapeWithBreaks(s string) string {
	if s == "" {
		return ""
	}
	parts := strings.Split(strings.TrimSpace(s), "\n\n")
	var b strings.Builder
	for i, p := range parts {
		if i > 0 {
			b.WriteString(`<div style="height:12px;line-height:12px;">&nbsp;</div>`)
		}
		// Keep single newlines as <br>.
		lines := strings.Split(p, "\n")
		for j, ln := range lines {
			if j > 0 {
				b.WriteString("<br>")
			}
			b.WriteString(Esc(ln))
		}
	}
	return b.String()
}

func hostFromURL(u string) string {
	if u == "" {
		return ""
	}
	s := strings.TrimPrefix(strings.TrimPrefix(u, "https://"), "http://")
	if i := strings.IndexAny(s, "/?#"); i >= 0 {
		s = s[:i]
	}
	return s
}
