package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	pf    = pathfinding{}
	one   = tile{cost: 1}
	nine  = tile{cost: 9}
	enemy = tile{unit: &Unit{Owner: &Player{}}}
)

func TestRow(t *testing.T) {
	b := Board{
		Tiles: [][]tile{
			{one, nine, one},
		},
	}
	way, dist, _ := pf.findShortestPath(&b, &Unit{Movement: 10}, coord{0, 0}, coord{0, 2})
	assert.Equal(t, 10, dist)
	assert.Equal(t, 2, len(way))
}

func TestColumn(t *testing.T) {
	b := Board{
		Tiles: [][]tile{
			{one},
			{nine},
			{one},
		},
	}
	_, dist, _ := pf.findShortestPath(&b, &Unit{Movement: 10}, coord{0, 0}, coord{2, 0})
	assert.Equal(t, 10, dist)
}

func TestEasy(t *testing.T) {
	b := Board{
		Tiles: [][]tile{
			{one, nine, one},
			{one, nine, one},
			{one, one, one},
		},
	}
	_, dist, _ := pf.findShortestPath(&b, &Unit{Movement: 4}, coord{0, 0}, coord{2, 2})
	assert.Equal(t, 4, dist)
	_, dist2, _ := pf.findShortestPath(&b, &Unit{Movement: 8}, coord{0, 0}, coord{0, 2})
	assert.Equal(t, 6, dist2)
}

func TestMedium(t *testing.T) {
	b := Board{
		Tiles: [][]tile{
			{one, nine, nine},
			{one, nine, one},
			{nine, nine, one},
		},
	}
	_, dist, _ := pf.findShortestPath(&b, &Unit{Movement: 9999}, coord{0, 0}, coord{2, 2})
	assert.Equal(t, 12, dist)
}

func TestHard(t *testing.T) {
	b := Board{
		Tiles: [][]tile{
			{one, nine, one, one, one},
			{one, nine, one, nine, one},
			{one, nine, one, nine, one},
			{one, one, one, nine, one},
		},
	}
	_, dist, _ := pf.findShortestPath(&b, &Unit{Movement: 9999}, coord{0, 0}, coord{3, 4})
	assert.Equal(t, 13, dist)
}

func TestUnitMaxRangeOverstep(t *testing.T) {
	b := Board{
		Tiles: [][]tile{
			{one, nine, nine},
			{one, nine, one},
			{nine, nine, one},
		},
	}
	_, _, canMoveHere := pf.findShortestPath(&b, &Unit{Movement: 11}, coord{0, 0}, coord{2, 2})
	assert.False(t, canMoveHere)
}

func TestImpassableUnits(t *testing.T) {
	b := Board{
		Tiles: [][]tile{{one, enemy, one}, {one, enemy, one}, {one, enemy, one}},
		pf:    pf,
	}
	_, _, canMoveHere := pf.findShortestPath(&b, &Unit{Movement: 25}, coord{0, 0}, coord{2, 2})
	assert.False(t, canMoveHere)
}

func TestPassableUnits(t *testing.T) {
	p1 := Player{}
	b := Board{
		Tiles: [][]tile{
			{one, enemy, one},
			{one, tile{unit: &Unit{Owner: &p1}}, one},
			{one, enemy, one},
		},
	}
	_, _, canMoveHere := pf.findShortestPath(&b, &Unit{Owner: &p1, Movement: 25}, coord{0, 0}, coord{2, 2})
	assert.True(t, canMoveHere)
}
