package game

type BoardFactory struct{}

var (
	tf = TileFactory{}
	uf = UnitFactory{}
)

func (b *BoardFactory) OneVsOne(p1 *Player, p2 *Player) Board {
	newBoard := Board{
		Tiles: [][]tile{
			{tf.Field(), tf.Field(), tf.Field()},
			{tf.Field(), tf.Field(), tf.Field()},
			{tf.Field(), tf.Field(), tf.Field()},
		},
	}
	i := uf.Infantry(p1, 0, 0)
	j := uf.Infantry(p2, 2, 2)
	newBoard.addUnits(&i, &j)
	return newBoard
}

func (b *BoardFactory) Tutorial1VNone(p1 *Player) Board {
	newBoard := Board{
		Tiles: [][]tile{
			{tf.Field(), tf.Field(), tf.Field(), tf.Field(), tf.Field(), tf.Ocean(), tf.Field(), tf.Field(), tf.Field()},
			{tf.Field(), tf.Field(), tf.Field(), tf.Ocean(), tf.Ocean(), tf.Ocean(), tf.Field(), tf.Field(), tf.Field()},
			{tf.Field(), tf.Field(), tf.Field(), tf.Ocean(), tf.Mountain(), tf.Mountain(), tf.Field(), tf.Field(), tf.Field()},
			{tf.Field(), tf.Field(), tf.Field(), tf.Ocean(), tf.Field(), tf.Field(), tf.Field(), tf.Field(), tf.Field()},
			{tf.Ocean(), tf.Ocean(), tf.Ocean(), tf.Ocean(), tf.Field(), tf.Field(), tf.Field(), tf.Field(), tf.Field()},
			{tf.Field(), tf.Field(), tf.Field(), tf.Field(), tf.Field(), tf.Field(), tf.Field(), tf.Field(), tf.Field()},
			{tf.Field(), tf.Field(), tf.Field(), tf.Field(), tf.Field(), tf.Field(), tf.Field(), tf.Field(), tf.Field()},
			{tf.Field(), tf.Field(), tf.Field(), tf.Field(), tf.Field(), tf.Field(), tf.Field(), tf.Field(), tf.Field()},
		},
	}
	i := uf.Infantry(p1, 5, 4)
	newBoard.addUnits(&i)
	return newBoard
}
