package gpu

import "github.com/Teshima-Tatsuya/GoBoy/pkg/types"

// Object Attribute Memory
// FE00-FE9F
type OAM struct {
	Objs [40]*Obj
	Buf  [0xA0]byte
}

type Obj struct {
	y, x, tile, addr byte
}

func NewOAM() *OAM {
	var objs [40]*Obj

	for i := 0; i < 40; i++ {
		objs[i] = &Obj{}
	}
	return &OAM{
		Objs: objs,
	}
}

// WIP
// @see https://gbdev.io/pandocs/OAM.html#writing-data-to-oam
func (o *OAM) Write(addr types.Addr, data byte) {
	o.Buf[addr] = data
}

func (o *OAM) Read(addr types.Addr) byte {
	return o.Buf[addr]
}
