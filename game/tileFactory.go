package game

type TileFactory struct{}

func (u *TileFactory) Field() Tile {
	return Tile{
		Img:     Plain,
		terrain: ground,
	}
}

func (u *TileFactory) Ocean() Tile {
	return Tile{
		Img:     Ocean,
		terrain: sea,
	}
}

func (u *TileFactory) Mountain() Tile {
	return Tile{
		Img:     Mountain,
		terrain: difficult,
	}
}
