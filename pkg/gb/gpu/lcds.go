package gpu

import "github.com/Teshima-Tatsuya/GoBoy/pkg/util"

/*
Addresss is $FF41
 LCDC is [Liquid Crystal Display Status]
 @see https://gbdev.io/pandocs/STAT.html#lcd-status-register
*/
type LCDS struct {
	Data byte
}

func NewLCDS(data byte) *LCDS {
	return &LCDS{
		Data: data,
	}
}

// Bit 6
// LY STAT Interrupt source
func (l *LCDS) LYC() bool {
	return util.Bit(l.Data, 6) == 1
}

// Bit 5
// OAM STAT Interrupt source
func (l *LCDS) Mode2() bool {
	return util.Bit(l.Data, 5) == 1
}

// Bit 4
// OAM STAT Interrupt source
func (l *LCDS) Mode1() bool {
	return util.Bit(l.Data, 4) == 1
}

// Bit 3
// OAM STAT Interrupt source
func (l *LCDS) Mode0() bool {
	return util.Bit(l.Data, 3) == 1
}

// Bit 2
// OAM STAT Interrupt source
func (l *LCDS) LYCLY() bool {
	return util.Bit(l.Data, 2) == 1
}

// Bit 1-0
// 0: HBlank
// 1: VBlank
// 2: Searching OAM
// 3: Transferring Data to LCD Controller
func (l *LCDS) Mode() Mode {
	return Mode(l.Data & 0x03)
}
