package io

import "github.com/Teshima-Tatsuya/GoBoy/pkg/types"

type IO struct {
	buf []byte
}

func New(size int) *IO {
	buf := make([]byte, size)

	return &IO{
		buf: buf,
	}

}

func (r *IO) Read(addr types.Addr) byte {
	return r.buf[addr]
}

func (r *IO) Write(addr types.Addr, value byte) {
	r.buf[addr] = value
}
