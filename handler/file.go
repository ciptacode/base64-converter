package handler

import (
	"encoding/base64"
	"fmt"
	"io"
	"mime"
	"os"
	"path/filepath"
	"strings"
	"time"

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

func Base64ToFile(c *fiber.Ctx) error {
	return c.Render("base64-to-file", fiber.Map{})
}

func Base64ToFileSubmit(c *fiber.Ctx) error {
	payload := struct {
		Base64 string `json:"base64"`
	}{}

	if err := c.BodyParser(&payload); err != nil {
		return c.Render("base64-to-file-result", fiber.Map{
			"Base64Err": "Base64 is invalid",
		})
	}
	parts := strings.Split(payload.Base64, ",")
	if len(parts) != 2 {
		return c.Render("base64-to-file-result", fiber.Map{
			"Base64Err": "invalid base64 format",
		})
	}

	mimeParts := strings.Split(parts[0], ":")
	if len(mimeParts) != 2 {
		return c.Render("base64-to-file-result", fiber.Map{
			"Base64Err": "invalid MIME type",
		})
	}

	mimeType := mimeParts[1]
	mimeTypeParts := strings.Split(mimeType, "/")
	if len(mimeTypeParts) != 2 {
		return c.Render("base64-to-file-result", fiber.Map{
			"Base64Err": "invalid MIME type",
		})
	}

	ext := mimeTypeParts[1]

	extParts := strings.Split(ext, ";")
	ext = extParts[0]

	data, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		return c.Render("base64-to-file-result", fiber.Map{
			"Base64Err": "Base64 is invalid",
		})
	}

	filename := fmt.Sprintf("%d.%s", time.Now().UnixNano(), ext)
	dir := filepath.Join("public", "result")
	filepath := filepath.Join(dir, filename)

	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return c.Render("base64-to-file-result", fiber.Map{
			"Base64Err": "Base64 is invalid",
		})
	}

	err = os.WriteFile(filepath, data, 0644)
	if err != nil {
		return c.Render("base64-to-file-result", fiber.Map{
			"Base64Err": "Base64 is invalid",
		})
	}

	return c.Render("base64-to-file-result", fiber.Map{
		"Base64": payload.Base64,
		"Result": filename,
	})
}
