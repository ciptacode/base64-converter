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
		Views:     engine,
		BodyLimit: 10 * 1024 * 1024,
	})

	app.Get("/", handler.Text)
	app.Post("/text", handler.TextSubmit)

	app.Get("/file-to-base64", handler.File)
	app.Post("/file-to-base64", handler.FileSubmit)

	app.Listen(":" + os.Getenv("PORT"))
}

func LoadEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error load .env file")
	}
}
