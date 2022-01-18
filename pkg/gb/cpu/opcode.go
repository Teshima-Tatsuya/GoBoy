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
	R1, R2   int
	Size     uint8
	Cycles   uint8
	Handler  func(*CPU, int, int)
}

var opCodes = []*OpCode{
	{0x00, "NOP", 0, 0, 0, 1, nop},
	{0x01, "LD BC,d16", BC, 0, 2, 3, ldr16d16},
	{0x02, "LD (BC),A", BC, A, 0, 2, ldm16r},
	{0x03, "INC BC", BC, 0, 0, 2, incr16},
	{0x04, "INC B", B, 0, 0, 1, incr},
	{0x05, "DEC B", B, 0, 0, 1, decr},
	{0x06, "LD B,d8", B, 0, 1, 2, ldrd},
	{0x07, "RLCA", 0, 0, 0, 1, rlca},
	{0x08, "LD (a16),SP", 0, SP, 2, 5, lda16r16},
	{0x09, "ADD HL,BC", HL, BC, 0, 2, addr16r16},
	{0x0A, "LD A,(BC)", A, BC, 0, 2, ldrm16},
	{0x0B, "DEC BC", BC, 0, 0, 2, decr16},
	{0x0C, "INC C", C, 0, 1, 1, incr},
	{0x0D, "DEC C", C, 0, 0, 1, decr},
	{0x0E, "LD C,d8", C, 0, 1, 2, ldrd},
	{0x0F, "RRCA", 0, 0, 0, 1, rrca},
	{0x10, "STOP 0", 0, 0, 0, 1, stop},
	{0x11, "LD DE,d16", DE, 0, 2, 3, ldr16d16},
	{0x12, "LD (DE),A", DE, A, 0, 2, ldm16r},
	{0x13, "INC DE", DE, 0, 0, 2, incr16},
	{0x14, "INC D", D, 0, 0, 1, incr},
	{0x15, "DEC D", D, 0, 0, 1, decr},
	{0x16, "LD D,d8", D, 0, 1, 2, ldrd},
	{0x17, "RLA", 0, 0, 0, 1, rla},
	{0x18, "JR r8", 0, 0, 1, 3, jrr8},
	{0x19, "ADD HL,DE", HL, DE, 0, 2, addr16r16},
	{0x1A, "LD A,(DE)", A, DE, 0, 2, ldrm16},
	{0x1B, "DEC DE", DE, 0, 0, 2, decr16},
	{0x1C, "INC E", E, 0, 0, 1, incr},
	{0x1D, "DEC E", E, 0, 0, 1, decr},
	{0x1E, "LD E,d8", E, 0, 1, 2, ldrd},
	{0x1F, "RRA", 0, 0, 0, 1, rra},
	{0x20, "JR NZ,r8", flagZ, 0, 1, 2, jrnfr8},
	{0x21, "LD HL,d16", HL, 0, 2, 3, ldr16d16},
	{0x22, "LD (HL+),A", HLI, A, 0, 2, ldm16r},
	{0x23, "INC HL", HL, 0, 0, 2, incr16},
	{0x24, "INC H", H, 0, 0, 1, incr},
	{0x25, "DEC H", H, 0, 0, 1, decr},
	{0x26, "LD H,d8", H, 0, 1, 2, ldrd},
	{0x27, "DAA", 0, 0, 0, 1, daa},
	{0x28, "JR Z,r8", flagZ, 0, 1, 2, jrfr8},
	{0x29, "ADD HL,HL", HL, HL, 0, 2, addr16r16},
	{0x2A, "LD A,(HL+)", A, HLI, 0, 2, ldrm16},
	{0x2B, "DEC HL", HL, 0, 0, 2, decr16},
	{0x2C, "INC L", L, 0, 0, 1, incr},
	{0x2D, "DEC L", L, 0, 0, 1, decr},
	{0x2E, "LD L,d8", L, 0, 0, 2, ldrd},
	{0x2F, "CPL", 0, 0, 0, 1, cpl},
	{0x30, "JR NC,r8", flagC, 0, 1, 2, jrnfr8},
	{0x31, "LD SP,d16", SP, 0, 2, 3, ldr16d16},
	{0x32, "LD (HL-),A", HLD, A, 0, 2, ldm16r},
	{0x33, "INC SP", SP, 0, 0, 2, incr16},
	{0x34, "INC (HL)", HL, 0, 0, 3, incm16},
	{0x35, "DEC (HL)", HL, 0, 0, 3, decm16},
	{0x36, "LD (HL),d8", HL, 0, 1, 3, ldm16d},
	{0x37, "SCF", 0, 0, 0, 1, scf},
	{0x38, "JR C,r8", flagC, 0, 1, 2, jrfr8},
	{0x39, "ADD HL,SP", HL, SP, 0, 2, addr16r16},
	{0x3A, "LD A,(HL-)", A, HLD, 0, 2, ldrm16},
	{0x3B, "DEC SP", SP, 0, 0, 2, decr16},
	{0x3C, "INC A", A, 0, 0, 1, incr},
	{0x3D, "DEC A", A, 0, 0, 1, decr},
	{0x3E, "LD A,d8", A, 0, 1, 2, ldrd},
	{0x3F, "CCF", 0, 0, 0, 1, ccf},
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
	{0x76, "HALT", 0, 0, 0, 1, halt},
	{0x77, "LD (HL), A", HL, A, 0, 2, ldm16r},
	{0x78, "LD A, B", A, B, 0, 1, ldrr},
	{0x79, "LD A, C", A, C, 0, 1, ldrr},
	{0x7A, "LD A, D", A, D, 0, 1, ldrr},
	{0x7B, "LD A, E", A, E, 0, 1, ldrr},
	{0x7C, "LD A, H", A, H, 0, 1, ldrr},
	{0x7D, "LD A, L", A, L, 0, 1, ldrr},
	{0x7E, "LD A, (HL)", A, HL, 0, 2, ldrm16},
	{0x7F, "LD A, A", A, A, 0, 1, ldrr},
	{0x80, "ADD A, B", A, B, 0, 1, addr},
	{0x81, "ADD A, C", A, C, 0, 1, addr},
	{0x82, "ADD A, D", A, D, 0, 1, addr},
	{0x83, "ADD A, E", A, E, 0, 1, addr},
	{0x84, "ADD A, H", A, H, 0, 1, addr},
	{0x85, "ADD A, L", A, L, 0, 1, addr},
	{0x86, "ADD A, (HL)", A, HL, 0, 2, addHL},
	{0x87, "ADD A, A", A, A, 0, 1, addr},
	{0x88, "ADC A, B", A, B, 0, 1, adcr},
	{0x89, "ADC A, C", A, C, 0, 1, adcr},
	{0x8A, "ADC A, D", A, D, 0, 1, adcr},
	{0x8B, "ADC A, E", A, E, 0, 1, adcr},
	{0x8C, "ADC A, H", A, H, 0, 1, adcr},
	{0x8D, "ADC A, L", A, L, 0, 1, adcr},
	{0x8E, "ADC A, (HL)", A, HL, 0, 2, adcm16},
	{0x8F, "ADC A, A", A, A, 0, 1, adcr},
	{0x90, "SUB B", B, 0, 0, 1, subr},
	{0x91, "SUB C", C, 0, 0, 1, subr},
	{0x92, "SUB D", D, 0, 0, 1, subr},
	{0x93, "SUB E", E, 0, 0, 1, subr},
	{0x94, "SUB H", H, 0, 0, 1, subr},
	{0x95, "SUB L", L, 0, 0, 1, subr},
	{0x96, "SUB (HL)", HL, 0, 0, 2, subHL},
	{0x97, "SUB A", A, 0, 0, 1, subr},
	{0x98, "SBC A, B", A, B, 0, 1, sbcr},
	{0x99, "SBC A, C", A, C, 0, 1, sbcr},
	{0x9A, "SBC A, D", A, D, 0, 1, sbcr},
	{0x9B, "SBC A, E", A, E, 0, 1, sbcr},
	{0x9C, "SBC A, H", A, H, 0, 1, sbcr},
	{0x9D, "SBC A, L", A, L, 0, 1, sbcr},
	{0x9E, "SBC A, (HL)", A, HL, 0, 2, sbcm16},
	{0x9F, "SBC A, A", A, A, 0, 1, sbcr},
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
	{0xC7, "RST 00H", 0x00, 0, 0, 4, rst},
	{0xC8, "RET Z", flagZ, 0, 0, 2, retf},
	{0xC9, "RET", 0, 0, 0, 4, ret},
	{0xCA, "JP Z,a16", flagZ, 0, 2, 3, jpfa16},
	{0xCB, "PREFIX CB", 0, 0, 0, 1, notimplemented},
	{0xCC, "CALL Z,a16", flagZ, 0, 2, 3, callf},
	{0xCD, "CALL a16", 0, 0, 2, 6, call},
	{0xCE, "ADC A,d8", A, 0, 1, 2, adcd},
	{0xCF, "RST 08H", 0x08, 0, 0, 4, rst},
	{0xD0, "RET NC", flagC, 0, 0, 2, retnf},
	{0xD1, "POP DE", DE, 0, 0, 3, pop},
	{0xD2, "JP NC,a16", flagC, 0, 2, 3, jpnfa16},
	{0xD3, "EMPTY", 0, 0, 0, 0, notimplemented},
	{0xD4, "CALL NC,a16", flagC, 0, 2, 3, callnf},
	{0xD5, "PUSH DE", DE, 0, 0, 4, push},
	{0xD6, "SUB d8", 0, 0, 1, 2, subd8},
	{0xD7, "RST 10H", 0x10, 0, 0, 4, rst},
	{0xD8, "RET C", flagC, 0, 0, 2, retf},
	{0xD9, "RETI", 0, 0, 0, 4, reti},
	{0xDA, "JP C,a16", flagC, 0, 2, 3, jpfa16},
	{0xDB, "EMPTY", 0, 0, 0, 0, notimplemented},
	{0xDC, "CALL C,a16", flagC, 0, 2, 3, callf},
	{0xDD, "EMPTY", 0, 0, 0, 0, notimplemented},
	{0xDE, "SBC A,d8", A, 0, 0, 2, sbcd},
	{0xDF, "RST 18H", 0x18, 0, 0, 4, rst},
	{0xE0, "LDH (a8),A", 0, A, 0, 3, ldar},
	{0xE1, "POP HL", HL, 0, 0, 3, pop},
	{0xE2, "LD (C),A", C, A, 1, 2, ldmr},
	{0xE3, "EMPTY", 0, 0, 0, 0, notimplemented},
	{0xE4, "EMPTY", 0, 0, 0, 0, notimplemented},
	{0xE5, "PUSH HL", HL, 0, 0, 4, push},
	{0xE6, "AND d8", 0, 0, 1, 2, andd8},
	{0xE7, "RST 20H", 0x20, 0, 0, 4, rst},
	{0xE8, "ADD SP,r8", SP, 0, 0, 4, addr16d},
	{0xE9, "JP (HL)", HL, 0, 0, 1, jpm16},
	{0xEA, "LD (a16),A", 0, A, 0, 4, lda16r},
	{0xEB, "EMPTY", 0, 0, 0, 0, notimplemented},
	{0xEC, "EMPTY", 0, 0, 0, 0, notimplemented},
	{0xED, "EMPTY", 0, 0, 0, 0, notimplemented},
	{0xEE, "XOR d8", 0, 0, 1, 2, xord8},
	{0xEF, "RST 28H", 0x28, 0, 0, 4, rst},
	{0xF0, "LDH A,(a8)", A, 0, 1, 3, ldra},
	{0xF1, "POP AF", AF, 0, 0, 3, pop},
	{0xF2, "LD A,(C)", A, C, 1, 2, ldrm},
	{0xF3, "DI", 0, 0, 0, 1, di},
	{0xF4, "EMPTY", 0, 0, 0, 0, notimplemented},
	{0xF5, "PUSH AF", AF, 0, 0, 4, push},
	{0xF6, "OR d8", 0, 0, 0, 2, ord8},
	{0xF7, "RST 30H", 0x30, 0, 0, 4, rst},
	{0xF8, "LD HL,SP+r8", HL, SP, 1, 3, ldr16r16d},
	{0xF9, "LD SP,HL", SP, HL, 0, 2, ldr16r16},
	{0xFA, "LD A,(a16)", A, 0, 2, 4, ldra16},
	{0xFB, "EI", 0, 0, 0, 1, ei},
	{0xFC, "EMPTY", 0, 0, 0, 0, notimplemented},
	{0xFD, "EMPTY", 0, 0, 0, 0, notimplemented},
	{0xFE, "CP d8", 0, 0, 1, 2, cpd8},
	{0xFF, "RST 38H", 0x38, 0, 0, 4, rst},
}

