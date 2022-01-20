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

// LY STAT Interrupt source
func (l *LCDS) LYC() bool {
	return util.Bit(l.Data, 6) == 1
}

func (l *LCDS) Mode() Mode {
	return Mode(l.Data & 0x03)
}
