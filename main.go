package main

import (
	"log"
	"os"

	"github.com/ciptacode/base64-converter/handler"
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

	app.Get("/", handler.Text)
	app.Post("/text", handler.TextSubmit)

	app.Get("/file", handler.File)

	app.Listen(":" + os.Getenv("PORT"))
}

func LoadEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error load .env file")
	}
}
