package gb

import (
	"github.com/DisgoOrg/log"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/cartridge"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/video"
)

type GB struct {
	Cartridge *cartridge.Cartridge
	Video     *video.Video
}

func NewGB(romData []byte) *GB {
	c := cartridge.New(romData)
	log.Info(c.Title)

	gb := &GB{
		Cartridge: c,
	}

	return gb
}

func (gb *GB) Draw() []byte {
	return gb.Video
}
