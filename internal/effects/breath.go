package effects

import (
	"context"
	"image/color"
	"time"

	"light-control/internal/adalight"
)

type Breath struct {
	color       color.Color
	speed       time.Duration
	c           context.Context
	strip       *adalight.Strip
	framesSpeed int
}

func NewBreath(c color.Color, speed time.Duration, ctx context.Context) (r *Breath) {
	return &Breath{color: c, speed: speed, c: ctx}
}

func (r *Breath) Init(s *adalight.Strip, fps int) string {
	r.strip = s
	r.framesSpeed = int(float64(fps)*r.speed.Seconds()) / 2
	return "Breath"
}

func (r *Breath) Frame(num int) bool {
	alpha := uint8((float64(num%r.framesSpeed) / float64(r.framesSpeed)) * 255.0)
	if num/r.framesSpeed%2 == 1 {
		alpha = 255 - alpha
	}

	rgb := color.NRGBAModel.Convert(r.color).(color.NRGBA)
	rgb.A = alpha

	select {
	case <-r.c.Done():
		r.strip.SetAllRGB(color.RGBA{})
		return true
	default:
		r.strip.SetAll(rgb)
		return false
	}
}
