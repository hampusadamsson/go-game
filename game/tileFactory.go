package game

type TileFactory struct{}

func (u *TileFactory) Field() Tile {
	return Tile{
		Img:     Plain,
		terrain: ground,
		cost:    1,
	}
}

func (u *TileFactory) Ocean() Tile {
	return Tile{
		Img:     Ocean,
		terrain: sea,
		cost:    1,
	}
}

func (u *TileFactory) Mountain() Tile {
	return Tile{
		Img:     Mountain,
		terrain: difficult,
		cost:    3,
	}
}
