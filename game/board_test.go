package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAdjacent(t *testing.T) {
	b := Board{
		Tiles: [][]tile{
			{tile{}, tile{}, tile{}},
			{tile{}, tile{}, tile{}},
			{tile{}, tile{}, tile{}},
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
	u := &Unit{X: 0, Y: 1}
	til := tile{Unit: u}
	tf := TileFactory{}
	b := Board{
		Tiles: [][]tile{
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
	til := tile{}
	uni := &Unit{Owner: p1, X: 0, Y: 1}
	til.AddUnit(uni)
	tf := TileFactory{}
	b := Board{
		Tiles: [][]tile{
			{tf.Field(), til, tf.Field()},
		},
	}
	u, _ := b.getUnit(0, 1)
	assert.Equal(t, uni, u)
}

func TestGetUnits(t *testing.T) {
	p1 := &Player{}
	til := tile{}
	uni := &Unit{Owner: p1}
	til.AddUnit(uni)
	tf := TileFactory{}
	b := Board{
		Tiles: [][]tile{
			{tf.Field(), til, tf.Field()},
		},
	}
	u := b.getUnits(p1)
	assert.Equal(t, 1, len(u))
	assert.Equal(t, uni, u[0])
}

func TestGetUnitsMultiple(t *testing.T) {
	p1 := &Player{Name: "a"}
	p2 := &Player{Name: "a"}
	p3 := &Player{Name: "c"}
	til1 := tile{Unit: &Unit{Owner: p1}}
	til2 := tile{Unit: &Unit{Owner: p1}}
	til3 := tile{Unit: &Unit{Owner: p2}}
	tf := TileFactory{}
	b := Board{
		Tiles: [][]tile{
			{tf.Field(), til1, tf.Field()},
			{til2, til3},
		},
	}
	assert.Equal(t, 2, len(b.getUnits(p1)))
	assert.Equal(t, 1, len(b.getUnits(p2)))
	assert.Equal(t, 0, len(b.getUnits(p3)))
}
