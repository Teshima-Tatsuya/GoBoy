package emulator

import (
	"image"
	"math"

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
	return 160 + (16 * 8), 192
}

func (e *Emulator) Update() error {
	e.GB.Step()
	return nil
}

func (e *Emulator) Draw(screen *ebiten.Image) {
	image, tiles := e.GB.Display()

	screen.ReplacePixels(conbine(image, tiles))
}

func conbine(d1, d2 *image.RGBA) []byte {

	s1 := d1.Rect.Size()
	s2 := d2.Rect.Size()
	yMax := math.Max(float64(s1.Y), float64(s2.Y))

	data := image.NewRGBA(image.Rect(0, 0, s1.X+s2.X, int(yMax)))
	for i := 0; i < s1.X; i++ {
		for j := 0; j < s1.Y; j++ {
			data.SetRGBA(i, j, d1.RGBAAt(i, j))
		}
	}
	for i := 0; i < s2.X; i++ {
		for j := 0; j < s2.Y; j++ {
			data.SetRGBA(s1.X+i, j, d2.RGBAAt(i, j))
		}
	}

	return data.Pix
}
