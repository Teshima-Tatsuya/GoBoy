package gpu

import "github.com/Teshima-Tatsuya/GoBoy/pkg/types"

const SPRITE_NUM = 40
const (
	SCREEN_WIDTH  = 160
	SCREEN_HEIGHT = 144
)

type Mode byte

const (
	Mode_HBlank Mode = iota
	Mode_VBlank
	Mode_SearchingOAM
	Mode_TransferringData
)

type WindowTypeMapArea types.Addr

const (
	WindowTileMapArea0 WindowTypeMapArea = 0x9800
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
