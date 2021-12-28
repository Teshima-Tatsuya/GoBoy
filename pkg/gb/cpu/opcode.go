package cpu

import (
	"fmt"

	"github.com/Teshima-Tatsuya/GoBoy/pkg/types"
)

type OpCode struct {
	Code    byte
	R1, R2  byte
	Size    uint8
	Cycles  uint8
	Handler func(*CPU, byte, byte)
}

var opCodes = []*OpCode{
	{0x00, 0, 0, 0, 1, nop},
	{0x01, BC, 0, 2, 3, ldr16d16},
	{0x02, BC, A, 0, 2, ldm16r},
	{0x03, BC, 0, 0, 2, inc16},
	{0x04, B, 0, 0, 1, incr},
	{0x05, 0, 0, 0, 1, notimplemented},
	{0x06, B, 0, 1, 2, ldr8d8},
	{0x07, 0, 0, 0, 1, notimplemented},
	{0x08, 0, 0, 0, 1, notimplemented},
	{0x09, 0, 0, 0, 1, notimplemented},
	{0x0A, A, BC, 0, 2, ldrm16},
	{0x0B, 0, 0, 0, 1, notimplemented},
	{0x0C, C, 0, 1, 2, incr},
	{0x0D, 0, 0, 0, 1, notimplemented},
	{0x0E, C, 0, 1, 2, ldr8d8},
	{0x0F, 0, 0, 0, 1, notimplemented},
	{0x10, 0, 0, 0, 1, notimplemented},
	{0x11, DE, 0, 2, 3, ldr16d16},
	{0x12, DE, A, 0, 2, ldm16r8},
	{0x13, DE, 0, 0, 2, inc16},
	{0x14, D, 0, 0, 1, incr},
	{0x15, 0, 0, 0, 1, notimplemented},
	{0x16, D, 0, 1, 2, ldr8d8},
	{0x17, 0, 0, 0, 1, notimplemented},
	{0x18, 0, 0, 0, 1, notimplemented},
	{0x19, 0, 0, 0, 1, notimplemented},
	{0x1A, A, DE, 0, 2, ldrm16},
	{0x1B, 0, 0, 0, 1, notimplemented},
	{0x1C, E, 0, 0, 1, incr},
	{0x1D, 0, 0, 0, 1, notimplemented},
	{0x1E, E, 0, 1, 2, ldr8d8},
	{0x1F, 0, 0, 0, 1, notimplemented},
	{0x20, flagZ, 0, 1, 2, jpnfr8},
	{0x21, HL, 0, 2, 3, ldr16d16},
	{0x22, HLI, A, 0, 2, ldm16r},
	{0x23, HL, 0, 0, 2, inc16},
	{0x24, H, 0, 0, 1, incr},
	{0x25, 0, 0, 0, 1, notimplemented},
	{0x26, H, 0, 1, 2, ldr8d8},
	{0x27, 0, 0, 0, 1, notimplemented},
	{0x28, 0, 0, 0, 1, notimplemented},
	{0x29, 0, 0, 0, 1, notimplemented},
	{0x2A, A, HLI, 0, 2, ldrm16},
	{0x2B, 0, 0, 0, 1, notimplemented},
	{0x2C, L, 0, 0, 1, incr},
	{0x2D, 0, 0, 0, 1, notimplemented},
	{0x2E, L, 0, 0, 1, ldr8d8},
	{0x2F, 0, 0, 0, 1, notimplemented},
	{0x30, flagC, 0, 1, 2, jpnfr8},
	{0x31, SP, 0, 2, 3, ldr16d16},
	{0x32, HLD, A, 0, 2, ldm16r},
	{0x33, SP, 0, 0, 2, inc16},
	{0x34, 0, 0, 0, 1, notimplemented},
	{0x35, 0, 0, 0, 1, notimplemented},
	{0x36, HL, 0, 1, 3, ldr16d16},
	{0x37, 0, 0, 0, 1, notimplemented},
	{0x38, 0, 0, 0, 1, notimplemented},
	{0x39, 0, 0, 0, 1, notimplemented},
	{0x3A, A, HLD, 0, 2, ldrm16},
	{0x3B, 0, 0, 0, 1, notimplemented},
	{0x3C, A, 0, 0, 1, incr},
	{0x3D, 0, 0, 0, 1, notimplemented},
	{0x3E, A, 0, 1, 2, ldr8d8},
	{0x3F, 0, 0, 0, 1, notimplemented},
	{0x40, B, B, 0, 1, ldrr},
	{0x41, B, C, 0, 1, ldrr},
	{0x42, B, D, 0, 1, ldrr},
	{0x43, B, E, 0, 1, ldrr},
	{0x44, B, H, 0, 1, ldrr},
	{0x45, B, L, 0, 1, ldrr},
	{0x46, B, HL, 0, 2, ldrm16},
	{0x47, B, A, 0, 1, ldrr},
	{0x48, C, B, 0, 1, ldrr},
	{0x49, C, C, 0, 1, ldrr},
	{0x4A, C, D, 0, 1, ldrr},
	{0x4B, C, E, 0, 1, ldrr},
	{0x4C, C, H, 0, 1, ldrr},
	{0x4D, C, L, 0, 1, ldrr},
	{0x4E, C, HL, 0, 2, ldrm16},
	{0x4F, C, A, 0, 1, ldrr},
	{0x50, D, B, 0, 1, ldrr},
	{0x51, D, C, 0, 1, ldrr},
	{0x52, D, D, 0, 1, ldrr},
	{0x53, D, E, 0, 1, ldrr},
	{0x54, D, H, 0, 1, ldrr},
	{0x55, D, L, 0, 1, ldrr},
	{0x56, D, HL, 0, 2, ldrm16},
	{0x57, D, A, 0, 1, ldrr},
	{0x58, E, B, 0, 1, ldrr},
	{0x59, E, C, 0, 1, ldrr},
	{0x5A, E, D, 0, 1, ldrr},
	{0x5B, E, E, 0, 1, ldrr},
	{0x5C, E, H, 0, 1, ldrr},
	{0x5D, E, L, 0, 1, ldrr},
	{0x5E, E, HL, 0, 2, ldrm16},
	{0x5F, E, A, 0, 1, ldrr},
	{0x60, H, B, 0, 1, ldrr},
	{0x61, H, C, 0, 1, ldrr},
	{0x62, H, D, 0, 1, ldrr},
	{0x63, H, E, 0, 1, ldrr},
	{0x64, H, H, 0, 1, ldrr},
	{0x65, H, L, 0, 1, ldrr},
	{0x66, H, HL, 0, 2, ldrm16},
	{0x67, H, A, 0, 1, ldrr},
	{0x68, L, B, 0, 1, ldrr},
	{0x69, L, C, 0, 1, ldrr},
	{0x6A, L, D, 0, 1, ldrr},
	{0x6B, L, E, 0, 1, ldrr},
	{0x6C, L, H, 0, 1, ldrr},
	{0x6D, L, L, 0, 1, ldrr},
	{0x6E, L, HL, 0, 2, ldrm16},
	{0x6F, L, A, 0, 1, ldrr},
	{0x70, HL, B, 0, 2, ldm16r},
	{0x71, HL, C, 0, 2, ldm16r},
	{0x72, HL, D, 0, 2, ldm16r},
	{0x73, HL, E, 0, 2, ldm16r},
	{0x74, HL, H, 0, 2, ldm16r},
	{0x75, HL, L, 0, 2, ldm16r},
	{0x76, 0, 0, 0, 1, notimplemented},
	{0x77, HL, A, 0, 2, ldm16r},
	{0x78, A, B, 0, 1, ldrr},
	{0x79, A, C, 0, 1, ldrr},
	{0x7A, A, D, 0, 1, ldrr},
	{0x7B, A, E, 0, 1, ldrr},
	{0x7C, A, H, 0, 1, ldrr},
	{0x7D, A, L, 0, 1, ldrr},
	{0x7E, A, HL, 0, 2, ldrm16},
	{0x7F, A, A, 0, 1, ldrr},
	{0x80, 0, 0, 0, 1, notimplemented},
	{0x81, 0, 0, 0, 1, notimplemented},
	{0x82, 0, 0, 0, 1, notimplemented},
	{0x83, 0, 0, 0, 1, notimplemented},
	{0x84, 0, 0, 0, 1, notimplemented},
	{0x85, 0, 0, 0, 1, notimplemented},
	{0x86, 0, 0, 0, 1, notimplemented},
	{0x87, 0, 0, 0, 1, notimplemented},
	{0x88, 0, 0, 0, 1, notimplemented},
	{0x89, 0, 0, 0, 1, notimplemented},
	{0x8A, 0, 0, 0, 1, notimplemented},
	{0x8B, 0, 0, 0, 1, notimplemented},
	{0x8C, 0, 0, 0, 1, notimplemented},
	{0x8D, 0, 0, 0, 1, notimplemented},
	{0x8E, 0, 0, 0, 1, notimplemented},
	{0x8F, 0, 0, 0, 1, notimplemented},
	{0x90, 0, 0, 0, 1, notimplemented},
	{0x91, 0, 0, 0, 1, notimplemented},
	{0x92, 0, 0, 0, 1, notimplemented},
	{0x93, 0, 0, 0, 1, notimplemented},
	{0x94, 0, 0, 0, 1, notimplemented},
	{0x95, 0, 0, 0, 1, notimplemented},
	{0x96, 0, 0, 0, 1, notimplemented},
	{0x97, 0, 0, 0, 1, notimplemented},
	{0x98, 0, 0, 0, 1, notimplemented},
	{0x99, 0, 0, 0, 1, notimplemented},
	{0x9A, 0, 0, 0, 1, notimplemented},
	{0x9B, 0, 0, 0, 1, notimplemented},
	{0x9C, 0, 0, 0, 1, notimplemented},
	{0x9D, 0, 0, 0, 1, notimplemented},
	{0x9E, 0, 0, 0, 1, notimplemented},
	{0x9F, 0, 0, 0, 1, notimplemented},
	{0xA0, 0, 0, 0, 1, notimplemented},
	{0xA1, 0, 0, 0, 1, notimplemented},
	{0xA2, 0, 0, 0, 1, notimplemented},
	{0xA3, 0, 0, 0, 1, notimplemented},
	{0xA4, 0, 0, 0, 1, notimplemented},
	{0xA5, 0, 0, 0, 1, notimplemented},
	{0xA6, 0, 0, 0, 1, notimplemented},
	{0xA7, 0, 0, 0, 1, notimplemented},
	{0xA8, 0, 0, 0, 1, notimplemented},
	{0xA9, 0, 0, 0, 1, notimplemented},
	{0xAA, 0, 0, 0, 1, notimplemented},
	{0xAB, 0, 0, 0, 1, notimplemented},
	{0xAC, 0, 0, 0, 1, notimplemented},
	{0xAD, 0, 0, 0, 1, notimplemented},
	{0xAE, 0, 0, 0, 1, notimplemented},
	{0xAF, 0, 0, 0, 1, notimplemented},
	{0xB0, 0, 0, 0, 1, notimplemented},
	{0xB1, 0, 0, 0, 1, notimplemented},
	{0xB2, 0, 0, 0, 1, notimplemented},
	{0xB3, 0, 0, 0, 1, notimplemented},
	{0xB4, 0, 0, 0, 1, notimplemented},
	{0xB5, 0, 0, 0, 1, notimplemented},
	{0xB6, 0, 0, 0, 1, notimplemented},
	{0xB7, 0, 0, 0, 1, notimplemented},
	{0xB8, 0, 0, 0, 1, notimplemented},
	{0xB9, 0, 0, 0, 1, notimplemented},
	{0xBA, 0, 0, 0, 1, notimplemented},
	{0xBB, 0, 0, 0, 1, notimplemented},
	{0xBC, 0, 0, 0, 1, notimplemented},
	{0xBD, 0, 0, 0, 1, notimplemented},
	{0xBE, 0, 0, 0, 1, notimplemented},
	{0xBF, 0, 0, 0, 1, notimplemented},
	{0xC0, flagZ, 0, 2, 2, retncc},
	{0xC1, 0, 0, 0, 1, notimplemented},
	{0xC2, flagZ, 0, 2, 3, jpnfa16},
	{0xC3, 0, 0, 2, 4, jpa16},
	{0xC4, 0, 0, 0, 1, notimplemented},
	{0xC5, 0, 0, 0, 1, notimplemented},
	{0xC6, 0, 0, 0, 1, notimplemented},
	{0xC7, 0, 0, 0, 1, notimplemented},
	{0xC8, 0, 0, 0, 1, notimplemented},
	{0xC9, 0, 0, 0, 4, ret},
	{0xCA, flagZ, 0, 2, 3, jpfa16},
	{0xCB, 0, 0, 0, 1, notimplemented},
	{0xCC, 0, 0, 0, 1, notimplemented},
	{0xCD, 0, 0, 0, 1, notimplemented},
	{0xCE, 0, 0, 0, 1, notimplemented},
	{0xCF, 0, 0, 0, 1, notimplemented},
	{0xD0, 0, 0, 0, 1, notimplemented},
	{0xD1, 0, 0, 0, 1, notimplemented},
	{0xD2, flagC, 0, 2, 3, jpnfa16},
	{0xD3, 0, 0, 0, 1, notimplemented},
	{0xD4, 0, 0, 0, 1, notimplemented},
	{0xD5, 0, 0, 0, 1, notimplemented},
	{0xD6, 0, 0, 0, 1, notimplemented},
	{0xD7, 0, 0, 0, 1, notimplemented},
	{0xD8, 0, 0, 0, 1, notimplemented},
	{0xD9, 0, 0, 0, 1, notimplemented},
	{0xDA, flagC, 0, 2, 3, jpfa16},
	{0xDB, 0, 0, 0, 1, notimplemented},
	{0xDC, 0, 0, 0, 1, notimplemented},
	{0xDD, 0, 0, 0, 1, notimplemented},
	{0xDE, 0, 0, 0, 1, notimplemented},
	{0xDF, 0, 0, 0, 1, notimplemented},
	{0xE0, 0, 0, 0, 1, notimplemented},
	{0xE1, 0, 0, 0, 1, notimplemented},
	{0xE2, 0, 0, 0, 1, notimplemented},
	{0xE3, 0, 0, 0, 1, notimplemented},
	{0xE4, 0, 0, 0, 1, notimplemented},
	{0xE5, 0, 0, 0, 1, notimplemented},
	{0xE6, 0, 0, 0, 1, notimplemented},
	{0xE7, 0, 0, 0, 1, notimplemented},
	{0xE8, 0, 0, 0, 1, notimplemented},
	{0xE9, HL, 0, 0, 1, jpm16},
	{0xEA, 0, 0, 0, 1, notimplemented},
	{0xEB, 0, 0, 0, 1, notimplemented},
	{0xEC, 0, 0, 0, 1, notimplemented},
	{0xED, 0, 0, 0, 1, notimplemented},
	{0xEE, 0, 0, 0, 1, notimplemented},
	{0xEF, 0, 0, 0, 1, notimplemented},
	{0xF0, 0, 0, 0, 1, notimplemented},
	{0xF1, 0, 0, 0, 1, notimplemented},
	{0xF2, 0, 0, 0, 1, notimplemented},
	{0xF3, 0, 0, 0, 1, notimplemented},
	{0xF4, 0, 0, 0, 1, notimplemented},
	{0xF5, 0, 0, 0, 1, notimplemented},
	{0xF6, 0, 0, 0, 1, notimplemented},
	{0xF7, 0, 0, 0, 1, notimplemented},
	{0xF8, 0, 0, 0, 1, notimplemented},
	{0xF9, 0, 0, 0, 1, notimplemented},
	{0xFA, 0, 0, 0, 1, notimplemented},
	{0xFB, 0, 0, 0, 1, notimplemented},
	{0xFC, 0, 0, 0, 1, notimplemented},
	{0xFD, 0, 0, 0, 1, notimplemented},
	{0xFE, 0, 0, 0, 1, notimplemented},
	{0xFF, 0, 0, 0, 1, notimplemented},
}

