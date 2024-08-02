package handler

import (
	"encoding/base64"
	"fmt"
	"io"
	"mime"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

func File(c *fiber.Ctx) error {
	return c.Render("file", fiber.Map{})
}

func FileSubmit(c *fiber.Ctx) error {
	form, err := c.MultipartForm()

	if err != nil {
		return c.Render("file-result", fiber.Map{
			"Base64Err": "File invalid",
		})
	}

	files := form.File["file"]
	if len(files) == 0 {
		return c.Render("file-result", fiber.Map{
			"Base64Err": "No file uploaded",
		})
	}

	file := files[0]
	fileContent, err := file.Open()

	if err != nil {
		return c.Render("file-result", fiber.Map{
			"Base64Err": "File invalid",
		})
	}

	defer fileContent.Close()

	content, err := io.ReadAll(fileContent)
	if err != nil {
		return c.Render("file-result", fiber.Map{
			"Base64Err": "File invalid",
		})
	}

	ext := filepath.Ext(file.Filename)
	mimeType := mime.TypeByExtension(ext)

	if mimeType == "" {
		mimeType = "application/octet-stream"
	}

	base64 := base64.StdEncoding.EncodeToString(content)

	dataURL := fmt.Sprintf("data:%s;base64,%s", mimeType, base64)

	return c.Render("file-result", fiber.Map{
		"Base64": dataURL,
	})
}
