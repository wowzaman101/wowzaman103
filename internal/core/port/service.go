package port

import "coding-games/internal/core/domain"

type Service interface {
	DoYouTrustInLuck(request *domain.Request) (domain.Response, error)
}

// type Card struct {
// 	Number Number `json:"number"`
// 	Suit   Suit   `json:"suit"`
// }

// type Request struct {
// 	PlayHands  [][]Card `json:"playHands"`
// 	KnownHands [][]Card `json:"knownHands,omitempty"`
// }

// type Number int
// type Suit string

// const (
// 	Hearts   Suit   = "hearts"
// 	Diamonds Suit   = "diamonds"
// 	Clubs    Suit   = "clubs"
// 	Spades   Suit   = "spades"
// 	Ace      Number = iota + 1
// 	Two
// 	Three
// 	Four
// 	Five
// 	Six
// 	Seven
// 	Eight
// 	Nine
// 	Ten
// 	Jack
// 	Queen
// 	King
// )

// type Response struct {
// 	Response []string `json:"response"`
// }
