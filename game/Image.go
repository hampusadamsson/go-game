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
	Mountain = imageMeta{x: 252, y: 1548, size: 15}
	Woods    = imageMeta{x: 444, y: 1567, size: 15}
	Ocean    = imageMeta{x: 340, y: 1567, size: 15}
	Plain    = imageMeta{x: 217, y: 1567, size: 15}
)

var (
	Infantry = imageMeta{x: 252, y: 1444, size: 15, animation: true}
)
