package lucksvc

import (
	"coding-games/internal/core/domain"
	"coding-games/internal/core/port"
	"math"
	"math/rand"
	"time"
)

type DRAW string

const (
	HIT   DRAW = "hit"
	STAND DRAW = "stand"
)

func (d DRAW) String() string {
	return string(d)
}

type service struct {
}

func New() port.Service {
	return &service{}
}

func (s *service) DoYouTrustInLuck(request *domain.Request) (domain.Response, error) {
	res := make([]string, len(request.PlayHands))
	for i, hand := range request.PlayHands {
		res[i] = calculateHit(hand).String()
	}
	return domain.Response{
		Response: res,
	}, nil
}

func calculateHit(hand []domain.Card) DRAW {
	if len(hand) != 2 {
		return HIT
	}

	n1 := normalizeNumber(int(hand[0].Number))
	n2 := normalizeNumber(int(hand[1].Number))

	sum := (n1 + n2) % 10
	sameSuit := hand[0].Suit == hand[1].Suit
	sameNumber := hand[0].Number == hand[1].Number

	if sum == 8 || sum == 9 {
		return STAND
	}

	if isFaceCard(int(hand[0].Number)) && isFaceCard(int(hand[1].Number)) {
		return HIT
	}
	if isPossibleStraight(int(hand[0].Number), int(hand[1].Number)) && sum < 7 && randomSeed(200) {
		return HIT
	}
	if isPossibleStraight(int(hand[0].Number), int(hand[1].Number)) && sum < 4 && randomSeed(700) {
		return HIT
	}
	if sameSuit && sum < 6 && randomSeed(700) {
		return HIT
	}
	if sameSuit && sum < 4 && randomSeed(900) {
		return HIT
	}
	if sameNumber && sum < 8 && randomSeed(300) {
		return HIT
	}
	if sameNumber && sum < 5 && randomSeed(800) {
		return HIT
	}

	return STAND
}

func normalizeNumber(n int) int {
	if n > 10 {
		return 0
	}
	return n
}

func isFaceCard(n int) bool {
	return n == 11 || n == 12 || n == 13
}

func isPossibleStraight(a, b int) bool {
	diff := int(math.Abs(float64(a - b)))
	return diff == 1 || diff == 2
}

func randomSeed(input int) bool {
	randomNum := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(1000)
	return input > randomNum
}
