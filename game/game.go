package game

import (
	"errors"
)

type Game struct {
	Board   Board
	round   int
	Players []*Player
	turn    *Player
	// gameOver bool
}

func NewGame() *Game {
	return &Game{} //Fill out
}

func (g *Game) Move(p *Player, from, to coord) (bool, error) {
	if g.turn == p {
		if u, err := g.Board.getUnit(from.x, from.y); err == nil {
			if u.canMove() {
				if _, _, canMoveThere := g.Board.getPath(u, to.x, to.y); canMoveThere == true {
					if notOccupiedCoord, err2 := g.Board.move(u, to.x, to.y); notOccupiedCoord == true {
						u.ExhaustedMove = true
						return true, nil
					} else {
						return false, err2 // no path to target
					}
				} else {
					return false, errors.New("no path to destination")
				}
			} else {
				return false, errors.New("unit is exhaused")
			}
		} else {
			return false, err // no unit at location
		}
	}
	return false, errors.New("not your turn")
}

func (g *Game) Attack(p *Player, attacker coord, defender coord) (bool, error) {
	if g.turn == p {
		if u, err := g.Board.getUnit(attacker.x, attacker.y); err == nil {
			if u.canAttack() {
				if u.canAttackUnit(defender) {
					u.ExhaustedAttack = true
					return true, nil
				} else {
					return false, errors.New("target not in range")
				}
			} else {
				return false, errors.New("unit is exhaused")
			}
		} else {
			return false, err // no unit at location
		}
	}
	return false, errors.New("not your turn")
}

func (g *Game) refreshAllUnits(p *Player) {
	units := g.Board.getUnits(p)
	for i := 0; i < len(units); i++ {
		units[i].refresh()
	}
}

func (g *Game) ChangeTurn(p *Player) bool {
	if g.turn == p {
		g.refreshAllUnits(p)
		g.round++
		g.turn = g.Players[g.round%len(g.Players)]
		return true
	}
	return false
}
