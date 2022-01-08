package video

import "github.com/Teshima-Tatsuya/GoBoy/pkg/types"

// Object Attribute Memory
// FE00-FE9F
type OAM struct {
	// Sprite’s vertical position on the screen + 16
	Y byte
	// Sprite’s horizontal position on the screen + 8
	X         byte
	TileIndex byte
	AttrFlags byte
	Buf       [0xA0]byte
}

func NewOAM() *OAM {
	return &OAM{}
}

// WIP
// @see https://gbdev.io/pandocs/OAM.html#writing-data-to-oam
func (o *OAM) Write(addr types.Addr, data byte) {
	o.Buf[addr] = data
}

func (o *OAM) Read(addr types.Addr) byte {
	return o.Buf[addr]
}
