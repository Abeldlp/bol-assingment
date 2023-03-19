package util

import (
	"log"

	"github.com/Abeldlp/bol-assignment/mancala-api/model"
)

func ReverseBoard(board []int) []int {
	for i := 0; i < len(board)/2; i++ {
		j := len(board) - i - 1
		board[i], board[j] = board[j], board[i]
	}
	return board
}

func LogBoard(game model.MancalaGame) {
	next := func(game *model.MancalaGame) string {
		if game.Turn%2 == 0 {
			return "player1"
		}
		return "player2"
	}

	log.Default().Println("player2: ", game.Player2.Bucket, ReverseBoard(game.Player2.Holes))
	log.Default().Println("player1:   ", game.Player1.Holes, game.Player1.Bucket)
	log.Default().Println("now moving: ", next(&game))
}
