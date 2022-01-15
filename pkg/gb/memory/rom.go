package memory

type ROM struct {
	// bank is 1-256
	Bank uint8
	// buf is 16kb(0x4000-0x7FFF)
	Buf [][]byte
}
