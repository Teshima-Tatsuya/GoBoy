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
		{name: "all1", args: args{0b11111111}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lcdc := NewLCDC(tt.args.value)
			assert.Equal(t, true, lcdc.LCDPPUEnable())
			assert.Equal(t, WindowTileMapArea(WindowTileMapArea1), lcdc.WinTileMapArea())
			assert.Equal(t, true, lcdc.WindowEnable())
			assert.Equal(t, BGWindowTileDataArea(BGWindowTileDataArea1), lcdc.BGWinTileDataArea())
			assert.Equal(t, BGTileMapArea(BGTileMapArea1), lcdc.BGTileMapArea())
			assert.Equal(t, uint8(1), lcdc.OBJSize())
			assert.Equal(t, true, lcdc.OBJEnable())
			assert.Equal(t, true, lcdc.BGWinEnable())
		})
	}
}
