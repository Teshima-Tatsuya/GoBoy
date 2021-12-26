package cpu

const (
	A = iota
	B
	C
	D
	E
	H
	L
	F // Flag Register
)

const (
	flagZ = 7
	flagN = 6
	flagH = 5
	flagC = 4
)

type Register struct {
	R  [8]byte
	SP uint16
	PC uint16
}

func (r *Register) reset() {
	r.R[A] = 0x01
	r.R[B] = 0x00
	r.R[C] = 0x13
	r.R[D] = 0x00
	r.R[E] = 0xD8
	r.R[H] = 0x01
	r.R[F] = 0x4D
	r.PC = 0x100 // Gameboy Start Addr
}
