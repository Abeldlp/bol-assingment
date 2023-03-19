package util

import (
	"log"

	"github.com/Abeldlp/bol-assignment/mancala-api/model"
)

func LogBoard(game model.MancalaGame) {
	reverseBoard := func(board []int) []int {
		for i := 0; i < len(board)/2; i++ {
			j := len(board) - i - 1
			board[i], board[j] = board[j], board[i]
		}
		return board
	}

	lastPlayerMove := func(game *model.MancalaGame) string {
		currentPlayer := game.GetCurentPlayer()

		if &currentPlayer == &game.Player1 {
			return "player2"
		}
		return "player1"
	}

	log.Default().Println("player2: ", game.Player2.Bucket, reverseBoard(game.Player2.Holes))
	log.Default().Println("player1:   ", game.Player1.Holes, game.Player1.Bucket)
	log.Default().Println("last turn: ", lastPlayerMove(&game))
}
