package pad

import (
	"github.com/Teshima-Tatsuya/GoBoy/pkg/types"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/util"
)

const (
	PadAddr types.Addr = 0x00
)

type Button byte

type Pad struct {
	p1    byte
	state Button
}

const (
	// A is the A button on the GameBoy.
	A Button = 0x01
	// B is the B button on the GameBoy.
	B Button = 0x02
	// Select is the select button on the GameBoy.
	Select Button = 0x04
	// Start is the start button on the GameBoy.
	Start Button = 0x08
	// Right is the right pad direction on the GameBoy.
	Right Button = 0x10
	// Left is the left pad direction on the GameBoy.
	Left Button = 0x20
	// Up is the up pad direction on the GameBoy.
	Up Button = 0x40
	// Down is the down pad direction on the GameBoy.
	Down Button = 0x80
)

func NewPad() *Pad {
	return &Pad{
		p1:    0xCF, // all buttuns are not pressed
		state: 0x00,
	}
}

func (p *Pad) Read(addr types.Addr) byte {
	if p.buttonPressed() {
		return p.p1 & ^byte(p.state&0x0F) | 0xC0
	}

	if p.directionPressed() {
		return p.p1 & ^byte(p.state>>4) | 0xC0
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

func (pad *Pad) Press(button Button) {
	pad.state |= button
}

func (pad *Pad) Release(button Button) {
	pad.state &= ^button
}
