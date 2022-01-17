package io

import "github.com/Teshima-Tatsuya/GoBoy/pkg/types"

type Pad struct {
	Data byte
}

func NewPad() *Pad {
	return &Pad{
		Data: 0x3F, // all buttuns are not pressed
	}
}

func (p *Pad) Read(addr types.Addr) byte {
	return p.Data
}

func (p *Pad) Write(addr types.Addr, value byte) {
	p.Data = value
}
