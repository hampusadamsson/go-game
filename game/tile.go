package game

import "errors"

type Tile struct {
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

func (t *Tile) GetUnit() (*Unit, error) {
	if t.isOccupied() {
		return t.unit, nil
	}
	return nil, errors.New("no unit at location")
}

func (t *Tile) isOccupied() bool {
	return t.unit != nil
}

func (t *Tile) AddUnit(u *Unit) {
	t.unit = u
}

func (t *Tile) RemoveUnit() {
	var x *Unit
	t.unit = x
}

func GetExampleMap() [][]Tile {
	tf := TileFactory{}
	uf := UnitFactory{}
	p := tf.Field()
	i := uf.Infantry()
	p.AddUnit(&i)
	return [][]Tile{
		{tf.Field(), tf.Field(), tf.Field(), tf.Field(), tf.Field(), tf.Ocean(), tf.Field(), tf.Field(), tf.Field()},
		{tf.Field(), tf.Field(), tf.Field(), tf.Ocean(), tf.Ocean(), tf.Ocean(), tf.Field(), tf.Field(), tf.Field()},
		{tf.Field(), tf.Field(), tf.Field(), tf.Ocean(), tf.Mountain(), tf.Mountain(), tf.Field(), tf.Field(), tf.Field()},
		{tf.Field(), tf.Field(), tf.Field(), tf.Ocean(), tf.Field(), tf.Field(), tf.Field(), tf.Field(), tf.Field()},
		{tf.Ocean(), tf.Ocean(), tf.Ocean(), tf.Ocean(), tf.Field(), tf.Field(), tf.Field(), tf.Field(), tf.Field()},
		{tf.Field(), tf.Field(), tf.Field(), tf.Field(), p, tf.Field(), tf.Field(), tf.Field(), tf.Field()},
		{tf.Field(), tf.Field(), tf.Field(), tf.Field(), tf.Field(), tf.Field(), tf.Field(), tf.Field(), tf.Field()},
		{tf.Field(), tf.Field(), tf.Field(), tf.Field(), tf.Field(), tf.Field(), tf.Field(), tf.Field(), tf.Field()},
	}
}
