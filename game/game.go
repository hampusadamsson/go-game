package game

type Game struct {
	Board Board
	// round    int
	Players []*Player
	// turn     *Player
	// gameOver bool
}

func NewGame() *Game {
	return &Game{} //Fill out
}
