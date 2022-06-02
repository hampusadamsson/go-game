package game

type Game struct {
	board board
	// round    int
	players []*Player
	// turn     *Player
	// gameOver bool
}

func NewGame() *Game {
	return &Game{} //Fill out
}
