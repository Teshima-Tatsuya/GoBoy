package memory

import "github.com/Teshima-Tatsuya/GoBoy/pkg/types"

type RAM struct {
	Buf []byte
}

func NewRAM(size int) *RAM {
	buf := make([]byte, size)

	return &RAM{
		Buf: buf,
	}
}

func (r *RAM) Read(addr types.Addr) byte {
	return r.Buf[addr]
}

func (r *RAM) Write(addr types.Addr, value byte) {
	r.Buf[addr] = value
}
