package bus

import (
	"fmt"

	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/cartridge"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/io"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/irq"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/ram"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/types"
)

type Bus struct {
	Cart  *cartridge.Cartridge
	VRAM  *ram.RAM
	WRAM  *ram.RAM
	WRAM2 *ram.RAM
	HRAM  *ram.RAM
	ERAM  *ram.RAM
	IO    *io.IO
	IRQ   *irq.IRQ
}

func New(cart *cartridge.Cartridge, vram *ram.RAM, wram *ram.RAM, wram2 *ram.RAM, hram *ram.RAM, io *io.IO, irq *irq.IRQ) *Bus {
	eram := ram.New(wram.Size)
	return &Bus{
		Cart:  cart,
		VRAM:  vram,
		WRAM:  wram,
		WRAM2: wram2,
		HRAM:  hram,
		ERAM:  eram,
		IO:    io,
		IRQ:   irq,
	}
}

// TODO: IF, IE
func (b *Bus) ReadByte(addr types.Addr) byte {
	switch {
	case addr >= 0x0000 && addr <= 0x7FFF:
		return b.Cart.ReadByte(addr)
	case addr >= 0x8000 && addr <= 0x9FFF:
		return b.VRAM.Read(addr - 0x8000)
	case addr >= 0xC000 && addr <= 0xCFFF:
		return b.WRAM.Read(addr - 0xC000)
	case addr >= 0xD000 && addr <= 0xDFFF:
		return b.WRAM2.Read(addr - 0xD000)
	case addr >= 0xE000 && addr <= 0xFDFF:
		return b.ERAM.Read(addr - 0xE000)
	case addr == 0xFF0F:
		return b.IRQ.Read(addr - 0xFF00)
	case addr >= 0xFF00 && addr <= 0xFF7F:
		return b.IO.Read(addr - 0xFF00)
	case addr >= 0xFF80 && addr <= 0xFFFE:
		return b.HRAM.Read(addr - 0xFF80)
	case addr == 0xFFFF:
		return b.IRQ.Read(addr - 0xFF00)
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
		b.VRAM.Write(addr-0x8000, value)
	case addr >= 0xC000 && addr <= 0xCFFF:
		b.WRAM.Write(addr-0xC000, value)
	case addr >= 0xD000 && addr <= 0xDFFF:
		b.WRAM2.Write(addr-0xD000, value)
	case addr >= 0xE000 && addr <= 0xFDFF:
		b.ERAM.Write(addr-0xE000, value)
	case addr == 0xFF0F:
		b.IRQ.Write(addr-0xFF00, value)
	case addr >= 0xFF00 && addr <= 0xFF7F:
		b.IO.Write(addr-0xFF00, value)
	case addr >= 0xFF80 && addr <= 0xFFFE:
		b.HRAM.Write(addr-0xFF80, value)
	case addr == 0xFFFF:
		b.IRQ.Write(addr-0xFF00, value)
	default:
		panic(fmt.Sprintf("Addr:0x%4x is not implemented", addr))
	}
}
