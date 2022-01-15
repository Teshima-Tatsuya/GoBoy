package gb

import (
	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/cartridge"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/memory"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/timer"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/video"
)

type GB struct {
	Cartridge *cartridge.Cartridge

	// memory
	VRAM *memory.RAM
	WRAM *memory.RAM
	HRAM *memory.RAM

	Video *video.Video
	Timer *timer.Timer
}

func NewGB(romData []byte) *GB {
	gb := &GB{
		Cartridge: cartridge.New(romData),
		VRAM:      memory.NewRAM(0x2000),
		WRAM:      memory.NewRAM(0x2000),
		HRAM:      memory.NewRAM(0x0080),
		Timer:     timer.NewTimer(),
	}

	return gb
}

func (gb *GB) Draw() []byte {
	return gb.Video.Display().Pix
}
