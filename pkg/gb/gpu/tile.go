package gpu

type Color byte

const (
	White Color = iota
	LightGray
	DarkGray
	Black
)

type Tile struct {
	// 8 x 8 pixels
	Data [8][8]Color
}

func NewTile() *Tile {
	return &Tile{}
}
