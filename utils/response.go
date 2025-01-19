package utils

import (
	"github.com/gofiber/fiber/v2"
)

// ResponseFormat menggunakan generik T
type ResponseFormat[T any] struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    T      `json:"data"` // Tipe generik
}

// ResponseFormatError digunakan untuk menangani error
type ResponseFormatError struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message interface{} `json:"message"` // Pesan error bisa berupa tipe apa saja
}

// Fungsi untuk mengirimkan respons JSON
func SendJSONResponse[T any](c *fiber.Ctx, code int, status string, message string, data T) error {
	// Buat response menggunakan struct ResponseFormat
	response := ResponseFormat[T]{
		Code:    code,
		Status:  status,
		Message: message,
		Data:    data,
	}
	return c.Status(code).JSON(response)
}

// Fungsi untuk mengirimkan respons error dalam format JSON
func SendJSONResponseError(c *fiber.Ctx, code int, status string, message interface{}) error {
	// Buat response error menggunakan struct ResponseFormatError
	response := ResponseFormatError{
		Code:    code,
		Status:  status,
		Message: message,
	}
	return c.Status(code).JSON(response)
}
