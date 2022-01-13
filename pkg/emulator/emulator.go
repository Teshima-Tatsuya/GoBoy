package emulator

import (
	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb"
	"github.com/apex/log"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
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
	if e.RomData[0xff02] == 0x81 {
		log.Info(string(e.RomData[0xff01]))
		e.RomData[0xff02] = 0x0
	}
	return nil
}

func (e *Emulator) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "hello, world")
	if e.RomData[0xff02] == 0x81 {
		log.Info(string(e.RomData[0xff01]))
		e.RomData[0xff02] = 0x0
	}
}