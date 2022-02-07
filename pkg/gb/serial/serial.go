package serial

import (
	"fmt"

	"github.com/Teshima-Tatsuya/GoBoy/pkg/types"
)

const (
	SBAddr = 0x01
	SCAddr = 0x02
)

type Serial struct {
	SB byte
	SC byte
}

func NewSerial() *Serial {
	return &Serial{
		SB: 0x00,
		SC: 0x00,
	}
}

func (s *Serial) Read(addr types.Addr) byte {
	switch {
	case addr == SBAddr:
		return s.SB
	case addr == SCAddr:
		return s.SC | 0x7E
	default:
		msg := fmt.Sprintf("Sereal doesn't support addr 0x%02X", addr)
		panic(msg)
	}
}

func (s *Serial) Write(addr types.Addr, value byte) {
	switch {
	case addr == SBAddr:
		s.SB = value
	case addr == SCAddr:
		s.SC = value & 0x83
	default:
		msg := fmt.Sprintf("Sereal doesn't support addr 0x%02X", addr)
		panic(msg)
	}
}
