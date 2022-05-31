package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUnit(t *testing.T) {
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
