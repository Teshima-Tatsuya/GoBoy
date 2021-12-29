package cpu

import (
	"testing"

	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/bus"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/cartridge"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/ie"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/io"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/ram"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/types"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/util"
	"github.com/stretchr/testify/assert"
)

func setupCPU() *CPU {
	romData := make([]byte, 0x8000)
	cart := cartridge.New(romData)
	vram := ram.New(0x2000)
	wram := ram.New(0x2000)
	wram2 := ram.New(0x2000)
	hram := ram.New(0x0080)
	io := io.New(0x0080)
	ie := ie.New()
	bus := bus.New(cart, vram, wram, wram2, hram, io, ie)

	return New(bus)
}

func (c *CPU) regreset() {
	c.Reg.R[A] = 0x01
	c.Reg.R[B] = 0x02
	c.Reg.R[C] = 0x03
	c.Reg.R[D] = 0x04
	c.Reg.R[E] = 0x05
	c.Reg.R[H] = 0x06
	c.Reg.R[L] = 0x07
	c.Reg.PC = 0x0100
	c.Reg.SP = 0xFFFE
}

func TestOpeCode_nop(t *testing.T) {

}

// test 0x40-0x6F (except 0xX6, 0xXE)
func TestOpeCode_ldrr(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
		r1     byte
		r2     byte
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "LD B, B",
			args: args{0x40, B, B},
		},
		{
			name: "LD B, C",
			args: args{0x41, B, C},
		},
		{
			name: "LD B, D",
			args: args{0x42, B, D},
		},
		{
			name: "LD B, E",
			args: args{0x43, B, E},
		},
		{
			name: "LD B, H",
			args: args{0x44, B, H},
		},
		{
			name: "LD B, L",
			args: args{0x45, B, L},
		},
		{
			name: "LD B, A",
			args: args{0x47, B, A},
		},
		{
			name: "LD C, B",
			args: args{0x48, C, B},
		},
		{
			name: "LD C, C",
			args: args{0x49, C, C},
		},
		{
			name: "LD C, D",
			args: args{0x4A, C, D},
		},
		{
			name: "LD C, E",
			args: args{0x4B, C, E},
		},
		{
			name: "LD C, H",
			args: args{0x4C, C, H},
		},
		{
			name: "LD C, L",
			args: args{0x4D, C, L},
		},
		{
			name: "LD C, A",
			args: args{0x4F, C, A},
		},
		{
			name: "LD D, B",
			args: args{0x50, D, B},
		},
		{
			name: "LD D, C",
			args: args{0x51, D, C},
		},
		{
			name: "LD D, D",
			args: args{0x52, D, D},
		},
		{
			name: "LD D, E",
			args: args{0x53, D, E},
		},
		{
			name: "LD D, H",
			args: args{0x54, D, H},
		},
		{
			name: "LD D, L",
			args: args{0x55, D, L},
		},
		{
			name: "LD D, A",
			args: args{0x57, D, A},
		},
		{
			name: "LD E, B",
			args: args{0x58, E, B},
		},
		{
			name: "LD E, C",
			args: args{0x59, E, C},
		},
		{
			name: "LD E, D",
			args: args{0x5A, E, D},
		},
		{
			name: "LD E, E",
			args: args{0x5B, E, E},
		},
		{
			name: "LD E, H",
			args: args{0x5C, E, H},
		},
		{
			name: "LD E, L",
			args: args{0x5D, E, L},
		},
		{
			name: "LD E, A",
			args: args{0x5F, E, A},
		},
		{
			name: "LD H, B",
			args: args{0x60, H, B},
		},
		{
			name: "LD H, C",
			args: args{0x61, H, C},
		},
		{
			name: "LD H, D",
			args: args{0x62, H, D},
		},
		{
			name: "LD H, E",
			args: args{0x63, H, E},
		},
		{
			name: "LD H, H",
			args: args{0x64, H, H},
		},
		{
			name: "LD H, L",
			args: args{0x65, H, L},
		},
		{
			name: "LD H, A",
			args: args{0x67, H, A},
		},
		{
			name: "LD L, B",
			args: args{0x68, L, B},
		},
		{
			name: "LD L, C",
			args: args{0x69, L, C},
		},
		{
			name: "LD L, D",
			args: args{0x6A, L, D},
		},
		{
			name: "LD L, E",
			args: args{0x6B, L, E},
		},
		{
			name: "LD L, H",
			args: args{0x6C, L, H},
		},
		{
			name: "LD L, L",
			args: args{0x6D, L, L},
		},
		{
			name: "LD L, A",
			args: args{0x6F, L, A},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := opCodes[tt.args.opcode]
			want := c.Reg.R[tt.args.r2]
			op.Handler(c, byte(op.R1), byte(op.R2))

			assert.Equal(t, op.R1, tt.args.r1)
			assert.Equal(t, op.R2, tt.args.r2)
			assert.Equal(t, c.Reg.R[tt.args.r1], want)
			assert.Equal(t, c.Reg.R[tt.args.r2], want)
		})
	}
}

