package ram

import "github.com/Teshima-Tatsuya/GoBoy/pkg/types"

type RAM struct {
	data []byte
}

func New(size int) *RAM {
	data := make([]byte, size)
	return &RAM{
		data: data,
	}
}

func (r *RAM) Read(addr types.Addr) byte {
	return r.data[addr]
}

func (r *RAM) Write(addr types.Addr, value byte) {
	r.data[addr] = value
}
