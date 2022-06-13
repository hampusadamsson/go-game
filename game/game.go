package game

import (
	"errors"
	"fmt"
	"log"
)

type Game struct {
	Board   Board
	round   int
	Players []*Player
	turn    *Player
	// gameOver bool
}

func (g *Game) Run() {
	for i := 0; i < len(g.Players); i++ {
		go g.playerEventHandler(g.Players[i])
	}
}

func (g *Game) playerEventHandler(p *Player) {
	for {
		action := <-p.Act
		switch action.ActionType {
		case ActionMove:
			fmt.Println(action)
			_, err := g.move(p, action.From, action.To)
			fmt.Println(err)
		case ActionAttack:
			_, err := g.attack(p, action.From, action.To)
			fmt.Println(err)
		case ActionEnd:
			ok := g.changeTurn(p)
			fmt.Println(ok)
		}
	}
}

func (g *Game) move(p *Player, from, to Coord) (bool, error) {
	log.Println(from, to)
	if g.turn == p {
		if u, err := g.Board.getUnit(from.X, from.Y); err == nil {
			log.Println("COORD", u.X, u.Y)
			if u.canMove() {
				if path, cost, canMoveThere := g.Board.getPath(u, to.X, to.Y); canMoveThere == true {
					log.Println(cost, path)
					if notOccupiedCoord, err2 := g.Board.move(u, to.X, to.Y); notOccupiedCoord == true {
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

func (g *Game) attack(p *Player, attacker Coord, defender Coord) (bool, error) {
	if g.turn == p {
		if u, err := g.Board.getUnit(attacker.X, attacker.Y); err == nil {
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

func (g *Game) changeTurn(p *Player) bool {
	fmt.Println(g.turn, p)
	if g.turn == p {
		g.refreshAllUnits(p)
		g.round++
		g.turn = g.Players[g.round%len(g.Players)]
		return true
	}
	return false
}
