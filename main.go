package main

import (
	"blackjackgame/bj"
	"fmt"
)

func main() {
	var gs bj.GameState
	gs = bj.Shuffle(gs)
	var input string
gameLoop:
	// Can't continue if theres less than half a deck in shoe to prevent card counting
	for input != "q" && len(gs.Deck) > 31 {
		gs = bj.Deal(gs)
		pScore, dScore := gs.Player.Score(), gs.Dealer.Score()
		// Determines if either player or dealer wins off rip
		if pScore == 21 || dScore == 21 {
			gs = bj.EndHand(gs)

		} else {
			// Player choices
			for gs.State == bj.StatePlayerTurn {
				fmt.Println("Player:", gs.Player)
				fmt.Println("Dealer:", gs.Dealer.DealerString())
				fmt.Println("Do You want to (h)it, (s)tand, (q)uit")
				fmt.Scanf("%s\n", &input)
				switch input {
				case "h":
					gs = bj.Hit(gs)
				case "s":
					gs = bj.Stand(gs)
				case "q":
					break gameLoop
				default:
					fmt.Println("Invalid Option: ", input)
				}
			}
			// Dealer choice conditions
			for gs.State == bj.StateDealerTurn {
				if gs.Dealer.Score() <= 16 || (gs.Dealer.Score() == 17 && gs.Dealer.MinScore() != 17) {
					gs = bj.Hit(gs)
				} else {
					gs = bj.Stand(gs)
				}
			}

			gs = bj.EndHand(gs)
		}
		fmt.Println()
	}
	if len(gs.Deck) <= 31 {
		fmt.Println("Ran out of cards, please restart the game!")
	}
}
