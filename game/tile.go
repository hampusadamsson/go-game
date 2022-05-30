package game

import "errors"

type Tile struct {
	Img  ImageMeta
	unit *Unit
}

func (t *Tile) GetUnit() (*Unit, error) {
	if t.HasUnit() {
		return t.unit, nil
	}
	return nil, errors.New("no unit")
}

func (t *Tile) HasUnit() bool {
	return t.unit != nil
}

func (t *Tile) AddUnit(u *Unit) {
	t.unit = u
}

func GetExampleMap() [][]Tile {
	tf := TileFactory{}
	uf := UnitFactory{}
	p := tf.Plain()
	i := uf.Infantry()
	p.AddUnit(&i)
	return [][]Tile{
		{tf.Plain(), tf.Plain(), tf.Plain(), tf.Plain(), tf.Plain(), tf.Ocean(), tf.Plain(), tf.Plain(), tf.Plain()},
		{tf.Plain(), tf.Plain(), tf.Plain(), tf.Ocean(), tf.Ocean(), tf.Ocean(), tf.Plain(), tf.Plain(), tf.Plain()},
		{tf.Plain(), tf.Plain(), tf.Plain(), tf.Ocean(), tf.Mountain(), tf.Mountain(), tf.Plain(), tf.Plain(), tf.Plain()},
		{tf.Plain(), tf.Plain(), tf.Plain(), tf.Ocean(), tf.Plain(), tf.Plain(), tf.Plain(), tf.Plain(), tf.Plain()},
		{tf.Ocean(), tf.Ocean(), tf.Ocean(), tf.Ocean(), tf.Plain(), tf.Plain(), tf.Plain(), tf.Plain(), tf.Plain()},
		{tf.Plain(), tf.Plain(), tf.Plain(), tf.Plain(), p, tf.Plain(), tf.Plain(), tf.Plain(), tf.Plain()},
		{tf.Plain(), tf.Plain(), tf.Plain(), tf.Plain(), tf.Plain(), tf.Plain(), tf.Plain(), tf.Plain(), tf.Plain()},
		{tf.Plain(), tf.Plain(), tf.Plain(), tf.Plain(), tf.Plain(), tf.Plain(), tf.Plain(), tf.Plain(), tf.Plain()},
	}
}
