package gamehdl

import (
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
	for _, hand := range req.PlayHands {
		if shouldHit(hand) {
			decisions = append(decisions, "hit")
		} else {
			decisions = append(decisions, "stand")
		}
	}

	for _, hand := range req.KnownHands {
		if shouldHit(hand) {
			decisions = append(decisions, "hit")
		} else {
			decisions = append(decisions, "stand")
		}
	}

	// Example handler logic
	// This could be replaced with actual game logic
	return ctx.JSON(fiber.Map{
		"response": decisions,
	})
}

// shouldHit returns true if the expected value of hitting exceeds standing.
func shouldHit(hand []Card) bool {
	deck := fullDeck()
	removeCards(deck, hand)

	// EV of standing
	initType := handType(hand)
	initWinProb := winProbability(hand, deck)
	evStand := initWinProb * float64(initType)

	// EV of hitting: average over all possible third cards
	var evHitSum float64
	candidates := deckKeys(deck)
	for _, c := range candidates {
		// simulate draw
		newHand := append(hand, c)
		deck2 := copyDeck(deck)
		delete(deck2, c)
		wp := winProbability(newHand, deck2)
		evHitSum += wp * float64(handType(newHand))
	}
	evHit := evHitSum / float64(len(candidates))

	return evHit > evStand
}

// winProbability computes P(win) vs all 2-card opponent hands from deck.
func winProbability(myHand []Card, deck map[Card]bool) float64 {
	cards := deckKeys(deck)
	total := 0
	win := 0
	for i := 0; i < len(cards); i++ {
		for j := i + 1; j < len(cards); j++ {
			opp := []Card{cards[i], cards[j]}
			total++
			switch compareHands(myHand, opp) {
			case 1:
				win++
			}
		}
	}
	if total == 0 {
		return 0
	}
	return float64(win) / float64(total)
}

// compareHands returns 1 if a wins, -1 if b wins, 0 if tie.
func compareHands(a, b []Card) int {
	at, as := handType(a), handType(b)
	if at > as {
		return 1
	} else if at < as {
		return -1
	}
	// same type: tie-break by highest number or by score mod 10
	if len(a) == 2 && a[0].Number == a[1].Number {
		// pair: compare number
		if a[0].Number > b[0].Number {
			return 1
		} else if a[0].Number < b[0].Number {
			return -1
		}
		return 0
	}
	// else compare mod10 score
	sa, sb := scoreMod10(a), scoreMod10(b)
	if sa > sb {
		return 1
	} else if sa < sb {
		return -1
	}
	return 0
}

// handType returns payout multiplier: set=5, strange(flush)=3, pair=2, normal=1.
func handType(h []Card) int {
	if len(h) == 3 {
		// set?
		if h[0].Number == h[1].Number && h[1].Number == h[2].Number {
			return 5
		}
		// flush?
		if h[0].Suit == h[1].Suit && h[1].Suit == h[2].Suit {
			return 3
		}
	}
	if len(h) == 2 && h[0].Number == h[1].Number {
		return 2
	}
	return 1
}

// scoreMod10 sums pip values mod10.
func scoreMod10(h []Card) int {
	sum := 0
	for _, c := range h {
		v := c.Number
		if v > 10 {
			v = 0
		}
		sum += v
	}
	return sum % 10
}

// fullDeck builds a map of all 52 cards.
func fullDeck() map[Card]bool {
	suits := []string{"hearts", "diamonds", "clubs", "spades"}
	deck := make(map[Card]bool, 52)
	for _, s := range suits {
		for n := 1; n <= 13; n++ {
			deck[Card{n, s}] = true
		}
	}
	return deck
}

// removeCards removes each c from deck.
func removeCards(deck map[Card]bool, cards []Card) {
	for _, c := range cards {
		delete(deck, c)
	}
}

// copyDeck shallow-copies deck.
func copyDeck(deck map[Card]bool) map[Card]bool {
	new := make(map[Card]bool, len(deck))
	for c := range deck {
		new[c] = true
	}
	return new
}

// deckKeys returns slice of cards present.
func deckKeys(deck map[Card]bool) []Card {
	out := make([]Card, 0, len(deck))
	for c := range deck {
		out = append(out, c)
	}
	return out
}
