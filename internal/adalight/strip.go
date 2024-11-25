package adalight

import (
	"fmt"
	"image/color"
)

type Strip struct {
	length int
	pixels []uint8
}

func (s *Strip) Length() int {
	return s.length
}

func (s *Strip) SetRGB(i int, rgb color.RGBA) {
	if i < 0 || i > s.length {
		return
	}
	s.pixels[(3*i)+0] = rgb.R
	s.pixels[(3*i)+1] = rgb.G
	s.pixels[(3*i)+2] = rgb.B
}

func (s *Strip) Set(i int, c color.Color) {
	rgb := color.RGBAModel.Convert(c).(color.RGBA)
	s.SetRGB(i, rgb)
}

func (s *Strip) SetAllRGB(rgb color.RGBA) {
	for i := 0; i < s.length; i++ {
		s.SetRGB(i, rgb)
	}
}

func (s *Strip) SetAll(c color.Color) {
	rgb := color.RGBAModel.Convert(c).(color.RGBA)
	s.SetAllRGB(rgb)
}

func newStrip(length int) *Strip {
	if length < 1 {
		fmt.Println("Invalid Strip length:", length)
		panic(fmt.Sprintf("%v", length))
	}
	return &Strip{length: length, pixels: make([]uint8, length*3)}
}
