package cartridge

type Bank struct {
	Current byte
	Size    byte
}

func NewBank(size byte) *Bank {
	return &Bank{
		Current: 0x00,
		Size:    size,
	}
}
