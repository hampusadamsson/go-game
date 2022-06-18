package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGameTestMoveInSequence(t *testing.T) {
	p1 := &Player{"a", nil}
	p2 := &Player{"b", nil}

	u1 := &Unit{X: 0, Y: 0, Owner: p1, Movement: 10}
	u2 := &Unit{X: 0, Y: 3, Owner: p2, Movement: 10}
	tile1 := tile{Unit: u1}
	tile2 := tile{Unit: u2}

	tf := TileFactory{}
	b := Board{
		Tiles: [][]tile{
			{tile1, tf.Field(), tf.Field(), tile2},
		},
	}

	g := Game{
		Players: []*Player{p1, p2},
		Board:   b,
		Turn:    p1,
	}

	_, err := g.move(p2, Coord{0, 3}, Coord{0, 2})
	assert.NotNil(t, err) // not your turn

	ok, err := g.move(p1, Coord{0, 0}, Coord{0, 1})
	assert.True(t, ok)

	fail, _ := g.move(p1, Coord{0, 0}, Coord{0, 1})
	assert.False(t, fail) // no unit there any longer

	g.changeTurn(p1)

	cantMove, _ := g.move(p2, Coord{0, 3}, Coord{0, 0})
	assert.False(t, cantMove) // Blocked by enemy - no path

	canMove, _ := g.move(p2, Coord{0, 3}, Coord{0, 2})
	assert.True(t, canMove)

	unitIsExahausted, _ := g.move(p2, Coord{0, 2}, Coord{0, 3})
	assert.False(t, unitIsExahausted)
}

func TestAttack(t *testing.T) {
	p1 := &Player{"a", nil}
	p2 := &Player{"b", nil}

	u1 := &Unit{X: 0, Y: 0, Owner: p1, Movement: 10, attackRange: 1, CanMoveAttack: true}
	u2 := &Unit{X: 0, Y: 3, Owner: p2, Movement: 10, attackRange: 1, CanMoveAttack: true}
	tile1 := tile{Unit: u1}
	tile2 := tile{Unit: u2}

	tf := TileFactory{}
	b := Board{
		Tiles: [][]tile{
			{tile1, tf.Field(), tf.Field(), tile2},
		},
	}

	g := Game{
		Players: []*Player{p1, p2},
		Board:   b,
		Turn:    p1,
	}

	successp2, _, _ := g.attack(p2, Coord{0, 3}, Coord{0, 0})
	assert.False(t, successp2) // not your turn

	successp1, _, _ := g.attack(p1, Coord{0, 3}, Coord{0, 0})
	assert.False(t, successp1) // not in range
	g.move(p1, Coord{0, 0}, Coord{0, 2})
	successp12, _, _ := g.attack(p1, Coord{0, 2}, Coord{0, 3})
	assert.True(t, successp12)
}

func TestGameChangeTurn(t *testing.T) {
	p1 := &Player{"a", nil}
	p2 := &Player{"b", nil}
	g := Game{
		Players: []*Player{p1, p2},
		Turn:    p1,
	}
	assert.Equal(t, g.Turn, p1)
	g.changeTurn(p1)
	assert.Equal(t, g.Turn, p2)
	g.changeTurn(p1)
	assert.Equal(t, g.Turn, p2)
	g.changeTurn(p2)
	assert.Equal(t, g.Turn, p1)
}
