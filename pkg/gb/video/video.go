package video

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
