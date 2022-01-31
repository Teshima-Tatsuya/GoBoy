package gpu

import (
	"testing"

	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/io"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/types"
	"github.com/Teshima-Tatsuya/GoBoy/test/mock"
	"github.com/stretchr/testify/assert"
)

func setupGPU() *GPU {
	g := New()
	b := mock.NewMockBus()
	bytes1 := []byte{0xFF, 0x00, 0x7E, 0xFF, 0x85, 0x81, 0x89, 0x83, 0x93, 0x85, 0xA5, 0x8B, 0xC9, 0x97, 0x7E, 0xFF}
	bytes2 := []byte{0x00, 0x00, 0x7E, 0xFF, 0x85, 0x81, 0x89, 0x83, 0x93, 0x85, 0xA5, 0x8B, 0xC9, 0x97, 0x7E, 0xFF}

	for tile := 0; tile < 384; tile++ {
		addr := types.Addr(0x8800) + types.Addr(tile*16)
		if tile%2 == 0 {
			for i := 0; i < len(bytes1); i++ {
				b.WriteByte(addr, bytes1[i])
			}
		} else {
			for i := 0; i < len(bytes2); i++ {
				b.WriteByte(addr, bytes2[i])
			}
		}
	}
	g.Init(b, io.NewIRQ().Request)

	return g
}

func TestGPU_loadTile(t *testing.T) {
	t.Run("load tile", func(t *testing.T) {
		g := setupGPU()

		colors1 := [8][8]Color{
			{LightGray, LightGray, LightGray, LightGray, LightGray, LightGray, LightGray, LightGray},
			{DarkGray, Black, Black, Black, Black, Black, Black, DarkGray},
			{Black, White, White, White, White, LightGray, White, Black},
			{Black, White, White, White, LightGray, White, DarkGray, Black},
			{Black, White, White, LightGray, White, DarkGray, LightGray, Black},
			{Black, White, LightGray, White, DarkGray, LightGray, DarkGray, Black},
			{Black, LightGray, White, DarkGray, LightGray, DarkGray, DarkGray, Black},
			{DarkGray, Black, Black, Black, Black, Black, Black, DarkGray},
		}
		colors2 := [8][8]Color{
			{White, White, White, White, White, White, White, White},
			{DarkGray, Black, Black, Black, Black, Black, Black, DarkGray},
			{Black, White, White, White, White, LightGray, White, Black},
			{Black, White, White, White, LightGray, White, DarkGray, Black},
			{Black, White, White, LightGray, White, DarkGray, LightGray, Black},
			{Black, White, LightGray, White, DarkGray, LightGray, DarkGray, Black},
			{Black, LightGray, White, DarkGray, LightGray, DarkGray, DarkGray, Black},
			{DarkGray, Black, Black, Black, Black, Black, Black, DarkGray},
		}

		t.Run("tile data 0", func(t *testing.T) {
			g.LCDC.Data = 0x00 // tile data starts 0x8800
			g.loadTile()

			for i := 0; i < 384; i++ {
				if i%2 == 0 {
					assert.Equal(t, colors1, g.tiles[i])
				} else {
					assert.Equal(t, colors2, g.tiles[i])
				}
			}
		})
	})
}
