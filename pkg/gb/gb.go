package gb

import (
	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/cartridge"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/ram"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/rom"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/timer"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/video"
)

type GB struct {
	Cartridge *cartridge.Cartridge

	// memory
	ROM  *rom.ROM
	VRAM *ram.RAM
	WRAM *ram.RAM
	HRAM *ram.RAM

	Video *video.Video
	Timer *timer.Timer
}

func NewGB(romData []byte) *GB {
	gb := &GB{
		Cartridge: cartridge.New(romData),
		VRAM:      ram.New(0x2000),
		WRAM:      ram.New(0x2000),
		HRAM:      ram.New(0x0080),
		Timer:     timer.NewTimer(),
	}

	gb.loadRom(romData)

	return gb
}

func (gb *GB) Draw() []byte {
	return gb.Video.Display().Pix
}

// load romData to ROM
func (gb *GB) loadRom(romData []byte) {
	bankSize := int(gb.Cartridge.Bank.Size)

	for bank := 0; bank < bankSize; bank++ {
		for i := 0; i < 0x4000; i++ {
			gb.ROM.Buf[bank][i] = romData[bank*0x4000+i]
		}
	}
}
