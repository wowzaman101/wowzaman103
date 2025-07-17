package gamehdl

import "github.com/gofiber/fiber/v3"

type handler struct {
}

type Handler interface {
	Test(ctx fiber.Ctx) error
}

func New() Handler {
	return &handler{}
}

func (h *handler) Test(ctx fiber.Ctx) error {
	// Example handler logic
	return ctx.JSON(fiber.Map{
		"message": "Test endpoint hit",
	})
}
