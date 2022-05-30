package game

type Unit struct {
	Img ImageMeta

	Owner  *Player
	Damage int
	Range  int
	HP     int
}
