package cpu

import (
	"fmt"

	"github.com/Teshima-Tatsuya/GoBoy/pkg/types"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/util"
	"github.com/apex/log"
)

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
	AF = iota + 8 // To prevent collide register index
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
	r.R[B] = 0xFF
	r.R[C] = 0x13
	r.R[D] = 0x00
	r.R[E] = 0xC1
	r.R[H] = 0x84
	r.R[L] = 0x03
	r.R[F] = 0x00
	r.PC = 0x0100
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

func (r *Register) setF(flag int, b bool) {
	if b {
		r.R[F] |= (1 << uint(flag))
	} else {
		r.R[F] &= ^(1 << uint(flag))
	}
}

func (r *Register) setNH(n, h bool) {
	r.setF(flagN, n)
	r.setF(flagH, h)
}

func (r *Register) setZNH(z, n, h bool) {
	r.setNH(n, h)
	r.setF(flagZ, z)
}

func (r *Register) setNHC(n, h, c bool) {
	r.setNH(n, h)
	r.setF(flagC, c)
}

func (r *Register) setZNHC(z, n, h, c bool) {
	r.setNHC(n, h, c)
	r.setF(flagZ, z)
}

func (r *Register) isSet(flag int) bool {
	return r.R[F]&byte(1<<uint(flag)) == 1<<uint(flag)
}

func (r *Register) Dump() {
	log.Info(fmt.Sprintf("    Regs  A:%02X B:%02X C:%02X D:%02X E:%02X H:%02X L:%02X", r.R[A], r.R[B], r.R[C], r.R[D], r.R[E], r.R[H], r.R[L]))
	log.Info(fmt.Sprintf("    Flags Z:%d N:%d H:%d C:%d", util.Bool2Int8(r.isSet(flagZ)), util.Bool2Int8(r.isSet(flagN)), util.Bool2Int8(r.isSet(flagH)), util.Bool2Int8(r.isSet(flagC))))
}
