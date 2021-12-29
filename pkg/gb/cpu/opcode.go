package cpu

import (
	"fmt"

	"github.com/Teshima-Tatsuya/GoBoy/pkg/types"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/util"
	"github.com/apex/log"
)

type OpCode struct {
	Code     byte
	Mnemonic string
	R1, R2   byte
	Size     uint8
	Cycles   uint8
	Handler  func(*CPU, byte, byte)
}

var opCodes = []*OpCode{
	{0x00, "NOP", 0, 0, 0, 1, nop},
	{0x01, "LD BC,d16", BC, 0, 2, 3, ldr16d16},
	{0x02, "LD (BC),A", BC, A, 0, 2, ldm16r},
	{0x03, "INC BC", BC, 0, 0, 2, incr16},
	{0x04, "INC B", B, 0, 0, 1, incr},
	{0x05, "DEC B", B, 0, 0, 1, decr},
	{0x06, "LD B,d8", B, 0, 1, 2, ldrd},
	{0x07, "RLCA", 0, 0, 0, 1, notimplemented},
	{0x08, "LD (a16),SP", 0, SP, 2, 5, lda16r16},
	{0x09, "ADD HL,BC", 0, 0, 0, 1, notimplemented},
	{0x0A, "LD A,(BC)", A, BC, 0, 2, ldrm16},
	{0x0B, "DEC BC", BC, 0, 0, 2, decr16},
	{0x0C, "INC C", C, 0, 1, 2, incr},
	{0x0D, "DEC C", C, 0, 0, 1, decr},
	{0x0E, "LD C,d8", C, 0, 1, 2, ldrd},
	{0x0F, "RRCA", 0, 0, 0, 1, notimplemented},
	{0x10, "STOP 0", 0, 0, 0, 1, stop},
	{0x11, "LD DE,d16", DE, 0, 2, 3, ldr16d16},
	{0x12, "LD (DE),A", DE, A, 0, 2, ldm16r},
	{0x13, "INC DE", DE, 0, 0, 2, incr16},
	{0x14, "INC D", D, 0, 0, 1, incr},
	{0x15, "DEC D", D, 0, 0, 1, decr},
	{0x16, "LD D,d8", D, 0, 1, 2, ldrd},
	{0x17, "RLA", 0, 0, 0, 1, notimplemented},
	{0x18, "JR r8", 0, 0, 1, 3, jrr8},
	{0x19, "ADD HL,DE", 0, 0, 0, 1, notimplemented},
	{0x1A, "LD A,(DE)", A, DE, 0, 2, ldrm16},
	{0x1B, "DEC DE", DE, 0, 0, 2, decr16},
	{0x1C, "INC E", E, 0, 0, 1, incr},
	{0x1D, "DEC E", E, 0, 0, 1, decr},
	{0x1E, "LD E,d8", E, 0, 1, 2, ldrd},
	{0x1F, "RRA", 0, 0, 0, 1, notimplemented},
	{0x20, "JR NZ,r8", flagZ, 0, 1, 2, jrnfr8},
	{0x21, "LD HL,d16", HL, 0, 2, 3, ldr16d16},
	{0x22, "LD (HL+),A", HLI, A, 0, 2, ldm16r},
	{0x23, "INC HL", HL, 0, 0, 2, incr16},
	{0x24, "INC H", H, 0, 0, 1, incr},
	{0x25, "DEC H", H, 0, 0, 1, decr},
	{0x26, "LD H,d8", H, 0, 1, 2, ldrd},
	{0x27, "DAA", 0, 0, 0, 1, notimplemented},
	{0x28, "JR Z,r8", flagZ, 0, 1, 2, jrfr8},
	{0x29, "ADD HL,HL", 0, 0, 0, 1, notimplemented},
	{0x2A, "LD A,(HL+)", A, HLI, 0, 2, ldrm16},
	{0x2B, "DEC HL", HL, 0, 0, 1, decr16},
	{0x2C, "INC L", L, 0, 0, 1, incr},
	{0x2D, "DEC L", L, 0, 0, 1, decr},
	{0x2E, "LD L,d8", L, 0, 0, 1, ldrd},
	{0x2F, "CPL", 0, 0, 0, 1, notimplemented},
	{0x30, "JR NC,r8", flagC, 0, 1, 2, jrnfr8},
	{0x31, "LD SP,d16", SP, 0, 2, 3, ldr16d16},
	{0x32, "LD (HL-),A", HLD, A, 0, 2, ldm16r},
	{0x33, "INC SP", SP, 0, 0, 2, incr16},
	{0x34, "INC (HL)", HL, 0, 0, 3, incm16},
	{0x35, "DEC (HL)", HL, 0, 0, 3, decm16},
	{0x36, "LD (HL),d8", HL, 0, 1, 3, ldm16d},
	{0x37, "SCF", 0, 0, 0, 1, notimplemented},
	{0x38, "JR C,r8", flagC, 0, 1, 2, jrfr8},
	{0x39, "ADD HL,SP", 0, 0, 0, 1, notimplemented},
	{0x3A, "LD A,(HL-)", A, HLD, 0, 2, ldrm16},
	{0x3B, "DEC SP", SP, 0, 0, 1, decr16},
	{0x3C, "INC A", A, 0, 0, 1, incr},
	{0x3D, "DEC A", A, 0, 0, 1, decr},
	{0x3E, "LD A,d8", A, 0, 1, 2, ldrd},
	{0x3F, "CCF", 0, 0, 0, 1, notimplemented},
	{0x40, "LD B, B", B, B, 0, 1, ldrr},
	{0x41, "LD B, C", B, C, 0, 1, ldrr},
	{0x42, "LD B, D", B, D, 0, 1, ldrr},
	{0x43, "LD B, E", B, E, 0, 1, ldrr},
	{0x44, "LD B, H", B, H, 0, 1, ldrr},
	{0x45, "LD B, L", B, L, 0, 1, ldrr},
	{0x46, "LD B, HL", B, HL, 0, 2, ldrm16},
	{0x47, "LD B, A", B, A, 0, 1, ldrr},
	{0x48, "LD C, B", C, B, 0, 1, ldrr},
	{0x49, "LD C, C", C, C, 0, 1, ldrr},
	{0x4A, "LD C, D", C, D, 0, 1, ldrr},
	{0x4B, "LD C, E", C, E, 0, 1, ldrr},
	{0x4C, "LD C, H", C, H, 0, 1, ldrr},
	{0x4D, "LD C, L", C, L, 0, 1, ldrr},
	{0x4E, "LD C, HL", C, HL, 0, 2, ldrm16},
	{0x4F, "LD C, A", C, A, 0, 1, ldrr},
	{0x50, "LD D, B", D, B, 0, 1, ldrr},
	{0x51, "LD D, C", D, C, 0, 1, ldrr},
	{0x52, "LD D, D", D, D, 0, 1, ldrr},
	{0x53, "LD D, E", D, E, 0, 1, ldrr},
	{0x54, "LD D, H", D, H, 0, 1, ldrr},
	{0x55, "LD D, L", D, L, 0, 1, ldrr},
	{0x56, "LD D, HL", D, HL, 0, 2, ldrm16},
	{0x57, "LD D, A", D, A, 0, 1, ldrr},
	{0x58, "LD E, B", E, B, 0, 1, ldrr},
	{0x59, "LD E, C", E, C, 0, 1, ldrr},
	{0x5A, "LD E, D", E, D, 0, 1, ldrr},
	{0x5B, "LD E, E", E, E, 0, 1, ldrr},
	{0x5C, "LD E, H", E, H, 0, 1, ldrr},
	{0x5D, "LD E, L", E, L, 0, 1, ldrr},
	{0x5E, "LD E, HL", E, HL, 0, 2, ldrm16},
	{0x5F, "LD E, A", E, A, 0, 1, ldrr},
	{0x60, "LD H, B", H, B, 0, 1, ldrr},
	{0x61, "LD H, C", H, C, 0, 1, ldrr},
	{0x62, "LD H, D", H, D, 0, 1, ldrr},
	{0x63, "LD H, E", H, E, 0, 1, ldrr},
	{0x64, "LD H, H", H, H, 0, 1, ldrr},
	{0x65, "LD H, L", H, L, 0, 1, ldrr},
	{0x66, "LD H, HL", H, HL, 0, 2, ldrm16},
	{0x67, "LD H, A", H, A, 0, 1, ldrr},
	{0x68, "LD L, B", L, B, 0, 1, ldrr},
	{0x69, "LD L, C", L, C, 0, 1, ldrr},
	{0x6A, "LD L, D", L, D, 0, 1, ldrr},
	{0x6B, "LD L, E", L, E, 0, 1, ldrr},
	{0x6C, "LD L, H", L, H, 0, 1, ldrr},
	{0x6D, "LD L, L", L, L, 0, 1, ldrr},
	{0x6E, "LD L, HL", L, HL, 0, 2, ldrm16},
	{0x6F, "LD L, A", L, A, 0, 1, ldrr},
	{0x70, "LD (HL), B", HL, B, 0, 2, ldm16r},
	{0x71, "LD (HL), C", HL, C, 0, 2, ldm16r},
	{0x72, "LD (HL), D", HL, D, 0, 2, ldm16r},
	{0x73, "LD (HL), E", HL, E, 0, 2, ldm16r},
	{0x74, "LD (HL), H", HL, H, 0, 2, ldm16r},
	{0x75, "LD (HL), L", HL, L, 0, 2, ldm16r},
	{0x76, "HALT", 0, 0, 0, 1, notimplemented},
	{0x77, "LD (HL), A", HL, A, 0, 2, ldm16r},
	{0x78, "LD A, B", A, B, 0, 1, ldrr},
	{0x79, "LD A, C", A, C, 0, 1, ldrr},
	{0x7A, "LD A, D", A, D, 0, 1, ldrr},
	{0x7B, "LD A, E", A, E, 0, 1, ldrr},
	{0x7C, "LD A, H", A, H, 0, 1, ldrr},
	{0x7D, "LD A, L", A, L, 0, 1, ldrr},
	{0x7E, "LD A, HL", A, HL, 0, 2, ldrm16},
	{0x7F, "LD A, A", A, A, 0, 1, ldrr},
	{0x80, "ADD A, B", A, B, 0, 1, addr},
	{0x81, "ADD A, C", A, C, 0, 1, addr},
	{0x82, "ADD A, D", A, D, 0, 1, addr},
	{0x83, "ADD A, E", A, E, 0, 1, addr},
	{0x84, "ADD A, H", A, H, 0, 1, addr},
	{0x85, "ADD A, L", A, L, 0, 1, addr},
	{0x86, "ADD A, (HL)", A, HL, 0, 2, addHL},
	{0x87, "ADD A, A", A, A, 0, 1, addr},
	{0x88, "ADC A, B", A, B, 0, 1, notimplemented},
	{0x89, "ADC A, C", A, C, 0, 1, notimplemented},
	{0x8A, "ADC A, D", A, D, 0, 1, notimplemented},
	{0x8B, "ADC A, E", A, E, 0, 1, notimplemented},
	{0x8C, "ADC A, H", A, H, 0, 1, notimplemented},
	{0x8D, "ADC A, L", A, L, 0, 1, notimplemented},
	{0x8E, "ADC A, (HL)", A, HL, 0, 1, notimplemented},
	{0x8F, "ADC A, A", A, A, 0, 1, notimplemented},
	{0x90, "SUB B", B, 0, 0, 1, subr},
	{0x91, "SUB C", C, 0, 0, 1, subr},
	{0x92, "SUB D", D, 0, 0, 1, subr},
	{0x93, "SUB E", E, 0, 0, 1, subr},
	{0x94, "SUB H", H, 0, 0, 1, subr},
	{0x95, "SUB L", L, 0, 0, 1, subr},
	{0x96, "SUB (HL)", HL, 0, 0, 2, subHL},
	{0x97, "SUB A", A, 0, 0, 1, subr},
	{0x98, "SBC A, B", A, B, 0, 1, notimplemented},
	{0x99, "SBC A, C", A, C, 0, 1, notimplemented},
	{0x9A, "SBC A, D", A, D, 0, 1, notimplemented},
	{0x9B, "SBC A, E", A, E, 0, 1, notimplemented},
	{0x9C, "SBC A, H", A, H, 0, 1, notimplemented},
	{0x9D, "SBC A, L", A, L, 0, 1, notimplemented},
	{0x9E, "SBC A, (HL)", A, HL, 0, 1, notimplemented},
	{0x9F, "SBC A, A", A, A, 0, 1, notimplemented},
	{0xA0, "AND B", B, 0, 0, 1, andr},
	{0xA1, "AND C", C, 0, 0, 1, andr},
	{0xA2, "AND D", D, 0, 0, 1, andr},
	{0xA3, "AND E", E, 0, 0, 1, andr},
	{0xA4, "AND H", H, 0, 0, 1, andr},
	{0xA5, "AND L", L, 0, 0, 1, andr},
	{0xA6, "AND (HL)", HL, 0, 0, 2, andHL},
	{0xA7, "AND A", A, 0, 0, 1, andr},
	{0xA8, "XOR B", B, 0, 0, 1, xorr},
	{0xA9, "XOR C", C, 0, 0, 1, xorr},
	{0xAA, "XOR D", D, 0, 0, 1, xorr},
	{0xAB, "XOR E", E, 0, 0, 1, xorr},
	{0xAC, "XOR H", H, 0, 0, 1, xorr},
	{0xAD, "XOR L", L, 0, 0, 1, xorr},
	{0xAE, "XOR (HL)", HL, 0, 0, 2, xorHL},
	{0xAF, "XOR A", A, 0, 0, 1, xorr},
	{0xB0, "OR B", B, 0, 0, 1, orr},
	{0xB1, "OR C", C, 0, 0, 1, orr},
	{0xB2, "OR D", D, 0, 0, 1, orr},
	{0xB3, "OR E", E, 0, 0, 1, orr},
	{0xB4, "OR H", H, 0, 0, 1, orr},
	{0xB5, "OR L", L, 0, 0, 1, orr},
	{0xB6, "OR (HL)", HL, 0, 0, 2, orHL},
	{0xB7, "OR A", A, 0, 0, 1, orr},
	{0xB8, "CP B", B, 0, 0, 1, cpr},
	{0xB9, "CP C", C, 0, 0, 1, cpr},
	{0xBA, "CP D", D, 0, 0, 1, cpr},
	{0xBB, "CP E", E, 0, 0, 1, cpr},
	{0xBC, "CP H", H, 0, 0, 1, cpr},
	{0xBD, "CP L", L, 0, 0, 1, cpr},
	{0xBE, "CP (HL)", HL, 0, 0, 2, cpHL},
	{0xBF, "CP A", A, 0, 0, 1, cpr},
	{0xC0, "RET NZ", flagZ, 0, 2, 2, retnf},
	{0xC1, "POP BC", BC, 0, 0, 3, pop},
	{0xC2, "JP NZ,a16", flagZ, 0, 2, 3, jpnfa16},
	{0xC3, "JP a16", 0, 0, 2, 4, jpa16},
	{0xC4, "CALL NZ,a16", flagZ, 0, 2, 3, callnf},
	{0xC5, "PUSH BC", BC, 0, 0, 4, push},
	{0xC6, "ADD A,d8", A, 0, 1, 2, addd8},
	{0xC7, "RST 00H", 0x00, 0, 0, 1, rst},
	{0xC8, "RET Z", flagZ, 0, 0, 2, retf},
	{0xC9, "RET", 0, 0, 0, 4, ret},
	{0xCA, "JP Z,a16", flagZ, 0, 2, 3, jpfa16},
	{0xCB, "PREFIX CB", 0, 0, 0, 1, notimplemented},
	{0xCC, "CALL Z,a16", flagZ, 0, 2, 3, callf},
	{0xCD, "CALL a16", 0, 0, 2, 4, call},
	{0xCE, "ADC A,d8", 0, 0, 0, 1, notimplemented},
	{0xCF, "RST 08H", 0x08, 0, 0, 1, rst},
	{0xD0, "RET NC", flagC, 0, 0, 2, retnf},
	{0xD1, "POP DE", DE, 0, 0, 3, pop},
	{0xD2, "JP NC,a16", flagC, 0, 2, 3, jpnfa16},
	{0xD3, "EMPTY", 0, 0, 0, 1, notimplemented},
	{0xD4, "CALL NC,a16", flagC, 0, 2, 3, callnf},
	{0xD5, "PUSH DE", DE, 0, 0, 1, push},
	{0xD6, "SUB d8", 0, 0, 1, 2, subd8},
	{0xD7, "RST 10H", 0x10, 0, 0, 1, rst},
	{0xD8, "RET C", flagC, 0, 0, 2, retf},
	{0xD9, "RETI", 0, 0, 0, 4, reti},
	{0xDA, "JP C,a16", flagC, 0, 2, 3, jpfa16},
	{0xDB, "EMPTY", 0, 0, 0, 1, notimplemented},
	{0xDC, "CALL C,a16", flagC, 0, 2, 3, callf},
	{0xDD, "EMPTY", 0, 0, 0, 1, notimplemented},
	{0xDE, "SBC A,d8", A, 0, 0, 1, notimplemented},
	{0xDF, "RST 18H", 0x18, 0, 0, 1, rst},
	{0xE0, "LDH (a8),A", 0, A, 0, 1, ldar},
	{0xE1, "POP HL", HL, 0, 0, 3, pop},
	{0xE2, "LD (C),A", C, A, 1, 2, ldmr},
	{0xE3, "EMPTY", 0, 0, 0, 1, notimplemented},
	{0xE4, "EMPTY", 0, 0, 0, 1, notimplemented},
	{0xE5, "PUSH HL", HL, 0, 0, 1, push},
	{0xE6, "AND d8", 0, 0, 1, 2, andd8},
	{0xE7, "RST 20H", 0x20, 0, 0, 1, rst},
	{0xE8, "ADD SP,r8", SP, 0, 0, 1, notimplemented},
	{0xE9, "JP (HL)", HL, 0, 0, 1, jpm16},
	{0xEA, "LD (a16),A", 0, A, 0, 1, lda16r},
	{0xEB, "EMPTY", 0, 0, 0, 1, notimplemented},
	{0xEC, "EMPTY", 0, 0, 0, 1, notimplemented},
	{0xED, "EMPTY", 0, 0, 0, 1, notimplemented},
	{0xEE, "XOR d8", 0, 0, 1, 2, xord8},
	{0xEF, "RST 28H", 0x28, 0, 0, 1, rst},
	{0xF0, "LDH A,(a8)", A, 0, 1, 3, ldra},
	{0xF1, "POP AF", AF, 0, 0, 3, pop},
	{0xF2, "LD A,(C)", A, C, 1, 2, ldrm},
	{0xF3, "DI", 0, 0, 0, 1, di},
	{0xF4, "EMPTY", 0, 0, 0, 1, notimplemented},
	{0xF5, "PUSH AF", AF, 0, 0, 1, push},
	{0xF6, "OR d8", 0, 0, 0, 1, ord8},
	{0xF7, "RST 30H", 0x30, 0, 0, 1, rst},
	{0xF8, "LD HL,SP+r8", HL, SP, 1, 3, ldr16r16d},
	{0xF9, "LD SP,HL", SP, HL, 0, 2, ldr16r16},
	{0xFA, "LD A,(a16)", A, 0, 2, 4, ldra16},
	{0xFB, "EI", 0, 0, 0, 1, ei},
	{0xFC, "EMPTY", 0, 0, 0, 1, notimplemented},
	{0xFD, "EMPTY", 0, 0, 0, 1, notimplemented},
	{0xFE, "CP d8", 0, 0, 1, 2, cpd8},
	{0xFF, "RST 38H", 0x38, 0, 0, 1, rst},
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

// func ldr(r, m, m16, d, a, a16)

// LD R1, R2
// Write R2 into R1
func ldrr(c *CPU, R1 byte, R2 byte) {
	c.Reg.R[R1] = c.Reg.R[R2]
}

// LD r1, (r2)
// Write r2 value into r1
func ldrm(c *CPU, R1 byte, R2 byte) {
	c.Reg.R[R1] = c.Bus.ReadByte(types.Addr(0xFF00 | types.Addr(c.Reg.R[R2])))
}

// LD r1, (r2)
// Write r2 value into r1
func ldrm16(c *CPU, R1 byte, R2 byte) {
	c.Reg.R[R1] = c.Bus.ReadByte(c.Reg.R16(int(R2)))
}

// LD r1, d8
func ldrd(c *CPU, r8 byte, _ byte) {
	c.Reg.R[r8] = c.fetch()
}

// LDH R,(a8)
func ldra(c *CPU, R1 byte, _ byte) {
	c.Reg.R[R1] = c.Bus.ReadByte(types.Addr(0xff00 | types.Addr(c.fetch())))
}

func ldra16(c *CPU, R1 byte, _ byte) {
	c.Reg.R[R1] = c.Bus.ReadByte(c.fetch16())
}

// func ldr16(r16, r16d, d16)

// LD r1, r2
func ldr16r16(c *CPU, R1 byte, R2 byte) {
	c.Reg.setR16(int(R1), c.Reg.R16(int(R2)))
}

// LD r1, r2+d
func ldr16r16d(c *CPU, R1 byte, R2 byte) {
	d := c.fetch()
	v := int32(c.Reg.R16(SP)) + int32(d) // Cast to int32 considers carry
	carryBits := uint32(c.Reg.R16(SP)) ^ uint32(d) ^ uint32(v)
	c.Reg.setR16(int(R1), c.Reg.R16(int(R2)))

	c.Reg.setFlag(flagZ)
	c.Reg.setFlag(flagN)

	if v < int32(c.Reg.SP) {
		c.Reg.setFlag(flagC)
	} else {
		c.Reg.clearFlag(flagC)
	}
	if carryBits&0x1000 == 0x1000 {
		c.Reg.setFlag(flagH)
	} else {
		c.Reg.clearFlag(flagH)
	}
}

// LD r1, d16
func ldr16d16(c *CPU, R1 byte, _ byte) {
	c.Reg.setR16(int(R1), c.fetch16())
}

// func ldm(r)

// LD (C), A
func ldmr(c *CPU, R1 byte, R2 byte) {
	addr := util.Byte2Addr(0xFF, c.Reg.R[R1])
	c.Bus.WriteByte(addr, c.Reg.R[R2])
}

// func ldm16(r, d)

// LD (r1), r2
func ldm16r(c *CPU, R1 byte, R2 byte) {
	c.Bus.WriteByte(c.Reg.R16(int(R1)), c.Reg.R[R2])
}

// LD (HL),d8
func ldm16d(c *CPU, R1 byte, R2 byte) {
	c.Bus.WriteByte(c.Reg.R16(int(R1)), c.fetch())
}

// func lda(r)

func ldar(c *CPU, _ byte, R2 byte) {
	c.Bus.WriteByte(util.Byte2Addr(0xFF, c.fetch()), c.Reg.R[R2])
}

// func lda16(r, r16)

func lda16r(c *CPU, _ byte, R2 byte) {
	c.Bus.WriteByte(c.fetch16(), c.Reg.R[R2])
}

func lda16r16(c *CPU, _ byte, R2 byte) {
	addr := c.fetch16()
	r16 := c.Reg.R16(int(R2))
	c.Bus.WriteByte(addr, util.ExtractLower(r16))
	c.Bus.WriteByte(addr+1, util.ExtractUpper(r16))
}

func retf(c *CPU, R1 byte, _ byte) {
	if c.Reg.isSet(R1) {
		c.popPC()
	}
}

func retnf(c *CPU, R1 byte, _ byte) {
	if !c.Reg.isSet(R1) {
		c.popPC()
	}
}

// arithmetic
func incr(c *CPU, r8 byte, _ byte) {
	r := c.Reg.R[r8]

	incremented := r + 0x01
	c.Reg.clearFlag(flagN) // not subtract
	if incremented == 0 {
		c.Reg.setFlag(flagZ)
	} else {
		c.Reg.clearFlag(flagZ)
	}

	// Harf Carry
	if (incremented^0x01^c.Reg.R[r8])&0x10 == 0x10 {
		c.Reg.setFlag(flagH)
	} else {
		c.Reg.clearFlag(flagH)
	}

	c.Reg.R[r8] = incremented
}

func incr16(c *CPU, r16 byte, _ byte) {
	c.Reg.setR16(int(r16), types.Addr(c.Reg.R16(int(r16))+1))
}

func incm16(c *CPU, r16 byte, _ byte) {
	d := c.Bus.ReadByte(c.Reg.R16(int(r16)))
	v := d + 1
	c.Reg.setFlagZ(v)
	c.Reg.clearFlag(flagN)
	carryBits := d ^ 1 ^ v

	if carryBits&0x10 == 0x10 {
		c.Reg.setFlag(flagH)
	} else {
		c.Reg.clearFlag(flagH)
	}

	c.Bus.WriteByte(c.Reg.R16(int(r16)), byte(v))
}

func decr(c *CPU, r8 byte, _ byte) {
	r := c.Reg.R[r8]

	decremented := r - 0x01
	c.Reg.setFlag(flagN) // subtract
	if decremented == 0 {
		c.Reg.setFlag(flagZ)
	} else {
		c.Reg.clearFlag(flagZ)
	}

	// Harf Carry
	if (decremented^0x01^c.Reg.R[r8])&0x10 == 0x10 {
		c.Reg.setFlag(flagH)
	} else {
		c.Reg.clearFlag(flagH)
	}

	c.Reg.R[r8] = decremented
}

func decr16(c *CPU, r16 byte, _ byte) {
	c.Reg.setR16(int(r16), types.Addr(c.Reg.R16(int(r16))-1))
}

func decm16(c *CPU, r16 byte, _ byte) {
	d := c.Bus.ReadByte(c.Reg.R16(int(r16)))
	v := d - 1
	c.Reg.setFlagZ(v)
	c.Reg.setFlag(flagN)
	carryBits := d ^ 1 ^ v

	if carryBits&0x10 == 0x10 {
		c.Reg.setFlag(flagH)
	} else {
		c.Reg.clearFlag(flagH)
	}

	c.Bus.WriteByte(c.Reg.R16(int(r16)), byte(v))
}

func _and(c *CPU, buf byte) {
	c.Reg.R[A] &= buf

	if c.Reg.R[A] == 0 {
		c.Reg.setFlag(flagZ)
	} else {
		c.Reg.clearFlag(flagZ)
	}
	c.Reg.clearFlag(flagN)
	c.Reg.setFlag(flagH)
	c.Reg.clearFlag(flagC)
}

func andr(c *CPU, r8 byte, _ byte) {
	buf := c.Reg.R[r8]
	_and(c, buf)
}

func andHL(c *CPU, r16 byte, _ byte) {
	buf := c.Bus.ReadByte(c.Reg.R16(int(r16)))
	_and(c, buf)
}

func andd8(c *CPU, _ byte, _ byte) {
	buf := c.fetch()
	_and(c, buf)
}

func _or(c *CPU, buf byte) {
	c.Reg.R[A] |= buf

	if c.Reg.R[A] == 0 {
		c.Reg.setFlag(flagZ)
	} else {
		c.Reg.clearFlag(flagZ)
	}
	c.Reg.clearFlag(flagN)
	c.Reg.clearFlag(flagH)
	c.Reg.clearFlag(flagC)
}

func orr(c *CPU, r8 byte, _ byte) {
	buf := c.Reg.R[r8]
	_or(c, buf)
}

func orHL(c *CPU, r16 byte, _ byte) {
	buf := c.Bus.ReadByte(c.Reg.R16(int(r16)))
	_or(c, buf)
}

func ord8(c *CPU, r8 byte, _ byte) {
	buf := c.fetch()
	_or(c, buf)
}

func _xor(c *CPU, buf byte) {
	c.Reg.R[A] ^= buf

	if c.Reg.R[A] == 0 {
		c.Reg.setFlag(flagZ)
	} else {
		c.Reg.clearFlag(flagZ)
	}
	c.Reg.clearFlag(flagN)
	c.Reg.clearFlag(flagH)
	c.Reg.clearFlag(flagC)
}

func xorr(c *CPU, r8 byte, _ byte) {
	buf := c.Reg.R[r8]
	_xor(c, buf)
}

func xorHL(c *CPU, r16 byte, _ byte) {
	buf := c.Bus.ReadByte(c.Reg.R16(int(r16)))
	_xor(c, buf)
}

func xord8(c *CPU, _ byte, _ byte) {
	buf := c.fetch()
	_xor(c, buf)
}

func _cp(c *CPU, v byte) {
	if c.Reg.R[A] == v {
		c.Reg.setFlag(flagZ)
	} else {
		c.Reg.clearFlag(flagZ)
	}
	if c.Reg.R[A] < v {
		c.Reg.setFlag(flagC)
	} else {
		c.Reg.clearFlag(flagC)
	}
	c.Reg.setFlagH(v)
	c.Reg.setFlag(flagN)
}

func cpr(c *CPU, r8 byte, _ byte) {
	v := c.Reg.R[r8]
	_cp(c, v)
}

func cpHL(c *CPU, r16 byte, _ byte) {
	v := c.Bus.ReadByte(c.Reg.R16(int(r16)))
	_cp(c, v)
}

func cpd8(c *CPU, _ byte, _ byte) {
	v := c.fetch()
	_cp(c, v)
}

func _add(c *CPU, b byte) {
	a := c.Reg.R[A]
	v := uint16(a) + uint16(b)
	carryBits := uint16(a) ^ uint16(b) ^ v
	flag_h := carryBits&(1<<4) != 0
	flag_c := carryBits&(1<<8) != 0

	c.Reg.R[A] = byte(v)
	c.Reg.clearFlag(flagN)
	c.Reg.setFlagZ(byte(v))

	if flag_h {
		c.Reg.setFlag(flagH)
	} else {
		c.Reg.clearFlag(flagH)
	}
	if flag_c {
		c.Reg.setFlag(flagC)
	} else {
		c.Reg.clearFlag(flagC)
	}
}

func addr(c *CPU, _ byte, r8 byte) {
	r := c.Reg.R[r8]
	_add(c, r)
}

func addHL(c *CPU, _ byte, r16 byte) {
	v := c.Bus.ReadByte(c.Reg.R16(int(r16)))
	_add(c, v)
}

func addd8(c *CPU, _ byte, r8 byte) {
	v := c.fetch()
	_add(c, v)
}

func _sub(c *CPU, b byte) {
	a := c.Reg.R[A]
	v := a - b
	carryBits := a ^ b ^ v
	flag_h := carryBits&(1<<4) != 0
	flag_c := a < v

	c.Reg.R[A] = v
	c.Reg.setFlag(flagN)
	c.Reg.setFlagZ(byte(v))

	if flag_h {
		c.Reg.setFlag(flagH)
	} else {
		c.Reg.clearFlag(flagH)
	}

	if flag_c {
		c.Reg.setFlag(flagC)
	} else {
		c.Reg.clearFlag(flagC)
	}
}

func subr(c *CPU, r8 byte, _ byte) {
	v := c.Reg.R[r8]
	_sub(c, v)
}

func subHL(c *CPU, r16 byte, _ byte) {
	v := c.Bus.ReadByte(c.Reg.R16(int(r16)))
	_sub(c, v)
}

func subd8(c *CPU, _ byte, _ byte) {
	r := c.fetch()
	_sub(c, r)
}

// ret
func ret(c *CPU, _ byte, _ byte) {
	c.popPC()
}

func reti(c *CPU, _ byte, _ byte) {
	c.popPC()

	// TODO irq enable
}

// -----jp-----

func _jp(c *CPU, addr types.Addr) {
	c.Reg.PC = addr
}

// JP a16
func jpa16(c *CPU, _ byte, _ byte) {
	_jp(c, c.fetch16())
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
func _jr(c *CPU, addr int8) {
	c.Reg.PC = types.Addr(int32(c.Reg.PC) + int32(addr))
}

// r8 is a signed data, which are added to PC
func jrr8(c *CPU, _ byte, _ byte) {
	n := c.fetch()
	_jr(c, int8(n))
}

// r8 is a signed data, which are added to PC
func jrfr8(c *CPU, flag byte, _ byte) {
	n := c.fetch()
	// flag is not set
	if c.Reg.isSet(flag) {
		_jr(c, int8(n))
	}
}

// r8 is a signed data, which are added to PC
func jrnfr8(c *CPU, flag byte, _ byte) {
	n := c.fetch()
	// flag is not set
	if !c.Reg.isSet(flag) {
		_jr(c, int8(n))
	}
}

// -----rst------

// RST n
// push and jump to n
func rst(c *CPU, n byte, _ byte) {
	log.Debug("TODO: implement")
	c.pushPC()
	c.Reg.PC = types.Addr(n)
}

// -----push-----
func push(c *CPU, r16 byte, _ byte) {
	buf := c.Reg.R16(int(r16))
	upper := util.ExtractUpper(types.Addr(buf))
	lower := util.ExtractLower(types.Addr(buf))
	c.push(upper)
	c.push(lower)
}

// -----pop------
func pop(c *CPU, r16 byte, _ byte) {
	var lower byte
	if r16 != AF {
		lower = c.pop()
	} else {
		lower = c.pop() & 0xf0 // extract only flag
	}

	upper := c.pop()

	c.Reg.setR16(int(r16), types.Addr(int16(upper)<<8|int16(lower)))
}

// -----call-----
func _call(c *CPU, dest types.Addr) {
	c.pushPC()
	c.Reg.PC = dest

}

func call(c *CPU, _ byte, _ byte) {
	dest := c.fetch16()
	_call(c, dest)
}

func callf(c *CPU, flag byte, _ byte) {
	dest := c.fetch16()
	if c.Reg.isSet(flag) {
		_call(c, dest)
	}
}

func callnf(c *CPU, flag byte, _ byte) {
	dest := c.fetch16()
	if !c.Reg.isSet(flag) {
		_call(c, dest)
	}
}

// -----misc-----

func stop(c *CPU, _ byte, _ byte) {
	log.Debug("TODO: implement")
	c.Reg.PC++
}

// desable interrupt
func di(c *CPU, _ byte, _ byte) {
	log.Debug("TODO: implement")
}

// enable interrupt
func ei(c *CPU, _ byte, _ byte) {
	log.Debug("TODO: implement")
}

func notimplemented(c *CPU, _ byte, _ byte) {
	c.Reg.PC--
	panic(fmt.Sprintf("OpCode 0x%2x is not implemented", c.fetch()))
}

// Prefix CB
var cbOpCodes = []*OpCode{
	{0x00, "RLC B", B, 0, 1, 2, rlcr},
	{0x01, "RLC C", C, 0, 1, 2, rlcr},
	{0x02, "RLC D", D, 0, 1, 2, rlcr},
	{0x03, "RLC E", E, 0, 1, 2, rlcr},
	{0x04, "RLC H", H, 0, 1, 2, rlcr},
	{0x05, "RLC L", L, 0, 1, 2, rlcr},
	{0x06, "RLC (HL)", HL, 0, 1, 4, rlcm16},
	{0x07, "RLC A", A, 0, 1, 2, rlcr},
	{0x08, "RRC B", B, 0, 1, 2, rrcr},
	{0x09, "RRC C", C, 0, 1, 2, rrcr},
	{0x0A, "RRC D", D, 0, 1, 2, rrcr},
	{0x0B, "RRC E", E, 0, 1, 2, rrcr},
	{0x0C, "RRC H", H, 0, 1, 2, rrcr},
	{0x0D, "RRC L", L, 0, 1, 2, rrcr},
	{0x0E, "RRC (HL)", HL, 0, 1, 4, rrcm16},
	{0x0F, "RRC A", A, 0, 1, 2, rrcr},
	{0x10, "RL B", B, 0, 1, 2, rlr},
	{0x11, "RL C", C, 0, 1, 2, rlr},
	{0x12, "RL D", D, 0, 1, 2, rlr},
	{0x13, "RL E", E, 0, 1, 2, rlr},
	{0x14, "RL H", H, 0, 1, 2, rlr},
	{0x15, "RL L", L, 0, 1, 2, rlr},
	{0x16, "RL (HL)", HL, 0, 1, 4, rlm16},
	{0x17, "RL A", A, 0, 1, 2, rlr},
	{0x18, "RR B", B, 0, 1, 2, rrr},
	{0x19, "RR C", C, 0, 1, 2, rrr},
	{0x1A, "RR D", D, 0, 1, 2, rrr},
	{0x1B, "RR E", E, 0, 1, 2, rrr},
	{0x1C, "RR H", H, 0, 1, 2, rrr},
	{0x1D, "RR L", L, 0, 1, 2, rrr},
	{0x1E, "RR (HL)", HL, 0, 1, 4, rrm16},
	{0x1F, "RR A", A, 0, 1, 2, rrr},
	{0x20, "SLA B", B, 0, 1, 2, notimplemented},
	{0x21, "SLA C", C, 0, 1, 2, notimplemented},
	{0x22, "SLA D", D, 0, 1, 2, notimplemented},
	{0x23, "SLA E", E, 0, 1, 2, notimplemented},
	{0x24, "SLA H", H, 0, 1, 2, notimplemented},
	{0x25, "SLA L", L, 0, 1, 2, notimplemented},
	{0x26, "SLA (HL)", HL, 0, 1, 4, notimplemented},
	{0x27, "SLA A", A, 0, 1, 2, notimplemented},
	{0x28, "SRA B", B, 0, 1, 2, notimplemented},
	{0x29, "SRA C", C, 0, 1, 2, notimplemented},
	{0x2A, "SRA D", D, 0, 1, 2, notimplemented},
	{0x2B, "SRA E", E, 0, 1, 2, notimplemented},
	{0x2C, "SRA H", H, 0, 1, 2, notimplemented},
	{0x2D, "SRA L", L, 0, 1, 2, notimplemented},
	{0x2E, "SRA (HL)", HL, 0, 1, 4, notimplemented},
	{0x2F, "SRA A", A, 0, 1, 2, notimplemented},
	{0x30, "SWAP B", B, 0, 1, 2, notimplemented},
	{0x31, "SWAP C", C, 0, 1, 2, notimplemented},
	{0x32, "SWAP D", D, 0, 1, 2, notimplemented},
	{0x33, "SWAP E", E, 0, 1, 2, notimplemented},
	{0x34, "SWAP H", H, 0, 1, 2, notimplemented},
	{0x35, "SWAP L", L, 0, 1, 2, notimplemented},
	{0x36, "SWAP (HL)", HL, 0, 1, 4, notimplemented},
	{0x37, "SWAP A", A, 0, 1, 2, notimplemented},
	{0x38, "SRL B", B, 0, 1, 2, notimplemented},
	{0x39, "SRL C", C, 0, 1, 2, notimplemented},
	{0x3A, "SRL D", D, 0, 1, 2, notimplemented},
	{0x3B, "SRL E", E, 0, 1, 2, notimplemented},
	{0x3C, "SRL H", H, 0, 1, 2, notimplemented},
	{0x3D, "SRL L", L, 0, 1, 2, notimplemented},
	{0x3E, "SRL (HL)", HL, 0, 1, 4, notimplemented},
	{0x3F, "SRL A", A, 0, 1, 2, notimplemented},
	{0x40, "BIT 0,B", 0, B, 1, 2, notimplemented},
	{0x41, "BIT 0,C", 0, C, 1, 2, notimplemented},
	{0x42, "BIT 0,D", 0, D, 1, 2, notimplemented},
	{0x43, "BIT 0,E", 0, E, 1, 2, notimplemented},
	{0x44, "BIT 0,H", 0, H, 1, 2, notimplemented},
	{0x45, "BIT 0,L", 0, L, 1, 2, notimplemented},
	{0x46, "BIT 0,(HL)", 0, HL, 1, 4, notimplemented},
	{0x47, "BIT 0,A", 0, A, 1, 2, notimplemented},
	{0x48, "BIT 1,B", 1, B, 1, 2, notimplemented},
	{0x49, "BIT 1,C", 1, C, 1, 2, notimplemented},
	{0x4A, "BIT 1,D", 1, D, 1, 2, notimplemented},
	{0x4B, "BIT 1,E", 1, E, 1, 2, notimplemented},
	{0x4C, "BIT 1,H", 1, H, 1, 2, notimplemented},
	{0x4D, "BIT 1,L", 1, L, 1, 2, notimplemented},
	{0x4E, "BIT 1,(HL)", 1, HL, 1, 4, notimplemented},
	{0x4F, "BIT 1,A", 1, A, 1, 2, notimplemented},
	{0x50, "BIT 2,B", 2, B, 1, 2, notimplemented},
	{0x51, "BIT 2,C", 2, C, 1, 2, notimplemented},
	{0x52, "BIT 2,D", 2, D, 1, 2, notimplemented},
	{0x53, "BIT 2,E", 2, E, 1, 2, notimplemented},
	{0x54, "BIT 2,H", 2, H, 1, 2, notimplemented},
	{0x55, "BIT 2,L", 2, L, 1, 2, notimplemented},
	{0x56, "BIT 2,(HL)", 2, HL, 1, 4, notimplemented},
	{0x57, "BIT 2,A", 2, A, 1, 2, notimplemented},
	{0x58, "BIT 3,B", 3, B, 1, 2, notimplemented},
	{0x59, "BIT 3,C", 3, C, 1, 2, notimplemented},
	{0x5A, "BIT 3,D", 3, D, 1, 2, notimplemented},
	{0x5B, "BIT 3,E", 3, E, 1, 2, notimplemented},
	{0x5C, "BIT 3,H", 3, H, 1, 2, notimplemented},
	{0x5D, "BIT 3,L", 3, L, 1, 2, notimplemented},
	{0x5E, "BIT 3,(HL)", 3, HL, 1, 4, notimplemented},
	{0x5F, "BIT 3,A", 3, A, 1, 2, notimplemented},
	{0x60, "BIT 4,B", 4, B, 1, 2, notimplemented},
	{0x61, "BIT 4,C", 4, C, 1, 2, notimplemented},
	{0x62, "BIT 4,D", 4, D, 1, 2, notimplemented},
	{0x63, "BIT 4,E", 4, E, 1, 2, notimplemented},
	{0x64, "BIT 4,H", 4, H, 1, 2, notimplemented},
	{0x65, "BIT 4,L", 4, L, 1, 2, notimplemented},
	{0x66, "BIT 4,(HL)", 4, HL, 1, 4, notimplemented},
	{0x67, "BIT 4,A", 4, A, 1, 2, notimplemented},
	{0x68, "BIT 5,B", 5, B, 1, 2, notimplemented},
	{0x69, "BIT 5,C", 5, C, 1, 2, notimplemented},
	{0x6A, "BIT 5,D", 5, D, 1, 2, notimplemented},
	{0x6B, "BIT 5,E", 5, E, 1, 2, notimplemented},
	{0x6C, "BIT 5,H", 5, H, 1, 2, notimplemented},
	{0x6D, "BIT 5,L", 5, L, 1, 2, notimplemented},
	{0x6E, "BIT 5,(HL)", 5, HL, 1, 4, notimplemented},
	{0x6F, "BIT 5,A", 5, A, 1, 2, notimplemented},
	{0x70, "BIT 6,B", 6, B, 1, 2, notimplemented},
	{0x71, "BIT 6,C", 6, C, 1, 2, notimplemented},
	{0x72, "BIT 6,D", 6, D, 1, 2, notimplemented},
	{0x73, "BIT 6,E", 6, E, 1, 2, notimplemented},
	{0x74, "BIT 6,H", 6, H, 1, 2, notimplemented},
	{0x75, "BIT 6,L", 6, L, 1, 2, notimplemented},
	{0x76, "BIT 6,(HL)", 6, HL, 1, 4, notimplemented},
	{0x77, "BIT 6,A", 6, A, 1, 2, notimplemented},
	{0x78, "BIT 7,B", 7, B, 1, 2, notimplemented},
	{0x79, "BIT 7,C", 7, C, 1, 2, notimplemented},
	{0x7A, "BIT 7,D", 7, D, 1, 2, notimplemented},
	{0x7B, "BIT 7,E", 7, E, 1, 2, notimplemented},
	{0x7C, "BIT 7,H", 7, H, 1, 2, notimplemented},
	{0x7D, "BIT 7,L", 7, L, 1, 2, notimplemented},
	{0x7E, "BIT 7,(HL)", 7, HL, 1, 4, notimplemented},
	{0x7F, "BIT 7,A", 7, A, 1, 2, notimplemented},
	{0x80, "RES 0,B", 0, B, 1, 2, notimplemented},
	{0x81, "RES 0,C", 0, C, 1, 2, notimplemented},
	{0x82, "RES 0,D", 0, D, 1, 2, notimplemented},
	{0x83, "RES 0,E", 0, E, 1, 2, notimplemented},
	{0x84, "RES 0,H", 0, H, 1, 2, notimplemented},
	{0x85, "RES 0,L", 0, L, 1, 2, notimplemented},
	{0x86, "RES 0,(HL)", 0, HL, 1, 4, notimplemented},
	{0x87, "RES 0,A", 0, A, 1, 2, notimplemented},
	{0x88, "RES 1,B", 1, B, 1, 2, notimplemented},
	{0x89, "RES 1,C", 1, C, 1, 2, notimplemented},
	{0x8A, "RES 1,D", 1, D, 1, 2, notimplemented},
	{0x8B, "RES 1,E", 1, E, 1, 2, notimplemented},
	{0x8C, "RES 1,H", 1, H, 1, 2, notimplemented},
	{0x8D, "RES 1,L", 1, L, 1, 2, notimplemented},
	{0x8E, "RES 1,(HL)", 1, HL, 1, 4, notimplemented},
	{0x8F, "RES 1,A", 1, A, 1, 2, notimplemented},
	{0x90, "RES 2,B", 2, B, 1, 2, notimplemented},
	{0x91, "RES 2,C", 2, C, 1, 2, notimplemented},
	{0x92, "RES 2,D", 2, D, 1, 2, notimplemented},
	{0x93, "RES 2,E", 2, E, 1, 2, notimplemented},
	{0x94, "RES 2,H", 2, H, 1, 2, notimplemented},
	{0x95, "RES 2,L", 2, L, 1, 2, notimplemented},
	{0x96, "RES 2,(HL)", 2, HL, 1, 4, notimplemented},
	{0x97, "RES 2,A", 2, A, 1, 2, notimplemented},
	{0x98, "RES 3,B", 3, B, 1, 2, notimplemented},
	{0x99, "RES 3,C", 3, C, 1, 2, notimplemented},
	{0x9A, "RES 3,D", 3, D, 1, 2, notimplemented},
	{0x9B, "RES 3,E", 3, E, 1, 2, notimplemented},
	{0x9C, "RES 3,H", 3, H, 1, 2, notimplemented},
	{0x9D, "RES 3,L", 3, L, 1, 2, notimplemented},
	{0x9E, "RES 3,(HL)", 3, HL, 1, 4, notimplemented},
	{0x9F, "RES 3,A", 3, A, 1, 2, notimplemented},
	{0xA0, "RES 4,B", 4, B, 1, 2, notimplemented},
	{0xA1, "RES 4,C", 4, C, 1, 2, notimplemented},
	{0xA2, "RES 4,D", 4, D, 1, 2, notimplemented},
	{0xA3, "RES 4,E", 4, E, 1, 2, notimplemented},
	{0xA4, "RES 4,H", 4, H, 1, 2, notimplemented},
	{0xA5, "RES 4,L", 4, L, 1, 2, notimplemented},
	{0xA6, "RES 4,(HL)", 4, HL, 1, 4, notimplemented},
	{0xA7, "RES 4,A", 4, A, 1, 2, notimplemented},
	{0xA8, "RES 5,B", 5, B, 1, 2, notimplemented},
	{0xA9, "RES 5,C", 5, C, 1, 2, notimplemented},
	{0xAA, "RES 5,D", 5, D, 1, 2, notimplemented},
	{0xAB, "RES 5,E", 5, E, 1, 2, notimplemented},
	{0xAC, "RES 5,H", 5, H, 1, 2, notimplemented},
	{0xAD, "RES 5,L", 5, L, 1, 2, notimplemented},
	{0xAE, "RES 5,(HL)", 5, HL, 1, 4, notimplemented},
	{0xAF, "RES 5,A", 5, A, 1, 2, notimplemented},
	{0xB0, "RES 6,B", 6, B, 1, 2, notimplemented},
	{0xB1, "RES 6,C", 6, C, 1, 2, notimplemented},
	{0xB2, "RES 6,D", 6, D, 1, 2, notimplemented},
	{0xB3, "RES 6,E", 6, E, 1, 2, notimplemented},
	{0xB4, "RES 6,H", 6, H, 1, 2, notimplemented},
	{0xB5, "RES 6,L", 6, L, 1, 2, notimplemented},
	{0xB6, "RES 6,(HL)", 6, HL, 1, 4, notimplemented},
	{0xB7, "RES 6,A", 6, A, 1, 2, notimplemented},
	{0xB8, "RES 7,B", 7, B, 1, 2, notimplemented},
	{0xB9, "RES 7,C", 7, C, 1, 2, notimplemented},
	{0xBA, "RES 7,D", 7, D, 1, 2, notimplemented},
	{0xBB, "RES 7,E", 7, E, 1, 2, notimplemented},
	{0xBC, "RES 7,H", 7, H, 1, 2, notimplemented},
	{0xBD, "RES 7,L", 7, L, 1, 2, notimplemented},
	{0xBE, "RES 7,(HL)", 7, HL, 1, 4, notimplemented},
	{0xBF, "RES 7,A", 7, A, 1, 2, notimplemented},
	{0xC0, "SET 0,B", 0, B, 1, 2, notimplemented},
	{0xC1, "SET 0,C", 0, C, 1, 2, notimplemented},
	{0xC2, "SET 0,D", 0, D, 1, 2, notimplemented},
	{0xC3, "SET 0,E", 0, E, 1, 2, notimplemented},
	{0xC4, "SET 0,H", 0, H, 1, 2, notimplemented},
	{0xC5, "SET 0,L", 0, L, 1, 2, notimplemented},
	{0xC6, "SET 0,(HL)", 0, HL, 1, 4, notimplemented},
	{0xC7, "SET 0,A", 0, A, 1, 2, notimplemented},
	{0xC8, "SET 1,B", 1, B, 1, 2, notimplemented},
	{0xC9, "SET 1,C", 1, C, 1, 2, notimplemented},
	{0xCA, "SET 1,D", 1, D, 1, 2, notimplemented},
	{0xCB, "SET 1,E", 1, E, 1, 2, notimplemented},
	{0xCC, "SET 1,H", 1, H, 1, 2, notimplemented},
	{0xCD, "SET 1,L", 1, L, 1, 2, notimplemented},
	{0xCE, "SET 1,(HL)", 1, HL, 1, 4, notimplemented},
	{0xCF, "SET 1,A", 1, A, 1, 2, notimplemented},
	{0xD0, "SET 2,B", 2, B, 1, 2, notimplemented},
	{0xD1, "SET 2,C", 2, C, 1, 2, notimplemented},
	{0xD2, "SET 2,D", 2, D, 1, 2, notimplemented},
	{0xD3, "SET 2,E", 2, E, 1, 2, notimplemented},
	{0xD4, "SET 2,H", 2, H, 1, 2, notimplemented},
	{0xD5, "SET 2,L", 2, L, 1, 2, notimplemented},
	{0xD6, "SET 2,(HL)", 2, HL, 1, 4, notimplemented},
	{0xD7, "SET 2,A", 2, A, 1, 2, notimplemented},
	{0xD8, "SET 3,B", 3, B, 1, 2, notimplemented},
	{0xD9, "SET 3,C", 3, C, 1, 2, notimplemented},
	{0xDA, "SET 3,D", 3, D, 1, 2, notimplemented},
	{0xDB, "SET 3,E", 3, E, 1, 2, notimplemented},
	{0xDC, "SET 3,H", 3, H, 1, 2, notimplemented},
	{0xDD, "SET 3,L", 3, L, 1, 2, notimplemented},
	{0xDE, "SET 3,(HL)", 3, HL, 1, 4, notimplemented},
	{0xDF, "SET 3,A", 3, A, 1, 2, notimplemented},
	{0xE0, "SET 4,B", 4, B, 1, 2, notimplemented},
	{0xE1, "SET 4,C", 4, C, 1, 2, notimplemented},
	{0xE2, "SET 4,D", 4, D, 1, 2, notimplemented},
	{0xE3, "SET 4,E", 4, E, 1, 2, notimplemented},
	{0xE4, "SET 4,H", 4, H, 1, 2, notimplemented},
	{0xE5, "SET 4,L", 4, L, 1, 2, notimplemented},
	{0xE6, "SET 4,(HL)", 4, HL, 1, 4, notimplemented},
	{0xE7, "SET 4,A", 4, A, 1, 2, notimplemented},
	{0xE8, "SET 5,B", 5, B, 1, 2, notimplemented},
	{0xE9, "SET 5,C", 5, C, 1, 2, notimplemented},
	{0xEA, "SET 5,D", 5, D, 1, 2, notimplemented},
	{0xEB, "SET 5,E", 5, E, 1, 2, notimplemented},
	{0xEC, "SET 5,H", 5, H, 1, 2, notimplemented},
	{0xED, "SET 5,L", 5, L, 1, 2, notimplemented},
	{0xEE, "SET 5,(HL)", 5, HL, 1, 4, notimplemented},
	{0xEF, "SET 5,A", 5, A, 1, 2, notimplemented},
	{0xF0, "SET 6,B", 6, B, 1, 2, notimplemented},
	{0xF1, "SET 6,C", 6, C, 1, 2, notimplemented},
	{0xF2, "SET 6,D", 6, D, 1, 2, notimplemented},
	{0xF3, "SET 6,E", 6, E, 1, 2, notimplemented},
	{0xF4, "SET 6,H", 6, H, 1, 2, notimplemented},
	{0xF5, "SET 6,L", 6, L, 1, 2, notimplemented},
	{0xF6, "SET 6,(HL)", 6, HL, 1, 4, notimplemented},
	{0xF7, "SET 6,A", 6, A, 1, 2, notimplemented},
	{0xF8, "SET 7,B", 7, B, 1, 2, notimplemented},
	{0xF9, "SET 7,C", 7, C, 1, 2, notimplemented},
	{0xFA, "SET 7,D", 7, D, 1, 2, notimplemented},
	{0xFB, "SET 7,E", 7, E, 1, 2, notimplemented},
	{0xFC, "SET 7,H", 7, H, 1, 2, notimplemented},
	{0xFD, "SET 7,L", 7, L, 1, 2, notimplemented},
	{0xFE, "SET 7,(HL)", 7, HL, 1, 4, notimplemented},
	{0xFF, "SET 7,A", 7, A, 1, 2, notimplemented},
}

func _rlc(c *CPU, v byte) byte {
	// check Bit 7 is set
	if v&0x80 == 0x80 {
		c.Reg.setFlag(flagC)
	} else {
		c.Reg.clearFlag(flagC)
	}

	v = v << 1
	c.Reg.setFlagZ(v)
	c.Reg.clearFlag(flagN)
	c.Reg.clearFlag(flagH)

	return v
}

// RLC r
func rlcr(c *CPU, r8 byte, _ byte) {
	r := c.Reg.R[r8]
	c.Reg.R[r8] = _rlc(c, r)

}

// RLC (HL)
func rlcm16(c *CPU, r16 byte, _ byte) {
	addr := c.Reg.R16(int(r16))
	r := c.Bus.ReadByte(addr)
	c.Bus.WriteByte(addr, _rlc(c, r))
}

func _rrc(c *CPU, v byte) byte {
	// check Bit 0 is set
	if v&0x01 == 0x01 {
		c.Reg.setFlag(flagC)
	} else {
		c.Reg.clearFlag(flagC)
	}

	v = v >> 1
	c.Reg.setFlagZ(v)
	c.Reg.clearFlag(flagN)
	c.Reg.clearFlag(flagH)

	return v
}

// RLC r
func rrcr(c *CPU, r8 byte, _ byte) {
	r := c.Reg.R[r8]
	c.Reg.R[r8] = _rrc(c, r)
}

// RLC (HL)
func rrcm16(c *CPU, r16 byte, _ byte) {
	addr := c.Reg.R16(int(r16))
	r := c.Bus.ReadByte(addr)
	c.Bus.WriteByte(addr, _rrc(c, r))
}

func _rl(c *CPU, v byte) byte {
	// check Bit 0 is set
	if v&0x80 == 0x80 {
		c.Reg.setFlag(flagC)
	} else {
		c.Reg.clearFlag(flagC)
	}

	v = v << 1
	if c.Reg.isSet(flagC) {
		v++
	}
	c.Reg.setFlagZ(v)
	c.Reg.clearFlag(flagN)
	c.Reg.clearFlag(flagH)

	return v
}

// RLC r
func rlr(c *CPU, r8 byte, _ byte) {
	r := c.Reg.R[r8]
	c.Reg.R[r8] = _rl(c, r)
}

// RLC (HL)
func rlm16(c *CPU, r16 byte, _ byte) {
	addr := c.Reg.R16(int(r16))
	r := c.Bus.ReadByte(addr)
	c.Bus.WriteByte(addr, _rl(c, r))
}

func _rr(c *CPU, v byte) byte {
	// check Bit 0 is set
	if v&0x01 == 0x01 {
		c.Reg.setFlag(flagC)
	} else {
		c.Reg.clearFlag(flagC)
	}

	v = v >> 1
	if c.Reg.isSet(flagC) {
		v |= 0x80
	}
	c.Reg.setFlagZ(v)
	c.Reg.clearFlag(flagN)
	c.Reg.clearFlag(flagH)

	return v
}

// RLC r
func rrr(c *CPU, r8 byte, _ byte) {
	r := c.Reg.R[r8]
	c.Reg.R[r8] = _rr(c, r)
}

// RLC (HL)
func rrm16(c *CPU, r16 byte, _ byte) {
	addr := c.Reg.R16(int(r16))
	r := c.Bus.ReadByte(addr)
	c.Bus.WriteByte(addr, _rr(c, r))
}
