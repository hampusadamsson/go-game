package game

import (
	"errors"
	"log"
)

type Unit struct {
	Img             ImageMeta // Metadata for image and rendering
	X               int       // x location in game tiles
	Y               int       // y location in game tiles
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
	return u.ExhaustedMove == false && u.ExhaustedAttack == false
}

func (u *Unit) canAttack() bool {
	return u.ExhaustedAttack == false && ((u.ExhaustedMove && u.CanMoveAttack) || u.ExhaustedMove == false)
}

func (u *Unit) GetAllAttackCoords() map[Coord]bool { // TODO broken
	pos := make(map[Coord]bool, 0)
	rng := u.attackRange
	for i := 0; i < rng+1; i++ {
		for j := -rng + i; j <= rng-i; j++ {
			if u.X == u.X+i && u.Y == u.Y-j { // your own tile
			} else {
				pos[Coord{u.X + i, u.Y - j}] = true
				pos[Coord{u.X - i, u.Y - j}] = true
			}
		}
	}
	return pos
}

func (u *Unit) canAttackUnit(target Coord) bool {
	_, coordInMap := u.GetAllAttackCoords()[target]
	log.Println(coordInMap)
	return coordInMap
}

func (u *Unit) fight(u2 *Unit) (bool, error) {
	if u.sameOwner(u2) {
		return false, errors.New("same owner")
	}
	u2.HP = u2.HP - u.Damage
	u.ExhaustedAttack = true
	return u2.HP <= 0, nil
}
