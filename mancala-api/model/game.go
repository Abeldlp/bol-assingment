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
	Turn      int  `json:"turn" gorm:"default:0"`
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

	lasBucket := player.PlayRoundAgainstOpponent(holeIndex, &opponent)

	g.SetMancalaGameUser(player.ID, player)
	g.SetMancalaGameUser(opponent.ID, opponent)

	g.NextTurn(lasBucket)

	g.GameOver = g.IsGameOver()

	if g.GameOver {
		g.Player1.Bucket += g.Player1.GetHolesSum()
		player.EmptyAllHoles()

		g.Player2.Bucket += g.Player2.GetHolesSum()
		opponent.EmptyAllHoles()
	}
}

func (g *MancalaGame) NextTurn(lastStoneInBucket bool) {
	if lastStoneInBucket {
		return
	}
	g.Turn++
}

func (g *MancalaGame) IsGameOver() bool {
	if g.Player1.GetHolesSum() == 0 || g.Player2.GetHolesSum() == 0 {
		return true
	}
	return false
}

func (g *MancalaGame) SetMancalaGameUser(id uint, player Player) {
	if g.Player1.ID == id {
		g.Player1 = player
		return
	}
	g.Player2 = player
}
