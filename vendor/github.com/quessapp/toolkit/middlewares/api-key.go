package middlewares

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

type Config struct {
	Next func(c *fiber.Ctx) bool // Required
	// ErrorHandler defines a function which is executed for an invalid API key.
	// It may be used to define a custom error.
	ErrorHandler fiber.ErrorHandler
	// Header name to get value from request headers. Example: x-api-key
	//
	// Default is "api-key"
	HeaderName string
	// Key is the value of Header name/the key of the secret API key.
	Key string
	// WrongKeyMessage is the message error if provided key is invalid.
	//
	// Default is "Wrong API key"
	WrongKeyMessage string
	// MissingHeaderMessage is the message error if provided header name field is invalid.
	//
	// Default is "Missing API key"
	MissingHeaderMessage string
}

var ConfigDefault = Config{
	Next:                 nil,
	HeaderName:           "api-key",
	WrongKeyMessage:      "Wrong API key",
	MissingHeaderMessage: "Missing API key",
}

func New(config Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if config.Next(c) {
			return c.Next()
		}

		if config.Key == "" {
			return config.ErrorHandler(c, errors.New("no API key provided. please, provide a value for 'config.Key' field"))
		}

		if config.HeaderName == "" {
			config.HeaderName = ConfigDefault.HeaderName
		}

		if config.MissingHeaderMessage == "" {
			config.MissingHeaderMessage = ConfigDefault.MissingHeaderMessage
		}

		if config.WrongKeyMessage == "" {
			config.WrongKeyMessage = ConfigDefault.WrongKeyMessage
		}

		h := c.Get(config.HeaderName)

		if h == "" {
			return config.ErrorHandler(c, errors.New(config.MissingHeaderMessage))
		}

		if h != config.Key {
			return config.ErrorHandler(c, errors.New(config.WrongKeyMessage))
		}

		return c.Next()
	}
}
