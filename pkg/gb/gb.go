package gb

import (
	"github.com/Teshima-Tatsuya/GoBoy/pkg/cartridge"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/video"
)

type GB struct {
	Cartridge *cartridge.Cartridge
	Video     *video.Video
}

func NewGB(romData []byte) *GB {
	gb := &GB{
		Cartridge: cartridge.New(romData),
	}

	return gb
}

func (gb *GB) Draw() []byte {
	return gb.Video.Display().Pix
}
