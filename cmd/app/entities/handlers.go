package entities

import (
	"github.com/gofiber/fiber/v2"
)

// HandlersContext is a global model for handlers. It defines the fiber context, app context, etc..
// Use HandlersContext to avoid long function params.
type HandlersContext struct {
	// Context from fiber.
	C *fiber.Ctx
	// App config.
	AppCtx
}
