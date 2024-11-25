package main

import (
	"context"
	"light-control/internal/adalight"
	"light-control/internal/effects"
)

func main() {
	c := adalight.New("/dev/ttyUSB0", 115200, 60)

	// _ = c.Run(effects.NewFlash(color.RGBA{R: 255, G: 245, B: 182}, 2*time.Second))
	// _ = c.Run(effects.NewFlash(color.RGBA{R: 239, G: 192, B: 112}, 2*time.Second))
	// _ = c.Run(effects.NewFlash(color.RGBA{R: 228, G: 112, B: 37}, 10*time.Hour))
	// _ = c.Run(effects.NewFlash(color.RGBA{R: 247, G: 206, B: 226}, 10*time.Hour))
	// _ = c.Run(effects.NewFlash(color.RGBA{R: 247, G: 227, B: 204}, 10*time.Hour))

	// _ = c.Run(effects.NewFlash(color.NRGBA{B: 0x7F, A: 0xFF}, 2*time.Second))
	// _ = c.Run(effects.NewTrain(color.NRGBA{G: 0xFF, A: 0xFF}, 1, true))

	// _ = c.Run(effects.NewBreath(color.NRGBA{R: 0xFF}, 1900*time.Millisecond, context.TODO()))
	// _ = c.Run(effects.NewBreath(color.RGBA{R: 228, G: 112, B: 37}, 1900*time.Millisecond, context.TODO()))

	_ = c.Run(effects.NewLoad(context.TODO()))
}