func nop(c *CPU, _ byte, _ byte) {}

// -----LD----
// r   is Register Single
// r16 is Register Comprex
// m   is Read From Register Single
// m16 is Read From Register Complex
// d   is 8 bit data
// d16 is 16 bit data
// ex:
//    ldrr   is LD r8, r8
//    ldrr16 is LD r8, r16
//    ldrm   is LD r8, Read(r8)
//    ldrm16 is LD r8, Read(r16)

// LD R1, R2
// Write R2 into R1
func ldrr(c *CPU, R1 byte, R2 byte) {
	c.Reg.R[R1] = c.Reg.R[R2]
}

// LD (r1), r2
func ldm16r(c *CPU, R1 byte, R2 byte) {
	c.Bus.WriteByte(c.Reg.R16(int(R1)), c.Reg.R[R2])
}

// LD r1, (r2)
// Write r2 value into r1
func ldrm16(c *CPU, R1 byte, R2 byte) {
	c.Reg.R[R1] = c.Bus.ReadByte(c.Reg.R16(int(R2)))
}

// LD r1, d16
func ldr16d16(c *CPU, R1 byte, _ byte) {
	c.Reg.setR16(types.Addr(R1), c.fetch16())
}

// LD r1, d8
func ldr8d8(c *CPU, r8 byte, _ byte) {
	c.Reg.R[r8] = c.fetch()
}

