package bus

type IO interface {
	ReadByte(addr uint16) byte
}
