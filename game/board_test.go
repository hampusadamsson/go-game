package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAdjacent(t *testing.T) {
	b := Board{
		tiles: [][]Tile{
			{Tile{}, Tile{}, Tile{}},
			{Tile{}, Tile{}, Tile{}},
			{Tile{}, Tile{}, Tile{}},
		},
	}
	assert.Equal(t, 4, len(b.getAdjacent(1, 1)))
	assert.Equal(t, 2, len(b.getAdjacent(0, 0)))
	assert.Equal(t, 3, len(b.getAdjacent(0, 1)))
	assert.Equal(t, 0, len(b.getAdjacent(99, 99)))
	assert.Equal(t, 2, len(b.getAdjacent(2, 2)))
	assert.Equal(t, 0, len(b.getAdjacent(3, 3)))
}

func TestMove(t *testing.T) {
	u := &Unit{x: 0, y: 1}
	til := Tile{unit: u}
	tf := TileFactory{}
	b := Board{
		tiles: [][]Tile{
			{tf.Field(), til, tf.Field()},
		},
	}
	moved0, _ := b.move(u, 0, 1)
	assert.False(t, moved0)

	moved1, _ := b.move(u, 0, 0)
	assert.True(t, moved1)

	moved2, _ := b.move(u, 0, 0)
	assert.False(t, moved2)

	moved3, _ := b.move(u, 0, 1)
	assert.True(t, moved3)
}

func TestGetUnit(t *testing.T) {
	p1 := &Player{}
	til := Tile{}
	uni := &Unit{Owner: p1, x: 0, y: 1}
	til.AddUnit(uni)
	tf := TileFactory{}
	b := Board{
		tiles: [][]Tile{
			{tf.Field(), til, tf.Field()},
		},
	}
	u, _ := b.GetUnit(0, 1)
	assert.Equal(t, uni, u)
}

func TestGetUnits(t *testing.T) {
	p1 := &Player{}
	til := Tile{}
	uni := &Unit{Owner: p1}
	til.AddUnit(uni)
	tf := TileFactory{}
	b := Board{
		tiles: [][]Tile{
			{tf.Field(), til, tf.Field()},
		},
	}
	u := b.GetUnits(p1)
	assert.Equal(t, 1, len(u))
	assert.Equal(t, uni, u[0])
}

func TestGetUnitsMultiple(t *testing.T) {
	p1 := &Player{name: "a"}
	p2 := &Player{name: "a"}
	p3 := &Player{name: "c"}
	til1 := Tile{unit: &Unit{Owner: p1}}
	til2 := Tile{unit: &Unit{Owner: p1}}
	til3 := Tile{unit: &Unit{Owner: p2}}
	tf := TileFactory{}
	b := Board{
		tiles: [][]Tile{
			{tf.Field(), til1, tf.Field()},
			{til2, til3},
		},
	}
	assert.Equal(t, 2, len(b.GetUnits(p1)))
	assert.Equal(t, 1, len(b.GetUnits(p2)))
	assert.Equal(t, 0, len(b.GetUnits(p3)))
}
