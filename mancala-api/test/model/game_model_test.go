package model_test

import (
	"testing"

	"github.com/Abeldlp/bol-assignment/mancala-api/model"
)

func TestNewMancalaGame(t *testing.T) {
	player1 := model.NewPlayer()
	player2 := model.NewPlayer()

	game := model.NewMancalaGame(player1, player2)

	if game.Turn != 0 {
		t.Errorf("Expected turn to be 0 but got %d", game.Turn)
	}

	if game.GameOver {
		t.Error("Expected game_over to be false but got true")
	}

	if game.Player1ID != player1.ID {
		t.Errorf("Expected Player1ID to be %d but got %d", player1.ID, game.Player1ID)
	}

	if game.Player2ID != player2.ID {
		t.Errorf("Expected Player2ID to be %d but got %d", player2.ID, game.Player2ID)
	}
}

func TestGetCurrentPlayer(t *testing.T) {
	player1 := model.NewPlayer()
	player2 := model.NewPlayer()

	game := model.NewMancalaGame(player1, player2)

	if game.GetCurentPlayer().ID != player1.ID {
		t.Error("Expected current player to be player1")
	}

	game.Turn = 1

	if game.GetCurentPlayer().ID != player2.ID {
		t.Error("Expected current player to be player2")
	}
}

func TestGetOpponentPlayer(t *testing.T) {
	player1 := model.NewPlayer()
	player2 := model.NewPlayer()

	game := model.NewMancalaGame(player1, player2)

	if game.GetOpponent().ID != player2.ID {
		t.Error("Expected opponent player to be player2")
	}

	game.Turn = 1

	if game.GetOpponent().ID != player1.ID {
		t.Error("Expected opponent player to be player1")
	}
}

func TestPlayRound(t *testing.T) {
	player1 := model.NewPlayer()
	player2 := model.NewPlayer()

	game := model.NewMancalaGame(player1, player2)

	game.PlayRound(2)

	if game.Turn != 1 {
		t.Errorf("Expected turn to be 1 but got %d", game.Turn)
	}
}

func TestPlayRound_GameOver(t *testing.T) {
	player1 := model.NewPlayer()
	player2 := model.NewPlayer()

	game := model.NewMancalaGame(player1, player2)

	game.Player1.Holes = []int{0, 0, 0, 0, 0, 1}
	game.Player2.Holes = []int{0, 0, 0, 0, 0, 0}

	game.PlayRound(5)

	if game.GameOver != true {
		t.Errorf("expected game to be over, but got game not over")
	}
}

func TestNextTurn(t *testing.T) {
	player1 := model.NewPlayer()
	player2 := model.NewPlayer()

	game := model.NewMancalaGame(player1, player2)

	game.NextTurn(false)

	if game.Turn != 1 {
		t.Errorf("Expected turn to be 1 but got %d", game.Turn)
	}

	game.NextTurn(true)

	if game.Turn != 1 {
		t.Errorf("Expected turn to be 1 but got %d", game.Turn)
	}
}

func TestIsGameOver(t *testing.T) {
	player1 := model.NewPlayer()
	player2 := model.NewPlayer()

	game := model.NewMancalaGame(player1, player2)

	if game.IsGameOver() {
		t.Error("Expected game to not be over but got true")
	}

	player1.EmptyAllHoles()
	player2.EmptyAllHoles()

	if !game.IsGameOver() {
		t.Error("Expected game to be over but got false")
	}
}

func TestSetMancalaGameUser(t *testing.T) {
	player1 := model.NewPlayer()
	player2 := model.NewPlayer()

	player1.ID = 1

	game := model.NewMancalaGame(player1, player2)

	player1.Bucket = 1

	if game.Player1.Bucket != 0 {
		t.Errorf("Expected player1 bucket to be 0 but got %d", game.Player1.Bucket)
	}

	game.SetMancalaGameUser(1, *player1)

	if game.Player1.Bucket != 1 {
		t.Errorf("Expected player1 bucket to be 1 but got %d", game.Player1.Bucket)
	}
}
