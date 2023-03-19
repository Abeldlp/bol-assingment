package model

import (
	"github.com/Abeldlp/bol-assignment/mancala-api/custom_type"
	"gorm.io/gorm"
)

type Player struct {
	gorm.Model
	Holes         custom_type.IntArray `json:"holes" gorm:"type:integer[]"`
	Bucket        int                  `json:"bucket"`
	MancalaGameID uint                 `json:"mancala_game_id"`
}

func NewPlayer() *Player {
	return &Player{
		Holes:  custom_type.IntArray{4, 4, 4, 4, 4, 4},
		Bucket: 0,
	}
}

func (p *Player) IncrementHole(index int) {
	p.Holes[index] += 1
}

func (p *Player) EmptyHole(index int) {
	p.Holes[index] = 0
}

func (p *Player) IncrementBucket() {
	p.Bucket += 1
}

func (p *Player) GetOppositeHole(index int, opponent *Player) int {
	oppositeIndex := len(p.Holes) - 1 - index
	amount := opponent.Holes[oppositeIndex]
	opponent.EmptyHole(oppositeIndex)
	return amount
}

func (p *Player) IncrementHolesWithRemainingStones(onHand int, opponent *Player) (int, bool) {
	setStones := 0
	lastIsBucket := false

	for i := 0; i < onHand; i++ {
		if i == 6 {
			p.IncrementBucket()
			setStones++
			if onHand == setStones {
				lastIsBucket = true
			}
			return onHand - setStones, lastIsBucket
		}
		p.IncrementHole(i)
		setStones++
	}

	lastHoleSet := setStones - 1
	lastHoleSetValue := p.Holes[lastHoleSet]
	oppositeIndex := len(p.Holes) - 1 - lastHoleSet
	oppositeHole := opponent.Holes[oppositeIndex]

	if lastHoleSetValue == 1 && oppositeHole > 0 {
		p.EmptyHole(lastHoleSet)
		p.Bucket += oppositeHole + 1
		opponent.EmptyHole(oppositeIndex)
	}

	return 0, lastIsBucket
}

func (p *Player) MoveStonesUntilBucket(index int, opponent *Player) (int, bool) {
	onHand := p.Holes[index]
	p.EmptyHole(index)
	setStones := 0
	lastIsBucket := false

	for i := 1; i < onHand+1; i++ {
		if index+i == len(p.Holes) {
			setStones++
			p.IncrementBucket()

			if onHand == setStones {
				lastIsBucket = true
			}
			return onHand - setStones, lastIsBucket
		}

		p.IncrementHole(index + i)
		setStones++
	}

	lastHoleSet := index + setStones
	lastHoleSetValue := p.Holes[lastHoleSet]
	oppositeIndex := len(p.Holes) - 1 - lastHoleSet
	oppositeHole := opponent.Holes[oppositeIndex]

	if lastHoleSetValue == 1 && oppositeHole > 0 {
		p.EmptyHole(lastHoleSet)
		p.Bucket += oppositeHole + 1
		opponent.EmptyHole(oppositeIndex)
	}

	return 0, lastIsBucket
}

func (p *Player) MoveStonesUntilOpponentBucket(onHand int, opponent *Player) int {
	setStones := 0

	for i := 0; i < onHand; i++ {
		if i == len(opponent.Holes) {
			return onHand - setStones
		}
		opponent.IncrementHole(i)
		setStones++
	}
	return 0
}

func (p *Player) PlayRoundAgainstOpponent(selectedHole int, opponent *Player) bool {
	bucketHit := false
	left, lastOnBucket := p.MoveStonesUntilBucket(selectedHole, opponent)
	bucketHit = lastOnBucket

	roundContinues := true

	for roundContinues {

		if left > 0 {
			left = p.MoveStonesUntilOpponentBucket(left, opponent)
		}
		if left > 0 {
			left, lastOnBucket = p.IncrementHolesWithRemainingStones(left, opponent)
			bucketHit = lastOnBucket
		}

		roundContinues = left > 0
	}

	return bucketHit
}

func (p *Player) GetHolesSum() int {
	sum := 0
	for _, hole := range p.Holes {
		sum += hole
	}
	return sum
}
