package interfaces

import "github.com/Teshima-Tatsuya/GoBoy/pkg/types"

type ByteReader interface {
	ReadByte(addr types.Addr) byte
}

type ByteWriter interface {
	WriteByte(addr types.Addr, value byte)
}

type Bus interface {
	ByteReader
	ByteWriter
}