func nop(c *CPU, _ int, _ int) {}

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

// LD r1,r2
// Put value r2 into r1
func ldrr(c *CPU, r1 int, r2 int) {
	c.Reg.R[r1] = c.Reg.R[r2]
}

// LD r1, (r2)
// Write r2 value into r1
func ldrm(c *CPU, R1 int, R2 int) {
	c.Reg.R[R1] = c.Bus.ReadByte(types.Addr(0xFF00 | types.Addr(c.Reg.R[R2])))
}

// LD r1, (r2)
// Write r2 value into r1
func ldrm16(c *CPU, R1 int, R2 int) {
	c.Reg.R[R1] = c.Bus.ReadByte(c.Reg.R16(int(R2)))
}

// LD r1, d8
func ldrd(c *CPU, r8 int, _ int) {
	c.Reg.R[r8] = c.fetch()
}

// LDH R,(a8)
func ldra(c *CPU, R1 int, _ int) {
	c.Reg.R[R1] = c.Bus.ReadByte(util.Byte2Addr(byte(0xFF), c.fetch()))
}

func ldra16(c *CPU, R1 int, _ int) {
	c.Reg.R[R1] = c.Bus.ReadByte(c.fetch16())
}

// func ldr16(r16, r16d, d16)

// LD r1, r2
func ldr16r16(c *CPU, R1 int, R2 int) {
	c.Reg.setR16(int(R1), c.Reg.R16(int(R2)))
}

