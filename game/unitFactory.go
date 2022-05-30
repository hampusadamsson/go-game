package game

type UnitFactory struct{}

func (u *UnitFactory) Infantry() Unit {
	return Unit{
		Img: Infantry,
	}
}
