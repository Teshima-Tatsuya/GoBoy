package bus

import "github.com/Teshima-Tatsuya/GoBoy/pkg/types"

type IO interface {
	ReadByte(addr types.Addr) byte
	WriteByte(addr types.Addr, value byte)
}
