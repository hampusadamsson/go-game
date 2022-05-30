package game

type TileFactory struct{}

func (u *TileFactory) Plain() Tile {
	return Tile{
		Img: Plain,
	}
}

func (u *TileFactory) Ocean() Tile {
	return Tile{
		Img: Ocean,
	}
}

func (u *TileFactory) Mountain() Tile {
	return Tile{
		Img: Ocean,
	}
}
