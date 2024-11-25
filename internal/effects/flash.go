package effects

import (
	"image/color"
	"time"

	"light-control/internal/adalight"
)

type Flash struct {
	color  color.Color
	dur    time.Duration
	strip  *adalight.Strip
	frames int
}

func NewFlash(c color.Color, d time.Duration) (f *Flash) {
	return &Flash{color: c, dur: d}
}

func (f *Flash) Init(s *adalight.Strip, fps int) string {
	f.strip = s
	f.frames = int(float64(fps) * f.dur.Seconds())
	return "Flash"
}

func (f *Flash) Frame(num int) bool {
	if num < f.frames {
		f.strip.SetAll(f.color)
		return false
	} else {
		f.strip.SetAllRGB(color.RGBA{})
		return true
	}
}
