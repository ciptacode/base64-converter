package main

import (
	"encoding/base64"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

func main() {
	LoadEnv()

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", Welcome)
	app.Post("/", EncodeBase64)

	app.Listen(":" + os.Getenv("PORT"))
}

func LoadEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error load .env file")
	}
}

func EncodeBase64(c *fiber.Ctx) error {
	payload := struct {
		Text string `json:"text"`
	}{}

	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	result := base64.StdEncoding.EncodeToString([]byte(payload.Text))

	return c.Render("index", fiber.Map{
		"Text":   payload.Text,
		"Result": result,
	})
}

func Welcome(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Text":   "",
		"Result": "",
	})
}
