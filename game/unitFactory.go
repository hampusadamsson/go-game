package game

type UnitFactory struct{}

func (u *UnitFactory) Infantry(owner *Player, x, y int) Unit { // Require owner
	return Unit{
		Img:           *Infantry,
		Owner:         owner,
		CanMoveAttack: true,
		attackRange:   1,
		Movement:      3,
		Damage:        3,
		HP:            2,
		X:             x,
		Y:             y,
	}
}

func (u *UnitFactory) Mech(owner *Player, x, y int) Unit { // Require owner
	return Unit{
		Img:           *Mech,
		Owner:         owner,
		CanMoveAttack: true,
		attackRange:   1,
		Movement:      2,
		Damage:        4,
		HP:            3,
		X:             x,
		Y:             y,
	}
}

func (u *UnitFactory) Reckon(owner *Player, x, y int) Unit { // Require owner
	return Unit{
		Img:           *Reckon,
		Owner:         owner,
		CanMoveAttack: true,
		attackRange:   1,
		Movement:      8,
		Damage:        3,
		HP:            7,
		X:             x,
		Y:             y,
	}
}

func (u *UnitFactory) Supply(owner *Player, x, y int) Unit { // Require owner
	return Unit{
		Img:           *Supply,
		Owner:         owner,
		CanMoveAttack: true,
		attackRange:   0,
		Movement:      5,
		Damage:        0,
		HP:            8,
		X:             x,
		Y:             y,
	}
}

func (u *UnitFactory) Tank(owner *Player, x, y int) Unit { // Require owner
	return Unit{
		Img:           *Tank,
		Owner:         owner,
		CanMoveAttack: true,
		attackRange:   0,
		Movement:      5,
		Damage:        0,
		HP:            9,
		X:             x,
		Y:             y,
	}
}

func (u *UnitFactory) HeavyTank(owner *Player, x, y int) Unit { // Require owner
	return Unit{
		Img:           *HeavyTank,
		Owner:         owner,
		CanMoveAttack: true,
		attackRange:   0,
		Movement:      5,
		Damage:        0,
		HP:            10,
		X:             x,
		Y:             y,
	}
}

func (u *UnitFactory) Artilery(owner *Player, x, y int) Unit { // Require owner
	return Unit{
		Img:           *Artilery,
		Owner:         owner,
		CanMoveAttack: false,
		attackRange:   2,
		Movement:      5,
		Damage:        10,
		HP:            9,
		X:             x,
		Y:             y,
	}
}
