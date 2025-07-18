package gamehdl

import (
	"coding-games/internal/core/domain"
	"coding-games/internal/core/port"

	"github.com/gofiber/fiber/v3"
)

type handler struct {
	svc port.Service
}

type Handler interface {
	DoYouTrustInLuck(ctx fiber.Ctx) error
}

func New(
	svc port.Service,
) Handler {
	return &handler{
		svc: svc,
	}
}

func (h *handler) DoYouTrustInLuck(ctx fiber.Ctx) error {
	request := new(domain.Request)
	if err := ctx.Bind().Body(request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	response, err := h.svc.DoYouTrustInLuck(request)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to process request",
		})
	}
	return ctx.JSON(response.Response)
}
