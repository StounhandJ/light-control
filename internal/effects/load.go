package effects

import (
	"context"
	"image/color"
	"time"

	"light-control/internal/adalight"
	"light-control/internal/utils"
)

type Load struct {
	c           context.Context
	strip       *adalight.Strip
	avgPercent  float64
	framesSpeed int
}

func NewLoad(c context.Context) (f *Load) {
	return &Load{c: c}
}

func (f *Load) Init(s *adalight.Strip, fps int) string {
	f.strip = s
	f.framesSpeed = fps
	return "Flash"
}

func (f *Load) Frame(num int) bool {
	avgPercent := (utils.UseCPU(time.Millisecond) + utils.UsedMemoryPercent()) / 2
	f.avgPercent = f.avgPercent - ((f.avgPercent - avgPercent) / float64(f.framesSpeed))
	colorValue := uint8(f.avgPercent / 100 * 255)

	select {
	case <-f.c.Done():
		f.strip.SetAllRGB(color.RGBA{})
		return true
	default:
		f.strip.SetAll(color.RGBA{R: colorValue, G: 255 - colorValue, B: 0})
		return false
	}
}
