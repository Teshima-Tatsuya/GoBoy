package io

import (
	"github.com/Teshima-Tatsuya/GoBoy/pkg/types"
)

type IO struct {
	pad    *Pad
	serial *Serial
	Timer  *Timer
	IRQ    *IRQ
	buf    []byte
}

func NewIO(pad *Pad, serial *Serial, timer *Timer, irq *IRQ, size int) *IO {
	buf := make([]byte, size)

	return &IO{
		pad:    pad,
		serial: serial,
		Timer:  timer,
		IRQ:    irq,
		buf:    buf,
	}

}

func (r *IO) Read(addr types.Addr) byte {
	switch {
	case addr == PadAddr:
		return r.pad.Read(addr)
	case DIVAddr <= addr && addr <= TACAddr:
		return r.Timer.Read(addr)
	case addr == IFAddr:
		return r.IRQ.Read(addr)
	case SBAddr <= addr && addr <= SCAddr:
		return r.serial.Read(addr)
	case addr == 0x4D:
		return 0
	case addr == IEAddr:
		return r.IRQ.Read(addr)
	default:
		// debug.Fatal("Unsuported addr for IO Read 0x%04X", addr)
	}
	return 0
}

func (r *IO) Write(addr types.Addr, value byte) {
	switch {
	case addr == PadAddr:
		r.pad.Write(addr, value)
	case SBAddr <= addr && addr <= SCAddr:
		r.serial.Write(addr, value)
	case DIVAddr <= addr && addr <= TACAddr:
		r.Timer.Write(addr, value)
	case addr == IFAddr:
		r.IRQ.Write(addr, value)
	case addr == 0x4D:
		// TODO
	case addr == IEAddr:
		r.IRQ.Write(addr, value)
	default:
		// debug.Fatal("Unsuported addr for IO Write 0x%04X", addr)
	}
}
