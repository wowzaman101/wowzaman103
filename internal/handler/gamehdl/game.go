package gamehdl

import (
	"math/rand/v2"

	"github.com/gofiber/fiber/v3"
)

type handler struct {
}

type Handler interface {
	Test(ctx fiber.Ctx) error
}

func New() Handler {
	return &handler{}
}

type Card struct {
	Number int    // 1=A, 2–9, 10, 11=J, 12=Q, 13=K
	Suit   string // hearts, diamonds, clubs, spades (unused for value)
}
type Request struct {
	PlayHands  [][]Card `json:"playHands"`
	KnownHands [][]Card `json:"knownHands"` // unused in this example
}

type Response struct {
	Data []string `json:"response"`
}

func (h *handler) Test(ctx fiber.Ctx) error {
	var req Request
	if err := ctx.Bind().Body(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	decisions := make([]string, 0)
	for _, game := range req.PlayHands {
		decisions = append(decisions, decide(game))
	}

	for _, game := range req.KnownHands {
		decisions = append(decisions, decide(game))
	}

	// Example handler logic
	// This could be replaced with actual game logic
	return ctx.JSON(fiber.Map{
		"response": decisions,
	})
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
