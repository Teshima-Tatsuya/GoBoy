package memory

// ROM is a readonly memory
type ROM struct {
	Buf []byte
}

func NewROM(data []byte) *ROM {
	return &ROM{
		Buf: data,
	}
}

// max ROM size is 8MB > uint16
func (r *ROM) Read(addr uint32) byte {
	return r.Buf[addr]
}
