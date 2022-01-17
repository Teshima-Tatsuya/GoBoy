package io

import (
	"fmt"

	"github.com/Teshima-Tatsuya/GoBoy/pkg/types"
)

type Timer struct {
	counter uint16
	DIV     byte
	TIMA    byte
	TMA     byte
	TAC     byte
}

func (t *Timer) Tick(cycle int) bool {
	r := false
	for i := 0; i < cycle; i++ {
		t.counter += 4

		// TODO double speed for GBC
		if t.counter%16384 == 0 {
			t.DIV++
		}

		if !t.started() {
			continue
		}

		if uint32(t.counter)%t.getFreq() == 0 {
			t.TIMA++

			if t.TIMA == 0 {
				t.TIMA = t.TMA
				r = true
			}
		}

	}

	return r
}

func NewTimer() *Timer {
	return &Timer{
		counter: 0,
		DIV:     0x00,
		TIMA:    0x00,
		TMA:     0x00,
		TAC:     0x00,
	}
}

func (t *Timer) getFreq() uint32 {
	switch t.TAC & 0x03 {
	case 0x00:
		return 4096
	case 0x01:
		return 262144
	case 0x10:
		return 65536
	case 0x11:
		return 16384
	default:
		panic("Illegal TAC")
	}
}

func (t *Timer) Read(addr types.Addr) byte {
	switch addr {
	case DIVAddr:
		return t.DIV
	case TIMAAddr:
		return t.TIMA
	case TMAAddr:
		return t.TMA
	case TACAddr:
		return t.TAC
	default:
		panic(fmt.Sprintf("Non Supported addr 0x%04X", addr))
	}
}

func (t *Timer) Write(addr types.Addr, v byte) {
	switch addr {
	case DIVAddr:
		t.DIV = 0
		t.counter = 0
	case TIMAAddr:
		t.TIMA = v
	case TMAAddr:
		t.TMA = v
	case TACAddr:
		t.TAC = v
	default:
		panic(fmt.Sprintf("Non Supported addr 0x%04X", addr))
	}
}

func (t *Timer) started() bool {
	return t.TAC&0x04 == 0x04
}
