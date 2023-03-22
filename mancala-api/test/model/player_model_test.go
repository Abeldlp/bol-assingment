package model_test

import (
	"testing"

	"github.com/Abeldlp/bol-assignment/mancala-api/model"
)

func TestPlayer_NewPlayer(t *testing.T) {
	p := model.NewPlayer()
	if len(p.Holes) != 6 {
		t.Errorf("Expected player to have 6 holes but got %d", len(p.Holes))
	}
	if p.Bucket != 0 {
		t.Errorf("Expected player to have 0 stones in the bucket but got %d", p.Bucket)
	}

	for _, hole := range p.Holes {
		if hole != 4 {
			t.Errorf("Expected player to have 4 stones in each hole but got %v", p.Holes)
			break
		}
	}
}

func TestPlayer_EmptyAllHoles(t *testing.T) {
	p := model.NewPlayer()
	p.EmptyAllHoles()
	for _, hole := range p.Holes {
		if hole != 0 {
			t.Errorf("Expected player to have empty holes but got %v", p.Holes)
			break
		}
	}
}

func TestPlayer_IncrementHole(t *testing.T) {
	p := model.NewPlayer()
	p.IncrementHole(0)
	if p.Holes[0] != 5 {
		t.Errorf("Expected player to have 5 stones in the first hole but got %d", p.Holes[0])
	}
}

func TestPlayer_EmptyHole(t *testing.T) {
	p := model.NewPlayer()
	p.EmptyHole(0)
	if p.Holes[0] != 0 {
		t.Errorf("Expected player to have an empty first hole but got %d", p.Holes[0])
	}
}

func TestPlayer_IncrementBucket(t *testing.T) {
	p := model.NewPlayer()
	p.IncrementBucket()
	if p.Bucket != 1 {
		t.Errorf("Expected player to have 1 stone in the bucket but got %d", p.Bucket)
	}
}

func TestPlayer_IncrementHolesWithRemainingStones_UntilBucket(t *testing.T) {
	p1 := model.NewPlayer()
	p2 := model.NewPlayer()

	remainingStones, lastIsBucket := p1.IncrementHolesWithRemainingStones(7, p2)

	if remainingStones != 0 {
		t.Errorf("Expected to have 1 remaining stone but got %d", remainingStones)
	}

	if !lastIsBucket {
		t.Errorf("Expected the last stone to be in the bucket but got %v", lastIsBucket)
	}
}

func TestPlayer_IncrementHolesWithRemainingStones_BeforeBucket(t *testing.T) {
	p1 := model.NewPlayer()
	p2 := model.NewPlayer()

	remainingStones, lastIsBucket := p1.IncrementHolesWithRemainingStones(3, p2)

	if remainingStones != 0 {
		t.Errorf("Expected to no stones but got %d", remainingStones)
	}

	if lastIsBucket {
		t.Errorf("Expected the last stone to not be in the bucket but got %v", lastIsBucket)
	}
}

func TestPlayer_IncrementHolesWithRemainingStones_PastBucket(t *testing.T) {
	p1 := model.NewPlayer()
	p2 := model.NewPlayer()

	remainingStones, lastIsBucket := p1.IncrementHolesWithRemainingStones(8, p2)

	if remainingStones != 1 {
		t.Errorf("Expected to have 1 remaining stone but got %d", remainingStones)
	}

	if lastIsBucket {
		t.Errorf("Expected the last stone to not be in the bucket but got %v", lastIsBucket)
	}
}

func TestPlayer_IncrementHolesWithRemainingStones_GrabOpponentStones(t *testing.T) {
	p1 := model.NewPlayer()
	p2 := model.NewPlayer()
	p1.Holes[5] = 0

	if p1.Bucket != 0 {
		t.Errorf("Expected player 1 to have 0 stones in the bucket but got %d", p1.Bucket)
	}

	p1.IncrementHolesWithRemainingStones(6, p2)

	if p1.Bucket != 5 {
		t.Errorf("Expected player 1 to have 5 stones in the bucket but got %d", p1.Bucket)
	}
}

func TestPlayer_IncrementHolesWithRemainingStones_DontGrabOpponentStones(t *testing.T) {
	p1 := model.NewPlayer()
	p2 := model.NewPlayer()
	p1.Holes[5] = 0
	p2.Holes[0] = 0

	if p1.Bucket != 0 {
		t.Errorf("Expected player 1 to have 0 stones in the bucket but got %d", p1.Bucket)
	}

	p1.IncrementHolesWithRemainingStones(6, p2)

	if p1.Bucket != 0 {
		t.Errorf("Expected player 1 to have 0 stones in the bucket but got %d", p1.Bucket)
	}
}

func TestPlayer_MoveStonesUntilBucket(t *testing.T) {
	p1 := model.NewPlayer()
	p2 := model.NewPlayer()

	remainingStones, lastIsBucket := p1.MoveStonesUntilBucket(2, p2)

	if remainingStones != 0 {
		t.Errorf("Expected to have no remaining stones but got %d", remainingStones)
	}

	if !lastIsBucket {
		t.Errorf("Expected the last stone to be in the bucket but got %v", lastIsBucket)
	}
}

