package deck

type Suit uint8

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker
)

var Suits = [...]Suit{Spade, Diamond, Club, Heart}
