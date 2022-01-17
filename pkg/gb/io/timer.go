package io

import (
	"fmt"

	"github.com/Teshima-Tatsuya/GoBoy/pkg/types"
)

type Timer struct {
	counter int16
	DIV     byte
	TIMA    byte
	TMA     byte
	TAC     byte
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

func (t *Timer) Read(addr types.Addr) byte {
	switch addr {
	case DIVAddr:
		return byte(t.DIV)
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

func (t *Timer) Write(addr types.Addr, v byte) byte {
	switch addr {
	case DIVAddr:
		return byte(t.DIV)
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
