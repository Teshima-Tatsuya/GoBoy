package gpu

import (
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
	for y := 0; y < 8; y++ {
		lower := bytes16[y*2]
		upper := bytes16[y*2+1]

		// https://gbdev.io/pandocs/Tile_Data.html
		for x := 7; x >= 0; x-- {
			lb := util.Bit(lower, x)
			ub := util.Bit(upper, x)
			c := Color((ub << 1) + lb)

			data[y][7-x] = c
		}
	}
	return &Tile{
		Data: data,
	}
}

type TileMap struct {
	MapData [32][32]Tile
}
