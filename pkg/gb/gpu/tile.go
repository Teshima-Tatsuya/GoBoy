package gpu

import (
	"image/color"

	"github.com/Teshima-Tatsuya/GoBoy/pkg/util"
)

type Color byte

const (
	White Color = iota
	LightGray
	DarkGray
	Black
)

// https://gbdev.io/pandocs/Tile_Data.html
type Tile struct {
	// 8 x 8 pixels
	Data [8][8]Color
}

func NewTile(bytes16 []byte) *Tile {
	var data [8][8]Color
	for i := 0; i < 8; i++ {
		lower := bytes16[i*2]
		upper := bytes16[i*2+1]

		// https://gbdev.io/pandocs/Tile_Data.html
		for bit := 7; bit >= 0; bit-- {
			lb := util.Bit(lower, bit)
			ub := util.Bit(upper, bit)
			c := Color((ub << 1) + lb)

			data[i][7-bit] = c
		}
	}
	return &Tile{
		Data: data,
	}
}

// build BG Tiles on current horizontal line accoding to LY Register
func (t *Tile) buildBG(ly byte) []color.RGBA {
	tiles := make([]color.RGBA, SCREEN_WIDTH)

	// build tile row
	for i := 0; i < SCREEN_WIDTH; i++ {

	}

	return tiles
}

type TileMap struct {
	MapData [32][32]Tile
}