func TestPlayer_MoveStonesUntilBucket_BeforeBucket(t *testing.T) {
	p1 := model.NewPlayer()
	p2 := model.NewPlayer()

	remainingStones, lastIsBucket := p1.MoveStonesUntilBucket(1, p2)

	if remainingStones != 0 {
		t.Errorf("Expected to have no remaining stones but got %d", remainingStones)
	}

	if lastIsBucket {
		t.Errorf("Expected the last stone to not be in the bucket but got %v", lastIsBucket)
	}
}

func TestPlayer_MoveStonesUntilBucket_PastBucket(t *testing.T) {
	p1 := model.NewPlayer()
	p2 := model.NewPlayer()

	remainingStones, lastIsBucket := p1.MoveStonesUntilBucket(3, p2)

	if remainingStones != 1 {
		t.Errorf("Expected to have 1 remaining stone but got %d", remainingStones)
	}

	if lastIsBucket {
		t.Errorf("Expected the last stone to not be in the bucket but got %v", lastIsBucket)
	}
}

func TestPlayer_MoveStonesUntilBucket_GrabOpponentStones(t *testing.T) {
	p1 := model.NewPlayer()
	p2 := model.NewPlayer()
	p1.Holes[5] = 0

	if p1.Bucket != 0 {
		t.Errorf("Expected player 1 to have 0 stones in the bucket but got %d", p1.Bucket)
	}

	p1.MoveStonesUntilBucket(1, p2)

	if p1.Bucket != 5 {
		t.Errorf("Expected player 1 to have 5 stones in the bucket but got %d", p1.Bucket)
	}
}

func TestPlayer_MoveStonesUntilBucket_DontGrabOpponentStones(t *testing.T) {
	p1 := model.NewPlayer()
	p2 := model.NewPlayer()
	p1.Holes[5] = 0
	p2.Holes[0] = 0

	if p1.Bucket != 0 {
		t.Errorf("Expected player 1 to have 0 stones in the bucket but got %d", p1.Bucket)
	}

	p1.MoveStonesUntilBucket(1, p2)

	if p1.Bucket != 0 {
		t.Errorf("Expected player 1 to have 0 stones in the bucket but got %d", p1.Bucket)
	}
}

func TestPlayer_MoveStonesUntilOpponentBucket(t *testing.T) {
	p1 := model.NewPlayer()
	p2 := model.NewPlayer()

	p1.MoveStonesUntilOpponentBucket(2, p2)

	if p2.Holes[1] != 5 {
		t.Errorf("Expected player 2 to have 5 stone in hole 1 but got %d", p2.Holes[0])
	}
}

func TestPlayer_MoveStonesUntilOpponentBucket_PassingBucket(t *testing.T) {
	p1 := model.NewPlayer()
	p2 := model.NewPlayer()

	p1.MoveStonesUntilOpponentBucket(7, p2)

	if p2.Bucket != 0 {
		t.Errorf("Expected player 2 to have 0 stone in the bucket but got %d", p2.Bucket)
	}
}

func TestPlayer_GetSHolesSum(t *testing.T) {
	p1 := model.NewPlayer()

	sum := p1.GetHolesSum()

	if sum != 24 {
		t.Errorf("Expected the sum to be 36 but got %d", sum)
	}
}

func TestPlayer_PlayRound(t *testing.T) {
	p1 := model.NewPlayer()
	p2 := model.NewPlayer()
	p1.Holes[5] = 8

	p1.PlayRoundAgainstOpponent(5, p2)

	if p1.Holes[5] != 0 {
		t.Errorf("Expected player 1 to have 0 stones in hole 5 but got %d", p1.Holes[5])
	}

	if p1.Bucket != 1 {
		t.Errorf("Expected player 1 to have 1 stone in the bucket but got %d", p1.Bucket)
	}

	if p2.Holes[5] != 5 {
		t.Errorf("Expected player 2 to have 5 stones in hole 5 but got %d", p2.Holes[5])
	}

	if p2.Bucket != 0 {
		t.Errorf("Expected player 2 to have 0 stone in the bucket but got %d", p2.Bucket)
	}

	if p1.Holes[0] != 5 {
		t.Errorf("Expected player 1 to have 5 stones in hole 0 but got %d", p1.Holes[0])
	}

}

func TestPlayer_PlayRound_UnsupportedHole(t *testing.T) {
	p1 := model.NewPlayer()
	p2 := model.NewPlayer()
	p1.Holes[5] = 8

	repeatTurnOutOfRange := p1.PlayRoundAgainstOpponent(12, p2)

	if !repeatTurnOutOfRange {
		t.Errorf("Expected to repeat the turn but got %v", repeatTurnOutOfRange)
	}

	repeatTurnNegativeInt := p1.PlayRoundAgainstOpponent(-2, p2)
	if !repeatTurnNegativeInt {
		t.Errorf("Expected to repeat the turn but got %v", repeatTurnNegativeInt)
	}
}
