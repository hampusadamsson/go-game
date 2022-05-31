package game

import "errors"

type Tile struct {
	Img     ImageMeta
	unit    *Unit
	terrain Terrain
}

type Terrain int

const (
	sea Terrain = iota
	ground
	difficult
)

func (t *Tile) GetUnit() (*Unit, error) {
	if t.HasUnit() {
		return t.unit, nil
	}
	return nil, errors.New("no unit at location")
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
