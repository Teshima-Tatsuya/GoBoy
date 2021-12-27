package bus

import (
	"fmt"

	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/cartridge"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/ram"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/types"
)

type Bus struct {
	Cart *cartridge.Cartridge
	VRAM *ram.RAM
	WRAM *ram.RAM
	HRAM *ram.RAM
}

func New(cart *cartridge.Cartridge, vram *ram.RAM, wram *ram.RAM, hram *ram.RAM) *Bus {
	return &Bus{
		Cart: cart,
		VRAM: vram,
		WRAM: wram,
		HRAM: hram,
	}
}

func (b *Bus) ReadByte(addr types.Addr) byte {
	switch {
	case addr >= 0x0000 && addr <= 0x7FFF:
		return b.Cart.ReadByte(addr)
	case addr >= 0x8000 && addr <= 0x9FFF:
		return b.VRAM.Read(addr - 0x7FFF)
	case addr >= 0xC000 && addr <= 0xCFFF:
		return b.WRAM.Read(addr - 0xBFFF)
	case addr >= 0xFF80 && addr <= 0xFFFE:
		return b.HRAM.Read(addr - 0xFF7F)
	default:
		panic(fmt.Sprintf("Non Supported Read Addr 0x%4d", addr))
	}
}

func (b *Bus) WriteByte(addr types.Addr, value byte) {
	switch {
	case addr >= 0x0000 && addr <= 0x7FFF:
		b.Cart.WriteByte(addr, value)
	}
}
