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
	{0x06, "LD B,d8", B, 0, 1, 2, ldr8d8},
	{0x07, "RLCA", 0, 0, 0, 1, notimplemented},
	{0x08, "LD (a16),SP", 0, 0, 0, 1, notimplemented},
	{0x09, "ADD HL,BC", 0, 0, 0, 1, notimplemented},
	{0x0A, "LD A,(BC)", A, BC, 0, 2, ldrm16},
	{0x0B, "DEC BC", BC, 0, 0, 2, decr16},
	{0x0C, "INC C", C, 0, 1, 2, incr},
	{0x0D, "DEC C", C, 0, 0, 1, decr},
	{0x0E, "LD C,d8", C, 0, 1, 2, ldr8d8},
	{0x0F, "RRCA", 0, 0, 0, 1, notimplemented},
	{0x10, "STOP 0", 0, 0, 0, 1, stop},
	{0x11, "LD DE,d16", DE, 0, 2, 3, ldr16d16},
	{0x12, "LD (DE),A", DE, A, 0, 2, ldm16r8},
	{0x13, "INC DE", DE, 0, 0, 2, incr16},
	{0x14, "INC D", D, 0, 0, 1, incr},
	{0x15, "DEC D", D, 0, 0, 1, decr},
	{0x16, "LD D,d8", D, 0, 1, 2, ldr8d8},
	{0x17, "RLA", 0, 0, 0, 1, notimplemented},
	{0x18, "JR r8", 0, 0, 1, 3, jrr8},
	{0x19, "ADD HL,DE", 0, 0, 0, 1, notimplemented},
	{0x1A, "LD A,(DE)", A, DE, 0, 2, ldrm16},
	{0x1B, "DEC DE", DE, 0, 0, 2, decr16},
	{0x1C, "INC E", E, 0, 0, 1, incr},
	{0x1D, "DEC E", E, 0, 0, 1, decr},
	{0x1E, "LD E,d8", E, 0, 1, 2, ldr8d8},
	{0x1F, "RRA", 0, 0, 0, 1, notimplemented},
	{0x20, "JR NZ,r8", flagZ, 0, 1, 2, jrnfr8},
	{0x21, "LD HL,d16", HL, 0, 2, 3, ldr16d16},
	{0x22, "LD (HL+),A", HLI, A, 0, 2, ldm16r},
	{0x23, "INC HL", HL, 0, 0, 2, incr16},
	{0x24, "INC H", H, 0, 0, 1, incr},
	{0x25, "DEC H", H, 0, 0, 1, decr},
	{0x26, "LD H,d8", H, 0, 1, 2, ldr8d8},
	{0x27, "DAA", 0, 0, 0, 1, notimplemented},
	{0x28, "JR Z,r8", flagZ, 0, 1, 2, jrfr8},
	{0x29, "ADD HL,HL", 0, 0, 0, 1, notimplemented},
	{0x2A, "LD A,(HL+)", A, HLI, 0, 2, ldrm16},
	{0x2B, "DEC HL", HL, 0, 0, 1, decr16},
	{0x2C, "INC L", L, 0, 0, 1, incr},
	{0x2D, "DEC L", L, 0, 0, 1, decr},
	{0x2E, "LD L,d8", L, 0, 0, 1, ldr8d8},
	{0x2F, "CPL", 0, 0, 0, 1, notimplemented},
	{0x30, "JR NC,r8", flagC, 0, 1, 2, jrnfr8},
	{0x31, "LD SP,d16", SP, 0, 2, 3, ldr16d16},
	{0x32, "LD (HL-),A", HLD, A, 0, 2, ldm16r},
	{0x33, "INC SP", SP, 0, 0, 2, incr16},
	{0x34, "INC (HL)", HL, 0, 0, 1, notimplemented},
	{0x35, "DEC (HL)", HL, 0, 0, 1, notimplemented},
	{0x36, "LD (HL),d8", HL, 0, 1, 3, ldr16d16},
	{0x37, "SCF", 0, 0, 0, 1, notimplemented},
	{0x38, "JR C,r8", flagC, 0, 1, 2, jrfr8},
	{0x39, "ADD HL,SP", 0, 0, 0, 1, notimplemented},
	{0x3A, "LD A,(HL-)", A, HLD, 0, 2, ldrm16},
	{0x3B, "DEC SP", 0, 0, 0, 1, notimplemented},
	{0x3C, "INC A", A, 0, 0, 1, incr},
	{0x3D, "DEC A", A, 0, 0, 1, notimplemented},
	{0x3E, "LD A,d8", A, 0, 1, 2, ldr8d8},
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
	{0x80, "ADD A, B", A, B, 0, 1, notimplemented},
	{0x81, "ADD A, C", A, C, 0, 1, notimplemented},
	{0x82, "ADD A, D", A, D, 0, 1, notimplemented},
	{0x83, "ADD A, E", A, E, 0, 1, notimplemented},
	{0x84, "ADD A, H", A, H, 0, 1, notimplemented},
	{0x85, "ADD A, L", A, L, 0, 1, notimplemented},
	{0x86, "ADD A, (HL)", A, HL, 0, 1, notimplemented},
	{0x87, "ADD A, A", A, A, 0, 1, notimplemented},
	{0x88, "ADC A, B", A, B, 0, 1, notimplemented},
	{0x89, "ADC A, C", A, C, 0, 1, notimplemented},
	{0x8A, "ADC A, D", A, D, 0, 1, notimplemented},
	{0x8B, "ADC A, E", A, E, 0, 1, notimplemented},
	{0x8C, "ADC A, H", A, H, 0, 1, notimplemented},
	{0x8D, "ADC A, L", A, L, 0, 1, notimplemented},
	{0x8E, "ADC A, (HL)", A, HL, 0, 1, notimplemented},
	{0x8F, "ADC A, A", A, A, 0, 1, notimplemented},
	{0x90, "SUB B", B, 0, 0, 1, notimplemented},
	{0x91, "SUB C", C, 0, 0, 1, notimplemented},
	{0x92, "SUB D", D, 0, 0, 1, notimplemented},
	{0x93, "SUB E", E, 0, 0, 1, notimplemented},
	{0x94, "SUB H", H, 0, 0, 1, notimplemented},
	{0x95, "SUB L", L, 0, 0, 1, notimplemented},
	{0x96, "SUB (HL)", HL, 0, 0, 1, notimplemented},
	{0x97, "SUB A", A, 0, 0, 1, notimplemented},
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
	{0xA6, "AND (HL)", HL, 0, 0, 2, notimplemented},
	{0xA7, "AND A", A, 0, 0, 1, andr},
	{0xA8, "XOR B", B, 0, 0, 1, xorr},
	{0xA9, "XOR C", C, 0, 0, 1, xorr},
	{0xAA, "XOR D", D, 0, 0, 1, xorr},
	{0xAB, "XOR E", E, 0, 0, 1, xorr},
	{0xAC, "XOR H", H, 0, 0, 1, xorr},
	{0xAD, "XOR L", L, 0, 0, 1, xorr},
	{0xAE, "XOR (HL)", HL, 0, 0, 2, notimplemented},
	{0xAF, "XOR A", A, 0, 0, 1, xorr},
	{0xB0, "OR B", B, 0, 0, 1, orr},
	{0xB1, "OR C", C, 0, 0, 1, orr},
	{0xB2, "OR D", D, 0, 0, 1, orr},
	{0xB3, "OR E", E, 0, 0, 1, orr},
	{0xB4, "OR H", H, 0, 0, 1, orr},
	{0xB5, "OR L", L, 0, 0, 1, orr},
	{0xB6, "OR (HL)", HL, 0, 0, 2, orr},
	{0xB7, "OR A", A, 0, 0, 1, orr},
	{0xB8, "CP B", B, 0, 0, 1, notimplemented},
	{0xB9, "CP C", C, 0, 0, 1, notimplemented},
	{0xBA, "CP D", D, 0, 0, 1, notimplemented},
	{0xBB, "CP E", E, 0, 0, 1, notimplemented},
	{0xBC, "CP H", H, 0, 0, 1, notimplemented},
	{0xBD, "CP L", L, 0, 0, 1, notimplemented},
	{0xBE, "CP (HL)", HL, 0, 0, 1, notimplemented},
	{0xBF, "CP A", A, 0, 0, 1, notimplemented},
	{0xC0, "RET NZ", flagZ, 0, 2, 2, retncc},
	{0xC1, "POP BC", BC, 0, 0, 3, pop},
	{0xC2, "JP NZ,a16", flagZ, 0, 2, 3, jpnfa16},
	{0xC3, "JP a16", 0, 0, 2, 4, jpa16},
	{0xC4, "CALL NZ,a16", flagZ, 0, 2, 3, callnf},
	{0xC5, "PUSH BC", BC, 0, 0, 4, push},
	{0xC6, "ADD A,d8", A, 0, 0, 1, notimplemented},
	{0xC7, "RST 00H", 0x00, 0, 0, 1, rst},
	{0xC8, "RET Z", flagZ, 0, 0, 1, notimplemented},
	{0xC9, "RET", 0, 0, 0, 4, ret},
	{0xCA, "JP Z,a16", flagZ, 0, 2, 3, jpfa16},
	{0xCB, "PREFIX CB", 0, 0, 0, 1, notimplemented},
	{0xCC, "CALL Z,a16", flagZ, 0, 2, 3, callf},
	{0xCD, "CALL a16", 0, 0, 2, 4, call},
	{0xCE, "ADC A,d8", 0, 0, 0, 1, notimplemented},
	{0xCF, "RST 08H", 0x08, 0, 0, 1, rst},
	{0xD0, "RET NC", flagC, 0, 0, 1, notimplemented},
	{0xD1, "POP DE", DE, 0, 0, 3, pop},
	{0xD2, "JP NC,a16", flagC, 0, 2, 3, jpnfa16},
	{0xD3, "EMPTY", 0, 0, 0, 1, notimplemented},
	{0xD4, "CALL NC,a16", flagC, 0, 2, 3, callnf},
	{0xD5, "PUSH DE", DE, 0, 0, 1, push},
	{0xD6, "SUB d8", 0, 0, 0, 1, notimplemented},
	{0xD7, "RST 10H", 0x10, 0, 0, 1, rst},
	{0xD8, "RET C", flagC, 0, 0, 1, notimplemented},
	{0xD9, "RETI", 0, 0, 0, 1, notimplemented},
	{0xDA, "JP C,a16", flagC, 0, 2, 3, jpfa16},
	{0xDB, "EMPTY", 0, 0, 0, 1, notimplemented},
	{0xDC, "CALL C,a16", flagC, 0, 2, 3, callf},
	{0xDD, "EMPTY", 0, 0, 0, 1, notimplemented},
	{0xDE, "SBC A,d8", A, 0, 0, 1, notimplemented},
	{0xDF, "RST 18H", 0x18, 0, 0, 1, rst},
	{0xE0, "LDH (a8),A", 0, A, 0, 1, lda8r},
	{0xE1, "POP HL", HL, 0, 0, 3, pop},
	{0xE2, "LD (C),A", 0, 0, 0, 1, notimplemented},
	{0xE3, "EMPTY", 0, 0, 0, 1, notimplemented},
	{0xE4, "EMPTY", 0, 0, 0, 1, notimplemented},
	{0xE5, "PUSH HL", HL, 0, 0, 1, push},
	{0xE6, "AND d8", 0, 0, 0, 1, notimplemented},
	{0xE7, "RST 20H", 0x20, 0, 0, 1, rst},
	{0xE8, "ADD SP,r8", SP, 0, 0, 1, notimplemented},
	{0xE9, "JP (HL)", HL, 0, 0, 1, jpm16},
	{0xEA, "LD (a16),A", 0, A, 0, 1, lda16r},
	{0xEB, "EMPTY", 0, 0, 0, 1, notimplemented},
	{0xEC, "EMPTY", 0, 0, 0, 1, notimplemented},
	{0xED, "EMPTY", 0, 0, 0, 1, notimplemented},
	{0xEE, "XOR d8", 0, 0, 0, 1, notimplemented},
	{0xEF, "RST 28H", 0x28, 0, 0, 1, rst},
	{0xF0, "LDH A,(a8)", A, 0, 1, 3, ldra8},
	{0xF1, "POP AF", AF, 0, 0, 3, pop},
	{0xF2, "LD A,(C)", 0, 0, 0, 1, notimplemented},
	{0xF3, "DI", 0, 0, 0, 1, di},
	{0xF4, "EMPTY", 0, 0, 0, 1, notimplemented},
	{0xF5, "PUSH AF", AF, 0, 0, 1, push},
	{0xF6, "OR d8", 0, 0, 0, 1, notimplemented},
	{0xF7, "RST 30H", 0x30, 0, 0, 1, rst},
	{0xF8, "LD HL,SP+r8", 0, 0, 0, 1, notimplemented},
	{0xF9, "LD SP,HL", 0, 0, 0, 1, notimplemented},
	{0xFA, "LD A,(a16)", 0, 0, 0, 1, notimplemented},
	{0xFB, "EI", 0, 0, 0, 1, ei},
	{0xFC, "EMPTY", 0, 0, 0, 1, notimplemented},
	{0xFD, "EMPTY", 0, 0, 0, 1, notimplemented},
	{0xFE, "CP d8", 0, 0, 0, 1, notimplemented},
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

