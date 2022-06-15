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

func (b *BoardFactory) First(p1 *Player, p2 *Player) Board {
	newBoard := Board{
		Tiles: [][]tile{
			{tf.Field(), tf.Field(), tf.Field(), tf.Field(), tf.Field(), tf.Ocean(), tf.Field(), tf.Field(), tf.Field()},
			{tf.Field(), tf.Field(), tf.Field(), tf.Ocean(), tf.Ocean(), tf.Ocean(), tf.Field(), tf.Field(), tf.Field()},
			{tf.Field(), tf.Field(), tf.Field(), tf.Ocean(), tf.Mountain(), tf.Mountain(), tf.Field(), tf.Field(), tf.Field()},
			{tf.Field(), tf.Field(), tf.Field(), tf.Ocean(), tf.Mountain(), tf.Field(), tf.Field(), tf.Field(), tf.Field()},
			{tf.Ocean(), tf.Ocean(), tf.Ocean(), tf.Ocean(), tf.Mountain(), tf.Field(), tf.Field(), tf.Field(), tf.Field()},
			{tf.Field(), tf.Field(), tf.Field(), tf.Field(), tf.Mountain(), tf.Field(), tf.Field(), tf.Ocean(), tf.Field()},
			{tf.Field(), tf.Field(), tf.Field(), tf.Field(), tf.Mountain(), tf.Mountain(), tf.Field(), tf.Ocean(), tf.Field()},
			{tf.Field(), tf.Field(), tf.Field(), tf.Field(), tf.Field(), tf.Field(), tf.Field(), tf.Ocean(), tf.Field()},
		},
	}
	units := []Unit{
		//p1
		uf.Infantry(p1, 0, 0),
		uf.Mech(p1, 1, 0),
		uf.Reckon(p1, 2, 0),
		uf.Supply(p1, 3, 0),
		uf.HeavyTank(p1, 4, 0),
		//p2
		uf.Tank(p2, 4, 8),
		uf.Infantry(p2, 5, 8),
		uf.Mech(p2, 6, 8),
		uf.Reckon(p2, 7, 8),
	}

	for i, _ := range units {
		u := &units[i]
		newBoard.addUnits(u)
	}

	return newBoard
}