func ldm16r8(c *CPU, R1 byte, R2 byte) {
	c.Bus.WriteByte(c.Reg.R16(int(R1)), c.Reg.R[R2])
}

func retcc(c *CPU, R1 byte, _ byte) {
	if c.Reg.R[F]&(1<<R1) != 0 {
		c.pop2PC()
	}
}

func retncc(c *CPU, R1 byte, _ byte) {
	if c.Reg.R[F]&(1<<R1) == 0 {
		c.pop2PC()
	}
}

// arithmetic
func incr(c *CPU, r8 byte, _ byte) {
	r := c.Reg.R[r8]

	incremented := r + 0x01
	c.Reg.R[r8] = incremented
	c.Reg.clearFlag(flagN) // not subtract
	if incremented == 0 {
		c.Reg.setFlag(flagZ)
	} else {
		c.Reg.clearFlag(flagZ)
	}

	// Harf Carry
	if (incremented^0x01^r8)&0x10 == 0x10 {
		c.Reg.setFlag(flagH)
	} else {
		c.Reg.clearFlag(flagH)
	}
}

func inc16(c *CPU, r16 byte, _ byte) {
	c.Reg.setR16(types.Addr(r16), types.Addr(c.Reg.R16(int(r16))+1))
}

// special
func ret(c *CPU, _ byte, _ byte) {
	c.pop2PC()
}