// LD r1, r2+d
func ldr16r16d(c *CPU, R1 int, R2 int) {
	d := c.fetch()
	sp := c.Reg.R16(SP)
	v := sp + types.Addr((int8(d)))
	c.Reg.setR16(int(R1), types.Addr(v))

	carry := sp ^ types.Addr(d) ^ (sp + types.Addr(d))

	c.Reg.setZNHC(false, false, carry&0x10 == 0x10, carry&0x100 == 0x100)
}

// LD r1, d16
func ldr16d16(c *CPU, R1 int, _ int) {
	c.Reg.setR16(int(R1), c.fetch16())
}

// func ldm(r)

// LD (C), A
func ldmr(c *CPU, R1 int, R2 int) {
	addr := util.Byte2Addr(0xFF, c.Reg.R[R1])
	c.Bus.WriteByte(addr, c.Reg.R[R2])
}

// func ldm16(r, d)

// LD (r1), r2
func ldm16r(c *CPU, R1 int, R2 int) {
	c.Bus.WriteByte(c.Reg.R16(int(R1)), c.Reg.R[R2])
}

// LD (HL),d8
func ldm16d(c *CPU, R1 int, R2 int) {
	c.Bus.WriteByte(c.Reg.R16(int(R1)), c.fetch())
}

// func lda(r)

func ldar(c *CPU, _ int, R2 int) {
	addr := util.Byte2Addr(0xFF, c.fetch())
	c.Bus.WriteByte(addr, c.Reg.R[R2])
}

// func lda16(r, r16)

func lda16r(c *CPU, _ int, R2 int) {
	addr := c.fetch16()
	c.Bus.WriteByte(addr, c.Reg.R[R2])
}

func lda16r16(c *CPU, _ int, R2 int) {
	addr := c.fetch16()
	r16 := c.Reg.R16(int(R2))
	c.Bus.WriteByte(addr, util.ExtractLower(r16))
	c.Bus.WriteByte(addr+1, util.ExtractUpper(r16))
}

// arithmetic
func incr(c *CPU, r8 int, _ int) {
	r := c.Reg.R[r8]

	incremented := r + 0x01
	c.Reg.R[r8] = incremented

	h := (incremented^0x01^r)&0x10 == 0x10
	c.Reg.setZNH(incremented == 0, false, h)
}

func incr16(c *CPU, r16 int, _ int) {
	c.Reg.setR16(int(r16), types.Addr(c.Reg.R16(int(r16))+1))
}

func incm16(c *CPU, r16 int, _ int) {
	d := c.Bus.ReadByte(c.Reg.R16(int(r16)))
	v := d + 1
	carryBits := d ^ 1 ^ v
	c.Bus.WriteByte(c.Reg.R16(int(r16)), byte(v))

	c.Reg.setZNH(v == 0, false, carryBits&0x10 == 0x10)
}

func decr(c *CPU, r8 int, _ int) {
	r := c.Reg.R[r8]
	decremented := r - 0x01
	c.Reg.R[r8] = decremented

	h := (decremented^0x01^r)&0x10 == 0x10
	c.Reg.setZNH(decremented == 0, true, h)
}

func decr16(c *CPU, r16 int, _ int) {
	c.Reg.setR16(int(r16), types.Addr(c.Reg.R16(int(r16))-1))
}

func decm16(c *CPU, r16 int, _ int) {
	d := c.Bus.ReadByte(c.Reg.R16(int(r16)))
	v := d - 1
	carryBits := d ^ 1 ^ v
	c.Bus.WriteByte(c.Reg.R16(int(r16)), byte(v))

	c.Reg.setZNH(v == 0, true, carryBits&0x10 == 0x10)
}

func _and(c *CPU, buf byte) {
	c.Reg.R[A] &= buf
	c.Reg.setZNHC(c.Reg.R[A] == 0, false, true, false)
}

func andr(c *CPU, r8 int, _ int) {
	buf := c.Reg.R[r8]
	_and(c, buf)
}

func andHL(c *CPU, r16 int, _ int) {
	buf := c.Bus.ReadByte(c.Reg.R16(int(r16)))
	_and(c, buf)
}

func andd8(c *CPU, _ int, _ int) {
	buf := c.fetch()
	_and(c, buf)
}

func _or(c *CPU, buf byte) {
	c.Reg.R[A] |= buf
	c.Reg.setZNHC(c.Reg.R[A] == 0, false, false, false)
}

func orr(c *CPU, r8 int, _ int) {
	buf := c.Reg.R[r8]
	_or(c, buf)
}

func orHL(c *CPU, r16 int, _ int) {
	buf := c.Bus.ReadByte(c.Reg.R16(int(r16)))
	_or(c, buf)
}

func ord8(c *CPU, r8 int, _ int) {
	buf := c.fetch()
	_or(c, buf)
}

func _xor(c *CPU, buf byte) {
	c.Reg.R[A] ^= buf
	c.Reg.setZNHC(c.Reg.R[A] == 0, false, false, false)
}

func xorr(c *CPU, r8 int, _ int) {
	buf := c.Reg.R[r8]
	_xor(c, buf)
}

func xorHL(c *CPU, r16 int, _ int) {
	buf := c.Bus.ReadByte(c.Reg.R16(int(r16)))
	_xor(c, buf)
}

func xord8(c *CPU, _ int, _ int) {
	buf := c.fetch()
	_xor(c, buf)
}

func _cp(c *CPU, v byte) {
	a := c.Reg.R[A]
	c.Reg.setZNHC(a == v, true, a&0x0F < v&0x0F, a < v)
}

func cpr(c *CPU, r8 int, _ int) {
	v := c.Reg.R[r8]
	_cp(c, v)
}

func cpHL(c *CPU, r16 int, _ int) {
	v := c.Bus.ReadByte(c.Reg.R16(int(r16)))
	_cp(c, v)
}

