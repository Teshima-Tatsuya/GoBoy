package gpu

import (
	"image"
	"image/color"

	"github.com/Teshima-Tatsuya/GoBoy/pkg/interfaces"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/types"
)

type GPU struct {
	bus        interfaces.Bus
	requestIRQ func(byte)
	clock      uint
	imageData  [][]color.RGBA
	LCDC       *LCDC
	LCDS       *LCDS
	Scroll     *Scroll
	Palette    *Palette
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
		Palette:   NewPalette(),
		DMA:       0,
		BGP:       0,
		OBP0:      0,
		OBP1:      0,
	}

	return gpu
}

func (g *GPU) Init(bus interfaces.Bus, requestIRQ func(byte)) {
	g.bus = bus
	g.requestIRQ = requestIRQ
}

// gpu main process
func (g *GPU) Step(cycles uint) {
	g.clock += cycles

	if g.clock >= CyclePerLine {
		if g.Scroll.isVBlankStart() {
			g.requestIRQ(1) // 1 is io.VBlankFlag, prepend cycle import...

		} else if g.Scroll.isVBlankPeriod() {

		} else if g.Scroll.isHBlankPeriod() {
			g.loadTile()
			// first build BG
			// second build Window IF exists
			g.drawBGLine()

			if g.LCDC.WindowEnable() {
				g.drawWinLine()
			}

		} else {
			g.Scroll.LY = 0
			g.loadTile()
			g.drawBGLine()
		}

		if g.Scroll.LY == g.Scroll.SCY {
			g.requestIRQ(2)
		}
		g.Scroll.LY++
		g.clock -= CyclePerLine
	}
}

func (g *GPU) Display() (*image.RGBA, *image.RGBA) {
	i := image.NewRGBA(image.Rect(0, 0, SCREEN_WIDTH, SCREEN_HEIGHT))
	itile := image.NewRGBA(image.Rect(0, 0, 8*16, 8*24))
	for y := 0; y < 144; y++ {
		for x := 0; x < 160; x++ {
			i.SetRGBA(x, y, g.imageData[x][y])
		}
	}

	for y := 0; y < 24; y++ {
		for x := 0; x < 16; x++ {
			for col := 0; col < 8; col++ {
				for row := 0; row < 8; row++ {
					itile.SetRGBA(x*8+col, y*8+row, GetPalette(g.tiles[y*16+x].Data[row][col]))
				}
			}
		}
	}
	return i, itile
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
			bytes16[b] = g.bus.ReadByte(types.Addr(addr) + types.Addr(i*16+b))
		}
		g.tiles[i] = *NewTile(bytes16[:])
	}
}

// Step1: get tile id from tile map
// Step2: get color form tile id
// Step3: Store color to imageData
func (g *GPU) drawBGLine() {
	for x := 0; x < SCREEN_WIDTH; x++ {
		g.imageData[x][g.Scroll.LY] = g.getTileColor(x)
	}
}

func (g *GPU) drawWinLine() {
	for x := 0; x < SCREEN_WIDTH; x++ {
		// g.imageData[x][g.Scroll.LY] = g.getTileColor(x)
	}
}

func (g *GPU) getTileColor(LX int) color.RGBA {
	// https://gbdev.io/pandocs/pixel_fifo.html#get-tile

	// Step1. Find the number of tiles from top.
	// yPos is current pixel from top(0-255)
	yPos := (g.Scroll.LY + g.Scroll.SCY) & 255
	// yTile is Tile corresponding at yPos
	yTile := yPos / 8
	// xPos is current pixel from left(0-31)
	xPos := (LX + int(g.Scroll.SCX)) & 255
	xTile := xPos / 8

	baseAddr := g.LCDC.BGTileMapArea()
	addr := types.Addr(baseAddr) + types.Addr(yTile)*32 + types.Addr(xTile)
	tileIdx := g.bus.ReadByte(addr)

	return GetPalette(g.tiles[tileIdx].Data[yPos%8][xPos%8])
}

func (g *GPU) ImageData() ([][]color.RGBA, []Tile) {
	return g.imageData, g.tiles
}

func (g *GPU) Read(addr types.Addr) byte {
	switch addr {
	case LCDCAddr:
		return g.LCDC.Data
	case LCDSAddr:
		return g.LCDS.Data
	case SCYAddr, SCXAddr, LYAddr, LYCAddr, WXAddr, WYAddr:
		return g.Scroll.Read(addr)
	case DMAAddr:
		return g.DMA
	case BGPAddr:
		return g.BGP
	case OBP0Addr:
		return g.OBP0
	case OBP1Addr:
		return g.OBP1
	default:
		panic("GPU Read")
	}
}

func (g *GPU) Write(addr types.Addr, value byte) {
	switch addr {
	case LCDCAddr:
		g.LCDC.Data = value
	case LCDSAddr:
		g.LCDS.Data = value
	case SCYAddr, SCXAddr, LYAddr, LYCAddr, WXAddr, WYAddr:
		g.Scroll.Write(addr, value)
	case DMAAddr:
		g.DMA = value
	case BGPAddr:
		g.BGP = value
	case OBP0Addr:
		g.OBP0 = value
	case OBP1Addr:
		g.OBP1 = value
	default:
		panic("GPU Write")
	}
}
