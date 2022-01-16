package bus

import (
	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/cartridge"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/io"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/memory"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/video"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/types"
)

type Bus struct {
	Cart  *cartridge.Cartridge
	VRAM  *memory.RAM
	WRAM  *memory.RAM
	WRAM2 *memory.RAM
	HRAM  *memory.RAM
	ERAM  *memory.RAM
	OAM   *video.OAM
	IO    *io.IO
	IRQ   *io.IRQ
}

func New(cart *cartridge.Cartridge, vram *memory.RAM, wram *memory.RAM, wram2 *memory.RAM, hram *memory.RAM, io *io.IO, irq *io.IRQ) *Bus {
	eram := memory.NewRAM(0x2000)
	oam := video.NewOAM()
	return &Bus{
		Cart:  cart,
		VRAM:  vram,
		WRAM:  wram,
		WRAM2: wram2,
		HRAM:  hram,
		ERAM:  eram,
		OAM:   oam,
		IO:    io,
		IRQ:   irq,
	}
}

func (b *Bus) ReadByte(addr types.Addr) byte {
	switch {
	case addr >= 0x0000 && addr <= 0x7FFF:
		return b.Cart.ReadByte(addr)
	case addr >= 0x8000 && addr <= 0x9FFF:
		return b.VRAM.Read(addr - 0x8000)
	case addr >= 0xA000 && addr <= 0xBFFF:
		return b.Cart.ReadByte(addr)
	case addr >= 0xC000 && addr <= 0xCFFF:
		return b.WRAM.Read(addr - 0xC000)
	case addr >= 0xD000 && addr <= 0xDFFF:
		return b.WRAM2.Read(addr - 0xD000)
	case addr >= 0xE000 && addr <= 0xFDFF:
		return b.ERAM.Read(addr - 0xE000)
	case addr >= 0xFE00 && addr <= 0xFE9F:
		return b.OAM.Read(addr - 0xFE00)
	case addr == 0xFF0F:
		return b.IRQ.Read(addr - 0xFF00)
	case addr >= 0xFF00 && addr <= 0xFF7F:
		return b.IO.Read(addr - 0xFF00)
	case addr >= 0xFF80 && addr <= 0xFFFE:
		return b.HRAM.Read(addr - 0xFF80)
	case addr == 0xFFFF:
		return b.IRQ.Read(addr - 0xFF00)
	default:
		// panic(fmt.Sprintf("Non Supported Read Addr 0x%4d", addr))
	}

	return 0
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
	case addr >= 0xA000 && addr <= 0xBFFF:
		b.Cart.WriteByte(addr, value)
	case addr >= 0xC000 && addr <= 0xCFFF:
		b.WRAM.Write(addr-0xC000, value)
	case addr >= 0xD000 && addr <= 0xDFFF:
		b.WRAM2.Write(addr-0xD000, value)
	case addr >= 0xE000 && addr <= 0xFDFF:
		b.ERAM.Write(addr-0xE000, value)
	case addr >= 0xFE00 && addr <= 0xFE9F:
		b.OAM.Write(addr-0xFE00, value)
	case addr == 0xFF0F:
		b.IRQ.Write(addr-0xFF00, value)
	case addr >= 0xFF00 && addr <= 0xFF7F:
		b.IO.Write(addr-0xFF00, value)
	case addr >= 0xFF80 && addr <= 0xFFFE:
		b.HRAM.Write(addr-0xFF80, value)
	case addr == 0xFFFF:
		b.IRQ.Write(addr-0xFF00, value)
	default:
		// panic(fmt.Sprintf("Addr:0x%4x is not implemented", addr))
	}
}
