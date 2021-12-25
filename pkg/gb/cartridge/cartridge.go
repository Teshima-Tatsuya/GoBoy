package cartridge

type CartridgeType byte

const (
	MBC_0                 = 0x00
	MBC_1                 = 0x01
	MBC_1_RAM             = 0x02
	MBC_1_RAM_BATT        = 0x03
	MBC_3_RAM_BATT        = 0x13
	MBC_3_RAM_BATT_RTC    = 0x10
	MBC_5                 = 0x19
	MBC_5_RAM             = 0x1A
	MBC_5_RAM_BATT        = 0x1B
	MBC_5_RUMBLE          = 0x1C
	MBC_5_RAM_RUMBLE      = 0x1D
	MBC_5_RAM_BATT_RUMBLE = 0x1E
)

// @see https://gbdev.io/pandocs/The_Cartridge_Header.html
type Cartridge struct {
	NintendoLogo []byte
	// Game Title
	Title string
	// GameBoy Color Flag
	CGBFlag bool
	// game supports SGB functions
	SGBFlag bool
	Type    byte
	ROMSize byte
	RAMSize int
	ROMData []byte
}

func New(romData []byte) *Cartridge {
	nintendLogo := romData[0x0104:0x0133]
	title := string(romData[0x0134:0x0143])

	cgbflag := romData[0x0143] == 0x80 || romData[0x0143] == 0xC0
	sgbflag := romData[0x0146] != 0x00

	return &Cartridge{
		NintendoLogo: nintendLogo,
		Title:        string(title),
		CGBFlag:      cgbflag,
		SGBFlag:      sgbflag,
		ROMSize:      romData[0x0148],
		RAMSize:      getRamSize(romData[0x0149]),
		ROMData:      romData,
	}
}

func getRamSize(size byte) int {
	switch size {
	case 0x00, 0x01:
		return 0
	case 0x02:
		return 8 * (2 << 10)
	case 0x03:
		return 32 * (2 << 10)
	case 0x04:
		return 128 * (2 << 10)
	case 0x05:
		return 64 * (2 << 10)
	}
	return 0
}

func (c *Cartridge) ReadByte(addr uint16) byte {
	return c.ROMData[addr]
}
