package gamehdl

import (
	"coding-games/internal/core/domain"
	"coding-games/internal/core/port"
	"math/rand/v2"

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
	return ctx.JSON(response)
}

// calculateValue returns the baccarat point total of a hand.
func calculateValue(hand []Card) int {
	sum := 0
	for _, c := range hand {
		v := c.Number
		if v >= 10 { // 10, J, Q, K all count as zero
			v = 0
		}
		sum += v
	}
	return sum % 10
}

// decide applies the baccarat rule: draw on 0–5, stand on 6–9.
func decide(hand []Card) string {
	total := calculateValue(hand)
	if total < 5 || (total == 5 && rand.Float64() < 0.3) {
		return "hit"
	}
	return "stand"
}
