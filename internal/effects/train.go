package effects

import (
	"image/color"

	"light-control/internal/adalight"
)

type Train struct {
	color    color.Color
	repeats  int
	strip    *adalight.Strip
	infinite bool
}

func NewTrain(c color.Color, r int, infinite bool) (f *Train) {
	return &Train{color: c, repeats: r, infinite: infinite}
}

func (t *Train) Init(s *adalight.Strip, _ int) string {
	t.strip = s
	return "Train"
}

func (t *Train) Frame(num int) bool {
	t.strip.SetAllRGB(color.RGBA{})

	if num < t.repeats*t.strip.Length() || t.infinite {
		t.strip.Set(t.strip.Length()-1-num%t.strip.Length(), t.color)
		return false
	}
	return true
}
