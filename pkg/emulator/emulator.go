package emulator

import (
	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb"
	"github.com/hajimehoshi/ebiten/v2"
)

type Emulator struct {
	RomData []byte
	GB      *gb.GB
}

func New(romData []byte) *Emulator {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("60fps")
	ebiten.SetWindowSize(160*4, 144*4)

	gb := gb.NewGB(romData)
	e := &Emulator{
		RomData: romData,
		GB:      gb,
	}

	return e
}

func (e *Emulator) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	// @see https://gbdev.io/pandocs/Rendering.html
	return 160, 144
}

func (e *Emulator) Update() error {
	e.GB.Step()
	return nil
}

func (e *Emulator) Draw(screen *ebiten.Image) {
	imageData := e.GB.Display()

	for y := 0; y < 144; y++ {
		for x := 0; x < 160; x++ {
			screen.Set(x, y, imageData[x][y])
		}
	}
}
