package game

type ImageMeta struct {
	X         int
	Y         int
	Size      int
	Animation bool
}

func (t *ImageMeta) ToggleAnimation() {
	if t.Animation {
		t.Animation = false
	} else {
		t.Animation = true
	}
}

func (t *ImageMeta) GetImage() (*int, *int, *int, *bool) {
	return &t.X, &t.Y, &t.Size, &t.Animation
}

var (
	Mountain     = &ImageMeta{X: 252, Y: 1548, Size: 15, Animation: false}
	Woods        = &ImageMeta{X: 444, Y: 1567, Size: 15}
	Ocean        = &ImageMeta{X: 340, Y: 1567, Size: 15}
	Plain        = &ImageMeta{X: 217, Y: 1567, Size: 15}
	NoZone       = &ImageMeta{X: 51, Y: 1688, Size: 15, Animation: false}
	Picker       = &ImageMeta{X: 70, Y: 1684, Size: 30, Animation: false}
	PickerAttack = &ImageMeta{X: 101, Y: 1684, Size: 30, Animation: false}
	RedZone      = &ImageMeta{X: 51, Y: 1688, Size: 15, Animation: false}
	GreenZone    = &ImageMeta{X: 33, Y: 1688, Size: 15, Animation: false}
)

func NumberImage(i int) *ImageMeta {
	return &ImageMeta{X: 68 + (i)*9, Y: 1719, Size: 8, Animation: false}
}

var (
	Infantry  = &ImageMeta{X: 250 + (18 * 0), Y: 1444, Size: 15, Animation: true}
	Mech      = &ImageMeta{X: 250 + (18 * 1), Y: 1444, Size: 15, Animation: true}
	Reckon    = &ImageMeta{X: 250 + (18 * 2), Y: 1444, Size: 15, Animation: true}
	Supply    = &ImageMeta{X: 250 + (18 * 3), Y: 1444, Size: 15, Animation: true}
	Tank      = &ImageMeta{X: 250 + (18 * 4), Y: 1444, Size: 15, Animation: true}
	HeavyTank = &ImageMeta{X: 250 + (18 * 5), Y: 1444, Size: 15, Animation: true}
)
