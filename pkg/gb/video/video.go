package video

import (
	"image"
	"image/color"
	"math/rand"
)

type Color uint16

type Video struct {
	Palette [4]Color
}

func New() *Video {
	palette := [4]Color{
		// BGP
		// @see https://gbdev.io/pandocs/Palettes.html#ff47---bgp-bg-palette-data-rw---non-cgb-mode-only
		0x7fff, // -> 0b11111, 0b11111, 0b11111 (white)
		0x56b5, // -> 0b10101, 0b10101, 0b10101 (light gray)
		0x294a, // -> 0b01010, 0b01010, 0b01010 (dark gray)
		0x0000, // -> 0b00000, 0b00000, 0b00000 (black)
	}

	return &Video{
		Palette: palette,
	}
}

func (g *Video) Display() *image.RGBA {
	i := image.NewRGBA(image.Rect(0, 0, 160, 144))
	for y := 0; y < 144; y++ {
		for x := 0; x < 160; x++ {
			//			p := g.Renderer.outputBuffer[y*160+x]
			//			red, green, blue := byte((p&0b11111)*8), byte(((p>>5)&0b11111)*8), byte(((p>>10)&0b11111)*8)

			i.SetRGBA(x, y, color.RGBA{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), 0xff})
		}
	}
	return i
}
