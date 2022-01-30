package gpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTile(t *testing.T) {
	t.Run("Tile", func(t *testing.T) {
		//https://www.huderlem.com/demos/gameboy2bpp.html
		bytes := []byte{0xFF, 0x00, 0x7E, 0xFF, 0x85, 0x81, 0x89, 0x83, 0x93, 0x85, 0xA5, 0x8B, 0xC9, 0x97, 0x7E, 0xFF}
		tile := NewTile(bytes)

		colors := [8][8]Color{
			{LightGray, LightGray, LightGray, LightGray, LightGray, LightGray, LightGray, LightGray},
			{DarkGray, Black, Black, Black, Black, Black, Black, DarkGray},
			{Black, White, White, White, White, LightGray, White, Black},
			{Black, White, White, White, LightGray, White, DarkGray, Black},
			{Black, White, White, LightGray, White, DarkGray, LightGray, Black},
			{Black, White, LightGray, White, DarkGray, LightGray, DarkGray, Black},
			{Black, LightGray, White, DarkGray, LightGray, DarkGray, DarkGray, Black},
			{DarkGray, Black, Black, Black, Black, Black, Black, DarkGray},
		}

		assert.Equal(t, colors, tile.Data)
	})
}
