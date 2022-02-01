package cartridge

import (
	"fmt"

	"github.com/Teshima-Tatsuya/GoBoy/pkg/debug"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/memory"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/types"
)

// Mode
const (
	SimpleROMBankingMode                 = 0x00
	RAMBankingModeAdvancedROMBankingMode = 0x01
)

type MBC1 struct {
	ROM *memory.ROM
	RAM *memory.RAM
	// max 2MB
	romBank uint8
	// max 32KiB
	ramBank   uint8
	ramEnable bool
	mode      uint8
}

func NewMBC1(romData []byte, ramSize int) *MBC1 {
	m := &MBC1{
		ROM: memory.NewROM(romData),
		// The ROM Bank Number defaults to 01 at power-on.
		romBank: 1,
		ramBank: 0,
		mode:    SimpleROMBankingMode,
	}

	debug.Info("Cartridge RAM Size = %d", ramSize)
	if ramSize > 0 {
		m.RAM = memory.NewRAM(ramSize)
		m.ramBank = 0
		m.ramEnable = false
	}

	return m
}

func (m *MBC1) Read(addr types.Addr) byte {
	// @see https://gbdev.io/pandocs/MBC1.html
	// Implement Read address range
	switch {
	case addr < 0x4000:
		return m.ROM.Read(uint32(addr))
	case 0x4000 <= addr && addr < 0x8000:
		return m.ROM.Read(uint32(m.romBank-1)*0x4000 + uint32(addr))
	case 0xA000 <= addr && addr < 0xC000:
		addr = types.Addr(uint16(addr) + uint16(m.ramBank)*0x2000 - 0xA000)
		debug.Info("0x%04X", addr)
		return m.RAM.Read(addr)
	default:
		msg := fmt.Sprintf("Non Supported addr 0x%4X for Read MBC1", addr)
		panic(msg)
	}
}

func (m *MBC1) Write(addr types.Addr, value byte) {
	// @see https://gbdev.io/pandocs/MBC1.html
	// Implement Write address range
	switch {
	case addr < 0x2000:
		if value == 0x00 {
			m.ramEnable = false
		} else if value == 0x0A {
			m.ramEnable = true
		}
	case 0x2000 <= addr && addr < 0x4000:
		m.SwitchROMBank(uint16(value & 0x1F))
	case 0x4000 <= addr && addr < 0x6000:
		if m.mode == SimpleROMBankingMode {
			m.SwitchHiROMBank(value)
		} else if m.mode == RAMBankingModeAdvancedROMBankingMode {
			// lower 2bit
			debug.Info("Switch RAM Bank %d", value&0x03)
			m.SwitchRAMBank(value & 0x03)
		}
	case 0x6000 <= addr && addr < 0x8000:
		m.mode = value
	case 0xA000 <= addr && addr < 0xC000:
		if m.ramEnable {
			addr = types.Addr(uint16(addr) + uint16(m.ramBank)*0x2000 - 0xA000)
			m.RAM.Write(addr, value)
		}
	}
}

func (m *MBC1) SwitchROMBank(bank uint16) {
	if bank == 0x00 || bank == 0x20 || bank == 0x40 || bank == 0x60 {
		bank++
	}

	m.romBank = uint8(bank)
}

func (m *MBC1) SwitchHiROMBank(value byte) {
	// clear Hi bit
	m.romBank &= 0x1F

	// clear Low bit
	value &= 0xE0

	m.romBank |= value
}

func (m *MBC1) SwitchRAMBank(bank uint8) {
	m.ramBank = bank
}
