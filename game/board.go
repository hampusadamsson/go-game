package game

type Board struct {
	tiles [][]Tile
}

func (b *Board) getTile(x int, y int) *Tile {
	return &b.tiles[x][y]
}

func (b *Board) getTiles() chan *Tile {
	c := make(chan *Tile)
	go func() {
		for i := 0; i < len(b.tiles); i++ {
			for j := 0; j < len(b.tiles[i]); j++ {
				c <- &b.tiles[i][j]
			}
		}
		close(c)
	}()
	return c
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
	if u, err := b.getTile(x, y).GetUnit(); err == nil {
		return u, nil
	} else {
		return nil, err
	}
}
