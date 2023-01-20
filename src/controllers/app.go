package controllers

import "github.com/gofiber/fiber/v2"

// PingAppController prints a ``pong``.
func PingAppController(c *fiber.Ctx) error {
	return c.SendString("pong")
}
