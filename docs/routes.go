package docs

import (
	"fmt"

	"github.com/quessapp/core-go/configs"

	"github.com/gofiber/swagger"
)

// LoadRoutes loads all docs routes.
func LoadRoutes(AppCtx *configs.AppCtx) {
	g := AppCtx.App.Group("/docs")

	docsURI := fmt.Sprintf("%s%s/docs/doc.json", AppCtx.Cfg.App.ServerHost, AppCtx.Cfg.App.ServerPort)

	g.Get("/*", swagger.New(swagger.Config{
		URL:          docsURI,
		DeepLinking:  false,
		DocExpansion: "none",
	}))
}
