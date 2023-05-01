package i18n

import (
	"log"

	"github.com/quessapp/core-go/configs"
	toolkitI18n "github.com/quessapp/toolkit/i18n"
)

func getLang(handlerCtx *configs.HandlersCtx) string {
	accept := handlerCtx.C.Get("Accept-Language")

	if accept == "" {
		log.Printf("[WARNING!!] Accept-Language header not found, defaulting to [en-US]")
		accept = "en-US"
	}

	return accept
}

// Translate translates re
// It takes two parameters, a HandlerCtx and an key.
// It returns a string with the translated key.
func Translate(handlerCtx *configs.HandlersCtx, key string) string {
	lang := getLang(handlerCtx)
	return toolkitI18n.Translate(lang, key)
}
