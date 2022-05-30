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
	g := Game{
		tiles: [][]Tile{
			{Plain, til, Plain},
		},
		players: []*Player{p1},
	}
	u := g.GetUnits(p1)
	assert.Equal(t, 1, len(u))
	assert.Equal(t, uni, u[0])
}

func TestGetUnitsMultiple(t *testing.T) {
	p1 := &Player{}
	p2 := &Player{}
	p3 := &Player{}
	til1 := Tile{unit: &Unit{Owner: p1}}
	til2 := Tile{unit: &Unit{Owner: p1}}
	til3 := Tile{unit: &Unit{Owner: p2}}
	g := Game{
		tiles: [][]Tile{
			{Plain, til1, Plain},
			{til2, til3},
		},
		players: []*Player{p1, p2, p3},
	}
	assert.Equal(t, len(g.GetUnits(p1)), 2)
	assert.Equal(t, len(g.GetUnits(p2)), 1)
	assert.Equal(t, len(g.GetUnits(p3)), 0)
}
