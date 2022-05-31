package game

import "errors"

type Pathfinding struct {
	history map[coord]int
	paths   map[coord]coord
}

type PathResult struct {
	path []coord
	cost int
}

func newPathfinding() Pathfinding {
	return Pathfinding{
		history: make(map[coord]int),
		paths:   make(map[coord]coord),
	}
}

func (p *Pathfinding) getPath(b *Board, x1 int, y1 int, x2 int, y2 int) (*PathResult, error) {
	START_PENALTY := 99999
	start := coord{x: x1, y: y1}
	neighbours := []coord{start}
	p.history[start] = START_PENALTY
	for len(neighbours) != 0 {
		curTile := neighbours[0]
		neighbours = neighbours[1:]
		next := b.getAdjacent(curTile.x, curTile.y)
		for _, nextCord := range next {
			nextTile, _ := b.getTile(nextCord.x, nextCord.y)
			wayHereCost := p.history[curTile] + nextTile.cost
			if p.history[*nextCord] == 0 || wayHereCost < p.history[*nextCord] {
				p.history[*nextCord] = wayHereCost
				p.paths[*nextCord] = curTile
				neighbours = append(neighbours, *nextCord)
			}
		}
	}
	totalCost := p.history[coord{x2, y2}]
	if totalCost != 0 {
		return &PathResult{
			cost: totalCost - START_PENALTY,
			path: p.getWayBack(x2, y2),
		}, nil
	} else {
		return nil, errors.New("no way to destination")
	}
}

func (p *Pathfinding) getWayBack(x int, y int) []coord {
	prev := coord{x, y}
	path := make([]coord, 0)
	path = append(path, prev)
	for {
		if val, ok := p.paths[prev]; ok {
			path = append(path, val)
			prev = val
		} else {
			return path
		}
	}
}
