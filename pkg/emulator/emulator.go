package emulator

import (
	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb"
	"github.com/apex/log"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type GoBoy struct {
	RomData []byte
	GB      *gb.GB
}

func New(romData []byte) *GoBoy {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("60fps")
	ebiten.SetWindowSize(160*4, 144*4)

	gb := gb.NewGB(romData)
	g := &GoBoy{
		RomData: romData,
		GB:      gb,
	}

	log.Info(string(g.RomData[0xff01]))

	return g
}

func (gb *GoBoy) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	// @see https://gbdev.io/pandocs/Rendering.html
	return 160, 144
}

func (gb *GoBoy) Update() error {
	if gb.RomData[0xff02] == 0x81 {
		log.Info(string(gb.RomData[0xff01]))
		gb.RomData[0xff02] = 0x0
	}
	return nil
}

func (gb *GoBoy) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "hello, world")
	if gb.RomData[0xff02] == 0x81 {
		log.Info(string(gb.RomData[0xff01]))
		gb.RomData[0xff02] = 0x0
	}
	//	screen.ReplacePixels(gb.GB.Draw())
}
