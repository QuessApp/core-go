package i18n

import (
	"log"
	"strings"

	"github.com/quessapp/core-go/configs"
	"github.com/quessapp/core-go/pkg/i18n/locales"
)

type locale = string
type key = string
type message = string

var translations = map[locale]map[key]message{
	"en-US": *locales.GetAmericanEnglishTranslations(),
	"pt-BR": *locales.GetBrazilianPortugueseTranslation(),
	"es-ES": *locales.GetSpanishTranslations(),
}

func getLang(handlerCtx *configs.HandlersCtx) string {
	accept := handlerCtx.C.Get("Accept-Language")

	if accept == "" {
		log.Printf("[WARNING!!] Accept-Language header not found, defaulting to [en-US]")
		accept = "en-US"
	}

	return accept
}

func getTranslation(lang, key string) string {
	if translations[lang][key] == "" {
		// we noticed that some keys are not found because the error cames with a dot (.) in the key
		// so we try to find the key without the dot
		return translations[lang][strings.Split(key, ".")[0]]
	}

	return translations[lang][key]
}

// Translate translates re
// It takes two parameters, a HandlerCtx and an key.
// It returns a string with the translated key.
func Translate(handlerCtx *configs.HandlersCtx, key string) string {
	lang := getLang(handlerCtx)

	foundTranslation := getTranslation(lang, key)

	if foundTranslation != "" {
		return foundTranslation
	}

	if foundTranslation == "" {
		log.Printf("[WARNING!!] Key [%s] not found in locale [%s]", key, lang)

		fallback := getTranslation("en-US", key)

		// If the key is not found in the current locale, we try to find it in the fallback locale (en-US)
		if fallback == "" {
			log.Printf("[WARNING!!] Key [%s] not found in fallback locale (en-US)", key)
			return key
		}

		return fallback
	}

	return key
}
