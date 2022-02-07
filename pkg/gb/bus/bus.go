package bus

import (
	"github.com/Teshima-Tatsuya/GoBoy/pkg/debug"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/apu"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/cartridge"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/gpu"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/interrupt"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/memory"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/pad"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/serial"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/timer"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/types"
)

type Bus struct {
	Cart   *cartridge.Cartridge
	VRAM   *memory.RAM
	WRAM   *memory.RAM
	WRAM2  *memory.RAM
	HRAM   *memory.RAM
	ERAM   *memory.RAM
	oam    *memory.RAM
	apu    *apu.APU
	pad    *pad.Pad
	timer  *timer.Timer
	serial *serial.Serial
	gpu    *gpu.GPU
	irq    *interrupt.IRQ
}

func New(cart *cartridge.Cartridge, vram *memory.RAM, wram *memory.RAM, wram2 *memory.RAM, hram *memory.RAM, a *apu.APU, g *gpu.GPU, irq *interrupt.IRQ, pad *pad.Pad, timer *timer.Timer, serial *serial.Serial) *Bus {
	eram := memory.NewRAM(0x2000)
	oam := memory.NewRAM(0x00A0)
	return &Bus{
		Cart:   cart,
		VRAM:   vram,
		WRAM:   wram,
		WRAM2:  wram2,
		HRAM:   hram,
		ERAM:   eram,
		oam:    oam,
		apu:    a,
		gpu:    g,
		irq:    irq,
		pad:    pad,
		timer:  timer,
		serial: serial,
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
		return b.oam.Read(addr - 0xFE00)
	case addr >= 0xFEA0 && addr <= 0xFEFF:
		return 0
	case addr == 0xFF00:
		return b.pad.Read(addr - 0xFF00)
	case addr == 0xFF01 || addr == 0xFF02:
		return b.serial.Read(addr - 0xFF00)
	case addr >= 0xFF04 && addr <= 0xFF07:
		return b.timer.Read(addr - 0xFF00)
	case addr == 0xFF0F || addr == 0xFFFF:
		return b.irq.Read(addr - 0xFF00)
	case addr >= 0xFF10 && addr <= 0xFF3F:
		return b.apu.Read(addr - 0xFF00)
	case addr >= 0xFF40 && addr <= 0xFF4B:
		return b.gpu.Read(addr - 0xFF00)
	case addr >= 0xFF80 && addr <= 0xFFFE:
		return b.HRAM.Read(addr - 0xFF80)
	default:
		debug.Fatal("Non Supported Read Addr 0x%4d", addr)
	}

	return 0xFF
}

func (b *Bus) ReadAddr(addr types.Addr) types.Addr {
	return 0x0000
}

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
		b.oam.Write(addr-0xFE00, value)
	case addr >= 0xFEA0 && addr <= 0xFEFF:
		// Nintendo says use of this area is prohibited.
	case addr == 0xFF00:
		b.pad.Write(addr-0xFF00, value)
	case addr == 0xFF01 || addr == 0xFF02:
		b.serial.Write(addr-0xFF00, value)
	case addr >= 0xFF04 && addr <= 0xFF07:
		b.timer.Write(addr-0xFF00, value)
	case addr == 0xFF0F || addr == 0xFFFF:
		b.irq.Write(addr-0xFF00, value)
	case addr >= 0xFF10 && addr <= 0xFF3F:
		b.apu.Write(addr-0xFF00, value)
	case addr >= 0xFF40 && addr <= 0xFF4B:
		b.gpu.Write(addr-0xFF00, value)
	case addr >= 0xFF80 && addr <= 0xFFFE:
		b.HRAM.Write(addr-0xFF80, value)
	default:
		debug.Fatal("Addr:0x%4x is not implemented", addr)
	}
}
