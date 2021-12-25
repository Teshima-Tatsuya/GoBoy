package video

// Object Attribute Memory
// FE00-FE9F
type OAM struct {
	// Sprite’s vertical position on the screen + 16
	Y byte
	// Sprite’s horizontal position on the screen + 8
	X         byte
	TileIndex byte
	AttrFlags byte
}

// WIP
// @see https://gbdev.io/pandocs/OAM.html#writing-data-to-oam
func (o *OAM) Write(data byte) {

}
