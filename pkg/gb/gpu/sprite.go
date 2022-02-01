package gpu

import "github.com/Teshima-Tatsuya/GoBoy/pkg/util"

type Sprite struct {
	y, x, tileIdx, attr byte
}

func NewSprite(bytes4 []byte) *Sprite {
	s := &Sprite{}

	s.y = bytes4[0]
	s.x = bytes4[1]
	s.tileIdx = bytes4[2]
	s.attr = bytes4[3]

	return s
}

func (s *Sprite) TileIdx() byte {
	return s.tileIdx
}

// attr methods

func (s *Sprite) YFlip() bool {
	return util.Bit(s.attr, 6) == 1
}

func (s *Sprite) XFlip() bool {
	return util.Bit(s.attr, 5) == 1
}

// Non CGB Mode Only
// 0=OBP0, 1=OBP1
func (s *Sprite) MBGPalleteNo() byte {
	return util.Bit(s.attr, 4)
}

// CGB Mode Only

// 0=Bank 0, 1=Bank 1
func (s *Sprite) VRAMBank() byte {
	return util.Bit(s.attr, 3)
}

func (s *Sprite) CGBPaletteNo() byte {
	// bit 2-0
	return s.attr & 0x07
}