func lda16r(c *CPU, _ byte, R2 byte) {
	c.Bus.WriteByte(c.fetch16(), c.Reg.R[R2])
}

func lda8r(c *CPU, _ byte, R2 byte) {
	c.Bus.WriteByte(types.Addr(0xff00|types.Addr(c.fetch())), c.Reg.R[R2])
}

func ldra8(c *CPU, R1 byte, _ byte) {
	c.Reg.R[R1] = c.Bus.ReadByte(types.Addr(0xff00 | types.Addr(c.fetch())))
}

func retcc(c *CPU, R1 byte, _ byte) {
	if c.Reg.R[F]&(1<<R1) != 0 {
		c.popPC()
	}
}

func retncc(c *CPU, R1 byte, _ byte) {
	if c.Reg.R[F]&(1<<R1) == 0 {
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
	c.Reg.setR16(types.Addr(r16), types.Addr(c.Reg.R16(int(r16))+1))
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
	c.Reg.setR16(types.Addr(r16), types.Addr(c.Reg.R16(int(r16))-1))
}

func andr(c *CPU, r8 byte, _ byte) {
	c.Reg.R[A] &= c.Reg.R[r8]

	if c.Reg.R[A] == 0 {
		c.Reg.setFlag(flagZ)
	} else {
		c.Reg.clearFlag(flagZ)
	}
	c.Reg.clearFlag(flagN)
	c.Reg.setFlag(flagH)
	c.Reg.clearFlag(flagC)
}

func orr(c *CPU, r8 byte, _ byte) {
	c.Reg.R[A] |= c.Reg.R[r8]

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
	c.Reg.R[A] ^= c.Reg.R[r8]

	if c.Reg.R[A] == 0 {
		c.Reg.setFlag(flagZ)
	} else {
		c.Reg.clearFlag(flagZ)
	}
	c.Reg.clearFlag(flagN)
	c.Reg.clearFlag(flagH)
	c.Reg.clearFlag(flagC)
}

// special
func ret(c *CPU, _ byte, _ byte) {
	c.popPC()
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

	c.Reg.setR16(types.Addr(r16), types.Addr(int16(upper)<<8|int16(lower)))
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
