package io

import "github.com/Teshima-Tatsuya/GoBoy/pkg/types"

// offset lower 2bit
const (
	PadAddr types.Addr = 0x00
	SBAddr  types.Addr = 0x01
	SCAddr  types.Addr = 0x02
)

type IO struct {
	pad    *Pad
	serial *Serial
	timer  *Timer
	irq    *IRQ
	buf    []byte
}

func New(size int) *IO {
	buf := make([]byte, size)

	return &IO{
		buf: buf,
	}

}

func (r *IO) Read(addr types.Addr) byte {
	return r.buf[addr]
}

func (r *IO) Write(addr types.Addr, value byte) {
	r.buf[addr] = value
}
