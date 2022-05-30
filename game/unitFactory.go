package game

type UnitFactory struct{}

func (u *UnitFactory) Infantry() Unit { // Require owner
	return Unit{
		Img: Infantry,
	}
}