func TestOpeCode_ldrm(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
		r1     byte
		r2     byte
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "LD A,(C)",
			args: args{0xF2, A, C},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := opCodes[tt.args.opcode]
			addr := types.Addr(0xFF03)
			want := byte(0x12)
			c.Bus.WriteByte(addr, want)
			op.Handler(c, byte(op.R1), byte(op.R2))

			assert.Equal(t, op.R1, tt.args.r1)
			assert.Equal(t, op.R2, tt.args.r2)
			assert.Equal(t, want, c.Reg.R[A])
		})
	}
}

func TestOpeCode_ldrm16(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
		r1     byte
		r2     byte
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "LD A, (BC)",
			args: args{0x0A, A, BC},
		},
		{
			name: "LD A, (DE)",
			args: args{0x1A, A, DE},
		},
		{
			name: "LD A, (HL+)",
			args: args{0x2A, A, HLI},
		},
		{
			name: "LD A, (HL-)",
			args: args{0x3A, A, HLD},
		},
		{
			name: "LD B, (HL)",
			args: args{0x46, B, HL},
		},
		{
			name: "LD C, (HL)",
			args: args{0x4E, C, HL},
		},
		{
			name: "LD D, (HL)",
			args: args{0x56, D, HL},
		},
		{
			name: "LD E, (HL)",
			args: args{0x5E, E, HL},
		},
		{
			name: "LD H, (HL)",
			args: args{0x66, H, HL},
		},
		{
			name: "LD L, (HL)",
			args: args{0x6E, L, HL},
		},
		{
			name: "LD A, (HL)",
			args: args{0x7E, A, HL},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := opCodes[tt.args.opcode]
			want := byte(0x20)
			c.Bus.WriteByte(c.Reg.R16(int(tt.args.r2)), want)
			if op.R2 == HLI {
				c.Reg.setHL(c.Reg.HL() - 1)
			} else if op.R2 == HLD {
				c.Reg.setHL(c.Reg.HL() + 1)
			}
			op.Handler(c, byte(op.R1), byte(op.R2))

			assert.Equal(t, op.R1, tt.args.r1)
			assert.Equal(t, op.R2, tt.args.r2)
			assert.Equal(t, want, c.Reg.R[tt.args.r1])
		})
	}
}

func TestOpeCode_ldrd(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
		r1     byte
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "LD B, d8",
			args: args{0x06, B},
		},
		{
			name: "LD C, d8",
			args: args{0x0E, C},
		},
		{
			name: "LD D, d8",
			args: args{0x16, D},
		},
		{
			name: "LD E, d8",
			args: args{0x1E, E},
		},
		{
			name: "LD H, d8",
			args: args{0x26, H},
		},
		{
			name: "LD L, d8",
			args: args{0x2E, L},
		},
		{
			name: "LD A, d8",
			args: args{0x3E, A},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := opCodes[tt.args.opcode]
			want := byte(0x20)
			c.Bus.WriteByte(c.Reg.PC, want)
			op.Handler(c, byte(op.R1), byte(op.R2))

			assert.Equal(t, op.R1, tt.args.r1)
			assert.Equal(t, want, c.Reg.R[op.R1])
		})
	}
}

