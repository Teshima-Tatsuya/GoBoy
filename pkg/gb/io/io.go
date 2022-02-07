package io

import (
	"github.com/Teshima-Tatsuya/GoBoy/pkg/types"
)

type IO struct {
	pad    *Pad
	serial *Serial
	Timer  *Timer
	buf    []byte
}

func NewIO(pad *Pad, serial *Serial, timer *Timer) *IO {
	return &IO{
		pad:    pad,
		serial: serial,
		Timer:  timer,
	}

}

func (r *IO) Read(addr types.Addr) byte {
	switch {
	case addr == PadAddr:
		return r.pad.Read(addr)
	case DIVAddr <= addr && addr <= TACAddr:
		return r.Timer.Read(addr)
	case SBAddr <= addr && addr <= SCAddr:
		return r.serial.Read(addr)
	default:
		// debug.Fatal("Unsuported addr for IO Read 0x%04X", addr)
	}
	return 0xFF
}

func (r *IO) Write(addr types.Addr, value byte) {
	switch {
	case addr == PadAddr:
		r.pad.Write(addr, value)
	case SBAddr <= addr && addr <= SCAddr:
		r.serial.Write(addr, value)
	case DIVAddr <= addr && addr <= TACAddr:
		r.Timer.Write(addr, value)
	default:
		// debug.Fatal("Unsuported addr for IO Write 0x%04X", addr)
	}
}
