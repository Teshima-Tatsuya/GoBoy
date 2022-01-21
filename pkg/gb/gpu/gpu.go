package gpu

import (
	"image"
	"image/color"
	"math/rand"

	"github.com/Teshima-Tatsuya/GoBoy/pkg/interfaces/bus"
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
	bus        bus.IO
	requestIRQ func(byte)
	clock      uint
	imageData  [][]color.RGBA
	LCDC       *LCDC
	LCDS       *LCDS
	Scroll     *Scroll
	DMA        byte
	BGP        byte
	OBP0       byte
	OBP1       byte
	tiles      []Tile
}

func New() *GPU {
	imageData := make([][]color.RGBA, SCREEN_WIDTH)

	for i := 0; i < SCREEN_WIDTH; i++ {
		imageData[i] = make([]color.RGBA, SCREEN_HEIGHT)
	}

	gpu := &GPU{
		clock:     0,
		imageData: imageData,
		LCDC:      NewLCDC(0x00),
		LCDS:      NewLCDS(0x00),
		Scroll:    NewScroll(),
		DMA:       0,
		BGP:       0,
		OBP0:      0,
		OBP1:      0,
	}

	return gpu
}

func (g *GPU) Init(bus bus.IO, requestIRQ func(byte)) {
	g.bus = bus
	g.requestIRQ = requestIRQ
	g.loadTile()
}

// gpu main process
func (g *GPU) Step(cycles uint) {
	g.clock += cycles

	if g.clock >= CyclePerLine {
		if g.Scroll.isVBlankStart() {
			g.requestIRQ(1) // 1 is io.VBlankFlag, prepend cycle import...

		} else if g.Scroll.isVBlankPeriod() {

		} else if g.Scroll.isHBlankPeriod() {
			// first build BG
			// second build Window IF exists

		}
		g.Scroll.LY++
		g.clock -= CyclePerLine
	}
}

func (g *GPU) Display() *image.RGBA {
	i := image.NewRGBA(image.Rect(0, 0, SCREEN_WIDTH, SCREEN_HEIGHT))
	for y := 0; y < 144; y++ {
		for x := 0; x < 160; x++ {
			//			p := g.Renderer.outputBuffer[y*160+x]
			//			red, green, blue := byte((p&0b11111)*8), byte(((p>>5)&0b11111)*8), byte(((p>>10)&0b11111)*8)

			i.SetRGBA(x, y, color.RGBA{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), 0xff})
		}
	}
	return i
}

func (g *GPU) loadTile() {
	addr := g.LCDC.BGWinTileDataArea()
	// todo CGBMode
	tileNum := 384
	g.tiles = make([]Tile, tileNum)
	var bytes16 [16]byte

	// One tile occupies 16 bytes
	for i := 0; i < tileNum; i++ {
		for b := 0; b < 16; b++ {
			bytes16[b] = g.bus.ReadByte(types.Addr(addr) + types.Addr(tileNum*16+b))
		}
		g.tiles[i] = *NewTile(bytes16[:])
	}
}

func (g *GPU) Read(addr types.Addr) byte {
	switch addr {
	case LCDCAddr:
		return g.LCDC.Data
	case LCDSAddr:
		return g.LCDS.Data
	case SCYAddr, SCXAddr, LYAddr, LYCAddr, WXAddr, WYAddr:
		// return g.Scroll
	case DMAAddr:
		return g.DMA
	case BGPAddr:
		return g.BGP
	case OBP0Addr:
		return g.OBP0
	case OBP1Addr:
		return g.OBP1
	}
	return 0
}

func (g *GPU) Write(addr types.Addr, value byte) {
	switch addr {
	case LCDCAddr:
		g.LCDC.Data = value
	case LCDSAddr:
		g.LCDS.Data = value
	case SCYAddr, SCXAddr, LYAddr, LYCAddr, WXAddr, WYAddr:
	case DMAAddr:
		g.DMA = value
	case BGPAddr:
		g.BGP = value
	case OBP0Addr:
		g.OBP0 = value
	case OBP1Addr:
		g.OBP1 = value
	}
}
