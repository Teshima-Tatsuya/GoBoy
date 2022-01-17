package io

import (
	"fmt"

	"github.com/Teshima-Tatsuya/GoBoy/pkg/types"
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
		return s.SC
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
		s.SC = value
	default:
		msg := fmt.Sprintf("Sereal doesn't support addr 0x%02X", addr)
		panic(msg)
	}
}