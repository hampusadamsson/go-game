package game

import "errors"

type tile struct {
	Img     imageMeta
	unit    *Unit
	terrain terrain
	cost    int
}

type terrain int

const (
	sea terrain = iota
	ground
	difficult
)

func (t *tile) GetUnit() (*Unit, error) {
	if t.isOccupied() {
		return t.unit, nil
	}
	return nil, errors.New("no unit at location")
}

func (t *tile) isOccupied() bool {
	return t.unit != nil
}

func (t *tile) AddUnit(u *Unit) {
	t.unit = u
}

func (t *tile) RemoveUnit() {
	var x *Unit
	t.unit = x
}
