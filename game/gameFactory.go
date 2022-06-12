package game

type GameFactory struct{}

var (
	bf = BoardFactory{}
)

func (b *GameFactory) Tutorial(p1 Player) *Game {
	player1 := &Player{}
	players := []*Player{player1}
	return &Game{
		Players: players,
		Board:   bf.Tutorial1VNone(player1),
	}

}
