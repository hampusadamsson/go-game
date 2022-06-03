package game

import (
	"errors"
)

type Board struct {
	Tiles [][]tile
	pf    pathfinding
}

type coord struct {
	x int
	y int
}

func (b *Board) getCoord(x int, y int) (*coord, error) {
	if x < 0 || x > len(b.Tiles)-1 || y < 0 || y > len(b.Tiles[0])-1 {
		return nil, errors.New("out of bound")
	}
	return &coord{x, y}, nil
}

func (b *Board) getTile(x int, y int) (*tile, error) {
	c, err := b.getCoord(x, y)
	if err != nil {
		return nil, err
	} else {
		return &b.Tiles[c.x][c.y], nil

	}
}

func (b *Board) getTiles() chan *tile {
	c := make(chan *tile)
	go func() {
		for i := 0; i < len(b.Tiles); i++ {
			for j := 0; j < len(b.Tiles[i]); j++ {
				c <- &b.Tiles[i][j]
			}
		}
		close(c)
	}()
	return c
}

func (b *Board) move(u *Unit, x int, y int) (bool, error) {
	destTile, _ := b.getTile(x, y)
	if destTile.isOccupied() {
		return false, errors.New("destination tile is occupied")
	}
	fromTile, _ := b.getTile(u.x, u.y)
	fromTile.RemoveUnit()
	destTile.AddUnit(u)
	u.x = x
	u.y = y
	return true, nil
}

func (b *Board) getUnits(p *Player) []*Unit {
	units := make([]*Unit, 0)
	for elem := range b.getTiles() {
		if u, err := elem.GetUnit(); err == nil {
			if u.Owner == p {
				units = append(units, u)
			}
		}
	}
	return units
}

func (b *Board) getUnit(x int, y int) (*Unit, error) {
	t, _ := b.getTile(x, y)
	if u, err := t.GetUnit(); err == nil {
		return u, nil
	} else {
		return nil, err
	}
}

func (b *Board) getAdjacent(x int, y int) []*coord {
	ad := make([]*coord, 0)
	if t, err := b.getCoord(x-1, y); err == nil {
		ad = append(ad, t)
	}
	if t, err := b.getCoord(x+1, y); err == nil {
		ad = append(ad, t)
	}
	if t, err := b.getCoord(x, y-1); err == nil {
		ad = append(ad, t)
	}
	if t, err := b.getCoord(x, y+1); err == nil {
		ad = append(ad, t)
	}
	return ad
}

func (b *Board) getPath(u *Unit, x int, y int) ([]coord, int, bool) {
	return b.pf.findShortestPath(b, u, coord{u.x, u.y}, coord{x, y})
}
