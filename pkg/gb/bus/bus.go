package bus

import (
	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/cartridge"
)

type Bus struct {
	Cart *cartridge.Cartridge
}

func New(cart *cartridge.Cartridge) *Bus {
	return &Bus{
		Cart: cart,
	}
}

func (b *Bus) ReadByte(addr uint16) byte {
	switch {
	case addr >= 0x0000 && addr <= 0x7FFF:
		return b.Cart.ReadByte(addr)
	}

	return 0x00
}