func cpd8(c *CPU, _ int, _ int) {
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
	c.Reg.setZNHC(byte(v) == 0, false, flag_h, flag_c)
}

func addr(c *CPU, _ int, r8 int) {
	r := c.Reg.R[r8]
	_add(c, r)
}

func _addr16(c *CPU, v1 types.Addr, v2 types.Addr) types.Addr {
	v := v1 + v2

	c.Reg.setNHC(false, (v^v1^v2)&0x1000 == 0x1000, v < v1)
	return v
}

func addr16r16(c *CPU, r1 int, r2 int) {
	a := c.Reg.R16(int(r1))
	b := c.Reg.R16(int(r2))
	c.Reg.setR16(int(r1), _addr16(c, a, b))
}

func addr16d(c *CPU, r16 int, _ int) {
	v1 := c.Reg.R16(r16)
	v2 := int8(c.fetch())

	v := uint32(v1) + uint32(v2)

	carry := uint32(v1) ^ uint32(v2) ^ v

	c.Reg.setR16(r16, types.Addr(v))
	c.Reg.setZNHC(false, false, carry&0x10 == 0x10, carry&0x100 == 0x100)
}

func addHL(c *CPU, _ int, r16 int) {
	v := c.Bus.ReadByte(c.Reg.R16(int(r16)))
	_add(c, v)
}

func addd8(c *CPU, _ int, r8 int) {
	v := c.fetch()
	_add(c, v)
}

func _adc(c *CPU, r byte) {
	a := c.Reg.R[A]
	carry := c.Reg.isSet(flagC)

	v := a + r + byte(util.Bool2Int8(carry))

	c.Reg.R[A] = v
	flag_h := a&0x0F+r&0x0F+byte(util.Bool2Int8(carry)) > 0x0F
	flag_c := uint(a&0xFF)+uint(r&0xFF)+uint(util.Bool2Int8(carry)) > 0xFF
	c.Reg.setZNHC(v == 0, false, flag_h, flag_c)
}

// ADC A,R
func adcr(c *CPU, _ int, r2 int) {
	r := c.Reg.R[r2]
	_adc(c, r)
}

func adcm16(c *CPU, _ int, r2 int) {
	r := c.Bus.ReadByte(c.Reg.R16(int(r2)))
	_adc(c, r)
}

func adcd(c *CPU, _ int, _ int) {
	r := c.fetch()
	_adc(c, r)
}

func _sub(c *CPU, b byte) {
	a := c.Reg.R[A]
	v := a - b
	carryBits := a ^ b ^ v
	flag_h := carryBits&(1<<4) != 0
	flag_c := a < v

	c.Reg.R[A] = v
	c.Reg.setZNHC(byte(v) == 0, true, flag_h, flag_c)
}

func subr(c *CPU, r8 int, _ int) {
	v := c.Reg.R[r8]
	_sub(c, v)
}

func subHL(c *CPU, r16 int, _ int) {
	v := c.Bus.ReadByte(c.Reg.R16(int(r16)))
	_sub(c, v)
}

func subd8(c *CPU, _ int, _ int) {
	r := c.fetch()
	_sub(c, r)
}

func _sbc(c *CPU, r byte) {
	a := c.Reg.R[A]
	carry := util.Bool2Int8(c.Reg.isSet(flagC))

	v := a - (r + byte(carry))
	c.Reg.R[A] = byte(v)

	flag_h := a&0x0F < r&0x0F+byte(carry)
	flag_c := uint16(a) < uint16(r)+uint16(carry)
	c.Reg.setZNHC(byte(v) == 0, true, flag_h, flag_c)
}

// SBC A,R
func sbcr(c *CPU, _ int, r2 int) {
	r := c.Reg.R[r2]
	_sbc(c, r)
}

func sbcm16(c *CPU, _ int, r2 int) {
	r := c.getRegMem(r2)
	_sbc(c, r)
}

func sbcd(c *CPU, _ int, _ int) {
	r := c.fetch()
	_sbc(c, r)
}

// ret
func ret(c *CPU, _ int, _ int) {
	c.popPC()
}

func retf(c *CPU, R1 int, _ int) {
	if c.Reg.isSet(R1) {
		c.popPC()
	}
}

func retnf(c *CPU, R1 int, _ int) {
	if !c.Reg.isSet(R1) {
		c.popPC()
	}
}

func reti(c *CPU, _ int, _ int) {
	c.popPC()
	c.IRQ.Enable()
}

// -----jp-----

func _jp(c *CPU, addr types.Addr) {
	c.Reg.PC = addr
}

// JP a16
func jpa16(c *CPU, _ int, _ int) {
	_jp(c, c.fetch16())
}

// JP flag, a16
// jump when flag = 1
func jpfa16(c *CPU, flag int, _ int) {
	addr := c.fetch16()
	if c.Reg.isSet(flag) {
		_jp(c, addr)
	}
}

// JP Nflag, a16
// jump when flag = 0
func jpnfa16(c *CPU, flag int, _ int) {
	addr := c.fetch16()
	if !c.Reg.isSet(flag) {
		_jp(c, addr)
	}
}

// JP (r16)
func jpm16(c *CPU, R1 int, _ int) {
	_jp(c, c.Reg.R16(int(R1)))
}

// -----jr-----
func _jr(c *CPU, addr int8) {
	c.Reg.PC = types.Addr(int32(c.Reg.PC) + int32(addr))
}

// r8 is a signed data, which are added to PC
func jrr8(c *CPU, _ int, _ int) {
	n := c.fetch()
	_jr(c, int8(n))
}

// r8 is a signed data, which are added to PC
func jrfr8(c *CPU, flag int, _ int) {
	n := c.fetch()
	// flag is set
	if c.Reg.isSet(flag) {
		_jr(c, int8(n))
	}
}

// r8 is a signed data, which are added to PC
func jrnfr8(c *CPU, flag int, _ int) {
	n := c.fetch()
	// flag is not set
	if !c.Reg.isSet(flag) {
		_jr(c, int8(n))
	}
}

// -----rst------

// RST n
// push and jump to n
func rst(c *CPU, n int, _ int) {
	c.pushPC()
	c.Reg.PC = types.Addr(n)
}

// -----push-----
func push(c *CPU, r16 int, _ int) {
	buf := c.Reg.R16(int(r16))
	upper := util.ExtractUpper(types.Addr(buf))
	lower := util.ExtractLower(types.Addr(buf))
	c.push(upper)
	c.push(lower)
}

// -----pop------
func pop(c *CPU, r16 int, _ int) {
	lower := c.pop()
	upper := c.pop()

	if r16 == AF {
		lower &= 0xF0
	}

	c.Reg.setR16(int(r16), types.Addr(int16(upper)<<8|int16(lower)))
}

