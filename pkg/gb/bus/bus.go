package bus

import (
	"fmt"

	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/cartridge"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/ie"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/io"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/ram"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/types"
)

type Bus struct {
	Cart  *cartridge.Cartridge
	VRAM  *ram.RAM
	WRAM  *ram.RAM
	WRAM2 *ram.RAM
	HRAM  *ram.RAM
	IO    *io.IO
	IE    *ie.IE
}

func New(cart *cartridge.Cartridge, vram *ram.RAM, wram *ram.RAM, wram2 *ram.RAM, hram *ram.RAM, io *io.IO, ie *ie.IE) *Bus {
	return &Bus{
		Cart:  cart,
		VRAM:  vram,
		WRAM:  wram,
		WRAM2: wram2,
		HRAM:  hram,
		IO:    io,
		IE:    ie,
	}
}

// TODO: IF, IE
func (b *Bus) ReadByte(addr types.Addr) byte {
	switch {
	case addr >= 0x0000 && addr <= 0x7FFF:
		return b.Cart.ReadByte(addr)
	case addr >= 0x8000 && addr <= 0x9FFF:
		return b.VRAM.Read(addr - 0x7FFF)
	case addr >= 0xC000 && addr <= 0xCFFF:
		return b.WRAM.Read(addr - 0xBFFF)
	case addr >= 0xD000 && addr <= 0xDFFF:
		return b.WRAM2.Read(addr - 0xCFFF)
	case addr >= 0xFF00 && addr <= 0xFF7F:
		return b.IO.Read(addr - 0xFEFF)
	case addr >= 0xFF80 && addr <= 0xFFFE:
		return b.HRAM.Read(addr - 0xFF7F)
	case addr == 0xFFFF:
		return b.IE.Read()
	default:
		panic(fmt.Sprintf("Non Supported Read Addr 0x%4d", addr))
	}
}

func (b *Bus) ReadAddr(addr types.Addr) types.Addr {
	return 0x0000
}

// TODO: IF, IE
func (b *Bus) WriteByte(addr types.Addr, value byte) {
	switch {
	case addr >= 0x0000 && addr <= 0x7FFF:
		b.Cart.WriteByte(addr, value)
	case addr >= 0x8000 && addr <= 0x9FFF:
		b.VRAM.Write(addr-0x7FFF, value)
	case addr >= 0xC000 && addr <= 0xCFFF:
		b.WRAM.Write(addr-0xBFFF, value)
	case addr >= 0xD000 && addr <= 0xDFFF:
		b.WRAM2.Write(addr-0xCFFF, value)
	case addr >= 0xFF00 && addr <= 0xFF7F:
		b.IO.Write(addr-0xFEFF, value)
	case addr >= 0xFF80 && addr <= 0xFFFE:
		b.HRAM.Write(addr-0xFF7F, value)
	case addr == 0xFFFF:
		b.IE.Write(value)
	default:
		panic(fmt.Sprintf("Addr:0x%4x is not implemented", addr))
	}
}
