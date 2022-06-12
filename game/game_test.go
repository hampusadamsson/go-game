package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGameTestMoveInSequence(t *testing.T) {
	p1 := &Player{"a"}
	p2 := &Player{"b"}

	u1 := &Unit{x: 0, y: 0, Owner: p1, Movement: 10}
	u2 := &Unit{x: 0, y: 3, Owner: p2, Movement: 10}
	tile1 := tile{unit: u1}
	tile2 := tile{unit: u2}

	tf := TileFactory{}
	b := Board{
		Tiles: [][]tile{
			{tile1, tf.Field(), tf.Field(), tile2},
		},
	}

	g := Game{
		Players: []*Player{p1, p2},
		Board:   b,
		turn:    p1,
	}

	_, err := g.Move(p2, coord{0, 3}, coord{0, 2})
	assert.NotNil(t, err) // not your turn

	ok, err := g.Move(p1, coord{0, 0}, coord{0, 1})
	assert.True(t, ok)

	fail, _ := g.Move(p1, coord{0, 0}, coord{0, 1})
	assert.False(t, fail) // no unit there any longer

	g.ChangeTurn(p1)

	cantMove, _ := g.Move(p2, coord{0, 3}, coord{0, 0})
	assert.False(t, cantMove) // Blocked by enemy - no path

	canMove, _ := g.Move(p2, coord{0, 3}, coord{0, 2})
	assert.True(t, canMove)

	unitIsExahausted, _ := g.Move(p2, coord{0, 2}, coord{0, 3})
	assert.False(t, unitIsExahausted)
}

func TestAttack(t *testing.T) {
	p1 := &Player{"a"}
	p2 := &Player{"b"}

	u1 := &Unit{x: 0, y: 0, Owner: p1, Movement: 10, attackRange: 1, CanMoveAttack: true}
	u2 := &Unit{x: 0, y: 3, Owner: p2, Movement: 10, attackRange: 1, CanMoveAttack: true}
	tile1 := tile{unit: u1}
	tile2 := tile{unit: u2}

	tf := TileFactory{}
	b := Board{
		Tiles: [][]tile{
			{tile1, tf.Field(), tf.Field(), tile2},
		},
	}

	g := Game{
		Players: []*Player{p1, p2},
		Board:   b,
		turn:    p1,
	}

	successp2, _ := g.Attack(p2, coord{0, 3}, coord{0, 0})
	assert.False(t, successp2) // not your turn

	successp1, _ := g.Attack(p1, coord{0, 3}, coord{0, 0})
	assert.False(t, successp1) // not in range
	g.Move(p1, coord{0, 0}, coord{0, 2})
	successp12, _ := g.Attack(p1, coord{0, 2}, coord{0, 3})
	assert.True(t, successp12)
}

func TestGameChangeTurn(t *testing.T) {
	p1 := &Player{"a"}
	p2 := &Player{"b"}
	g := Game{
		Players: []*Player{p1, p2},
		turn:    p1,
	}
	assert.Equal(t, g.turn, p1)
	g.ChangeTurn(p1)
	assert.Equal(t, g.turn, p2)
	g.ChangeTurn(p1)
	assert.Equal(t, g.turn, p2)
	g.ChangeTurn(p2)
	assert.Equal(t, g.turn, p1)
}