// -----call-----
func _call(c *CPU, dest types.Addr) {
	c.pushPC()
	c.Reg.PC = dest

}

func call(c *CPU, _ int, _ int) {
	dest := c.fetch16()
	_call(c, dest)
}

func callf(c *CPU, flag int, _ int) {
	dest := c.fetch16()
	if c.Reg.isSet(flag) {
		_call(c, dest)
	}
}

func callnf(c *CPU, flag int, _ int) {
	dest := c.fetch16()
	if !c.Reg.isSet(flag) {
		_call(c, dest)
	}
}

// -----misc-----

func cpl(c *CPU, _ int, _ int) {
	c.Reg.R[A] = ^c.Reg.R[A]
	c.Reg.setNH(true, true)
}

func scf(c *CPU, _ int, _ int) {
	c.Reg.setNHC(false, false, true)
}

func ccf(c *CPU, _ int, _ int) {
	c.Reg.setNHC(false, false, !c.Reg.isSet(flagC))
}

// @see https://donkeyhacks.zouri.jp/html/En-Us/snes/apu/spc700/daa.html
func daa(c *CPU, _ int, _ int) {
	a := types.Addr(c.Reg.R[A])

	if !c.Reg.isSet(flagN) {
		if a&0x0F >= 0x0A || c.Reg.isSet(flagH) {
			a += 0x06
		}
		if a >= 0xA0 || c.Reg.isSet(flagC) {
			a += 0x60
		}
	} else {
		if c.Reg.isSet(flagH) {
			a = (a - 0x06) & 0xFF
		}
		if c.Reg.isSet(flagC) {
			a -= 0x60
		}
	}

	c.Reg.R[A] = byte(a)
	c.Reg.setF(flagZ, byte(a) == 0)
	c.Reg.setF(flagH, false)
	if a&0x100 == 0x100 {
		c.Reg.setF(flagC, true)
	}
}

func stop(c *CPU, _ int, _ int) {
	log.Debug("TODO: implement")
}

// desable interrupt
func di(c *CPU, _ int, _ int) {
	c.IRQ.Disable()
}

// enable interrupt
func ei(c *CPU, _ int, _ int) {
	c.IRQ.Enable()
}

func halt(c *CPU, _ int, _ int) {
	c.Halt = true
}

