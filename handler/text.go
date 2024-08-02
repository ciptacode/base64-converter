package handler

import (
	"encoding/base64"

	"github.com/gofiber/fiber/v2"
)

func Text(c *fiber.Ctx) error {
	return c.Render("text", fiber.Map{
		"Text":   "",
		"Base64": "",
	})
}

func TextSubmit(c *fiber.Ctx) error {
	payload := struct {
		Text   string `json:"text"`
		Base64 string `json:"base64"`
		Type   string `json:"type"`
	}{}

	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	if payload.Type == "encode" {
		base64 := base64.StdEncoding.EncodeToString([]byte(payload.Text))

		return c.Render("text-form", fiber.Map{
			"Text":   payload.Text,
			"Base64": base64,
		})
	}

	text, err := base64.StdEncoding.DecodeString(payload.Base64)

	if err != nil {
		text = []byte("")
	}

	return c.Render("text-form", fiber.Map{
		"Text":      string(text),
		"Base64":    payload.Base64,
		"Base64Err": err,
	})
}
