package irq

import (
	"fmt"

	"github.com/Teshima-Tatsuya/GoBoy/pkg/types"
)

// @see https://gbdev.io/pandocs/Interrupts.html#ffff---ie---interrupt-enable-rw
const (
	VBlankAddr   types.Addr = 0x0040
	LCD_STATAddr types.Addr = 0x0048
	TimerAddr    types.Addr = 0x0050
	SerialAddr   types.Addr = 0x0058
	JoypadAddr   types.Addr = 0x0060
)

const (
	VBlank   byte = 0x01
	LCD_STAT byte = 0x02
	Timer    byte = 0x04
	Serial   byte = 0x08
	Joypad   byte = 0x10
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

// @see https://gbdev.io/pandocs/Interrupts.html#interrupt-priorities
func (i *IRQ) InterruptAddr() types.Addr {
	idx := i.IF & i.IE
	switch {
	case idx&VBlank != 0:
		i.IF &= ^VBlank
		return VBlankAddr
	case idx&LCD_STAT != 0:
		i.IF &= ^LCD_STAT
		return LCD_STATAddr
	case idx&Timer != 0:
		i.IF &= ^Timer
		return TimerAddr
	case idx&Serial != 0:
		i.IF &= ^Serial
		return SerialAddr
	case idx&Joypad != 0:
		i.IF &= ^Joypad
		return JoypadAddr
	default:
		panic(fmt.Sprintf("Non Supported idx 0x%02x", idx))
	}
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
