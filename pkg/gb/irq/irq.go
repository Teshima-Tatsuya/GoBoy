package irq

import "github.com/Teshima-Tatsuya/GoBoy/pkg/types"

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

func (i *IRQ) Enable() {
	i.IME = true
	i.SetIF(0x60)
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
