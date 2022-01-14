package cartridge

type Bank struct {
	Current byte
	Size    uint16
}

func NewBank(romSize int) *Bank {
	return &Bank{
		Current: 0x00,
		Size:    uint16(romSize / (16 * 1024)),
	}
}
