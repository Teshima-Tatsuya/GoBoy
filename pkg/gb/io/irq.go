package io

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
	VBlankFlag   byte = 0x01
	LCD_STATFlag byte = 0x02
	TimerFlag    byte = 0x04
	SerialFlag   byte = 0x08
	JoypadFlag   byte = 0x10
)

type IRQ struct {
	IF  byte
	IE  byte
	IME bool
}

func NewIRQ() *IRQ {
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
	case idx&VBlankFlag != 0:
		i.IF &= ^VBlankFlag
		return VBlankAddr
	case idx&LCD_STATFlag != 0:
		i.IF &= ^LCD_STATFlag
		return LCD_STATAddr
	case idx&TimerFlag != 0:
		i.IF &= ^TimerFlag
		return TimerAddr
	case idx&SerialFlag != 0:
		i.IF &= ^SerialFlag
		return SerialAddr
	case idx&JoypadFlag != 0:
		i.IF &= ^JoypadFlag
		return JoypadAddr
	default:
		panic(fmt.Sprintf("Non Supported idx 0x%02x", idx))
	}
}

func (i *IRQ) Enable() {
	i.IME = true
}

func (i *IRQ) Enabled() bool {
	return i.IME
}

func (i *IRQ) Disable() {
	i.IME = false
}

func (i *IRQ) Write(addr types.Addr, v byte) {
	switch addr {
	case IEAddr:
		i.IE = v
	case IFAddr:
		i.IF = v
	default:
		panic("Can't Write addr")
	}
}

func (i *IRQ) Read(addr types.Addr) byte {
	switch addr {
	case IEAddr:
		return i.IE
	case IFAddr:
		return i.IF
	default:
		panic("Can't Write addr")
	}
}
