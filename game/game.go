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
