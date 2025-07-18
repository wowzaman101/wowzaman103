package server

import (
	"coding-games/internal/handler/gamehdl"

	"github.com/gofiber/fiber/v3"
)

func New(gh gamehdl.Handler) *fiber.App {

	app := fiber.New()
	// Health check route
	app.Get("/health", func(c fiber.Ctx) error {
		return c.SendString("OK")
	})

	// Game handler route
	// app.Get("/game/test", gh.Test)

	app.Post("/game/lucky", gh.DoYouTrustInLuck)

	return app
}