func notimplemented(c *CPU, _ int, _ int) {
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
	{0x20, "SLA B", B, 0, 1, 2, slar},
	{0x21, "SLA C", C, 0, 1, 2, slar},
	{0x22, "SLA D", D, 0, 1, 2, slar},
	{0x23, "SLA E", E, 0, 1, 2, slar},
	{0x24, "SLA H", H, 0, 1, 2, slar},
	{0x25, "SLA L", L, 0, 1, 2, slar},
	{0x26, "SLA (HL)", HL, 0, 1, 4, slam16},
	{0x27, "SLA A", A, 0, 1, 2, slar},
	{0x28, "SRA B", B, 0, 1, 2, srar},
	{0x29, "SRA C", C, 0, 1, 2, srar},
	{0x2A, "SRA D", D, 0, 1, 2, srar},
	{0x2B, "SRA E", E, 0, 1, 2, srar},
	{0x2C, "SRA H", H, 0, 1, 2, srar},
	{0x2D, "SRA L", L, 0, 1, 2, srar},
	{0x2E, "SRA (HL)", HL, 0, 1, 4, sram16},
	{0x2F, "SRA A", A, 0, 1, 2, srar},
	{0x30, "SWAP B", B, 0, 1, 2, swapr},
	{0x31, "SWAP C", C, 0, 1, 2, swapr},
	{0x32, "SWAP D", D, 0, 1, 2, swapr},
	{0x33, "SWAP E", E, 0, 1, 2, swapr},
	{0x34, "SWAP H", H, 0, 1, 2, swapr},
	{0x35, "SWAP L", L, 0, 1, 2, swapr},
	{0x36, "SWAP (HL)", HL, 0, 1, 4, swapm16},
	{0x37, "SWAP A", A, 0, 1, 2, swapr},
	{0x38, "SRL B", B, 0, 1, 2, srlr},
	{0x39, "SRL C", C, 0, 1, 2, srlr},
	{0x3A, "SRL D", D, 0, 1, 2, srlr},
	{0x3B, "SRL E", E, 0, 1, 2, srlr},
	{0x3C, "SRL H", H, 0, 1, 2, srlr},
	{0x3D, "SRL L", L, 0, 1, 2, srlr},
	{0x3E, "SRL (HL)", HL, 0, 1, 4, srlm16},
	{0x3F, "SRL A", A, 0, 1, 2, srlr},
	{0x40, "BIT 0,B", 0, B, 1, 2, bitr},
	{0x41, "BIT 0,C", 0, C, 1, 2, bitr},
	{0x42, "BIT 0,D", 0, D, 1, 2, bitr},
	{0x43, "BIT 0,E", 0, E, 1, 2, bitr},
	{0x44, "BIT 0,H", 0, H, 1, 2, bitr},
	{0x45, "BIT 0,L", 0, L, 1, 2, bitr},
	{0x46, "BIT 0,(HL)", 0, HL, 1, 3, bitm16},
	{0x47, "BIT 0,A", 0, A, 1, 2, bitr},
	{0x48, "BIT 1,B", 1, B, 1, 2, bitr},
	{0x49, "BIT 1,C", 1, C, 1, 2, bitr},
	{0x4A, "BIT 1,D", 1, D, 1, 2, bitr},
	{0x4B, "BIT 1,E", 1, E, 1, 2, bitr},
	{0x4C, "BIT 1,H", 1, H, 1, 2, bitr},
	{0x4D, "BIT 1,L", 1, L, 1, 2, bitr},
	{0x4E, "BIT 1,(HL)", 1, HL, 1, 3, bitm16},
	{0x4F, "BIT 1,A", 1, A, 1, 2, bitr},
	{0x50, "BIT 2,B", 2, B, 1, 2, bitr},
	{0x51, "BIT 2,C", 2, C, 1, 2, bitr},
	{0x52, "BIT 2,D", 2, D, 1, 2, bitr},
	{0x53, "BIT 2,E", 2, E, 1, 2, bitr},
	{0x54, "BIT 2,H", 2, H, 1, 2, bitr},
	{0x55, "BIT 2,L", 2, L, 1, 2, bitr},
	{0x56, "BIT 2,(HL)", 2, HL, 1, 3, bitm16},
	{0x57, "BIT 2,A", 2, A, 1, 2, bitr},
	{0x58, "BIT 3,B", 3, B, 1, 2, bitr},
	{0x59, "BIT 3,C", 3, C, 1, 2, bitr},
	{0x5A, "BIT 3,D", 3, D, 1, 2, bitr},
	{0x5B, "BIT 3,E", 3, E, 1, 2, bitr},
	{0x5C, "BIT 3,H", 3, H, 1, 2, bitr},
	{0x5D, "BIT 3,L", 3, L, 1, 2, bitr},
	{0x5E, "BIT 3,(HL)", 3, HL, 1, 3, bitm16},
	{0x5F, "BIT 3,A", 3, A, 1, 2, bitr},
	{0x60, "BIT 4,B", 4, B, 1, 2, bitr},
	{0x61, "BIT 4,C", 4, C, 1, 2, bitr},
	{0x62, "BIT 4,D", 4, D, 1, 2, bitr},
	{0x63, "BIT 4,E", 4, E, 1, 2, bitr},
	{0x64, "BIT 4,H", 4, H, 1, 2, bitr},
	{0x65, "BIT 4,L", 4, L, 1, 2, bitr},
	{0x66, "BIT 4,(HL)", 4, HL, 1, 3, bitm16},
	{0x67, "BIT 4,A", 4, A, 1, 2, bitr},
	{0x68, "BIT 5,B", 5, B, 1, 2, bitr},
	{0x69, "BIT 5,C", 5, C, 1, 2, bitr},
	{0x6A, "BIT 5,D", 5, D, 1, 2, bitr},
	{0x6B, "BIT 5,E", 5, E, 1, 2, bitr},
	{0x6C, "BIT 5,H", 5, H, 1, 2, bitr},
	{0x6D, "BIT 5,L", 5, L, 1, 2, bitr},
	{0x6E, "BIT 5,(HL)", 5, HL, 1, 3, bitm16},
	{0x6F, "BIT 5,A", 5, A, 1, 2, bitr},
	{0x70, "BIT 6,B", 6, B, 1, 2, bitr},
	{0x71, "BIT 6,C", 6, C, 1, 2, bitr},
	{0x72, "BIT 6,D", 6, D, 1, 2, bitr},
	{0x73, "BIT 6,E", 6, E, 1, 2, bitr},
	{0x74, "BIT 6,H", 6, H, 1, 2, bitr},
	{0x75, "BIT 6,L", 6, L, 1, 2, bitr},
	{0x76, "BIT 6,(HL)", 6, HL, 1, 3, bitm16},
	{0x77, "BIT 6,A", 6, A, 1, 2, bitr},
	{0x78, "BIT 7,B", 7, B, 1, 2, bitr},
	{0x79, "BIT 7,C", 7, C, 1, 2, bitr},
	{0x7A, "BIT 7,D", 7, D, 1, 2, bitr},
	{0x7B, "BIT 7,E", 7, E, 1, 2, bitr},
	{0x7C, "BIT 7,H", 7, H, 1, 2, bitr},
	{0x7D, "BIT 7,L", 7, L, 1, 2, bitr},
	{0x7E, "BIT 7,(HL)", 7, HL, 1, 3, bitm16},
	{0x7F, "BIT 7,A", 7, A, 1, 2, bitr},
	{0x80, "RES 0,B", 0, B, 1, 2, resr},
	{0x81, "RES 0,C", 0, C, 1, 2, resr},
	{0x82, "RES 0,D", 0, D, 1, 2, resr},
	{0x83, "RES 0,E", 0, E, 1, 2, resr},
	{0x84, "RES 0,H", 0, H, 1, 2, resr},
	{0x85, "RES 0,L", 0, L, 1, 2, resr},
	{0x86, "RES 0,(HL)", 0, HL, 1, 4, resm16},
	{0x87, "RES 0,A", 0, A, 1, 2, resr},
	{0x88, "RES 1,B", 1, B, 1, 2, resr},
	{0x89, "RES 1,C", 1, C, 1, 2, resr},
	{0x8A, "RES 1,D", 1, D, 1, 2, resr},
	{0x8B, "RES 1,E", 1, E, 1, 2, resr},
	{0x8C, "RES 1,H", 1, H, 1, 2, resr},
	{0x8D, "RES 1,L", 1, L, 1, 2, resr},
	{0x8E, "RES 1,(HL)", 1, HL, 1, 4, resm16},
	{0x8F, "RES 1,A", 1, A, 1, 2, resr},
	{0x90, "RES 2,B", 2, B, 1, 2, resr},
	{0x91, "RES 2,C", 2, C, 1, 2, resr},
	{0x92, "RES 2,D", 2, D, 1, 2, resr},
	{0x93, "RES 2,E", 2, E, 1, 2, resr},
	{0x94, "RES 2,H", 2, H, 1, 2, resr},
	{0x95, "RES 2,L", 2, L, 1, 2, resr},
	{0x96, "RES 2,(HL)", 2, HL, 1, 4, resm16},
	{0x97, "RES 2,A", 2, A, 1, 2, resr},
	{0x98, "RES 3,B", 3, B, 1, 2, resr},
	{0x99, "RES 3,C", 3, C, 1, 2, resr},
	{0x9A, "RES 3,D", 3, D, 1, 2, resr},
	{0x9B, "RES 3,E", 3, E, 1, 2, resr},
	{0x9C, "RES 3,H", 3, H, 1, 2, resr},
	{0x9D, "RES 3,L", 3, L, 1, 2, resr},
	{0x9E, "RES 3,(HL)", 3, HL, 1, 4, resm16},
	{0x9F, "RES 3,A", 3, A, 1, 2, resr},
	{0xA0, "RES 4,B", 4, B, 1, 2, resr},
	{0xA1, "RES 4,C", 4, C, 1, 2, resr},
	{0xA2, "RES 4,D", 4, D, 1, 2, resr},
	{0xA3, "RES 4,E", 4, E, 1, 2, resr},
	{0xA4, "RES 4,H", 4, H, 1, 2, resr},
	{0xA5, "RES 4,L", 4, L, 1, 2, resr},
	{0xA6, "RES 4,(HL)", 4, HL, 1, 4, resm16},
	{0xA7, "RES 4,A", 4, A, 1, 2, resr},
	{0xA8, "RES 5,B", 5, B, 1, 2, resr},
	{0xA9, "RES 5,C", 5, C, 1, 2, resr},
	{0xAA, "RES 5,D", 5, D, 1, 2, resr},
	{0xAB, "RES 5,E", 5, E, 1, 2, resr},
	{0xAC, "RES 5,H", 5, H, 1, 2, resr},
	{0xAD, "RES 5,L", 5, L, 1, 2, resr},
	{0xAE, "RES 5,(HL)", 5, HL, 1, 4, resm16},
	{0xAF, "RES 5,A", 5, A, 1, 2, resr},
	{0xB0, "RES 6,B", 6, B, 1, 2, resr},
	{0xB1, "RES 6,C", 6, C, 1, 2, resr},
	{0xB2, "RES 6,D", 6, D, 1, 2, resr},
	{0xB3, "RES 6,E", 6, E, 1, 2, resr},
	{0xB4, "RES 6,H", 6, H, 1, 2, resr},
	{0xB5, "RES 6,L", 6, L, 1, 2, resr},
	{0xB6, "RES 6,(HL)", 6, HL, 1, 4, resm16},
	{0xB7, "RES 6,A", 6, A, 1, 2, resr},
	{0xB8, "RES 7,B", 7, B, 1, 2, resr},
	{0xB9, "RES 7,C", 7, C, 1, 2, resr},
	{0xBA, "RES 7,D", 7, D, 1, 2, resr},
	{0xBB, "RES 7,E", 7, E, 1, 2, resr},
	{0xBC, "RES 7,H", 7, H, 1, 2, resr},
	{0xBD, "RES 7,L", 7, L, 1, 2, resr},
	{0xBE, "RES 7,(HL)", 7, HL, 1, 4, resm16},
	{0xBF, "RES 7,A", 7, A, 1, 2, resr},
	{0xC0, "SET 0,B", 0, B, 1, 2, setr},
	{0xC1, "SET 0,C", 0, C, 1, 2, setr},
	{0xC2, "SET 0,D", 0, D, 1, 2, setr},
	{0xC3, "SET 0,E", 0, E, 1, 2, setr},
	{0xC4, "SET 0,H", 0, H, 1, 2, setr},
	{0xC5, "SET 0,L", 0, L, 1, 2, setr},
	{0xC6, "SET 0,(HL)", 0, HL, 1, 4, setm16},
	{0xC7, "SET 0,A", 0, A, 1, 2, setr},
	{0xC8, "SET 1,B", 1, B, 1, 2, setr},
	{0xC9, "SET 1,C", 1, C, 1, 2, setr},
	{0xCA, "SET 1,D", 1, D, 1, 2, setr},
	{0xCB, "SET 1,E", 1, E, 1, 2, setr},
	{0xCC, "SET 1,H", 1, H, 1, 2, setr},
	{0xCD, "SET 1,L", 1, L, 1, 2, setr},
	{0xCE, "SET 1,(HL)", 1, HL, 1, 4, setm16},
	{0xCF, "SET 1,A", 1, A, 1, 2, setr},
	{0xD0, "SET 2,B", 2, B, 1, 2, setr},
	{0xD1, "SET 2,C", 2, C, 1, 2, setr},
	{0xD2, "SET 2,D", 2, D, 1, 2, setr},
	{0xD3, "SET 2,E", 2, E, 1, 2, setr},
	{0xD4, "SET 2,H", 2, H, 1, 2, setr},
	{0xD5, "SET 2,L", 2, L, 1, 2, setr},
	{0xD6, "SET 2,(HL)", 2, HL, 1, 4, setm16},
	{0xD7, "SET 2,A", 2, A, 1, 2, setr},
	{0xD8, "SET 3,B", 3, B, 1, 2, setr},
	{0xD9, "SET 3,C", 3, C, 1, 2, setr},
	{0xDA, "SET 3,D", 3, D, 1, 2, setr},
	{0xDB, "SET 3,E", 3, E, 1, 2, setr},
	{0xDC, "SET 3,H", 3, H, 1, 2, setr},
	{0xDD, "SET 3,L", 3, L, 1, 2, setr},
	{0xDE, "SET 3,(HL)", 3, HL, 1, 4, setm16},
	{0xDF, "SET 3,A", 3, A, 1, 2, setr},
	{0xE0, "SET 4,B", 4, B, 1, 2, setr},
	{0xE1, "SET 4,C", 4, C, 1, 2, setr},
	{0xE2, "SET 4,D", 4, D, 1, 2, setr},
	{0xE3, "SET 4,E", 4, E, 1, 2, setr},
	{0xE4, "SET 4,H", 4, H, 1, 2, setr},
	{0xE5, "SET 4,L", 4, L, 1, 2, setr},
	{0xE6, "SET 4,(HL)", 4, HL, 1, 4, setm16},
	{0xE7, "SET 4,A", 4, A, 1, 2, setr},
	{0xE8, "SET 5,B", 5, B, 1, 2, setr},
	{0xE9, "SET 5,C", 5, C, 1, 2, setr},
	{0xEA, "SET 5,D", 5, D, 1, 2, setr},
	{0xEB, "SET 5,E", 5, E, 1, 2, setr},
	{0xEC, "SET 5,H", 5, H, 1, 2, setr},
	{0xED, "SET 5,L", 5, L, 1, 2, setr},
	{0xEE, "SET 5,(HL)", 5, HL, 1, 4, setm16},
	{0xEF, "SET 5,A", 5, A, 1, 2, setr},
	{0xF0, "SET 6,B", 6, B, 1, 2, setr},
	{0xF1, "SET 6,C", 6, C, 1, 2, setr},
	{0xF2, "SET 6,D", 6, D, 1, 2, setr},
	{0xF3, "SET 6,E", 6, E, 1, 2, setr},
	{0xF4, "SET 6,H", 6, H, 1, 2, setr},
	{0xF5, "SET 6,L", 6, L, 1, 2, setr},
	{0xF6, "SET 6,(HL)", 6, HL, 1, 4, setm16},
	{0xF7, "SET 6,A", 6, A, 1, 2, setr},
	{0xF8, "SET 7,B", 7, B, 1, 2, setr},
	{0xF9, "SET 7,C", 7, C, 1, 2, setr},
	{0xFA, "SET 7,D", 7, D, 1, 2, setr},
	{0xFB, "SET 7,E", 7, E, 1, 2, setr},
	{0xFC, "SET 7,H", 7, H, 1, 2, setr},
	{0xFD, "SET 7,L", 7, L, 1, 2, setr},
	{0xFE, "SET 7,(HL)", 7, HL, 1, 4, setm16},
	{0xFF, "SET 7,A", 7, A, 1, 2, setr},
}

