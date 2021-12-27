package bus

import "github.com/Teshima-Tatsuya/GoBoy/pkg/types"

type IO interface {
	ReadByte(addr types.Addr) byte
	ReadAddr(addr types.Addr) types.Addr
	WriteByte(addr types.Addr, value byte)
}
