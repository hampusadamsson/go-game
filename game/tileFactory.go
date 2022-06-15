package game

type TileFactory struct{}

func (u *TileFactory) Field() tile {
	return tile{
		Img:     *Plain,
		terrain: ground,
		Cost:    1,
	}
}

func (u *TileFactory) Ocean() tile {
	return tile{
		Img:     *Ocean,
		terrain: sea,
		Cost:    2,
	}
}

func (u *TileFactory) Mountain() tile {
	return tile{
		Img:     *Mountain,
		terrain: difficult,
		Cost:    4,
	}
}
