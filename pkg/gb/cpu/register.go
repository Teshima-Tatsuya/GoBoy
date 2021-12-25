package cpu

type Register = byte

type Registers struct {
	A Register
	B Register
	C Register
	D Register
	E Register
	H Register
	L Register
	F Register // Flag Register
}

type bit bool

type Flag struct {
	Z bit
	N bit
	H bit
	C bit
}
