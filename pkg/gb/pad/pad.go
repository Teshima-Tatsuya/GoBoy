package pad

import (
	"github.com/Teshima-Tatsuya/GoBoy/pkg/types"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/util"
)

const (
	PadAddr types.Addr = 0x00
)

type Pad struct {
	p1    byte
	state byte
}

func NewPad() *Pad {
	return &Pad{
		p1:    0xCF, // all buttuns are not pressed
		state: 0x00,
	}
}

func (p *Pad) Read(addr types.Addr) byte {
	if p.buttonPressed() {
		return p.p1 & ^(p.state&0x0F) | 0xC0
	}

	if p.directionPressed() {
		return p.p1 & ^(p.state>>4) | 0xC0
	}

	// all not pressed
	return p.p1 | 0xCF
}

func (p *Pad) Write(addr types.Addr, value byte) {
	// because bit 3-0 is read only
	p.p1 = (p.p1 & 0xCF) | (value & 0x30)
}

// button is start, select, A, B
func (p *Pad) buttonPressed() bool {
	return util.Bit(p.p1, 5) == 0
}

// direction is up, down, left, right
func (p *Pad) directionPressed() bool {
	return util.Bit(p.p1, 4) == 0
}
