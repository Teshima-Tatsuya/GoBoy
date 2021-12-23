package goboy

import (
	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb"
	"github.com/hajimehoshi/ebiten/v2"
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

	return g
}

func (gb *GoBoy) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	// @see https://gbdev.io/pandocs/Rendering.html
	return 160, 144
}

func (gb *GoBoy) Update() error {
	return nil
}

func (gb *GoBoy) Draw(screen *ebiten.Image) {
	screen.ReplacePixels(gb.GB.Draw())
}
