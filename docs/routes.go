package docs

import (
	"core/configs"
	"fmt"

	"github.com/gofiber/swagger"
)

// LoadRoutes loads all docs routes.
func LoadRoutes(AppCtx *configs.AppCtx) {
	g := AppCtx.App.Group("/docs")

	docsURI := fmt.Sprintf("%s%s/docs/doc.json", AppCtx.Cfg.ServerHost, AppCtx.Cfg.ServerPort)

	fmt.Println(docsURI)
	g.Get("/*", swagger.New(swagger.Config{ // custom
		URL:         docsURI,
		DeepLinking: false,

		DocExpansion: "none",
	}))
}
