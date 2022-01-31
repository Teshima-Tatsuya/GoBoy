package gpu

import "github.com/Teshima-Tatsuya/GoBoy/pkg/types"

const SPRITE_NUM = 40
const CyclePerLine = 456
const (
	SCREEN_WIDTH  = 160
	SCREEN_HEIGHT = 144
)

// Offset is FF00
const (
	LCDCAddr types.Addr = 0x40
	LCDSAddr            = 0x41
	SCYAddr             = 0x42
	SCXAddr             = 0x43
	LYAddr              = 0x44
	LYCAddr             = 0x45
	DMAAddr             = 0x46
	BGPAddr             = 0x47
	OBP0Addr            = 0x48
	OBP1Addr            = 0x49
	WYAddr              = 0x4A
	WXAddr              = 0x4B
	BCPSAddr            = 0x68
	BCPDAddr            = 0x69
	OCPSAddr            = 0x6A
	OCPDAddr            = 0x6B
)

type Mode byte

const (
	Mode_HBlank Mode = iota
	Mode_VBlank
	Mode_SearchingOAM
	Mode_TransferringData
)

type WindowTileMapArea types.Addr

const (
	WindowTileMapArea0 WindowTileMapArea = 0x9800
	WindowTileMapArea1                   = 0x9C00
)

type BGWindowTileDataArea types.Addr

const (
	BGWindowTileDataArea0 BGWindowTileDataArea = 0x8800
	BGWindowTileDataArea1                      = 0x8000
)

type BGTileMapArea types.Addr

const (
	BGTileMapArea0 BGTileMapArea = 0x9800
	BGTileMapArea1               = 0x9C00
)
