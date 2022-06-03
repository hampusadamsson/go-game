package game

type TileFactory struct{}

func (u *TileFactory) Field() tile {
	return tile{
		Img:     Plain,
		terrain: ground,
		cost:    1,
	}
}

func (u *TileFactory) Ocean() tile {
	return tile{
		Img:     Ocean,
		terrain: sea,
		cost:    1,
	}
}

func (u *TileFactory) Mountain() tile {
	return tile{
		Img:     Mountain,
		terrain: difficult,
		cost:    3,
	}
}
