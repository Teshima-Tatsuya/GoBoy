package io

import "github.com/Teshima-Tatsuya/GoBoy/pkg/types"

// offset is FF00
const (
	PadAddr  types.Addr = 0x00
	SBAddr              = 0x01
	SCAddr              = 0x02
	DIVAddr             = 0x04
	TIMAAddr            = 0x05
	TMAAddr             = 0x06
	TACAddr             = 0x07
	IFAddr              = 0x0F
	NR10Addr            = 0x10
	NR52Addr            = 0x26
	KEY1Addr            = 0x4D
	IEAddr              = 0xFF
)
