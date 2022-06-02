package game

import "errors"

type Unit struct {
	Img        imageMeta // Metadata for image and rendering
	x          int       // x location in game tiles
	y          int       // y location in game tiles
	Owner      *Player   // Who controls the unit
	Exhausted  bool      // true = unit has an action
	Movement   int       // Number of tiles unit can move
	MoveAttack bool      // Can the unit move and attack
	Damage     int		
	Range      int
	HP         int
}

func (u *Unit) sameOwner(u2 *Unit) bool {
	return u.Owner == u2.Owner
}

// func (u *Unit) canAttack(u2 *Unit) bool {
// 	return !u.sameOwner(u2)
// }

func (u *Unit) fight(u2 *Unit) (bool, error) {
	if u.sameOwner(u2) {
		return false, errors.New("same owner")
	}
	u2.HP = u2.HP - u.Damage
	return u2.HP <= 0, nil
}