func _rlc(c *CPU, v byte) byte {
	// check Bit 7 is set
	flag_c := v&0x80 == 0x80
	if flag_c {
		v = v<<1 | 0x01
	} else {
		v = v << 1
	}
	c.Reg.setZNHC(v == 0, false, false, flag_c)

	return v
}

// RLCA
func rlca(c *CPU, _ int, _ int) {
	r := c.Reg.R[A]
	c.Reg.R[A] = _rlc(c, r)
	c.Reg.setF(flagZ, false)
}

// RLC r
func rlcr(c *CPU, r8 int, _ int) {
	r := c.Reg.R[r8]
	c.Reg.R[r8] = _rlc(c, r)
}

// RLC (HL)
func rlcm16(c *CPU, r16 int, _ int) {
	addr := c.Reg.R16(int(r16))
	r := c.Bus.ReadByte(addr)
	c.Bus.WriteByte(addr, _rlc(c, r))
}

func _rrc(c *CPU, v byte) byte {
	// check Bit 0 is set
	flag_c := v&0x01 == 0x01
	if flag_c {
		v = v>>1 | 0x80
	} else {
		v = v >> 1
	}

	c.Reg.setZNHC(v == 0, false, false, flag_c)

	return v
}

// RRCA
func rrca(c *CPU, _ int, _ int) {
	r := c.Reg.R[A]
	c.Reg.R[A] = _rrc(c, r)
	c.Reg.setF(flagZ, false)
}

