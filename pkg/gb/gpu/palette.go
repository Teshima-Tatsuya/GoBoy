package gpu

import "image/color"

var palette = []color.RGBA{
	{175, 197, 160, 255},
	{93, 147, 66, 255},
	{22, 63, 48, 255},
	{0, 40, 0, 255},
}

type Palette struct {
	// FF47
	BGP byte
}

func GetPalette(idx int) color.RGBA {
	return palette[idx]
}
