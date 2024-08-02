package handler

import "github.com/gofiber/fiber/v2"

func File(c *fiber.Ctx) error {
	return c.SendString("File")
}

func FileSubmit(c *fiber.Ctx) error {
	return c.SendString("File")
}
