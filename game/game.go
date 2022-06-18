package game

import (
	"errors"
	"fmt"
)

type Game struct {
	Board    Board
	round    int
	Players  []*Player
	Turn     *Player
	GameOver bool
}

func (g *Game) Run() {
	for i := 0; i < len(g.Players); i++ {
		go g.playerEventHandler(g.Players[i])
	}
}

func (g *Game) checkWinCondition(p *Player) { // TODO - make more than two players
	pUnits := g.Board.GetUnits(p)
	if len(pUnits) == 0 {
		g.GameOver = true
	}
}

func (g *Game) playerEventHandler(p *Player) {
	for {
		g.checkWinCondition(p) // TODO - don't need to this all the time
		action := <-p.Act
		switch action.ActionType {
		case ActionMove:
			_, err := g.move(p, action.From, action.To)
			if p.Name != "B" {
				fmt.Println(err)
			}
		case ActionAttack:
			_, _, err := g.attack(p, action.From, action.To)
			if p.Name != "B" {
				fmt.Println(err)
			}
		case ActionEnd:
			ok := g.changeTurn(p)
			if p.Name != "B" {
				fmt.Println(ok)
			}
		}

		if g.GameOver {
			return
		}
	}
}

func (g *Game) move(p *Player, from, to Coord) (bool, error) {
	// log.Println(from, to)
	if g.Turn == p {
		if u, err := g.Board.GetUnit(from.X, from.Y); err == nil {
			// log.Println("COORD", u.X, u.Y)
			if p == u.Owner {
				if u.canMove() {
					if _, _, canMoveThere := g.Board.GetShortestPath(u, to.X, to.Y); canMoveThere == true {
						// log.Println(cost, path)
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
				return false, errors.New("unit does not belong to player")
			}
		} else {
			return false, err // no unit at location
		}
	}
	return false, errors.New("not your turn")
}

// Attack returns - attackIsSuccesful, unitDidDie, error
func (g *Game) attack(p *Player, attacker Coord, defender Coord) (bool, bool, error) {
	if g.Turn == p {
		if _, err := g.Board.GetUnit(defender.X, defender.Y); err == nil {
			if u, err := g.Board.GetUnit(attacker.X, attacker.Y); err == nil {
				if p == u.Owner {
					if u.canAttack() {
						if u.canAttackUnit(defender) {
							success, unitDied, err := g.Board.attack(attacker, defender)
							return success, unitDied, err
						} else {
							return false, false, errors.New("target not in range")
						}
					} else {
						return false, false, errors.New("unit is exhaused")
					}
				} else {
					return false, false, errors.New("unit does not belong to player")
				}
			} else {
				return false, false, err // no unit at location attacker
			}
		} else {
			return false, false, err // no unit at location defender
		}
	}
	return false, false, errors.New("not your turn")
}

func (g *Game) refreshAllUnits(p *Player) {
	units := g.Board.GetUnits(p)
	for i := 0; i < len(units); i++ {
		units[i].refresh()
	}
}

func (g *Game) changeTurn(p *Player) bool {
	if g.Turn == p {
		g.refreshAllUnits(p)
		g.round++
		g.Turn = g.Players[g.round%len(g.Players)]
		return true
	}
	return false
}
