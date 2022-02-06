package mock

import "github.com/Teshima-Tatsuya/GoBoy/pkg/types"

type MockBus struct {
	buf [0xFFFF]byte
}

func NewMockBus() *MockBus {
	return &MockBus{}
}

func (b *MockBus) ReadByte(addr types.Addr) byte {
	return b.buf[addr]
}

// TODO: IF, IE
func (b *MockBus) WriteByte(addr types.Addr, value byte) {
	b.buf[addr] = value
}
