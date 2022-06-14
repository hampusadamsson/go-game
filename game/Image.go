package game

type imageMeta struct {
	x         int
	y         int
	size      int
	animation bool
}

func (t *imageMeta) ToggleAnimation() {
	if t.animation {
		t.animation = false
	} else {
		t.animation = true
	}
}

func (t *imageMeta) GetImage() (int, int, int, bool) {
	return t.x, t.y, t.size, t.animation
}

var (
	Mountain     = imageMeta{x: 252, y: 1548, size: 15, animation: false}
	Woods        = imageMeta{x: 444, y: 1567, size: 15}
	Ocean        = imageMeta{x: 340, y: 1567, size: 15}
	Plain        = imageMeta{x: 217, y: 1567, size: 15}
	NoZone       = imageMeta{x: 51, y: 1688, size: 15, animation: false}
	Picker       = imageMeta{x: 70, y: 1684, size: 30, animation: false}
	PickerAttack = imageMeta{x: 101, y: 1684, size: 30, animation: false}
)

var (
	Infantry = imageMeta{x: 252, y: 1444, size: 15, animation: true}
)
