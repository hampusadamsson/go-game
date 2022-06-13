package game

type Player struct {
	Name string
	Act  chan Action
}

type Action struct {
	ActionType ActionId
	From       Coord
	To         Coord
}

type ActionId int

const (
	ActionMove   = iota
	ActionAttack = iota
	ActionEnd    = iota
)
