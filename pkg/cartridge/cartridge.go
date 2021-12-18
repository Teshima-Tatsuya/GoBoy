package cartridge

// @see https://gbdev.io/pandocs/The_Cartridge_Header.html
type Cartridge struct {
	Title string
}

func New(romData []byte) *Cartridge {
	var title []byte

	for i := 0x134; i < 0x143; i++ {
		title = append(title, romData[i])
	}
	return &Cartridge{
		Title: string(title),
	}
}
