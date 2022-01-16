package cartridge

import (
	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/memory"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/types"
)

type MBC0 struct {
	ROM *memory.ROM
}

func NewMBC0(romData []byte) *MBC0 {
	return &MBC0{
		ROM: memory.NewROM(romData),
	}
}

func (m *MBC0) Read(addr types.Addr) byte {
	return m.ROM.Buf[addr]
}

// nop
func (m *MBC0) Write(addr types.Addr, value byte) {
}

// nop
func (m *MBC0) SwitchROMBank(bank uint16) {
}

// nop
func (m *MBC0) SwitchRAMBank(bank uint8) {
}
