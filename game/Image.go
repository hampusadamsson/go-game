package game

type ImageMeta struct {
	x         int
	y         int
	size      int
	animation bool
}

func (t *ImageMeta) ToggleAnimation() {
	if t.animation {
		t.animation = false
	} else {
		t.animation = true
	}
}

func (t *ImageMeta) GetImage() (int, int, int, bool) {
	return t.x, t.y, t.size, t.animation
}

var (
	Mountain = ImageMeta{x: 252, y: 1548, size: 15}
	Woods    = ImageMeta{x: 444, y: 1567, size: 15}
	Ocean    = ImageMeta{x: 340, y: 1567, size: 15}
	Plain    = ImageMeta{x: 217, y: 1567, size: 15}
)

var (
	Infantry = ImageMeta{x: 252, y: 1444, size: 15, animation: true}
)
