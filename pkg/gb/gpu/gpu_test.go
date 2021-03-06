package gpu

import (
	"testing"

	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/interrupt"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/types"
	"github.com/Teshima-Tatsuya/GoBoy/test/mock"
	"github.com/stretchr/testify/assert"
)

func setupGPU1() *GPU {
	g := New()
	b := mock.NewMockBus()
	bytes1 := []byte{0xFF, 0x00, 0x7E, 0xFF, 0x85, 0x81, 0x89, 0x83, 0x93, 0x85, 0xA5, 0x8B, 0xC9, 0x97, 0x7E, 0xFF}
	bytes2 := []byte{0x00, 0x00, 0x7E, 0xFF, 0x85, 0x81, 0x89, 0x83, 0x93, 0x85, 0xA5, 0x8B, 0xC9, 0x97, 0x7E, 0xFF}

	for tile := 0; tile < 255; tile++ {
		baseAddr := types.Addr(0x8800) + types.Addr(tile*16)
		if tile%2 == 0 {
			for i := 0; i < len(bytes1); i++ {
				b.WriteByte(baseAddr+types.Addr(i), bytes1[i])
			}
		} else {
			for i := 0; i < len(bytes2); i++ {
				b.WriteByte(baseAddr+types.Addr(i), bytes2[i])
			}
		}
	}
	g.Init(b, interrupt.NewIRQ().Request)

	return g
}

func setupGPU2() *GPU {
	g := New()
	b := mock.NewMockBus()
	bytes1 := []byte{0xFF, 0x00, 0x7E, 0xFF, 0x85, 0x81, 0x89, 0x83, 0x93, 0x85, 0xA5, 0x8B, 0xC9, 0x97, 0x7E, 0xFF}
	bytes2 := []byte{0x00, 0x00, 0x7E, 0xFF, 0x85, 0x81, 0x89, 0x83, 0x93, 0x85, 0xA5, 0x8B, 0xC9, 0x97, 0x7E, 0xFF}

	for tile := 0; tile < 255; tile++ {
		baseAddr := types.Addr(0x8000) + types.Addr(tile*16)
		if tile%2 == 0 {
			for i := 0; i < len(bytes2); i++ {
				b.WriteByte(baseAddr+types.Addr(i), bytes2[i])
			}
		} else {
			for i := 0; i < len(bytes1); i++ {
				b.WriteByte(baseAddr+types.Addr(i), bytes1[i])
			}
		}
	}
	g.Init(b, io.NewIRQ().Request)

	return g
}
func TestGPU_loadTile(t *testing.T) {
	t.Run("load tile", func(t *testing.T) {

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
			g := setupGPU1()
			g.LCDC.Data = 0x00 // tile data starts 0x8800
			g.loadTile()

			for i := 0; i < 255; i++ {
				if i <= 127 {
					if i%2 == 0 {
						assert.Equal(t, colors1, g.tiles[0][i].Data)
					} else {
						assert.Equal(t, colors2, g.tiles[0][i].Data)
					}
				} else {
					if i%2 == 0 {
						assert.Equal(t, colors1, g.tiles[1][i].Data)
					} else {
						assert.Equal(t, colors2, g.tiles[1][i].Data)
					}

				}
			}
		})
		t.Run("tile data 1", func(t *testing.T) {
			g := setupGPU2()
			g.LCDC.Data = 0x10 // tile data starts 0x8000
			g.loadTile()

			for i := 0; i < 255; i++ {
				if i%2 == 0 {
					assert.Equal(t, colors2, g.tiles[i].Data)
				} else {
					assert.Equal(t, colors1, g.tiles[i].Data)
				}
			}
		})
	})
}