// -----jp-----

func _jp(c *CPU, addr types.Addr) {
	c.Reg.PC = addr
}

// JP a16
func jpa16(c *CPU, _ byte, _ byte) {
	_jp(c, c.fetch16())
}

func jpcc(c *CPU, cc byte, _ byte) {
	c.Reg.PC = c.fetch16()

}

// JP flag, a16
// jump when flag = 1
func jpfa16(c *CPU, flag byte, _ byte) {
	if c.Reg.isSet(flag) {
		_jp(c, c.fetch16())
	}
}

// JP Nflag, a16
// jump when flag = 0
func jpnfa16(c *CPU, flag byte, _ byte) {
	if !c.Reg.isSet(flag) {
		_jp(c, c.fetch16())
	}
}

// JP (r16)
func jpm16(c *CPU, R1 byte, _ byte) {
	_jp(c, types.Addr(c.Bus.ReadByte(c.Reg.R16(int(R1)))))
}

// -----jr-----
func _jr(c *CPU, addr byte) {
	// -1 get rid of fetch()
	c.Reg.PC = types.Addr(int32(c.Reg.PC) + int32(addr) - 1)
}

// r8 is a signed data, which are added to PC
func jpnfr8(c *CPU, flag byte, _ byte) {
	if !c.Reg.isSet(flag) {
		_jr(c, c.fetch())
	}
}

func notimplemented(c *CPU, _ byte, _ byte) {
	c.Reg.PC--
	panic(fmt.Sprintf("OpCode 0x%2x is not implemented", c.fetch()))
}
