package model

import (
	"github.com/Abeldlp/bol-assignment/mancala-api/config"
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

	config.DB.Save(&player)
	config.DB.Save(&opponent)

	if player.GetHolesSum() == 0 {
		opponent.Bucket += opponent.GetHolesSum()
	}

	if opponent.GetHolesSum() == 0 {
		player.Bucket += player.GetHolesSum()
	}

	g.NextTurn(lasBucket)

	g.GameOver = g.IsGameOver()
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

func (g *MancalaGame) GetWinner() Player {
	if g.Player1.Bucket > g.Player2.Bucket {
		return g.Player1
	}
	return g.Player2
}

func (g *MancalaGame) SetMancalaGameUser(id uint, player Player) {
	if g.Player1.ID == id {
		g.Player1 = player
		return
	}
	g.Player2 = player
}