// RRC r
func rrcr(c *CPU, r8 int, _ int) {
	r := c.Reg.R[r8]
	c.Reg.R[r8] = _rrc(c, r)
}

// RRC (HL)
func rrcm16(c *CPU, r16 int, _ int) {
	addr := c.Reg.R16(int(r16))
	r := c.Bus.ReadByte(addr)
	c.Bus.WriteByte(addr, _rrc(c, r))
}

func _rl(c *CPU, v byte) byte {
	beforeCarry := c.Reg.isSet(flagC)
	// check Bit 0 is set
	flag_c := v&0x80 == 0x80
	v = v << 1
	if beforeCarry {
		v++
	}

	c.Reg.setZNHC(v == 0, false, false, flag_c)

	return v
}

// RLA
func rla(c *CPU, _ int, _ int) {
	r := c.Reg.R[A]
	c.Reg.R[A] = _rl(c, r)
	c.Reg.setF(flagZ, false)
}

// RL r
func rlr(c *CPU, r8 int, _ int) {
	r := c.Reg.R[r8]
	c.Reg.R[r8] = _rl(c, r)
}

// RL (HL)
func rlm16(c *CPU, r16 int, _ int) {
	addr := c.Reg.R16(int(r16))
	r := c.Bus.ReadByte(addr)
	c.Bus.WriteByte(addr, _rl(c, r))
}

func _rr(c *CPU, v byte) byte {
	beforeCarry := c.Reg.isSet(flagC)
	flag_c := v&0x01 == 0x01
	v = v >> 1
	if beforeCarry {
		v |= 0x80
	}
	c.Reg.setZNHC(v == 0, false, false, flag_c)

	return v
}

// RRA
func rra(c *CPU, _ int, _ int) {
	r := c.Reg.R[A]
	c.Reg.R[A] = _rr(c, r)
	c.Reg.setF(flagZ, false)
}

// RR r
func rrr(c *CPU, r8 int, _ int) {
	r := c.Reg.R[r8]
	c.Reg.R[r8] = _rr(c, r)
}

// RR (HL)
func rrm16(c *CPU, r16 int, _ int) {
	addr := c.Reg.R16(int(r16))
	r := c.Bus.ReadByte(addr)
	c.Bus.WriteByte(addr, _rr(c, r))
}

func _sla(c *CPU, v byte) byte {
	flag_c := v&0x80 == 0x80

	v = v << 1
	c.Reg.setZNHC(v == 0, false, false, flag_c)

	return v
}

// SLA r
func slar(c *CPU, r8 int, _ int) {
	r := c.Reg.R[r8]
	c.Reg.R[r8] = _sla(c, r)
}

// SLA (HL)
func slam16(c *CPU, r16 int, _ int) {
	addr := c.Reg.R16(int(r16))
	r := c.Bus.ReadByte(addr)
	c.Bus.WriteByte(addr, _sla(c, r))
}

func _sra(c *CPU, v byte) byte {
	flag_c := v&0x01 == 0x01
	bit7 := util.Bit(v, 7)

	v = v >> 1
	if bit7 == 1 {
		v |= 1 << 7
	}
	c.Reg.setZNHC(v == 0, false, false, flag_c)

	return v
}

// SRA r
func srar(c *CPU, r8 int, _ int) {
	r := c.Reg.R[r8]
	c.Reg.R[r8] = _sra(c, r)
}

// SRA (HL)
func sram16(c *CPU, r16 int, _ int) {
	addr := c.Reg.R16(int(r16))
	r := c.Bus.ReadByte(addr)
	c.Bus.WriteByte(addr, _sra(c, r))
}

func _swap(c *CPU, v byte) byte {
	upper := (v >> 4) & 0x0F
	lower := v & 0x0F

	b := byte(lower<<4) | upper
	c.Reg.setZNHC(b == 0, false, false, false)

	return b
}

// SWAP R
func swapr(c *CPU, r int, _ int) {
	c.Reg.R[r] = _swap(c, c.Reg.R[r])
}

// SWAP (HL)
func swapm16(c *CPU, r16 int, _ int) {
	addr := c.Reg.R16(int(r16))
	v := c.Bus.ReadByte(addr)
	c.Bus.WriteByte(addr, _swap(c, v))
}

func _srl(c *CPU, v byte) byte {
	flag_c := v&0x01 == 0x01

	v = v >> 1
	c.Reg.setZNHC(v == 0, false, false, flag_c)

	return v
}

func srlr(c *CPU, r8 int, _ int) {
	r := c.Reg.R[r8]
	c.Reg.R[r8] = _srl(c, r)
}

// SRA (HL)
func srlm16(c *CPU, r16 int, _ int) {
	addr := c.Reg.R16(int(r16))
	r := c.Bus.ReadByte(addr)
	c.Bus.WriteByte(addr, _srl(c, r))
}

func bitr(c *CPU, b int, r8 int) {
	r := c.Reg.R[r8]
	c.Reg.setZNH(util.Bit(r, int(b)) == 0, false, true)
}

func bitm16(c *CPU, b int, r16 int) {
	addr := c.Reg.R16(int(r16))
	r := c.Bus.ReadByte(addr)
	c.Reg.setZNH(util.Bit(r, int(b)) == 0, false, true)
}

func resr(c *CPU, b int, r8 int) {
	c.Reg.R[r8] = c.Reg.R[r8] & (^(1 << b))
}

func resm16(c *CPU, b int, r16 int) {
	addr := c.Reg.R16(int(r16))
	r := c.Bus.ReadByte(addr)

	c.Bus.WriteByte(addr, r&(^(1 << b)))
}

func setr(c *CPU, b int, r8 int) {
	c.Reg.R[r8] = c.Reg.R[r8] | (1 << b)
}

func setm16(c *CPU, b int, r16 int) {
	addr := c.Reg.R16(int(r16))
	r := c.Bus.ReadByte(addr)

	c.Bus.WriteByte(addr, r|(1<<b))
}
