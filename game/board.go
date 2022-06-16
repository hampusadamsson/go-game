package game

import (
	"errors"
)

type Board struct {
	Tiles       [][]tile
	pathfinding pathfinding
}

type Coord struct {
	X int
	Y int
}

func (b *Board) getCoord(x int, y int) (*Coord, error) {
	if x < 0 || x > len(b.Tiles)-1 || y < 0 || y > len(b.Tiles[0])-1 {
		return nil, errors.New("out of bound")
	}
	return &Coord{x, y}, nil
}

func (b *Board) getTile(x int, y int) (*tile, error) {
	c, err := b.getCoord(x, y)
	if err != nil {
		return nil, err
	} else {
		return &b.Tiles[c.X][c.Y], nil
	}
}

func (b *Board) addUnits(u ...*Unit) {
	for i := 0; i < len(u); i++ {
		tile, err := b.getTile(u[i].X, u[i].Y)
		if err != nil {
			panic("no such position")
		}
		tile.AddUnit(u[i])
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

func (b *Board) attack(attacker Coord, defender Coord) (bool, error) {
	uAttacker, err := b.GetUnit(attacker.X, attacker.Y)
	if err != nil {
		return false, err
	}
	uDefender, err := b.GetUnit(defender.X, defender.Y)
	if err != nil {
		return false, err
	}
	uAttacker.fight(uDefender) // TODO add tile defence etc
	if uDefender.HP <= 0 {     // TODO - attack back?
		td, _ := b.getTile(defender.X, defender.Y)
		td.RemoveUnit()
	}
	return true, nil
}

func (b *Board) move(u *Unit, x int, y int) (bool, error) {
	destTile, _ := b.getTile(x, y)
	if destTile.isOccupied() {
		return false, errors.New("destination tile is occupied")
	}
	fromTile, _ := b.getTile(u.X, u.Y)
	fromTile.RemoveUnit()
	destTile.AddUnit(u)
	u.X = x
	u.Y = y
	return true, nil
}

func (b *Board) GetUnits(p *Player) []*Unit {
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

func (b *Board) GetUnit(x int, y int) (*Unit, error) {
	t, err := b.getTile(x, y)
	if err != nil {
		return nil, err
	}
	if u, err := t.GetUnit(); err == nil {
		return u, nil
	} else {
		return nil, err
	}
}

func (b *Board) getAdjacent(x int, y int) []*Coord {
	ad := make([]*Coord, 0)
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

func (b *Board) GetShortestPath(u *Unit, x int, y int) ([]Coord, int, bool) {
	return b.pathfinding.findShortestPath(b, u, Coord{u.X, u.Y}, Coord{x, y})
}

func (b *Board) GetAllPaths(u *Unit) map[Coord]int {
	return b.pathfinding.GetAllPaths(b, u, Coord{u.X, u.Y})
}
