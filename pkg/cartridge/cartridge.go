package cartridge

// @see https://gbdev.io/pandocs/The_Cartridge_Header.html
type Cartridge struct {
	NintendoLogo []byte
	Title        string
}

func New(romData []byte) *Cartridge {
	var title []byte
	var nintendLogo []byte

	for i := 0x104; i < 0x133; i++ {
		nintendLogo = append(nintendLogo, romData[i])
	}

	for i := 0x134; i < 0x143; i++ {
		title = append(title, romData[i])
	}

	return &Cartridge{
		NintendoLogo: nintendLogo,
		Title:        string(title),
	}
}
