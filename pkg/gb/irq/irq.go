package irq

import "github.com/Teshima-Tatsuya/GoBoy/pkg/types"

// @see https://gbdev.io/pandocs/Interrupts.html#ffff---ie---interrupt-enable-rw
const (
	VBlank   types.Addr = 0x0040
	LCD_STAT types.Addr = 0x0048
	Timer    types.Addr = 0x0050
	Serial   types.Addr = 0x0058
	Joypad   types.Addr = 0x0060
)

type IRQ struct {
	IF  byte
	IE  byte
	IME bool
}

func New() *IRQ {
	return &IRQ{
		IF:  0x00,
		IE:  0x00,
		IME: false,
	}
}

func (i *IRQ) SetIF(v byte) {
	i.Write(0xFF, v)
}

func (i *IRQ) SetIE(v byte) {
	i.Write(0x0F, v)
}

// if any irq is enable
func (i *IRQ) Has() bool {
	return i.IF&i.IE != 0x00
}

func (i *IRQ) Enable() {
	i.IME = true
	i.SetIF(0x60)
}

func (i *IRQ) Enabled() bool {
	return i.IME
}

func (i *IRQ) Disable() {
	i.IME = false
	i.SetIF(0x00)
}

func (i *IRQ) Write(addr types.Addr, v byte) {
	switch addr {
	case 0xFF:
		i.IF = v
	case 0x0F:
		i.IE = v
	default:
		panic("Can't Write addr")
	}
}

func (i *IRQ) Read(addr types.Addr) byte {
	switch addr {
	case 0xFF:
		return i.IF
	case 0x0F:
		return i.IE
	default:
		panic("Can't Write addr")
	}
}
