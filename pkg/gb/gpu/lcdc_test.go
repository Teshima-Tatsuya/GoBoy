package gpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLCDC(t *testing.T) {

	type args struct {
		value byte
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "LCDC"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Run("when all 1", func(t *testing.T) {
				lcdc := NewLCDC(0xFF)
				assert.Equal(t, true, lcdc.LCDPPUEnable())
				assert.Equal(t, WindowTileMapArea(WindowTileMapArea1), lcdc.WinTileMapArea())
				assert.Equal(t, true, lcdc.WindowEnable())
				assert.Equal(t, BGWindowTileDataArea(BGWindowTileDataArea1), lcdc.BGWinTileDataArea())
				assert.Equal(t, BGTileMapArea(BGTileMapArea1), lcdc.BGTileMapArea())
				assert.Equal(t, uint8(1), lcdc.OBJSize())
				assert.Equal(t, true, lcdc.OBJEnable())
				assert.Equal(t, true, lcdc.BGWinEnable())
			})
			t.Run("when all 0", func(t *testing.T) {
				lcdc := NewLCDC(0x00)
				assert.Equal(t, false, lcdc.LCDPPUEnable())
				assert.Equal(t, WindowTileMapArea(WindowTileMapArea0), lcdc.WinTileMapArea())
				assert.Equal(t, false, lcdc.WindowEnable())
				assert.Equal(t, BGWindowTileDataArea(BGWindowTileDataArea0), lcdc.BGWinTileDataArea())
				assert.Equal(t, BGTileMapArea(BGTileMapArea0), lcdc.BGTileMapArea())
				assert.Equal(t, uint8(0), lcdc.OBJSize())
				assert.Equal(t, false, lcdc.OBJEnable())
				assert.Equal(t, false, lcdc.BGWinEnable())
			})
		})
	}
}
