package main

import (
	"blackjackgame/deck"
	"fmt"
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

type State int8

const (
	StatePlayerTurn State = iota
	StateDealerTurn
	StateHandOver
)

type GameState struct {
	Deck   []deck.Card
	State  State
	Player Hand
	Dealer Hand
}

// Games state of which player is drawing cards
func (gs *GameState) CurrentPlayer() *Hand {
	switch gs.State {
	case StatePlayerTurn:
		return &gs.Player
	case StateDealerTurn:
		return &gs.Dealer
	default:
		panic("It isn't currently any players turn")
	}
}

// Clones the game state
func clone(gs GameState) GameState {
	ret := GameState{
		Deck:   make([]deck.Card, len(gs.Deck)),
		State:  gs.State,
		Player: make(Hand, len(gs.Player)),
		Dealer: make(Hand, len(gs.Dealer)),
	}
	copy(ret.Deck, gs.Deck)
	copy(ret.Player, gs.Player)
	copy(ret.Dealer, gs.Dealer)
	return ret
}

// Draws the card and removes it from cards deck
func draw(cards []deck.Card) (deck.Card, []deck.Card) {
	return cards[0], cards[1:]
}

// Gets the min
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Creates a new shoe of cards with 8 decks in a shoe
func Shuffle(gs GameState) GameState {
	ret := clone(gs)
	ret.Deck = deck.New(deck.Deck(8), deck.Shuffle)
	return ret
}

// Deal the cards to each player and dealer
func Deal(gs GameState) GameState {
	ret := clone(gs)
	ret.Player = make(Hand, 0, 9)
	ret.Dealer = make(Hand, 0, 9)
	var card deck.Card
	for i := 0; i < 2; i++ {
		card, ret.Deck = draw(ret.Deck)
		ret.Player = append(ret.Player, card)
		card, ret.Deck = draw(ret.Deck)
		ret.Dealer = append(ret.Dealer, card)
	}
	ret.State = StatePlayerTurn
	return ret
}

// Removes random card from the top of the shoe
func Hit(gs GameState) GameState {
	ret := clone(gs)
	hand := ret.CurrentPlayer()
	var card deck.Card
	card, ret.Deck = draw(ret.Deck)
	*hand = append(*hand, card)
	if hand.Score() > 21 {
		return Stand(ret)
	}
	return ret
}

// Goes to the dealer or ends the game, changes the game state
func Stand(gs GameState) GameState {
	ret := clone(gs)
	ret.State++
	return ret
}

// Determines end game state where you win, lose or draw
func EndHand(gs GameState) GameState {
	ret := clone(gs)
	pScore, dScore := ret.Player.Score(), ret.Dealer.Score()
	fmt.Println("--- Final Hands ---")
	fmt.Println("Player:", ret.Player, "\nScore:", pScore)
	fmt.Println("Dealer:", ret.Dealer, "\nScore:", dScore)
	switch {
	case pScore > 21:
		fmt.Println("You Busted, You Lose")
	case dScore > 21:
		fmt.Println("Dealer Busted, You Win")
	case pScore > dScore:
		fmt.Println("You Win")
	case pScore < dScore:
		fmt.Println("You Lose")
	case pScore == dScore:
		fmt.Println("You Draw With The Dealer")
	}
	ret.Player = nil
	ret.Dealer = nil
	return ret
}

func main() {
	var gs GameState
	gs = Shuffle(gs)
	var input string
gameLoop:
	// Can't continue if theres less than half a deck in shoe to prevent card counting
	for input != "q" && len(gs.Deck) > 31 {
		gs = Deal(gs)
		pScore, dScore := gs.Player.Score(), gs.Dealer.Score()
		// Determines if either player or dealer wins off rip
		if pScore == 21 || dScore == 21 {
			gs = EndHand(gs)

		} else {
			// Player choices
			for gs.State == StatePlayerTurn {
				fmt.Println("Player:", gs.Player)
				fmt.Println("Dealer:", gs.Dealer.DealerString())
				fmt.Println("Do You want to (h)it, (s)tand, (q)uit")
				fmt.Scanf("%s\n", &input)
				switch input {
				case "h":
					gs = Hit(gs)
				case "s":
					gs = Stand(gs)
				case "q":
					break gameLoop
				default:
					fmt.Println("Invalid Option: ", input)
				}
			}

			for gs.State == StateDealerTurn {
				if gs.Dealer.Score() <= 16 || (gs.Dealer.Score() == 17 && gs.Dealer.MinScore() != 17) {
					gs = Hit(gs)
				} else {
					gs = Stand(gs)
				}
			}

			gs = EndHand(gs)
		}
		fmt.Println()
	}
	if len(gs.Deck) == 0 {
		fmt.Println("Ran out of cards, please restart the game!")
	}
}
