//go:generate stringer -type=Suit,Rank

package deck

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Card struct {
	Suit
	Rank
}

// ToString for each card to show it's Suit and Rank
func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}

// Creates new deck of cards based on constant values
func New(opts ...func([]Card) []Card) []Card {
	var cards []Card
	for _, suit := range Suits {
		for rank := MinRank; rank <= MaxRank; rank++ {
			cards = append(cards, Card{Suit: suit, Rank: rank})
		}
	}
	for _, opt := range opts {
		cards = opt(cards)
	}
	return cards
}

// Uses sort module for custom sorting
func Sort(less func(cards []Card) func(i, j int) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		sort.Slice(cards, less(cards))
		return cards
	}
}

// Uses sort module for default sorting
func DefaultSort(cards []Card) []Card {
	sort.Slice(cards, Less(cards))
	return cards
}

// Checks to see if index i card is less than index j card for sorting
func Less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return absRank(cards[i]) < absRank(cards[j])
	}
}

// Used for sorting via the rank of card
func absRank(c Card) int {
	return int(c.Suit)*int(MaxRank) + int(c.Rank)
}

// Shuffles the deck of cards in random order
func Shuffle(cards []Card) []Card {
	ret := make([]Card, len(cards))
	r := rand.New(rand.NewSource((time.Now().Unix())))
	perm := r.Perm((len(cards)))
	for i, j := range perm {
		ret[i] = cards[j]
	}
	return ret
}

// Creates a number n of Jokers in the deck
func Jokers(n int) func([]Card) []Card {
	return func(cards []Card) []Card {
		for i := 0; i < n; i++ {
			cards = append(cards, Card{
				Rank: Rank(i),
				Suit: Joker,
			})
		}
		return cards
	}
}

// Creates n many decks of cards together
func Deck(n int) func([]Card) []Card {
	return func(cards []Card) []Card {
		var ret []Card
		for i := 0; i < n; i++ {
			ret = append(ret, cards...)
		}
		return ret
	}
}
