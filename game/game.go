package game

type Game struct {
	board Board
	// round    int
	players []*Player
	// turn     *Player
	// gameOver bool
}

func NewGame() *Game {
	return &Game{} //Fill out
}

// func (g *Game) move(u *Unit, x int, y int) error {
// 	g.tiles[u.x][u.y].unit = nil
// 	if g.ge {

// 	}
// 	g.tiles[u.x][u.y].unit = nil
// }
