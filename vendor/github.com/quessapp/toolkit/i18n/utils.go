package i18n

import (
	"log"
	"strings"

	"github.com/quessapp/toolkit/i18n/locales"
)

type locale = string
type key = string
type message = string

var translations = map[locale]map[key]message{
	"en-US": *locales.GetAmericanEnglishTranslations(),
	"pt-BR": *locales.GetBrazilianPortugueseTranslation(),
	"es-ES": *locales.GetSpanishTranslations(),
}

func getTranslation(lang, key string) string {
	if translations[lang][key] == "" {
		// we noticed that some keys are not found because the error cames with a dot (.) in the key
		// so we try to find the key without the dot
		return translations[lang][strings.Split(key, ".")[0]]
	}

	return translations[lang][key]
}

// Translate is a function that translates a key to the current locale.
// If the key is not found in the current locale, it will try to find it in the fallback locale (en-US)
// If the key is not found in the fallback locale, it will return the key itself
func Translate(lang, key string) string {
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
