package io

import (
	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/gpu"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/types"
)

type IO struct {
	pad    *Pad
	serial *Serial
	timer  *Timer
	IRQ    *IRQ
	gpu    *gpu.GPU
	buf    []byte
}

func NewIO(pad *Pad, serial *Serial, timer *Timer, irq *IRQ, gpu *gpu.GPU, size int) *IO {
	buf := make([]byte, size)

	return &IO{
		pad:    pad,
		serial: serial,
		timer:  timer,
		IRQ:    irq,
		gpu:    gpu,
		buf:    buf,
	}

}

func (r *IO) Read(addr types.Addr) byte {
	switch {
	case addr == PadAddr:
		return r.pad.Read(addr)
	case DIVAddr <= addr && addr <= TACAddr:
		return r.timer.Read(addr)
	case addr == IFAddr:
		return r.IRQ.Read(addr)
	case SBAddr <= addr && addr <= SCAddr:
		return r.serial.Read(addr)
	case LCDCAddr <= addr && addr <= WXAddr:
		return r.gpu.Read(addr)
	case addr == IEAddr:
		return r.IRQ.Read(addr)
	}
	return r.buf[addr]
}

func (r *IO) Write(addr types.Addr, value byte) {
	switch {
	case addr == PadAddr:
		r.pad.Write(addr, value)
	case SBAddr <= addr && addr <= SCAddr:
		r.serial.Write(addr, value)
	case DIVAddr <= addr && addr <= TACAddr:
		r.timer.Write(addr, value)
	case addr == IFAddr:
		r.IRQ.Write(addr, value)
	case LCDCAddr <= addr && addr <= WXAddr:
		r.gpu.Write(addr, value)
	case addr == IEAddr:
		r.IRQ.Write(addr, value)
	}
	r.buf[addr] = value
}
