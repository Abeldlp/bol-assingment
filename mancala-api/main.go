package main

import (
	"fmt"

	"github.com/Abeldlp/mancala/entity"
)

func main() {
	player1 := entity.NewPlayer()
	player2 := entity.NewPlayer()

	player1.PlayRoundAgainstOpponent(2, player2)
	player1.PlayRoundAgainstOpponent(3, player2)

	player2.PlayRoundAgainstOpponent(2, player1)
	player2.PlayRoundAgainstOpponent(1, player1)

	player1.PlayRoundAgainstOpponent(4, player2)
	player1.PlayRoundAgainstOpponent(5, player2)

	//Play one
	player1.PlayRoundAgainstOpponent(0, player2)

	//Play two
	// player2.PlayRoundAgainstOpponent(0, player1)
	// player2.PlayRoundAgainstOpponent(1, player1)
	// player2.PlayRoundAgainstOpponent(2, player1)
	// player2.PlayRoundAgainstOpponent(3, player1)

	fmt.Printf("player1: %d\n", player1)
	fmt.Printf("player2: %d\n", player2)
}
