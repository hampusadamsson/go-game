package game

type GameFactory struct{}

var (
	bf = BoardFactory{}
)

func (b *GameFactory) OneVsOne(player1 *Player, player2 *Player) *Game {
	players := []*Player{player1, player2}
	return &Game{
		Players: players,
		Board:   bf.OneVsOne(player1, player2),
		Turn:    player1,
	}

}

func (b *GameFactory) OneVsOneFirstGame(player1 *Player, player2 *Player) *Game {
	players := []*Player{player1, player2}
	return &Game{
		Players: players,
		Board:   bf.First(player1, player2),
		Turn:    player1,
	}

}
