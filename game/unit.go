package game

type Unit struct {
	Img ImageMeta

	Owner     *Player
	Exhausted bool
	Damage    int
	Range     int
	HP        int
}
