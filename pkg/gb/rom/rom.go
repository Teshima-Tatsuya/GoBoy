package rom

type ROM struct {
	// bank is 1-256
	Bank byte
	// buf is 16kb(0x4000-0x7FFF)
	Buf [256][0x4000]byte
}
