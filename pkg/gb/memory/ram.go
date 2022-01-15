package memory

import "github.com/Teshima-Tatsuya/GoBoy/pkg/types"

type RAM struct {
	// max 16
	Bank uint8
	// [bank][idx] idx is 0x8000
	Buf  [][]byte
	data []byte
	Size int
}

func NewRAM(size int) *RAM {
	data := make([]byte, size)
	return &RAM{
		data: data,
		Size: size,
	}
}

func (r *RAM) Read(addr types.Addr) byte {
	return r.data[addr]
}

func (r *RAM) Write(addr types.Addr, value byte) {
	r.data[addr] = value
}
