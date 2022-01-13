package gb

import (
	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/cartridge"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/ram"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/timer"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/video"
)

type GB struct {
	Cartridge *cartridge.Cartridge
	Video     *video.Video
	VRAM      *ram.RAM
	WRAM      *ram.RAM
	HRAM      *ram.RAM
	Timer     *timer.Timer
}

func NewGB(romData []byte) *GB {
	gb := &GB{
		Cartridge: cartridge.New(romData),
		VRAM:      ram.New(0x2000),
		WRAM:      ram.New(0x2000),
		HRAM:      ram.New(0x0080),
		Timer:     timer.NewTimer(),
	}

	return gb
}

func (gb *GB) Draw() []byte {
	return gb.Video.Display().Pix
}
