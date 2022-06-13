package game

type GameFactory struct{}

var (
	bf = BoardFactory{}
)

func (b *GameFactory) Tutorial(player1 *Player) *Game {
	players := []*Player{player1}
	return &Game{
		Players: players,
		Board:   bf.Tutorial1VNone(player1),
		turn:    player1,
	}

}
