package i18n

import (
	"log"

	"github.com/quessapp/core-go/configs"
	"github.com/quessapp/core-go/pkg/i18n/locales"
)

type locale = string
type key = string
type message = string

var translations = map[locale]map[key]message{
	"en-US": *locales.GetAmericanEnglishTranslations(),
	"pt-BR": *locales.GetBrazilianPortugueseTranslation(),
}

func getDefaultLang(handlerCtx *configs.HandlersCtx) string {
	accept := handlerCtx.C.Get("Accept-Language")

	if accept == "" {
		log.Printf("[WARNING!!] Accept-Language header not found, defaulting to [en-US]")
		accept = "en-US"
	}

	return accept
}

// Translate translates an key into a human readable message.
// It takes two parameters, a HandlerCtx and an key.
// It returns a string with the translated key.
func Translate(handlerCtx *configs.HandlersCtx, key string) string {
	lang := getDefaultLang(handlerCtx)

	translation := translations[lang][key]

	if translation == "" {
		log.Printf("[WARNING!!] Error [%s] not found in locale [%s]", key, lang)
		return key
	}

	return translation
}
