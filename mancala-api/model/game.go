package model

import (
	"gorm.io/gorm"
)

type MancalaGame struct {
	gorm.Model
	Player1   Player `gorm:"foreignKey:ID"`
	Player1ID uint
	Player2   Player `gorm:"foreignKey:ID"`
	Player2ID uint
	Turn      int  `json:"turn" gorm:"default:1"`
	GameOver  bool `json:"game_over" gorm:"default:false"`
}

func NewMancalaGame(player1 *Player, player2 *Player) *MancalaGame {
	return &MancalaGame{
		Turn:      0,
		GameOver:  false,
		Player1:   *player1,
		Player1ID: player1.ID,
		Player2:   *player2,
		Player2ID: player2.ID,
	}
}

func (g *MancalaGame) GetCurentPlayer() Player {
	if g.Turn%2 == 0 {
		return g.Player1
	}
	return g.Player2
}

func (g *MancalaGame) GetOpponent() Player {
	if g.Turn%2 == 0 {
		return g.Player2
	}
	return g.Player1
}

func (g *MancalaGame) PlayRound(holeIndex int) {

	player := g.GetCurentPlayer()
	opponent := g.GetOpponent()

	player.PlayRoundAgainstOpponent(holeIndex, &opponent)

	if player.GetHolesSum() == 0 {
		opponent.Bucket += opponent.GetHolesSum()
	}

	if opponent.GetHolesSum() == 0 {
		player.Bucket += player.GetHolesSum()
	}

	g.GameOver = g.IsGameOver()
}

func (g *MancalaGame) IsGameOver() bool {
	if g.Player1.GetHolesSum() == 0 || g.Player2.GetHolesSum() == 0 {
		return true
	}

	return false
}

func (g *MancalaGame) GetWinner() Player {
	if g.Player1.Bucket > g.Player2.Bucket {
		return g.Player1
	}
	return g.Player2
}