func TestOpeCode_ldra(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
		r1     byte
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "LDH A,(a8)",
			args: args{0xF0, A},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := opCodes[tt.args.opcode]
			addr := types.Addr(0xFF12)
			a8 := byte(0x12)
			want := byte(0x34)

			// $0100 = 0x12
			c.Bus.WriteByte(c.Reg.PC, a8)

			// $FF12 = 0x34
			c.Bus.WriteByte(addr, want)

			// A = 0x34
			op.Handler(c, byte(op.R1), byte(op.R2))

			assert.Equal(t, op.R1, tt.args.r1)
			assert.Equal(t, want, c.Reg.R[A])
		})
	}
}

func TestOpeCode_ldra16(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
		r1     byte
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "LD A,(a16)",
			args: args{0xFA, A},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := opCodes[tt.args.opcode]
			upper := byte(0x12)
			lower := byte(0x34)
			want := byte(0x56)

			// $0100 = 0x12
			c.Bus.WriteByte(c.Reg.PC, lower)
			c.Bus.WriteByte(c.Reg.PC+1, upper)
			c.Bus.WriteByte(util.Byte2Addr(upper, lower), want)

			op.Handler(c, byte(op.R1), byte(op.R2))

			assert.Equal(t, op.R1, tt.args.r1)
			assert.Equal(t, want, c.Reg.R[A])
		})
	}
}

func TestOpeCode_ldr16d16(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
		r1     byte
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "LD BC, d16",
			args: args{0x01, BC},
		},
		{
			name: "LD DE, d16",
			args: args{0x11, DE},
		},
		{
			name: "LD HL, d16",
			args: args{0x21, HL},
		},
		{
			name: "LD SP, d16",
			args: args{0x31, SP},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := opCodes[tt.args.opcode]
			want := byte(0x20)
			c.Bus.WriteByte(c.Reg.PC, want)
			c.Bus.WriteByte(c.Reg.PC+1, want+1)
			op.Handler(c, byte(op.R1), byte(op.R2))

			assert.Equal(t, op.R1, tt.args.r1)
			assert.Equal(t, util.Byte2Addr(want+1, want), c.Reg.R16(int(tt.args.r1)))
		})
	}
}

func TestOpeCode_ldmr(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
		r1     byte
		r2     byte
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "LD (C), A",
			args: args{0xE2, C, A},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := opCodes[tt.args.opcode]

			addr := util.Byte2Addr(0xFF, c.Reg.R[tt.args.r1])
			want := c.Reg.R[tt.args.r2]
			op.Handler(c, byte(op.R1), byte(op.R2))

			assert.Equal(t, op.R1, tt.args.r1)
			assert.Equal(t, op.R2, tt.args.r2)
			assert.Equal(t, want, c.Bus.ReadByte(addr))
		})
	}
}

