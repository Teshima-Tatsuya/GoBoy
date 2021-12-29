package cpu

import "github.com/Teshima-Tatsuya/GoBoy/pkg/types"

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

// @see https://www.pastraiser.com/cpu/gameboy/gameboy_opcodes.html
const (
	AF = iota
	BC
	DE
	HL
	HLD
	HLI
	SP
	PC
)

const (
	flagZ = 7
	flagN = 6
	flagH = 5
	flagC = 4
)

type Register struct {
	R  [8]byte
	SP types.Addr
	PC types.Addr
}

func (r *Register) reset() {
	r.R[A] = 0x01
	r.R[B] = 0x00
	r.R[C] = 0x13
	r.R[D] = 0x00
	r.R[E] = 0xD8
	r.R[H] = 0x01
	r.R[L] = 0x4D
	r.R[F] = 0xB0
	r.PC = 0x0100 // Gameboy Start Addr
	r.SP = 0xfffe
}

func (r *Register) R16(i int) types.Addr {
	switch i {
	case AF:
		return r.AF()
	case BC:
		return r.BC()
	case DE:
		return r.DE()
	case HL:
		return r.HL()
	case HLD:
		hl := r.HL()
		r.setHL(hl - 1)
		return hl
	case HLI:
		hl := r.HL()
		r.setHL(hl + 1)
		return hl
	case SP:
		return r.SP
	case PC:
		return r.PC
	default:
		panic("Invalid Register!")
	}
}

func (r *Register) AF() types.Addr {
	return (types.Addr(r.R[A]) << 8) | types.Addr(r.R[F])
}

func (r *Register) BC() types.Addr {
	return (types.Addr(r.R[B]) << 8) | types.Addr(r.R[C])
}

func (r *Register) DE() types.Addr {
	return (types.Addr(r.R[D]) << 8) | types.Addr(r.R[E])
}

func (r *Register) HL() types.Addr {
	return (types.Addr(r.R[H]) << 8) | types.Addr(r.R[L])
}

func (r *Register) setR16(R int, value types.Addr) {
	switch R {
	case AF:
		r.setAF(value)
	case BC:
		r.setBC(value)
	case DE:
		r.setDE(value)
	case HL:
		r.setHL(value)
	case SP:
		r.SP = value
	case PC:
		r.PC = value
	default:
		panic("Unknown Register 16")
	}
}

func (r *Register) setAF(value types.Addr) {
	r.R[A], r.R[F] = byte(value>>8), byte(value)
}

func (r *Register) setBC(value types.Addr) {
	r.R[B], r.R[C] = byte(value>>8), byte(value)
}

func (r *Register) setDE(value types.Addr) {
	r.R[D], r.R[E] = byte(value>>8), byte(value)
}

func (r *Register) setHL(value types.Addr) {
	r.R[H], r.R[L] = byte(value>>8), byte(value)
}

func (r *Register) setFlagH(v byte) {
	if r.R[A]&0x0F < v&0xF {
		r.setFlag(flagH)
	} else {
		r.clearFlag(flagH)
	}
}

func (r *Register) setFlagZ(v byte) {
	if v == 0 {
		r.setFlag(flagZ)
	} else {
		r.clearFlag(flagZ)
	}
}

func (r *Register) setFlag(flag byte) {
	switch flag {
	case flagZ:
		r.R[F] = r.R[F] | (1 << uint(flagZ))
	case flagN:
		r.R[F] = r.R[F] | (1 << uint(flagN))
	case flagH:
		r.R[F] = r.R[F] | (1 << uint(flagH))
	case flagC:
		r.R[F] = r.R[F] | (1 << uint(flagC))
	}
}

func (r *Register) clearFlag(flag byte) {
	switch flag {
	case flagZ:
		r.R[F] = r.R[F] & ^byte((uint(1 << uint(flagZ))))
	case flagN:
		r.R[F] = r.R[F] & ^byte((uint(1 << uint(flagN))))
	case flagH:
		r.R[F] = r.R[F] & ^byte((uint(1 << uint(flagH))))
	case flagC:
		r.R[F] = r.R[F] & ^byte((uint(1 << uint(flagC))))
	}
}

func (r *Register) isSet(flag byte) bool {
	switch flag {
	case flagZ:
		return r.R[F]&byte(1<<uint(flagZ)) == 1<<uint(flagZ)
	case flagN:
		return r.R[F]&byte(1<<uint(flagN)) == 1<<uint(flagN)
	case flagH:
		return r.R[F]&byte(1<<uint(flagH)) == 1<<uint(flagH)
	case flagC:
		return r.R[F]&byte(1<<uint(flagC)) == 1<<uint(flagC)
	default:
		panic("Unknown Flag")
	}
}
