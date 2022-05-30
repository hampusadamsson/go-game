package game

type Game struct {
	tiles [][]Tile
	// round    int
	players []*Player
	// turn     *Player
	// gameOver bool
}

func NewGame() *Game {
	return &Game{} //Fill out
}

func (g *Game) GetUnits(p *Player) []*Unit {
	units := make([]*Unit, 0)
	for elem := range g.getTiles() {
		if u, err := elem.GetUnit(); err == nil {
			if u.Owner == p {
				units = append(units, u)
			}
		}
	}
	return units
}

func (g *Game) getTiles() chan *Tile {
	c := make(chan *Tile)
	go func() {
		for i := 0; i < len(g.tiles); i++ {
			for j := 0; j < len(g.tiles[i]); j++ {
				c <- &g.tiles[i][j]
			}
		}
		close(c)
	}()
	return c
}
