package gpu

import "github.com/Teshima-Tatsuya/GoBoy/pkg/util"

/*
Addresss is $FF40
 LCDC is [Liquid Crystal Display Controll]
 Its bits toggle what elements are displayed on the screen, and how.
 @see https://gbdev.io/pandocs/LCDC.html#lcd-control
*/
type LCDC struct {
	Data byte
}

func NewLCDC(data byte) *LCDC {
	return &LCDC{
		Data: data,
	}
}

// LCD and PPU enable
func (l *LCDC) LCDPPUEnable() bool {
	return util.Bit(l.Data, 7) == 1
}

/*
Window tile map area
 0=9800-9BFF, 1=9C00-9FFF
*/
func (l *LCDC) WinTileMapArea() WindowTypeMapArea {
	area := util.Bit(l.Data, 6)
	if area == 0 {
		return WindowTileMapArea0
	}

	return WindowTileMapArea1
}

// Window enable
func (l *LCDC) WindowEnable() bool {
	return util.Bit(l.Data, 5) == 1
}

/*
BG and Window tile data area
 0=8800-97FF, 1=8000-8FFF
*/
func (l *LCDC) BGWinTileDataArea() BGWindowTileDataArea {
	area := util.Bit(l.Data, 4)
	if area == 0 {
		return BGWindowTileDataArea1
	}

	return BGWindowTileDataArea0
}

/*
BG tile map area
 0=9800-9BFF, 1=9C00-9FFF
*/
func (l *LCDC) BGTileMapArea() BGTileMapArea {
	area := util.Bit(l.Data, 3)
	if area == 0 {
		return BGTileMapArea0
	}

	return BGTileMapArea1
}

/*
OBJ size
 0=8x8, 1=8x16
*/
func (l *LCDC) OBJSize() uint8 {
	return util.Bit(l.Data, 2)
}

// OBJ enable
func (l *LCDC) OBJEnable() bool {
	return util.Bit(l.Data, 1) == 1
}

// BG and Window enable/priority
func (l *LCDC) BGWinEnable() bool {
	return util.Bit(l.Data, 0) == 1
}
