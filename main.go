package main

import (
	"github.com/gofiber/fiber/v2"
	"os"
)

func main() {
	app := fiber.New()

	// Rota simples
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("🚀 Go Fiber no Railway!")
	})

	// Define a porta (Railway define automaticamente a porta como variável de ambiente)
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// Inicia o servidor na porta correta
	app.Listen(":" + port)
}
