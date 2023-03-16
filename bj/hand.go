package bj

import (
	"blackjackgame/deck"
	"strings"
)

type Hand []deck.Card

// ToString method for each Hand
func (h Hand) String() string {
	strs := make([]string, len(h))
	for i := range h {
		strs[i] = h[i].String()
	}
	return strings.Join(strs, ", ")
}

// The initial state shown of the Dealer Hand
func (h Hand) DealerString() string {
	return h[0].String() + ", *HIDDEN*"
}

// Levys upper quartile cards to a value of 10 i.e. Jack, Queen, King
func (h Hand) MinScore() int {
	score := 0
	for _, card := range h {
		score += min(int(card.Rank), 10)
	}
	return score
}

// Assigns most logical score for player based on Ace rule
func (h Hand) Score() int {
	minScore := h.MinScore()
	if minScore > 11 {
		return minScore
	}
	// Ace card handler for either 1 or 11
	for _, card := range h {
		if card.Rank == deck.Ace {
			return minScore + 10
		}
	}
	return minScore
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
