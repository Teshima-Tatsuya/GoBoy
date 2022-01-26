package apu

import "github.com/Teshima-Tatsuya/GoBoy/pkg/types"

type APU struct {
	NR10 byte
	NR11 byte
	NR12 byte
	NR13 byte
	NR14 byte
	NR21 byte
	NR22 byte
	NR23 byte
	NR24 byte
	NR30 byte
	NR31 byte
	NR32 byte
	NR33 byte
	NR34 byte
	NR41 byte
	NR42 byte
	NR43 byte
	NR44 byte
	NR50 byte
	NR51 byte
	NR52 byte
}

func NewAPU() *APU {
	return &APU{}
}

func (a *APU) Read(addr types.Addr) byte {
	switch {
	case addr == NR10Addr:
		return a.NR10
	case addr == NR11Addr:
		return a.NR11
	case addr == NR12Addr:
		return a.NR12
	case addr == NR13Addr:
		return a.NR13
	case addr == NR14Addr:
		return a.NR14
	case addr == NR21Addr:
		return a.NR21
	case addr == NR22Addr:
		return a.NR22
	case addr == NR23Addr:
		return a.NR23
	case addr == NR24Addr:
		return a.NR24
	case addr == NR30Addr:
		return a.NR30
	case addr == NR31Addr:
		return a.NR31
	case addr == NR32Addr:
		return a.NR32
	case addr == NR33Addr:
		return a.NR33
	case addr == NR34Addr:
		return a.NR34
	case addr == NR41Addr:
		return a.NR41
	case addr == NR42Addr:
		return a.NR42
	case addr == NR43Addr:
		return a.NR43
	case addr == NR44Addr:
		return a.NR44
	case addr == NR50Addr:
		return a.NR50
	case addr == NR51Addr:
		return a.NR51
	case addr == NR52Addr:
		return a.NR52
	case 0x30 <= addr && addr <= 0x3F:
		return 0
	default:
		panic("Unsupported addr for APU Read")
	}
}

func (a *APU) Write(addr types.Addr, value byte) {
	switch {
	case addr == NR10Addr:
		a.NR10 = value
	case addr == NR11Addr:
		a.NR11 = value
	case addr == NR12Addr:
		a.NR12 = value
	case addr == NR13Addr:
		a.NR13 = value
	case addr == NR14Addr:
		a.NR14 = value
	case addr == NR21Addr:
		a.NR21 = value
	case addr == NR22Addr:
		a.NR22 = value
	case addr == NR23Addr:
		a.NR23 = value
	case addr == NR24Addr:
		a.NR24 = value
	case addr == NR30Addr:
		a.NR30 = value
	case addr == NR31Addr:
		a.NR31 = value
	case addr == NR32Addr:
		a.NR32 = value
	case addr == NR33Addr:
		a.NR33 = value
	case addr == NR34Addr:
		a.NR34 = value
	case addr == NR41Addr:
		a.NR41 = value
	case addr == NR42Addr:
		a.NR42 = value
	case addr == NR43Addr:
		a.NR43 = value
	case addr == NR44Addr:
		a.NR44 = value
	case addr == NR50Addr:
		a.NR50 = value
	case addr == NR51Addr:
		a.NR51 = value
	case addr == NR52Addr:
		a.NR52 = value
	case 0x30 <= addr && addr <= 0x3F:
	default:
		panic("Unsupported addr for APU Read")
	}
}
