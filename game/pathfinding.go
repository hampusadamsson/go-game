package game

type Pathfinding struct {
}

func (p *Pathfinding) unitCanPass(u *Unit, t *Tile) bool {
	return true //TODO - add logic for passapble terrain etc.
}

func (p *Pathfinding) FindShortestPath(b *Board, unit *Unit, from coord, to coord) ([]coord, int, bool) {
	history := make(map[coord]int)
	paths := make(map[coord]coord)
	neighbours := []coord{from}
	for len(neighbours) != 0 {
		curTile := neighbours[0]
		neighbours = neighbours[1:]
		next := b.getAdjacent(curTile.x, curTile.y)
		for _, nextCord := range next {
			nextTile, _ := b.getTile(nextCord.x, nextCord.y)
			wayHereCost := history[curTile] + nextTile.cost
			if unit.Movement >= wayHereCost {
				if history[*nextCord] == 0 || wayHereCost < history[*nextCord] && p.unitCanPass(unit, nextTile) {
					history[*nextCord] = wayHereCost
					paths[*nextCord] = curTile
					neighbours = append(neighbours, *nextCord)
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

func (p *Pathfinding) getWayBack(paths map[coord]coord, from coord, to coord) []coord {
	path := make([]coord, 0)
	path = append(path, to)
	for {
		if val, ok := paths[to]; ok {
			if val.x == from.x && val.y == from.y {
				return path
			}
			path = append(path, val)
			to = val
		} else {
			return path
		}
	}
}
