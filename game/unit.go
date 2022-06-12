package game

import (
	"errors"
)

type Unit struct {
	Img             imageMeta // Metadata for image and rendering
	x               int       // x location in game tiles
	y               int       // y location in game tiles
	Owner           *Player   // Who controls the unit
	ExhaustedMove   bool      // true = unit has a move action
	ExhaustedAttack bool      // true = unit has an attack action
	Movement        int       // Number of tiles unit can move
	CanMoveAttack   bool      // Can the unit move and attack
	attackRange     int       // Range for attack - 1 means melee
	Damage          int
	HP              int
}

func (u *Unit) sameOwner(u2 *Unit) bool {
	return u.Owner == u2.Owner
}

func (u *Unit) refresh() {
	u.ExhaustedMove = false
	u.ExhaustedAttack = false
}

func (u *Unit) canMove() bool {
	return u.ExhaustedMove == false
}

func (u *Unit) canAttack() bool {
	return u.ExhaustedAttack == false && ((u.ExhaustedMove && u.CanMoveAttack) || u.ExhaustedMove == false)
}

func (u *Unit) getAllAttackCoords() map[coord]bool { // TODO broken
	pos := make(map[coord]bool, 0)
	rng := u.attackRange
	for i := 0; i < rng+1; i++ {
		for j := -rng + i; j <= rng-i; j++ {
			if u.x == u.x+i && u.y == u.y-j { // your own tile
			} else {
				pos[coord{u.x + i, u.y - j}] = true
				pos[coord{u.x - i, u.y - j}] = true
			}
		}
	}
	return pos
}

func (u *Unit) canAttackUnit(target coord) bool {
	_, coordInMap := u.getAllAttackCoords()[target]
	return coordInMap
}

func (u *Unit) fight(u2 *Unit) (bool, error) {
	if u.sameOwner(u2) {
		return false, errors.New("same owner")
	}
	u2.HP = u2.HP - u.Damage
	return u2.HP <= 0, nil
}
