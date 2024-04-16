package utils

import (
	"mime"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func IsEnabled(key bool) func(c *fiber.Ctx) bool {
	if key {
		return nil
	}

	return func(c *fiber.Ctx) bool { return true }
}

// utility funtion to help convert an array of string to convert to comma separeted string format
func StringArrayToStringConvertor(stringArray []string) string {
	var stringToBeReturned string
	for _, val := range stringArray {
		stringToBeReturned += (val + ", ")
	}
	return stringToBeReturned
}

// Get name of the file without extension
// abc.txt -> abc
func TrimFileExt(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}

// Generate S3 Object Name
func GenerateS3ObjectName(mediaID string, mimeType string) string {
	fileExt, _ := mime.ExtensionsByType(mimeType)

	return filepath.Join(mediaID + fileExt[0])
}
