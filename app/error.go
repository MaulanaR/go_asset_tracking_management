package app

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/gofiber/fiber/v2"
	"grest.dev/grest"
)

// Error returns a pointer to the errorUtil instance (eu).
// If eu is not initialized, it creates a new errorUtil instance and assigns it to eu.
// It ensures that only one instance of errorUtil is created and reused.
func Error() *errorUtil {
	if eu == nil {
		eu = &errorUtil{}
	}
	return eu
}

// eu is a pointer to an errorUtil instance.
// It is used to store and access the singleton instance of errorUtil.
var eu *errorUtil

// errorUtil represents an error utility.
// It embeds grest.Error, which indicates that errorUtil inherits from grest.Error.
type errorUtil struct {
	grest.Error
}

// New creates a new error instance based on the provided status code, message, and optional details.
// It returns the created error instance.
func (errorUtil) New(statusCode int, message string, detail ...any) error {
	return grest.NewError(statusCode, message, detail...)
}

// StatusCode retrieves the status code from an error.
// It checks if the error is an instance of grest.Error or fiber.Error.
// If the error is of either type, it returns the corresponding status code.
// Otherwise, it returns http.StatusInternalServerError.
func (errorUtil) StatusCode(err error) int {
	e, ok := err.(*grest.Error)
	if ok {
		return e.StatusCode()
	}
	f, ok := err.(*fiber.Error)
	if ok {
		return f.Code
	}
	return http.StatusInternalServerError
}

// Detail retrieves the details from an error.
// It checks if the error is an instance of grest.Error.
// If it is, it returns the body of the error.
// Otherwise, it returns nil.
func (errorUtil) Detail(err error) any {
	e, ok := err.(*grest.Error)
	if ok {
		respError := map[string]any{
			"code":    e.Code,
			"message": e.Message,
		}
		if e.Detail != nil {
			respError["data"] = e.Detail
		}
		return respError
	}
	return nil
}

// Trace retrieves the trace information from an error.
// It checks if the error is an instance of grest.Error.
// If it is, it returns the trace information.
// Otherwise, it returns nil.
func (errorUtil) Trace(err error) []map[string]any {
	e, ok := err.(*grest.Error)
	if ok {
		return e.Trace()
	}
	return nil
}

// TraceSimple retrieves simplified trace information from an error.
// It checks if the error is an instance of grest.Error.
// If it is, it returns the simplified trace information.
// Otherwise, it returns nil.
func (errorUtil) TraceSimple(err error) map[string]string {
	e, ok := err.(*grest.Error)
	if ok {
		return e.TraceSimple()
	}
	return nil
}

// Handler handles errors by processing them and returning an appropriate response.
// It retrieves the language from the context (c) and assigns it to lang.
// It checks if the error is an instance of grest.Error.
// If it is not, it sets the error code and message based on the received error.
// If the error status code is not in the 4xx or 5xx range, it sets the code to http.StatusInternalServerError.
// If the error status code is http.StatusInternalServerError, it translates the error message and assigns it to e.Message.
// It returns a JSON response with the error status code and body.
func (errorUtil) Handler(c *fiber.Ctx, err error) error {
	lang := "en"
	if ctx, ok := c.Locals("ctx").(*Ctx); ok && ctx != nil {
		lang = ctx.Lang
	}

	// Pastikan e SELALU non-nil
	var e *grest.Error
	if ge, ok := err.(*grest.Error); ok && ge != nil {
		e = ge
	} else {
		code := http.StatusInternalServerError
		if fe, ok := err.(*fiber.Error); ok && fe != nil {
			code = fe.Code
		}
		e = grest.NewError(code, err.Error(), nil)
	}

	// Normalisasi status code
	if e.StatusCode() < 400 || e.StatusCode() > 599 {
		e.Code = http.StatusInternalServerError
	}

	// Pesan 500 → pakai translator + simpan pesan asli ke detail bila kosong
	if e.StatusCode() == http.StatusInternalServerError {
		e.Message = Translator().Trans(lang, "500_internal_error")
		if e.Detail == nil {
			e.Detail = map[string]any{"message": err.Error()}
		}
	}

	respError := map[string]any{
		"code":    e.Code,
		"message": e.Message,
	}
	if e.Detail != nil {
		respError["data"] = e.Detail
	}

	return c.Status(e.StatusCode()).JSON(respError)
}

// Recover recovers from a panic during Fiber request processing.
// It uses a defer statement to catch and recover from panics.
// Inside the deferred function, there is a placeholder for saving logs and sending alerts.
func (errorUtil) Recover(c *fiber.Ctx) (err error) {
	defer func() {
		if r := recover(); r != nil {
			// pastikan kita ubah apa pun tipenya ke string
			var msg string
			switch v := r.(type) {
			case error:
				msg = v.Error()
			default:
				msg = fmt.Sprintf("%v", v)
			}

			// siapkan detail (trace optional)
			detail := map[string]any{
				"trace": string(debug.Stack()),
			}

			// bungkus sebagai grest.Error dan tulis response JSON
			_ = Error().Handler(c, grest.NewError(http.StatusInternalServerError, msg, detail))
		}
	}()
	return c.Next()
}
