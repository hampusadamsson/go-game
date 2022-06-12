package game

type BoardFactory struct{}

func (b *BoardFactory) Tutorial1VNone(p1 *Player) Board {
	tf := TileFactory{}
	uf := UnitFactory{}
	p := tf.Field()
	i := uf.Infantry(p1)
	p.AddUnit(&i)

	return Board{
		Tiles: [][]tile{
			{tf.Field(), tf.Field(), tf.Field(), tf.Field(), tf.Field(), tf.Ocean(), tf.Field(), tf.Field(), tf.Field()},
			{tf.Field(), tf.Field(), tf.Field(), tf.Ocean(), tf.Ocean(), tf.Ocean(), tf.Field(), tf.Field(), tf.Field()},
			{tf.Field(), tf.Field(), tf.Field(), tf.Ocean(), tf.Mountain(), tf.Mountain(), tf.Field(), tf.Field(), tf.Field()},
			{tf.Field(), tf.Field(), tf.Field(), tf.Ocean(), tf.Field(), tf.Field(), tf.Field(), tf.Field(), tf.Field()},
			{tf.Ocean(), tf.Ocean(), tf.Ocean(), tf.Ocean(), tf.Field(), tf.Field(), tf.Field(), tf.Field(), tf.Field()},
			{tf.Field(), tf.Field(), tf.Field(), tf.Field(), p, tf.Field(), tf.Field(), tf.Field(), tf.Field()},
			{tf.Field(), tf.Field(), tf.Field(), tf.Field(), tf.Field(), tf.Field(), tf.Field(), tf.Field(), tf.Field()},
			{tf.Field(), tf.Field(), tf.Field(), tf.Field(), tf.Field(), tf.Field(), tf.Field(), tf.Field(), tf.Field()},
		},
	}

}
