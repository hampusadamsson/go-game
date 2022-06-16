package game

import (
	"errors"
)

type tile struct {
	Img     ImageMeta
	Unit    *Unit
	terrain terrain
	Cost    int
}

type terrain int

const (
	sea terrain = iota
	ground
	difficult
)

func (t *tile) GetUnit() (*Unit, error) {
	if t.isOccupied() {
		return t.Unit, nil
	}
	return nil, errors.New("no unit at location")
}

func (t *tile) isOccupied() bool {
	return t.Unit != nil
}

func (t *tile) AddUnit(u *Unit) {
	t.Unit = u
}

func (t *tile) RemoveUnit() {
	var x *Unit
	t.Unit = x
}
