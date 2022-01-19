package gpu

import (
	"image"
	"image/color"
	"math/rand"

	"github.com/Teshima-Tatsuya/GoBoy/pkg/types"
)

// Offset is FF00
const (
	LCDCAddr types.Addr = 0x40
	LCDSAddr            = 0x41
	SCYAddr             = 0x42
	SCXAddr             = 0x43
	LYAddr              = 0x44
	LYCAddr             = 0x45
	DMAAddr             = 0x46
	BGPAddr             = 0x47
	OBP0Addr            = 0x48
	OBP1Addr            = 0x49
	WYAddr              = 0x4A
	WXAddr              = 0x4B
)

type GPU struct {
	LCDC *LCDC
	LCDS *LCDS
	SCY  byte
	SCX  byte
	LY   byte
	LYC  byte
	DMA  byte
	BGP  byte
	OBP0 byte
	OBP1 byte
	WX   byte
	WY   byte
}

func New() *GPU {
	return &GPU{
		LCDC: NewLCDC(0x00),
		LCDS: NewLCDS(0x00),
		SCY:  0,
		SCX:  0,
		LY:   0,
		LYC:  0,
		DMA:  0,
		BGP:  0,
		OBP0: 0,
		OBP1: 0,
		WX:   0,
		WY:   0,
	}
}

func (g *GPU) Display() *image.RGBA {
	i := image.NewRGBA(image.Rect(0, 0, 160, 144))
	for y := 0; y < 144; y++ {
		for x := 0; x < 160; x++ {
			//			p := g.Renderer.outputBuffer[y*160+x]
			//			red, green, blue := byte((p&0b11111)*8), byte(((p>>5)&0b11111)*8), byte(((p>>10)&0b11111)*8)

			i.SetRGBA(x, y, color.RGBA{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), 0xff})
		}
	}
	return i
}

func (g *GPU) Read(addr types.Addr) byte {
	switch addr {
	case LCDCAddr:
		return g.LCDC.Data
	case LCDSAddr:
		return g.LCDS.Data
	case SCYAddr:
		return g.SCY
	case SCXAddr:
		return g.SCX
	case LYAddr:
		return g.LY
	case LYCAddr:
		return g.LYC
	case DMAAddr:
		return g.DMA
	case BGPAddr:
		return g.BGP
	case OBP0Addr:
		return g.OBP0
	case OBP1Addr:
		return g.OBP1
	case WXAddr:
		return g.WX
	case WYAddr:
		return g.WY
	}
	return 0
}

func (g *GPU) Write(addr types.Addr, value byte) {
	switch addr {
	case LCDCAddr:
		g.LCDC.Data = value
	case LCDSAddr:
		g.LCDS.Data = value
	case SCYAddr:
		g.SCY = value
	case SCXAddr:
		g.SCX = value
	case LYAddr:
		g.LY = value
	case LYCAddr:
		g.LYC = value
	case DMAAddr:
		g.DMA = value
	case BGPAddr:
		g.BGP = value
	case OBP0Addr:
		g.OBP0 = value
	case OBP1Addr:
		g.OBP1 = value
	case WXAddr:
		g.WX = value
	case WYAddr:
		g.WY = value
	}
}
