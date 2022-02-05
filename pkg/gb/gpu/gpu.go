package gpu

import (
	"image"
	"image/color"

	"github.com/Teshima-Tatsuya/GoBoy/pkg/debug"
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
	palette    *Palette
	DMA        byte
	tiles      []Tile
	dmaStarted bool
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
		palette:   NewPalette(),
		DMA:       0,
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
		g.loadTile()
		if g.Scroll.isVBlankStart() {
			g.drawSplite()
			g.requestIRQ(1) // 1 is io.VBlankFlag, prepend cycle import...

		} else if g.Scroll.isVBlankPeriod() {

		} else if g.Scroll.isHBlankPeriod() {
			// first build BG
			// second build Window IF exists
			g.drawBGLine()

			if g.LCDC.WindowEnable() {
				g.drawWinLine()
			}

		} else {
			g.Scroll.LY = 0
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
					itile.SetRGBA(x*8+col, y*8+row, g.palette.GetPalette(g.tiles[y*16+x].Data[row][col]))
				}
			}
		}
	}
	return i, itile
}

func (g *GPU) loadTile() {
	// addr := g.LCDC.BGWinTileDataArea()
	addr := 0x8000
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
		g.imageData[x][g.Scroll.LY] = g.getBGTileColor(x)
	}
}

func (g *GPU) drawWinLine() {
	for x := 0; x < SCREEN_WIDTH; x++ {
		g.imageData[x][g.Scroll.LY] = g.getWinTileColor(x)
	}
}

func (g *GPU) drawSplite() {
	for i := 0; i < SPRITE_NUM; i++ {
		bytes4 := [4]byte{}
		for j := 0; j < 4; j++ {
			addr := OAMSTARTAddr + types.Addr(i*4) + types.Addr(j)
			// debug.Info("0x%04x", addr)
			bytes4[j] = g.bus.ReadByte(addr)
		}

		s := NewSprite(bytes4[:])

		var objHeight int
		if g.LCDC.OBJSize() == 1 {
			objHeight = 16
		} else {
			objHeight = 8
		}

		for x := 0; x < 8; x++ {
			for y := 0; y < objHeight; y++ {
				xPos := int(s.x) + x
				yPos := int(s.y) + y

				// ignore out of screen
				if (xPos < 0 || SCREEN_WIDTH <= xPos) ||
					(yPos < 0 || SCREEN_HEIGHT <= yPos) {
					continue
				}

				tile := g.tiles[s.tileIdx]

				if s.YFlip() {
					y = 7 - y
				}
				yPos = int(s.y) + y
				if s.XFlip() {
					x = 7 - x
				}
				xPos = int(s.x) + x

				c := tile.Data[x][y]

				if c != 0 {
					p := g.palette.GetObjPalette(c, uint(s.MBGPalleteNo()))
					g.imageData[xPos][yPos] = p
				}
			}
		}
	}

}

func (g *GPU) getBGTileColor(LX int) color.RGBA {
	// yPos is current pixel from top(0-255)
	yPos := (g.Scroll.LY + g.Scroll.SCY) & 255
	xPos := (LX + int(g.Scroll.SCX)) & 255
	baseAddr := g.LCDC.BGTileMapArea()

	return g.getTileColor(xPos, int(yPos), types.Addr(baseAddr))
}

func (g *GPU) getWinTileColor(LX int) color.RGBA {
	// yPos is current pixel from top(0-255)
	yPos := g.Scroll.WY
	xPos := g.Scroll.WX - 7
	baseAddr := g.LCDC.WinTileMapArea()

	return g.getTileColor(int(xPos), int(yPos), types.Addr(baseAddr))
}

func (g *GPU) getTileColor(xPos, yPos int, baseAddr types.Addr) color.RGBA {
	// https://gbdev.io/pandocs/pixel_fifo.html#get-tile

	// yTile is Tile corresponding at yPos
	yTile := yPos / 8
	// xPos is current pixel from left(0-31)
	xTile := xPos / 8

	addr := types.Addr(baseAddr) + types.Addr(yTile)*32 + types.Addr(xTile)
	tileIdx := g.bus.ReadByte(addr)

	if g.LCDC.BGWinTileDataArea() == 0x8800 {
		tileIdx += 128
	}

	return g.palette.GetPalette(g.tiles[tileIdx].Data[yPos%8][xPos%8])
}

func (g *GPU) ImageData() ([][]color.RGBA, []Tile) {
	return g.imageData, g.tiles
}

// WIP
func (g *GPU) TransferOAM() {
	for i := 0; i < 0xA0; i++ {
		addr := types.Addr(g.DMA) * 0x100
		b := g.bus.ReadByte(addr + types.Addr(i))
		g.bus.WriteByte(OAMSTARTAddr+types.Addr(i), b)
	}

	g.dmaStarted = false
}

func (g *GPU) IsDmaStarted() bool {
	return g.dmaStarted
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
	case BGPAddr, OBP0Addr, OBP1Addr:
		return g.palette.Read(addr)
	default:
		debug.Fatal("GPU Read 0x%04X", addr)
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
		g.Scroll.Write(addr, value)
	case DMAAddr:
		g.dmaStarted = true
		g.DMA = value
	case BGPAddr, OBP0Addr, OBP1Addr:
		g.palette.Write(addr, value)
	default:
		debug.Fatal("GPU Write 0x%04X", addr)
	}
}
