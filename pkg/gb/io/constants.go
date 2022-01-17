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
	LCDCAddr            = 0x40
	LCDSAddr            = 0x41
	SCYAddr             = 0x42
	SCXAddr             = 0x43
	LYAddr              = 0x44
	LYCAddr             = 0x45
	DMAAddr             = 0x46
	BGPAddr             = 0x47
	OBP0Addr            = 0x48
	OBP1Addr            = 0x49
	WYAddr              = 0x4A
	WXAddr              = 0x4B
	IEAddr              = 0xFF
)