func TestOpeCode_ldm16r(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
		r1     byte
		r2     byte
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "LD (BC), A",
			args: args{0x02, BC, A},
		},
		{
			name: "LD (DE), A",
			args: args{0x12, DE, A},
		},
		{
			name: "LD (HL+), A",
			args: args{0x22, HLI, A},
		},
		{
			name: "LD (HL-), A",
			args: args{0x32, HLD, A},
		},
		{
			name: "LD (HL), B",
			args: args{0x70, HL, B},
		},
		{
			name: "LD (HL), C",
			args: args{0x71, HL, C},
		},
		{
			name: "LD (HL), D",
			args: args{0x72, HL, D},
		},
		{
			name: "LD (HL), E",
			args: args{0x73, HL, E},
		},
		{
			name: "LD (HL), H",
			args: args{0x74, HL, H},
		},
		{
			name: "LD (HL), L",
			args: args{0x75, HL, L},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := opCodes[tt.args.opcode]
			want := c.Reg.R[tt.args.r2]
			op.Handler(c, byte(op.R1), byte(op.R2))

			assert.Equal(t, op.R1, tt.args.r1)
			assert.Equal(t, op.R2, tt.args.r2)
			if op.R1 == HLI {
				c.Reg.setHL(c.Reg.HL() - 1)
			} else if op.R1 == HLD {
				c.Reg.setHL(c.Reg.HL() + 1)
			}
			assert.Equal(t, want, c.Bus.ReadByte(c.Reg.R16(int(tt.args.r1))))
			assert.Equal(t, want, c.Reg.R[tt.args.r2])
		})
	}
}

func TestOpeCode_ldm16d8(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
		r1     byte
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "LD B, d8",
			args: args{0x06, B},
		},
		{
			name: "LD C, d8",
			args: args{0x0E, C},
		},
		{
			name: "LD D, d8",
			args: args{0x16, D},
		},
		{
			name: "LD E, d8",
			args: args{0x1E, E},
		},
		{
			name: "LD H, d8",
			args: args{0x26, H},
		},
		{
			name: "LD L, d8",
			args: args{0x2E, L},
		},
		{
			name: "LD A, d8",
			args: args{0x3E, A},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := opCodes[tt.args.opcode]
			want := byte(0x20)
			c.Bus.WriteByte(c.Reg.PC, want)
			op.Handler(c, byte(op.R1), byte(op.R2))

			assert.Equal(t, op.R1, tt.args.r1)
			assert.Equal(t, want, c.Reg.R[op.R1])
		})
	}
}

// -----jp-----
func TestOpeCode_jpa16(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
		addr   types.Addr
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "JP a16",
			args: args{0xC3, 0x1234},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := opCodes[tt.args.opcode]
			want := tt.args.addr
			c.Bus.WriteByte(c.Reg.PC, 0x34)
			c.Bus.WriteByte(c.Reg.PC+1, 0x12)
			op.Handler(c, byte(op.R1), byte(op.R2))

			assert.Equal(t, want, c.Reg.PC)
		})
	}
}

func TestOpeCode_jpfa16(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
		flag   int
		value  int
		addr   types.Addr
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "JP Z, a16 when zero",
			args: args{0xCA, flagZ, 0, 0x0100},
		},
		{
			name: "JP Z, a16 when non zero",
			args: args{0xCA, flagZ, 1, 0x1234},
		},
		{
			name: "JP C, a16 when zero",
			args: args{0xDA, flagC, 0, 0x0100},
		},
		{
			name: "JP C, a16 when non zero",
			args: args{0xDA, flagC, 1, 0x1234},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := opCodes[tt.args.opcode]
			want := tt.args.addr
			c.Bus.WriteByte(c.Reg.PC, 0x34)
			c.Bus.WriteByte(c.Reg.PC+1, 0x12)

			if tt.args.value == 1 {
				c.Reg.setFlag(byte(tt.args.flag))
			} else {
				c.Reg.clearFlag(byte(tt.args.flag))
			}
			op.Handler(c, byte(op.R1), byte(op.R2))

			assert.Equal(t, want, c.Reg.PC)
		})
	}
}

