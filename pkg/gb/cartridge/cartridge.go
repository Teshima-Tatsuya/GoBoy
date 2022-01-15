package cartridge

import (
	"fmt"

	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/memory"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/types"
)

type CartridgeType byte

// ROM Type
// @see https://gbdev.io/pandocs/The_Cartridge_Header.html#0147---cartridge-type
const (
	NO_MBC = iota
	MBC1
	MBC2
	MBC3
	MBC4
	MBC5
	MBC6
	MBC7
)

// ROM Size
// @see https://gbdev.io/pandocs/The_Cartridge_Header.html#0148---rom-size
const (
	ROM_32KB int = (1 << (iota + 5)) * 1024
	ROM_64KB
	ROM_128KB
	ROM_256KB
	ROM_512KB
	ROM_1024KB
	ROM_2048KB
	ROM_4096KB
	ROM_8192KB
)

// Mode
const (
	SimpleROMBankingMode                 = 0x00
	RAMBankingModeAdvancedROMBankingMode = 0x01
)

// RAM Size
// @see https://gbdev.io/pandocs/The_Cartridge_Header.html#0149---ram-size
const (
	NO_RAM    int = 0 // also UNUSED
	RAM_8KB   int = 8 * 1024
	RAM_32KB  int = 32 * 1024
	RAM_128KB int = 128 * 1024
	RAM_64KB  int = 64 * 1024
)

// @see https://gbdev.io/pandocs/The_Cartridge_Header.html
type Cartridge struct {
	NintendoLogo []byte
	// Game Title
	Title string
	// GameBoy Color Flag
	CGBFlag bool
	// game supports SGB functions
	SGBFlag bool
	Type    byte
	ROM     *memory.ROM
	RAM     *memory.RAM
	ROMSize int
	RAMSize int
	ROMData []byte
	Mode    byte
	Bank    *Bank
}

func New(romData []byte) *Cartridge {
	nintendLogo := romData[0x0104:0x0133]
	title := string(romData[0x0134:0x0143])

	cgbflag := romData[0x0143] == 0x80 || romData[0x0143] == 0xC0
	sgbflag := romData[0x0146] != 0x00
	romSize := getRomSize(romData[0x0148])

	c := &Cartridge{
		NintendoLogo: nintendLogo,
		Title:        string(title),
		CGBFlag:      cgbflag,
		SGBFlag:      sgbflag,
		Type:         getType(romData[0x0147]),
		ROMSize:      romSize,
		RAMSize:      getRamSize(romData[0x0149]),
		ROMData:      romData,
		Mode:         0x00,
		Bank:         NewBank(romSize),
	}

	c.loadRom(romData)

	return c
}

func getRamSize(s byte) int {
	switch s {
	case 0x00, 0x01:
		return NO_RAM
	case 0x02:
		return RAM_8KB
	case 0x03:
		return RAM_32KB
	case 0x04:
		return RAM_128KB
	case 0x05:
		return RAM_64KB
	default:
		errMsg := fmt.Sprintf("Non Supported RAM Type 0x%02X", s)
		panic(errMsg)
	}
}

func getType(t byte) byte {
	switch t {
	case 0x00:
		return NO_MBC
	case 0x01, 0x02, 0x03:
		return MBC1
	case 0x05, 0x06:
		return MBC2
	case 0x0F, 0x10, 0x11, 0x12, 0x13:
		return MBC3
	case 0x15, 0x16, 0x17:
		return MBC4
	case 0x19, 0x1A, 0x1B, 0x1C, 0x1D, 0x1E:
		return MBC5
	case 0x20:
		return MBC6
	case 0x22:
		return MBC7
	default:
		errMsg := fmt.Sprintf("Non Supported type 0x%02X", t)
		panic(errMsg)
	}
}

func getRomSize(s byte) int {
	return ROM_32KB << int(s)
}

func validateCheckSum(romData []byte) bool {
	x := byte(0x00)

	for i := 0x0134; i <= 0x014C; i++ {
		x = x - romData[i] - byte(0x01)
	}

	if x != romData[0x014D] {
		return false
	}

	return true
}

// TODO Read RAM
func (c *Cartridge) ReadByte(addr types.Addr) byte {
	switch {
	case addr < 0x4000:
		// Bank 0
		return c.ROM.Buf[0][addr]
	case 0x4000 <= addr && addr < 0x8000:
		return c.ROM.Buf[c.Bank.Current][addr-0x4000]
	}
	return c.ROMData[addr]
}

func (c *Cartridge) WriteByte(addr types.Addr, value byte) {
	switch {
	// @see https://gbdev.io/pandocs/MBC1.html#2000-3fff---rom-bank-number-write-only
	case 0x2000 <= addr && addr <= 0x3FFF:
		if c.Type == MBC1 {
			c.ROM.Bank = value & 0x1F
		}
	case 0x4000 <= addr && addr <= 0x5FFF:
		if c.Type == MBC1 {
			c.Mode = value
		}
	// @see https://gbdev.io/pandocs/MBC1.html#6000-7fff---banking-mode-select-write-only
	case 0x6000 <= addr && addr <= 0x7FFF:
		if c.Type == MBC1 {
			c.Mode = value
		}
	}
	c.ROMData[addr] = value
}

// load romData to ROM
func (c *Cartridge) loadRom(romData []byte) {
	bankSize := int(c.Bank.Size)
	c.ROM.Buf = make([][]byte, bankSize)

	for bank := 0; bank < bankSize; bank++ {
		c.ROM.Buf[bank] = make([]byte, 0x4000)
		for i := 0; i < 0x4000; i++ {
			c.ROM.Buf[bank][i] = romData[bank*0x4000+i]
		}
	}
}
