package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUnits(t *testing.T) {
	p1 := &Player{}
	til := Tile{}
	uni := &Unit{Owner: p1}
	til.AddUnit(uni)
	tf := TileFactory{}
	g := Game{
		tiles: [][]Tile{
			{tf.Plain(), til, tf.Plain()},
		},
		players: []*Player{p1},
	}
	u := g.GetUnits(p1)
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
	g := Game{
		tiles: [][]Tile{
			{tf.Plain(), til1, tf.Plain()},
			{til2, til3},
		},
		players: []*Player{p1, p2, p3},
	}
	assert.Equal(t, 2, len(g.GetUnits(p1)))
	assert.Equal(t, 1, len(g.GetUnits(p2)))
	assert.Equal(t, 0, len(g.GetUnits(p3)))
}
