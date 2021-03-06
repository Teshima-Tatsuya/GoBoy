package gpu

import (
	"image/color"

	"github.com/Teshima-Tatsuya/GoBoy/pkg/types"
)

var palette = []color.RGBA{
	{175, 197, 160, 255},
	{93, 147, 66, 255},
	{22, 63, 48, 255},
	{0, 40, 0, 255},
}

type Palette struct {
	// FF47
	BGP byte
	// FF48
	OBP0 byte
	// FF49
	OBP1 byte

	// CGB Only

	// FF68
	BCPS byte
	// FF69
	BCPD byte
	// FF6A
	OCPS byte
	// FF6B
	OCPD byte
}

func NewPalette() *Palette {
	return &Palette{}
}

func (p *Palette) GetPalette(idx Color) color.RGBA {
	// Bit 7-6 - Color for index 3
	// Bit 5-4 - Color for index 2
	// Bit 3-2 - Color for index 1
	// Bit 1-0 - Color for index 0
	c := (p.BGP >> (idx * 2)) & 0x03
	return palette[c]
}

func (p *Palette) GetObjPalette(idx Color, obp uint) color.RGBA {
	// Bit 7-6 - Color for index 3
	// Bit 5-4 - Color for index 2
	// Bit 3-2 - Color for index 1
	// Bit 1-0 - Color for index 0
	var c byte
	if obp == 1 {
		c = (p.OBP1 >> (idx * 2)) & 0x03
	} else {
		c = (p.OBP0 >> (idx * 2)) & 0x03
	}

	return palette[c]
}

func (p *Palette) Read(addr types.Addr) byte {
	switch addr {
	case BGPAddr:
		return p.BGP
	case OBP0Addr:
		return p.OBP0
	case OBP1Addr:
		return p.OBP1
	case BCPSAddr:
		return p.BCPS
	case BCPDAddr:
		return p.BCPD
	case OCPSAddr:
		return p.OCPS
	case OCPDAddr:
		return p.OCPD
	default:
		panic("Palette Read")
	}
}

func (p *Palette) Write(addr types.Addr, value byte) {
	switch addr {
	case BGPAddr:
		p.BGP = value
	case OBP0Addr:
		p.OBP0 = value
	case OBP1Addr:
		p.OBP1 = value
	case BCPSAddr:
		p.BCPS = value
	case BCPDAddr:
		p.BCPD = value
	case OCPSAddr:
		p.OCPS = value
	case OCPDAddr:
		p.OCPD = value
	default:
		panic("Palette Write")
	}
}
