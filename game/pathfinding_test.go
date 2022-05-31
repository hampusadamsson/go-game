package game

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	pf   = newPathfinding()
	one  = Tile{cost: 1}
	nine = Tile{cost: 9}
)

func TestEasy(t *testing.T) {
	b := Board{
		tiles: [][]Tile{
			{one, nine, one},
			{one, nine, one},
			{one, one, one},
		},
	}
	res, _ := pf.getPath(&b, 0, 0, 2, 2)
	assert.Equal(t, 4, res.cost)
	res2, _ := pf.getPath(&b, 0, 0, 0, 2)
	assert.Equal(t, 6, res2.cost)
}

func TestMedium(t *testing.T) {
	b := Board{
		tiles: [][]Tile{
			{one, nine, nine},
			{one, nine, one},
			{nine, nine, one},
		},
	}
	res, _ := pf.getPath(&b, 0, 0, 2, 2)
	assert.Equal(t, 12, res.cost)
}

func TestHard(t *testing.T) {
	b := Board{
		tiles: [][]Tile{
			{one, nine, one, one, one},
			{one, nine, one, nine, one},
			{one, nine, one, nine, one},
			{one, one, one, nine, one},
		},
	}
	res, err := pf.getPath(&b, 0, 0, 3, 4)
	fmt.Println(err)
	fmt.Println(res.path)
	assert.Equal(t, 13, res.cost)
}

func TestHard2(t *testing.T) {
	b := Board{
		tiles: [][]Tile{
			{one, nine, one, one, one, one, nine, one, one, one, one, nine, one, one, one, one, nine, one, one, one, one, nine, one, one, one},
			{one, nine, one, one, one, one, nine, one, one, one, one, nine, one, one, one, one, nine, one, one, one, one, nine, one, one, one},
			{one, nine, one, one, one, one, nine, one, one, one, one, nine, one, one, one, one, nine, one, one, one, one, nine, one, one, one},
			{one, nine, one, one, one, one, nine, one, one, one, one, nine, one, one, one, one, nine, one, one, one, one, nine, one, one, one},
			{one, nine, one, one, one, one, nine, one, one, one, one, nine, one, one, one, one, nine, one, one, one, one, nine, one, one, one},
			{one, nine, one, one, one, one, nine, one, one, one, one, nine, one, one, one, one, nine, one, one, one, one, nine, one, one, one},
			{one, nine, one, one, one, one, nine, one, one, one, one, nine, one, one, one, one, nine, one, one, one, one, nine, one, one, one},
			{one, nine, one, one, one, one, nine, one, one, one, one, nine, one, one, one, one, nine, one, one, one, one, nine, one, one, one},
			{one, nine, one, one, one, one, nine, one, one, one, one, nine, one, one, one, one, nine, one, one, one, one, nine, one, one, one},
			{one, nine, one, one, one, one, nine, one, one, one, one, nine, one, one, one, one, nine, one, one, one, one, nine, one, one, one},
			{one, nine, one, one, one, one, nine, one, one, one, one, nine, one, one, one, one, nine, one, one, one, one, nine, one, one, one},
			{one, nine, one, one, one, one, nine, one, one, one, one, nine, one, one, one, one, nine, one, one, one, one, nine, one, one, one},
			{one, nine, one, one, one, one, nine, one, one, one, one, nine, one, one, one, one, nine, one, one, one, one, nine, one, one, one},
			{one, nine, one, one, one, one, nine, one, one, one, one, nine, one, one, one, one, nine, one, one, one, one, nine, one, one, one},
			{one, nine, one, one, one, one, nine, one, one, one, one, nine, one, one, one, one, nine, one, one, one, one, nine, one, one, one},
			{one, nine, one, one, one, one, nine, one, one, one, one, nine, one, one, one, one, nine, one, one, one, one, nine, one, one, one},
			{one, nine, one, one, one, one, nine, one, one, one, one, nine, one, one, one, one, nine, one, one, one, one, nine, one, one, one},
			{one, nine, one, one, one, one, nine, one, one, one, one, nine, one, one, one, one, nine, one, one, one, one, nine, one, one, one},
		},
	}
	res, _ := pf.getPath(&b, 0, 0, len(b.tiles)-1, len(b.tiles[0])-1)
	fmt.Println(res.path)
}
