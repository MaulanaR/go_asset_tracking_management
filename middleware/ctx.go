package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"

	"github.com/maulanar/go_asset_tracking_management/app"
)

func Ctx() *ctxHandler {
	if ch == nil {
		ch = &ctxHandler{}
	}
	return ch
}

var ch *ctxHandler

type ctxHandler struct{}

func (*ctxHandler) New(c *fiber.Ctx) error {
	lang := c.Get("Accept-Language")
	if lang == "" || lang == "*" || strings.Contains(lang, ",") || strings.Contains(lang, ";") {
		lang = c.Get("Content-Language") // backward compatibility
		if lang == "" {
			lang = "en"
		}
	}
	action := app.Action{
		Method:   c.Method(),
		EndPoint: c.Path(),
		Referer:  c.Get("Referer"),
		BaseURL:  c.BaseURL(),
		Path:     c.Path(),
		URL:      c.Request().URI().String(),
		IP:       strings.Join(c.IPs(), ", "),
	}
	action.Referer, _, _ = strings.Cut(action.Referer, "?")
	path := strings.Split(action.EndPoint, "/")
	pathLen := len(path)
	if pathLen > 3 {
		action.EndPoint = path[3]
	}
	if pathLen > 4 {
		action.DataID = path[4]
	}

	requestID := app.Crypto().NewToken()
	ctx := app.Ctx{
		RequestID: requestID,
		Lang:      lang,
		Action:    action,
	}

	c.Locals("ctx", &ctx)

	return c.Next()
}
