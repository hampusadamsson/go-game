package game

type UnitFactory struct{}

func (u *UnitFactory) Infantry(owner *Player) Unit { // Require owner
	return Unit{
		Img:           Infantry,
		Owner:         owner,
		CanMoveAttack: true,
		attackRange:   1,
		Movement:      3,
		Damage:        3,
		HP:            10}
}
