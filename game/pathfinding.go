package game

type pathfinding struct {
}

func (p *pathfinding) unitCanPass(u *Unit, t *tile) bool {
	if u2, err := t.GetUnit(); err == nil {
		v := u.sameOwner(u2)
		return v
	} else {
		return true
	}
}

func (p *pathfinding) findShortestPath(b *Board, unit *Unit, from Coord, to Coord) ([]Coord, int, bool) {
	history := make(map[Coord]int)
	paths := make(map[Coord]Coord)
	neighbours := []Coord{from}
	for len(neighbours) != 0 {
		curTile := neighbours[0]
		neighbours = neighbours[1:]
		next := b.getAdjacent(curTile.X, curTile.Y)
		for _, nextCord := range next {
			nextTile, _ := b.getTile(nextCord.X, nextCord.Y)
			wayHereCost := history[curTile] + nextTile.Cost
			if unit.Movement >= wayHereCost {
				if p.unitCanPass(unit, nextTile) {
					if history[*nextCord] == 0 || wayHereCost < history[*nextCord] {
						history[*nextCord] = wayHereCost
						paths[*nextCord] = curTile
						neighbours = append(neighbours, *nextCord)
					}

				}
			}
		}
	}
	if _, ok := paths[to]; ok {
		return p.getWayBack(paths, from, to), history[to], true
	} else {
		return nil, 0, false
	}
}

func (p *pathfinding) getWayBack(paths map[Coord]Coord, from Coord, to Coord) []Coord {
	path := make([]Coord, 0)
	path = append(path, to)
	for {
		if val, ok := paths[to]; ok {
			if val.X == from.X && val.Y == from.Y {
				return path
			}
			path = append(path, val)
			to = val
		} else {
			return path
		}
	}
}
