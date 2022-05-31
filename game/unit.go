package game

type Unit struct {
	Img       ImageMeta // Metadata for image and rendering
	x         int       // x location in game tiles
	y         int       // y location in game tiles
	Owner     *Player   // Who controls the unit
	Exhausted bool      // true = unit has an action
	Damage    int
	Range     int
	HP        int
}