func TestOpeCode_jpnfa16(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
		flag   int
		value  int
		addr   types.Addr
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "JP NZ, a16 when zero",
			args: args{0xC2, flagZ, 0, 0x1234},
		},
		{
			name: "JP NZ, a16 when non zero",
			args: args{0xC2, flagZ, 1, 0x0100},
		},
		{
			name: "JP NC, a16 when zero",
			args: args{0xD2, flagC, 0, 0x1234},
		},
		{
			name: "JP NC, a16 when non zero",
			args: args{0xD2, flagC, 1, 0x0100},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := opCodes[tt.args.opcode]
			want := tt.args.addr
			c.Bus.WriteByte(c.Reg.PC, 0x34)
			c.Bus.WriteByte(c.Reg.PC+1, 0x12)

			if tt.args.value == 1 {
				c.Reg.setFlag(byte(tt.args.flag))
			} else {
				c.Reg.clearFlag(byte(tt.args.flag))
			}
			op.Handler(c, byte(op.R1), byte(op.R2))

			assert.Equal(t, want, c.Reg.PC)
		})
	}
}

func TestOpeCode_jpm16(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
		r16    int
		addr   types.Addr
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "JP (HL)",
			args: args{0xE9, HL, 0x34},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := opCodes[tt.args.opcode]
			want := tt.args.addr
			c.Bus.WriteByte(c.Reg.R16(tt.args.r16), 0x34)

			op.Handler(c, byte(op.R1), byte(op.R2))

			assert.Equal(t, want, c.Reg.PC)
		})
	}
}

// -----jr-----

func TestOpeCode_jrnfr8(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
		flag   int
		value  int
		addr   types.Addr
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "JR NZ, r8 when zero",
			args: args{0x20, flagZ, 0, 0x0111},
		},
		{
			name: "JR NZ, r8 when not zero",
			args: args{0x20, flagZ, 1, c.Reg.PC + 1},
		},
		{
			name: "JR NC, r8 when zero",
			args: args{0x30, flagC, 0, 0x0111},
		},
		{
			name: "JR NC, r8 when not zero",
			args: args{0x30, flagC, 1, c.Reg.PC + 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := opCodes[tt.args.opcode]
			want := tt.args.addr
			c.Bus.WriteByte(c.Reg.PC, 0x10)

			if tt.args.value == 1 {
				c.Reg.setFlag(byte(tt.args.flag))
			} else {
				c.Reg.clearFlag(byte(tt.args.flag))
			}
			op.Handler(c, byte(op.R1), byte(op.R2))

			assert.Equal(t, want, c.Reg.PC)
		})
	}
}

// ----push----
func TestOpeCode_push(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
		flag   int
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "PUSH BC",
			args: args{0xC5, BC},
		},
		{
			name: "PUSH DE",
			args: args{0xD5, DE},
		},
		{
			name: "PUSH HL",
			args: args{0xE5, HL},
		},
		{
			name: "PUSH AF",
			args: args{0xF5, AF},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := opCodes[tt.args.opcode]

			upper := util.ExtractUpper(c.Reg.R16(int(op.R1)))
			lower := util.ExtractLower(c.Reg.R16(int(op.R1)))
			before_sp := c.Reg.SP

			op.Handler(c, byte(op.R1), byte(op.R2))

			assert.Equal(t, before_sp-2, c.Reg.SP)
			assert.Equal(t, lower, c.Bus.ReadByte((c.Reg.SP)))
			assert.Equal(t, upper, c.Bus.ReadByte((c.Reg.SP + 1)))
		})
	}
}

func TestOpeCode_pop(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
		flag   int
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "POP BC",
			args: args{0xC1, BC},
		},
		{
			name: "POP DE",
			args: args{0xD1, DE},
		},
		{
			name: "POP HL",
			args: args{0xE1, HL},
		},
		{
			name: "POP AF",
			args: args{0xF1, AF},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := opCodes[tt.args.opcode]

			c.push(0x12) // upper
			c.push(0x34) // lower

			op.Handler(c, byte(op.R1), byte(op.R2))

			if op.R1 != AF {
				assert.Equal(t, types.Addr(0x1234), c.Reg.R16(int(op.R1)))
			} else {
				assert.Equal(t, types.Addr(0x1230), c.Reg.R16(int(op.R1)))
			}
		})
	}
}
