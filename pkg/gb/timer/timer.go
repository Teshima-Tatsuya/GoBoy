package timer

import (
	"fmt"

	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/interrupt"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/types"
)

const (
	DIVAddr  = 0x04
	TIMAAddr = 0x05
	TMAAddr  = 0x06
	TACAddr  = 0x07
)

type Timer struct {
	requestIRQ func(byte)
	counter    uint16
	DIV        byte
	TIMA       byte
	TMA        byte
	TAC        byte
}

func NewTimer() *Timer {
	return &Timer{
		counter: 0,
		DIV:     0x19,
		TIMA:    0x00,
		TMA:     0x00,
		TAC:     0x00,
	}
}

func (t *Timer) SetRequestIRQ(request func(byte)) {
	t.requestIRQ = request
}

func (t *Timer) Tick(cycle uint) {
	for i := uint(0); i < cycle; i++ {
		t.counter += 4

		// TODO double speed for GBC
		if t.counter%256 == 0 {
			t.DIV++
		}

		if !t.started() {
			continue
		}

		if t.TIMA == 0 {
			t.TIMA = t.TMA
			t.requestIRQ(interrupt.TimerFlag)
		}
		if t.counter%(1<<(t.getFreq()+1)) == 0 {
			t.TIMA++

		}

	}
}

func (t *Timer) getFreq() uint16 {
	switch t.TAC & 0x03 {
	case 0:
		return 9
	case 1:
		return 3
	case 2:
		return 5
	case 3:
		return 7
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
		return t.TAC | 0xF8
	default:
		panic(fmt.Sprintf("Non Supported addr 0x%04X", addr))
	}
}

func (t *Timer) Write(addr types.Addr, v byte) {
	switch addr {
	case DIVAddr:
		t.DIV = 0
		if t.counter>>t.getFreq()&0x01 == 1 {
			t.TIMA++
		}
		t.counter = 0
	case TIMAAddr:
		t.TIMA = v
	case TMAAddr:
		t.TMA = v
	case TACAddr:
		t.TAC = v & 0x07
	default:
		panic(fmt.Sprintf("Non Supported addr 0x%04X", addr))
	}
}

func (t *Timer) started() bool {
	return t.TAC&0x04 == 0x04
}
