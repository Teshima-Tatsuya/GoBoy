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

func TestOpeCode_ldr16r16(t *testing.T) {
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
			name: "LD SP,HL",
			args: args{0xF9, SP, HL},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := opCodes[tt.args.opcode]
			want := c.Reg.R16(int(tt.args.r2))
			op.Handler(c, byte(op.R1), byte(op.R2))

			assert.Equal(t, op.R1, tt.args.r1)
			assert.Equal(t, op.R2, tt.args.r2)
			assert.Equal(t, want, c.Reg.R16(int(tt.args.r1)))
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

func TestOpeCode_ldm16d(t *testing.T) {
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
			name: "LD (HL),d8",
			args: args{0x36, HL},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := opCodes[tt.args.opcode]

			want := byte(0x12)
			c.Bus.WriteByte(c.Reg.PC, 0x12)

			op.Handler(c, byte(op.R1), byte(op.R2))

			assert.Equal(t, op.R1, tt.args.r1)
			assert.Equal(t, want, c.Bus.ReadByte(c.Reg.R16(int(tt.args.r1))))
		})
	}
}

func TestOpeCode_ldar(t *testing.T) {
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
			name: "LDH (a8),A",
			args: args{0xE0, A},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := opCodes[tt.args.opcode]

			addr := byte(0x12)
			want := byte(0x01)
			c.Bus.WriteByte(c.Reg.PC, addr)

			op.Handler(c, byte(op.R1), byte(op.R2))

			assert.Equal(t, op.R2, tt.args.r1)
			assert.Equal(t, want, c.Bus.ReadByte(util.Byte2Addr(0xFF, addr)))
		})
	}
}

func TestOpeCode_lda16r(t *testing.T) {
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
			name: "LD (a16),A",
			args: args{0xEA, A},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := opCodes[tt.args.opcode]

			addr := types.Addr(0x1234)
			want := byte(0x01)
			c.Reg.R[A] = want
			c.Bus.WriteByte(c.Reg.PC, util.ExtractLower(addr))
			c.Bus.WriteByte(c.Reg.PC+1, util.ExtractUpper(addr))

			op.Handler(c, byte(op.R1), byte(op.R2))

			assert.Equal(t, op.R2, tt.args.r1)
			assert.Equal(t, want, c.Bus.ReadByte(addr))
		})
	}
}

func TestOpeCode_lda16r16(t *testing.T) {
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
			name: "LD (a16),SP",
			args: args{0x08, SP},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := opCodes[tt.args.opcode]

			addr := types.Addr(0x1234)
			want := types.Addr(0x4567)
			c.Reg.setR16(SP, want)
			c.Bus.WriteByte(c.Reg.PC, util.ExtractLower(addr))
			c.Bus.WriteByte(c.Reg.PC+1, util.ExtractUpper(addr))

			op.Handler(c, byte(op.R1), byte(op.R2))

			assert.Equal(t, op.R2, tt.args.r1)
			assert.Equal(t, util.ExtractLower(want), c.Bus.ReadByte(addr))
			assert.Equal(t, util.ExtractUpper(want), c.Bus.ReadByte(addr+1))
		})
	}
}

// -arithmetic-

func TestOpeCode_incr(t *testing.T) {
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
			name: "INC B",
			args: args{0x04, B},
		},
		{
			name: "INC C",
			args: args{0x0C, C},
		},
		{
			name: "INC D",
			args: args{0x14, D},
		},
		{
			name: "INC E",
			args: args{0x1C, E},
		},
		{
			name: "INC H",
			args: args{0x24, H},
		},
		{
			name: "INC L",
			args: args{0x2C, L},
		},
		{
			name: "INC A",
			args: args{0x3C, A},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := opCodes[tt.args.opcode]

			var want byte
			// Not Max
			c.Reg.R[tt.args.r1] = byte(0xF0)
			want = byte(0xF1)

			op.Handler(c, byte(op.R1), byte(op.R2))

			assert.Equal(t, tt.args.r1, op.R1)
			assert.Equal(t, want, c.Reg.R[op.R1])
			assert.Equal(t, false, c.Reg.isSet(flagZ))
			assert.Equal(t, false, c.Reg.isSet(flagN))
			assert.Equal(t, false, c.Reg.isSet(flagH))

			// Max
			c.Reg.R[tt.args.r1] = byte(0xFF)
			want = byte(0x00)

			op.Handler(c, byte(op.R1), byte(op.R2))

			assert.Equal(t, tt.args.r1, op.R1)
			assert.Equal(t, want, c.Reg.R[op.R1])
			assert.Equal(t, true, c.Reg.isSet(flagZ))
			assert.Equal(t, false, c.Reg.isSet(flagN))
			assert.Equal(t, true, c.Reg.isSet(flagH))
		})
	}
}

func TestOpeCode_incr16(t *testing.T) {
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
			name: "INC BC",
			args: args{0x03, BC},
		},
		{
			name: "INC DE",
			args: args{0x13, DE},
		},
		{
			name: "INC HL",
			args: args{0x23, HL},
		},
		{
			name: "INC SP",
			args: args{0x33, SP},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := opCodes[tt.args.opcode]

			c.Reg.setR16(int(op.R1), 0x1234)
			want := types.Addr(0x1235)

			op.Handler(c, byte(op.R1), byte(op.R2))

			assert.Equal(t, tt.args.r1, op.R1)
			assert.Equal(t, want, c.Reg.R16(int(op.R1)))
		})
	}
}

func TestOpeCode_incm16(t *testing.T) {
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
			name: "INC (HL)",
			args: args{0x34, HL},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := opCodes[tt.args.opcode]

			var want, act byte

			// zero
			c.Bus.WriteByte(c.Reg.R16(int(tt.args.r1)), byte(0xFF))
			want = byte(0x00)

			op.Handler(c, byte(op.R1), byte(op.R2))

			act = c.Bus.ReadByte(c.Reg.R16(int(op.R1)))

			assert.Equal(t, tt.args.r1, op.R1)
			assert.Equal(t, want, act)
			assert.Equal(t, true, c.Reg.isSet(flagZ))
			assert.Equal(t, false, c.Reg.isSet(flagN))
			assert.Equal(t, true, c.Reg.isSet(flagH))

			// not zero
			c.Bus.WriteByte(c.Reg.R16(int(tt.args.r1)), byte(0x00))
			want = byte(0x01)

			op.Handler(c, byte(op.R1), byte(op.R2))

			act = c.Bus.ReadByte(c.Reg.R16(int(op.R1)))

			assert.Equal(t, tt.args.r1, op.R1)
			assert.Equal(t, want, act)
			assert.Equal(t, false, c.Reg.isSet(flagZ))
			assert.Equal(t, false, c.Reg.isSet(flagN))
			assert.Equal(t, false, c.Reg.isSet(flagH))
		})
	}
}

func TestOpeCode_decr(t *testing.T) {
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
			name: "DEC B",
			args: args{0x05, B},
		},
		{
			name: "DEC C",
			args: args{0x0D, C},
		},
		{
			name: "DEC D",
			args: args{0x15, D},
		},
		{
			name: "DEC E",
			args: args{0x1D, E},
		},
		{
			name: "DEC H",
			args: args{0x25, H},
		},
		{
			name: "DEC L",
			args: args{0x2D, L},
		},
		{
			name: "DEC A",
			args: args{0x3D, A},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := opCodes[tt.args.opcode]

			var want byte
			// Not Min
			c.Reg.R[tt.args.r1] = byte(0x01)
			want = byte(0x00)

			op.Handler(c, byte(op.R1), byte(op.R2))

			assert.Equal(t, tt.args.r1, op.R1)
			assert.Equal(t, want, c.Reg.R[op.R1])
			assert.Equal(t, true, c.Reg.isSet(flagZ))
			assert.Equal(t, true, c.Reg.isSet(flagN))
			assert.Equal(t, false, c.Reg.isSet(flagH))

			// Min
			c.Reg.R[tt.args.r1] = byte(0x00)
			want = byte(0xFF)

			op.Handler(c, byte(op.R1), byte(op.R2))

			assert.Equal(t, tt.args.r1, op.R1)
			assert.Equal(t, want, c.Reg.R[op.R1])
			assert.Equal(t, false, c.Reg.isSet(flagZ))
			assert.Equal(t, true, c.Reg.isSet(flagN))
			assert.Equal(t, true, c.Reg.isSet(flagH))
		})
	}
}

func TestOpeCode_decr16(t *testing.T) {
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
			name: "DEC BC",
			args: args{0x0B, BC},
		},
		{
			name: "DEC DE",
			args: args{0x1B, DE},
		},
		{
			name: "DEC HL",
			args: args{0x2B, HL},
		},
		{
			name: "DEC SP",
			args: args{0x3B, SP},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := opCodes[tt.args.opcode]

			c.Reg.setR16(int(op.R1), 0x1234)
			want := types.Addr(0x1233)

			op.Handler(c, byte(op.R1), byte(op.R2))

			assert.Equal(t, tt.args.r1, op.R1)
			assert.Equal(t, want, c.Reg.R16(int(op.R1)))
		})
	}
}

func TestOpeCode_decm16(t *testing.T) {
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
			name: "DEC (HL)",
			args: args{0x35, HL},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := opCodes[tt.args.opcode]

			var want, act byte

			// zero
			c.Bus.WriteByte(c.Reg.R16(int(tt.args.r1)), byte(0x01))
			want = byte(0x00)

			op.Handler(c, byte(op.R1), byte(op.R2))

			act = c.Bus.ReadByte(c.Reg.R16(int(op.R1)))

			assert.Equal(t, tt.args.r1, op.R1)
			assert.Equal(t, want, act)
			assert.Equal(t, true, c.Reg.isSet(flagZ))
			assert.Equal(t, true, c.Reg.isSet(flagN))
			assert.Equal(t, false, c.Reg.isSet(flagH))

			// not zero
			c.Bus.WriteByte(c.Reg.R16(int(tt.args.r1)), byte(0x00))
			want = byte(0xFF)

			op.Handler(c, byte(op.R1), byte(op.R2))

			act = c.Bus.ReadByte(c.Reg.R16(int(op.R1)))

			assert.Equal(t, tt.args.r1, op.R1)
			assert.Equal(t, want, act)
			assert.Equal(t, false, c.Reg.isSet(flagZ))
			assert.Equal(t, true, c.Reg.isSet(flagN))
			assert.Equal(t, true, c.Reg.isSet(flagH))
		})
	}
}

func TestOpeCode_addr(t *testing.T) {
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
			name: "ADD A, B",
			args: args{0x80, A, B},
		},
		{
			name: "ADD A, C",
			args: args{0x81, A, C},
		},
		{
			name: "ADD A, D",
			args: args{0x82, A, D},
		},
		{
			name: "ADD A, E",
			args: args{0x83, A, E},
		},
		{
			name: "ADD A, H",
			args: args{0x84, A, H},
		},
		{
			name: "ADD A, L",
			args: args{0x85, A, L},
		},
		{
			name: "ADD A, A",
			args: args{0x87, A, A},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := opCodes[tt.args.opcode]

			assert.Equal(t, tt.args.r1, op.R1)
			assert.Equal(t, tt.args.r2, op.R2)

			t.Run("when no carry", func(t *testing.T) {
				if tt.args.r2 == A {
					t.Skip()
				}
				c.Reg.R[tt.args.r1] = 0xE1
				c.Reg.R[tt.args.r2] = 0x0E
				op.Handler(c, byte(op.R1), byte(op.R2))
				assert.Equal(t, byte(0xEF), c.Reg.R[A])
				assert.Equal(t, false, c.Reg.isSet(flagZ))
				assert.Equal(t, false, c.Reg.isSet(flagN))
				assert.Equal(t, false, c.Reg.isSet(flagH))
				assert.Equal(t, false, c.Reg.isSet(flagC))
			})
			t.Run("when Harf carry", func(t *testing.T) {
				if tt.args.r2 == A {
					t.Skip()
				}
				c.Reg.R[tt.args.r1] = 0xE1
				c.Reg.R[tt.args.r2] = 0x0F
				op.Handler(c, byte(op.R1), byte(op.R2))
				assert.Equal(t, byte(0xF0), c.Reg.R[A])
				assert.Equal(t, false, c.Reg.isSet(flagZ))
				assert.Equal(t, false, c.Reg.isSet(flagN))
				assert.Equal(t, true, c.Reg.isSet(flagH))
				assert.Equal(t, false, c.Reg.isSet(flagC))
			})
			t.Run("when carry and zero", func(t *testing.T) {
				if tt.args.r2 == A {
					t.Skip()
				}
				c.Reg.R[tt.args.r1] = 0xE1
				c.Reg.R[tt.args.r2] = 0x1F
				op.Handler(c, byte(op.R1), byte(op.R2))
				assert.Equal(t, byte(0x00), c.Reg.R[A])
				assert.Equal(t, true, c.Reg.isSet(flagZ))
				assert.Equal(t, false, c.Reg.isSet(flagN))
				assert.Equal(t, true, c.Reg.isSet(flagH))
				assert.Equal(t, true, c.Reg.isSet(flagC))
			})
		})
	}
}

func TestOpeCode_addr16r16(t *testing.T) {
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
			name: "ADD HL,BC",
			args: args{0x09, HL, BC},
		},
		{
			name: "ADD HL,DE",
			args: args{0x19, HL, DE},
		},
		{
			name: "ADD HL,HL",
			args: args{0x29, HL, HL},
		},
		{
			name: "ADD HL,SP",
			args: args{0x39, HL, SP},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := opCodes[tt.args.opcode]

			assert.Equal(t, tt.args.r1, op.R1)
			assert.Equal(t, tt.args.r2, op.R2)

			t.Run("when no carry", func(t *testing.T) {
				if tt.args.r2 == HL {
					t.Skip()
				}
				c.Reg.setR16(int(tt.args.r1), 0x00E1)
				c.Reg.setR16(int(tt.args.r2), 0x000E)
				op.Handler(c, byte(op.R1), byte(op.R2))
				assert.Equal(t, types.Addr(0x00EF), c.Reg.R16(int(op.R1)))
				assert.Equal(t, false, c.Reg.isSet(flagN))
				assert.Equal(t, false, c.Reg.isSet(flagH))
				assert.Equal(t, false, c.Reg.isSet(flagC))
			})
			t.Run("when Harf carry", func(t *testing.T) {
				if tt.args.r2 == HL {
					t.Skip()
				}
				c.Reg.setR16(int(tt.args.r1), 0x0FF1)
				c.Reg.setR16(int(tt.args.r2), 0x000F)
				op.Handler(c, byte(op.R1), byte(op.R2))
				assert.Equal(t, types.Addr(0x1000), c.Reg.R16(int(op.R1)))
				assert.Equal(t, false, c.Reg.isSet(flagN))
				assert.Equal(t, true, c.Reg.isSet(flagH))
				assert.Equal(t, false, c.Reg.isSet(flagC))
			})
			t.Run("when carry and zero", func(t *testing.T) {
				if tt.args.r2 == HL {
					t.Skip()
				}
				c.Reg.setR16(int(tt.args.r1), 0xFFF1)
				c.Reg.setR16(int(tt.args.r2), 0x000F)
				op.Handler(c, byte(op.R1), byte(op.R2))
				assert.Equal(t, types.Addr(0x00), c.Reg.R16(int(op.R1)))
				assert.Equal(t, false, c.Reg.isSet(flagN))
				assert.Equal(t, true, c.Reg.isSet(flagH))
				assert.Equal(t, true, c.Reg.isSet(flagC))
			})
		})
	}
}

func TestOpeCode_subr(t *testing.T) {
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
			name: "SUB B",
			args: args{0x90, B},
		},
		{
			name: "SUB C",
			args: args{0x91, C},
		},
		{
			name: "SUB D",
			args: args{0x92, D},
		},
		{
			name: "SUB E",
			args: args{0x93, E},
		},
		{
			name: "SUB H",
			args: args{0x94, H},
		},
		{
			name: "SUB L",
			args: args{0x95, L},
		},
		{
			name: "SUB A",
			args: args{0x97, A},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := opCodes[tt.args.opcode]

			assert.Equal(t, tt.args.r1, op.R1)

			t.Run("when no carry", func(t *testing.T) {
				if tt.args.r1 == A {
					t.Skip()
				}
				c.Reg.R[A] = 0xE1
				c.Reg.R[tt.args.r1] = 0x01
				op.Handler(c, byte(op.R1), byte(op.R2))
				assert.Equal(t, byte(0xE0), c.Reg.R[A])
				assert.Equal(t, false, c.Reg.isSet(flagZ))
				assert.Equal(t, true, c.Reg.isSet(flagN))
				assert.Equal(t, false, c.Reg.isSet(flagH))
				assert.Equal(t, false, c.Reg.isSet(flagC))
			})
			t.Run("when Harf carry", func(t *testing.T) {
				if tt.args.r1 == A {
					t.Skip()
				}
				c.Reg.R[A] = 0xE1
				c.Reg.R[tt.args.r1] = 0x02
				op.Handler(c, byte(op.R1), byte(op.R2))
				assert.Equal(t, byte(0xDF), c.Reg.R[A])
				assert.Equal(t, false, c.Reg.isSet(flagZ))
				assert.Equal(t, true, c.Reg.isSet(flagN))
				assert.Equal(t, true, c.Reg.isSet(flagH))
				assert.Equal(t, false, c.Reg.isSet(flagC))
			})
			t.Run("when zero", func(t *testing.T) {
				if tt.args.r1 == A {
					t.Skip()
				}
				c.Reg.R[A] = 0xE1
				c.Reg.R[tt.args.r1] = 0xE1
				op.Handler(c, byte(op.R1), byte(op.R2))
				assert.Equal(t, byte(0x00), c.Reg.R[A])
				assert.Equal(t, true, c.Reg.isSet(flagZ))
				assert.Equal(t, true, c.Reg.isSet(flagN))
				assert.Equal(t, false, c.Reg.isSet(flagH))
				assert.Equal(t, false, c.Reg.isSet(flagC))
			})
			t.Run("when carry", func(t *testing.T) {
				if tt.args.r1 == A {
					t.Skip()
				}
				c.Reg.R[A] = 0xE1
				c.Reg.R[tt.args.r1] = 0xE2
				op.Handler(c, byte(op.R1), byte(op.R2))
				assert.Equal(t, byte(0xFF), c.Reg.R[A])
				assert.Equal(t, false, c.Reg.isSet(flagZ))
				assert.Equal(t, true, c.Reg.isSet(flagN))
				assert.Equal(t, true, c.Reg.isSet(flagH))
				assert.Equal(t, true, c.Reg.isSet(flagC))
			})
		})
	}
}

// test andr andHL andd8
func TestOpeCode_and(t *testing.T) {
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
			name: "AND B",
			args: args{0xA0, B},
		},
		{
			name: "AND C",
			args: args{0xA1, C},
		},
		{
			name: "AND D",
			args: args{0xA2, D},
		},
		{
			name: "AND E",
			args: args{0xA3, E},
		},
		{
			name: "AND H",
			args: args{0xA4, H},
		},
		{
			name: "AND L",
			args: args{0xA5, L},
		},
		{
			name: "AND (HL)",
			args: args{0xA6, HL},
		},
		{
			name: "AND A",
			args: args{0xA7, A},
		},
		{
			name: "AND d8",
			args: args{0xE6, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := opCodes[tt.args.opcode]
			assert.Equal(t, tt.args.r1, op.R1)

			if tt.args.r1 == A {
				t.Skip()
			}
			t.Run("when oposite", func(t *testing.T) {
				c.Reg.R[A] = 0b11110000
				val := byte(0b00001111)
				c.Bus.WriteByte(c.Reg.PC, val)
				if tt.args.r1 == HL {
					c.Bus.WriteByte(c.Reg.R16(int(HL)), val)
				}
				c.Reg.R[tt.args.r1] = val
				op.Handler(c, byte(op.R1), byte(op.R2))
				assert.Equal(t, byte(0b00000000), c.Reg.R[A])
			})
			t.Run("when equal", func(t *testing.T) {
				c.Reg.R[A] = 0b11110000
				val := byte(0b11110000)
				c.Bus.WriteByte(c.Reg.PC, val)
				if tt.args.r1 == HL {
					c.Bus.WriteByte(c.Reg.R16(int(HL)), val)
				}
				c.Reg.R[tt.args.r1] = val
				op.Handler(c, byte(op.R1), byte(op.R2))
				assert.Equal(t, byte(0b11110000), c.Reg.R[A])
			})
			t.Run("when other", func(t *testing.T) {
				c.Reg.R[A] = 0b11110000
				val := byte(0b10100000)
				c.Bus.WriteByte(c.Reg.PC, val)
				if tt.args.r1 == HL {
					c.Bus.WriteByte(c.Reg.R16(int(HL)), val)
				}
				c.Reg.R[tt.args.r1] = val
				op.Handler(c, byte(op.R1), byte(op.R2))
				assert.Equal(t, byte(0b10100000), c.Reg.R[A])
			})
		})
	}
}

// test orr orHL ord8
func TestOpeCode_or(t *testing.T) {
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
			name: "OR B",
			args: args{0xB0, B},
		},
		{
			name: "OR C",
			args: args{0xB1, C},
		},
		{
			name: "OR D",
			args: args{0xB2, D},
		},
		{
			name: "OR E",
			args: args{0xB3, E},
		},
		{
			name: "OR H",
			args: args{0xB4, H},
		},
		{
			name: "OR L",
			args: args{0xB5, L},
		},
		{
			name: "OR (HL)",
			args: args{0xB6, HL},
		},
		{
			name: "OR A",
			args: args{0xB7, A},
		},
		{
			name: "OR d8",
			args: args{0xF6, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := opCodes[tt.args.opcode]
			assert.Equal(t, tt.args.r1, op.R1)

			if tt.args.r1 == A {
				t.Skip()
			}
			t.Run("when oposite", func(t *testing.T) {
				c.Reg.R[A] = 0b11110000
				val := byte(0b00001111)
				c.Bus.WriteByte(c.Reg.PC, val)
				if tt.args.r1 == HL {
					c.Bus.WriteByte(c.Reg.R16(int(HL)), val)
				}
				c.Reg.R[tt.args.r1] = val
				op.Handler(c, byte(op.R1), byte(op.R2))
				assert.Equal(t, byte(0b11111111), c.Reg.R[A])
			})
			t.Run("when equal", func(t *testing.T) {
				c.Reg.R[A] = 0b11110000
				val := byte(0b11110000)
				c.Bus.WriteByte(c.Reg.PC, val)
				if tt.args.r1 == HL {
					c.Bus.WriteByte(c.Reg.R16(int(HL)), val)
				}
				c.Reg.R[tt.args.r1] = val
				op.Handler(c, byte(op.R1), byte(op.R2))
				assert.Equal(t, byte(0b11110000), c.Reg.R[A])
			})
			t.Run("when other", func(t *testing.T) {
				c.Reg.R[A] = 0b11110000
				val := byte(0b00000101)
				if tt.args.r1 == HL {
					c.Bus.WriteByte(c.Reg.R16(int(HL)), val)
				}
				c.Reg.R[tt.args.r1] = val
				op.Handler(c, byte(op.R1), byte(op.R2))
				assert.Equal(t, byte(0b11110101), c.Reg.R[A])
			})
		})
	}
}

// test xorr xorHL xord8
func TestOpeCode_xor(t *testing.T) {
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
			name: "XOR B",
			args: args{0xA8, B},
		},
		{
			name: "XOR C",
			args: args{0xA9, C},
		},
		{
			name: "XOR D",
			args: args{0xAA, D},
		},
		{
			name: "XOR E",
			args: args{0xAB, E},
		},
		{
			name: "XOR H",
			args: args{0xAC, H},
		},
		{
			name: "XOR L",
			args: args{0xAD, L},
		},
		{
			name: "XOR (HL)",
			args: args{0xAE, HL},
		},
		{
			name: "XOR A",
			args: args{0xAF, A},
		},
		{
			name: "XOR d8",
			args: args{0xEE, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := opCodes[tt.args.opcode]
			assert.Equal(t, tt.args.r1, op.R1)

			if tt.args.r1 == A {
				t.Skip()
			}
			t.Run("when oposite", func(t *testing.T) {
				c.Reg.R[A] = 0b11110000
				val := byte(0b00001111)
				c.Bus.WriteByte(c.Reg.PC, val)
				if tt.args.r1 == HL {
					c.Bus.WriteByte(c.Reg.R16(int(HL)), val)
				}
				c.Reg.R[tt.args.r1] = val
				op.Handler(c, byte(op.R1), byte(op.R2))
				assert.Equal(t, byte(0b11111111), c.Reg.R[A])
			})
			t.Run("when equal", func(t *testing.T) {
				c.Reg.R[A] = 0b11110000
				val := byte(0b11110000)
				c.Bus.WriteByte(c.Reg.PC, val)
				if tt.args.r1 == HL {
					c.Bus.WriteByte(c.Reg.R16(int(HL)), val)
				}
				c.Reg.R[tt.args.r1] = val
				op.Handler(c, byte(op.R1), byte(op.R2))
				assert.Equal(t, byte(0b00000000), c.Reg.R[A])
			})
			t.Run("when other", func(t *testing.T) {
				c.Reg.R[A] = 0b11110000
				val := byte(0b01010101)
				if tt.args.r1 == HL {
					c.Bus.WriteByte(c.Reg.R16(int(HL)), val)
				}
				c.Reg.R[tt.args.r1] = val
				op.Handler(c, byte(op.R1), byte(op.R2))
				assert.Equal(t, byte(0b10100101), c.Reg.R[A])
			})
		})
	}
}

func TestOpeCode_cpr(t *testing.T) {
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
			name: "CP B",
			args: args{0xB8, B},
		},
		{
			name: "CP C",
			args: args{0xB9, C},
		},
		{
			name: "CP D",
			args: args{0xBA, D},
		},
		{
			name: "CP E",
			args: args{0xBB, E},
		},
		{
			name: "CP H",
			args: args{0xBC, H},
		},
		{
			name: "CP L",
			args: args{0xBD, L},
		},
		{
			name: "CP A",
			args: args{0xBF, A},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := opCodes[tt.args.opcode]
			// prepare
			c.Reg.R[A] = 0x12

			t.Run("when equal", func(t *testing.T) {
				c.Reg.R[tt.args.r1] = 0x12
				op.Handler(c, byte(op.R1), byte(op.R2))
				assert.Equal(t, true, c.Reg.isSet(flagZ))
				assert.Equal(t, true, c.Reg.isSet(flagN))
				assert.Equal(t, false, c.Reg.isSet(flagH))
				assert.Equal(t, false, c.Reg.isSet(flagC))
			})
			t.Run("when greater than A", func(t *testing.T) {
				if tt.args.r1 == A {
					t.Skip()
				}
				c.Reg.R[tt.args.r1] = 0x13
				op.Handler(c, byte(op.R1), byte(op.R2))
				assert.Equal(t, false, c.Reg.isSet(flagZ))
				assert.Equal(t, true, c.Reg.isSet(flagN))
				assert.Equal(t, true, c.Reg.isSet(flagH))
				assert.Equal(t, true, c.Reg.isSet(flagC))
			})
			t.Run("when less than A with borrow", func(t *testing.T) {
				if tt.args.r1 == A {
					t.Skip()
				}
				c.Reg.R[tt.args.r1] = 0x03
				op.Handler(c, byte(op.R1), byte(op.R2))
				assert.Equal(t, false, c.Reg.isSet(flagZ))
				assert.Equal(t, true, c.Reg.isSet(flagN))
				assert.Equal(t, true, c.Reg.isSet(flagH))
				assert.Equal(t, false, c.Reg.isSet(flagC))
			})
			t.Run("when less than A no borrow", func(t *testing.T) {
				if tt.args.r1 == A {
					t.Skip()
				}
				c.Reg.R[tt.args.r1] = 0x02
				op.Handler(c, byte(op.R1), byte(op.R2))
				assert.Equal(t, false, c.Reg.isSet(flagZ))
				assert.Equal(t, true, c.Reg.isSet(flagN))
				assert.Equal(t, false, c.Reg.isSet(flagH))
				assert.Equal(t, false, c.Reg.isSet(flagC))
			})
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

// PREFIX CB

func TestOpeCode_rlc(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
		r1     int
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "RLC B",
			args: args{0x00, B},
		},
		{
			name: "RLC C",
			args: args{0x01, C},
		},
		{
			name: "RLC D",
			args: args{0x02, D},
		},
		{
			name: "RLC E",
			args: args{0x03, E},
		},
		{
			name: "RLC H",
			args: args{0x04, H},
		},
		{
			name: "RLC L",
			args: args{0x05, L},
		},
		{
			name: "RLC HL",
			args: args{0x06, HL},
		},
		{
			name: "RLC A",
			args: args{0x07, A},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := cbOpCodes[tt.args.opcode]

			assert.Equal(t, uint8(tt.args.r1), op.R1)

			t.Run("when bit7 = 1", func(t *testing.T) {
				before := byte(0b10010000)
				c.Reg.R[tt.args.r1] = before
				c.Bus.WriteByte(c.Reg.R16(int(HL)), before)
				op.Handler(c, byte(op.R1), byte(op.R2))
				assert.Equal(t, false, c.Reg.isSet(flagZ))
				assert.Equal(t, false, c.Reg.isSet(flagN))
				assert.Equal(t, false, c.Reg.isSet(flagH))
				assert.Equal(t, true, c.Reg.isSet(flagC))

				want := byte(0b00100001)

				if tt.name != "RLC HL" {
					assert.Equal(t, want, c.Reg.R[op.R1])
				} else {
					assert.Equal(t, want, c.Bus.ReadByte(c.Reg.R16(int(HL))))
				}
			})
			t.Run("when bit7 = 0 and other = 0", func(t *testing.T) {
				before := byte(0b00000000)
				c.Reg.R[tt.args.r1] = before
				c.Bus.WriteByte(c.Reg.R16(int(HL)), before)
				op.Handler(c, byte(op.R1), byte(op.R2))
				assert.Equal(t, true, c.Reg.isSet(flagZ))
				assert.Equal(t, false, c.Reg.isSet(flagN))
				assert.Equal(t, false, c.Reg.isSet(flagH))
				assert.Equal(t, false, c.Reg.isSet(flagC))

				want := byte(0b00000000)

				if tt.name != "RLC HL" {
					assert.Equal(t, want, c.Reg.R[op.R1])
				} else {
					assert.Equal(t, want, c.Bus.ReadByte(c.Reg.R16(int(HL))))
				}
			})
			t.Run("when bit7 = 0 and other != 0", func(t *testing.T) {
				before := byte(0b00100000)
				c.Reg.R[tt.args.r1] = before
				c.Bus.WriteByte(c.Reg.R16(int(HL)), before)
				op.Handler(c, byte(op.R1), byte(op.R2))
				assert.Equal(t, false, c.Reg.isSet(flagZ))
				assert.Equal(t, false, c.Reg.isSet(flagN))
				assert.Equal(t, false, c.Reg.isSet(flagH))
				assert.Equal(t, false, c.Reg.isSet(flagC))

				want := byte(0b01000000)

				if tt.name != "RLC HL" {
					assert.Equal(t, want, c.Reg.R[op.R1])
				} else {
					assert.Equal(t, want, c.Bus.ReadByte(c.Reg.R16(int(HL))))
				}
			})
		})
	}
}

func TestOpeCode_rrc(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
		r1     int
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "RRC B",
			args: args{0x08, B},
		},
		{
			name: "RRC C",
			args: args{0x09, C},
		},
		{
			name: "RRC D",
			args: args{0x0A, D},
		},
		{
			name: "RRC E",
			args: args{0x0B, E},
		},
		{
			name: "RRC H",
			args: args{0x0C, H},
		},
		{
			name: "RRC L",
			args: args{0x0D, L},
		},
		{
			name: "RRC HL",
			args: args{0x0E, HL},
		},
		{
			name: "RRC A",
			args: args{0x0F, A},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := cbOpCodes[tt.args.opcode]

			assert.Equal(t, uint8(tt.args.r1), op.R1)

			t.Run("when bit0 = 1", func(t *testing.T) {
				before := byte(0b10010001)
				c.Reg.R[tt.args.r1] = before
				c.Bus.WriteByte(c.Reg.R16(int(HL)), before)
				op.Handler(c, byte(op.R1), byte(op.R2))
				assert.Equal(t, false, c.Reg.isSet(flagZ))
				assert.Equal(t, false, c.Reg.isSet(flagN))
				assert.Equal(t, false, c.Reg.isSet(flagH))
				assert.Equal(t, true, c.Reg.isSet(flagC))

				want := byte(0b11001000)

				if tt.name != "RRC HL" {
					assert.Equal(t, want, c.Reg.R[op.R1])
				} else {
					assert.Equal(t, want, c.Bus.ReadByte(c.Reg.R16(int(HL))))
				}
			})
			t.Run("when bit0 = 0 and other = 0", func(t *testing.T) {
				before := byte(0b00000000)
				c.Reg.R[tt.args.r1] = before
				c.Bus.WriteByte(c.Reg.R16(int(HL)), before)
				op.Handler(c, byte(op.R1), byte(op.R2))
				assert.Equal(t, true, c.Reg.isSet(flagZ))
				assert.Equal(t, false, c.Reg.isSet(flagN))
				assert.Equal(t, false, c.Reg.isSet(flagH))
				assert.Equal(t, false, c.Reg.isSet(flagC))

				want := byte(0b00000000)

				if tt.name != "RRC HL" {
					assert.Equal(t, want, c.Reg.R[op.R1])
				} else {
					assert.Equal(t, want, c.Bus.ReadByte(c.Reg.R16(int(HL))))
				}
			})
			t.Run("when bit0 = 0 and other != 0", func(t *testing.T) {
				before := byte(0b00100000)
				c.Reg.R[tt.args.r1] = before
				c.Bus.WriteByte(c.Reg.R16(int(HL)), before)
				op.Handler(c, byte(op.R1), byte(op.R2))
				assert.Equal(t, false, c.Reg.isSet(flagZ))
				assert.Equal(t, false, c.Reg.isSet(flagN))
				assert.Equal(t, false, c.Reg.isSet(flagH))
				assert.Equal(t, false, c.Reg.isSet(flagC))

				want := byte(0b00010000)

				if tt.name != "RRC HL" {
					assert.Equal(t, want, c.Reg.R[op.R1])
				} else {
					assert.Equal(t, want, c.Bus.ReadByte(c.Reg.R16(int(HL))))
				}
			})
		})
	}
}

func TestOpeCode_rl(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
		r1     int
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "RL B",
			args: args{0x10, B},
		},
		{
			name: "RL C",
			args: args{0x11, C},
		},
		{
			name: "RL D",
			args: args{0x12, D},
		},
		{
			name: "RL E",
			args: args{0x13, E},
		},
		{
			name: "RL H",
			args: args{0x14, H},
		},
		{
			name: "RL L",
			args: args{0x15, L},
		},
		{
			name: "RL HL",
			args: args{0x16, HL},
		},
		{
			name: "RL A",
			args: args{0x17, A},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := cbOpCodes[tt.args.opcode]

			assert.Equal(t, uint8(tt.args.r1), op.R1)

			t.Run("when bit7 = 1", func(t *testing.T) {
				before := byte(0b10010000)
				t.Run("and carry = 1", func(t *testing.T) {
					c.Reg.R[tt.args.r1] = before
					c.Bus.WriteByte(c.Reg.R16(int(HL)), before)
					c.Reg.setFlag(flagC)
					op.Handler(c, byte(op.R1), byte(op.R2))
					assert.Equal(t, false, c.Reg.isSet(flagZ))
					assert.Equal(t, false, c.Reg.isSet(flagN))
					assert.Equal(t, false, c.Reg.isSet(flagH))
					assert.Equal(t, true, c.Reg.isSet(flagC))

					want := byte(0b00100001)

					if tt.name != "RL HL" {
						assert.Equal(t, want, c.Reg.R[op.R1])
					} else {
						assert.Equal(t, want, c.Bus.ReadByte(c.Reg.R16(int(HL))))
					}
				})
				t.Run("and carry = 0", func(t *testing.T) {
					c.Reg.R[tt.args.r1] = before
					c.Bus.WriteByte(c.Reg.R16(int(HL)), before)
					c.Reg.clearFlag(flagC)
					op.Handler(c, byte(op.R1), byte(op.R2))
					assert.Equal(t, false, c.Reg.isSet(flagZ))
					assert.Equal(t, false, c.Reg.isSet(flagN))
					assert.Equal(t, false, c.Reg.isSet(flagH))
					assert.Equal(t, true, c.Reg.isSet(flagC))

					want := byte(0b00100000)

					if tt.name != "RL HL" {
						assert.Equal(t, want, c.Reg.R[op.R1])
					} else {
						assert.Equal(t, want, c.Bus.ReadByte(c.Reg.R16(int(HL))))
					}
				})
			})
			t.Run("when bit7 = 0 and other = 0", func(t *testing.T) {
				before := byte(0b00000000)
				t.Run("and carry = 1", func(t *testing.T) {
					c.Reg.R[tt.args.r1] = before
					c.Bus.WriteByte(c.Reg.R16(int(HL)), before)
					c.Reg.setFlag(flagC)
					op.Handler(c, byte(op.R1), byte(op.R2))
					assert.Equal(t, false, c.Reg.isSet(flagZ))
					assert.Equal(t, false, c.Reg.isSet(flagN))
					assert.Equal(t, false, c.Reg.isSet(flagH))
					assert.Equal(t, false, c.Reg.isSet(flagC))

					want := byte(0b00000001)

					if tt.name != "RL HL" {
						assert.Equal(t, want, c.Reg.R[op.R1])
					} else {
						assert.Equal(t, want, c.Bus.ReadByte(c.Reg.R16(int(HL))))
					}
				})
				t.Run("and carry = 0", func(t *testing.T) {
					c.Reg.R[tt.args.r1] = before
					c.Bus.WriteByte(c.Reg.R16(int(HL)), before)
					c.Reg.clearFlag(flagC)
					op.Handler(c, byte(op.R1), byte(op.R2))
					assert.Equal(t, true, c.Reg.isSet(flagZ))
					assert.Equal(t, false, c.Reg.isSet(flagN))
					assert.Equal(t, false, c.Reg.isSet(flagH))
					assert.Equal(t, false, c.Reg.isSet(flagC))

					want := byte(0b00000000)

					if tt.name != "RL HL" {
						assert.Equal(t, want, c.Reg.R[op.R1])
					} else {
						assert.Equal(t, want, c.Bus.ReadByte(c.Reg.R16(int(HL))))
					}
				})
			})
			t.Run("when bit7 = 0 and other != 0", func(t *testing.T) {
				before := byte(0b00100000)
				t.Run("and carry = 1", func(t *testing.T) {
					c.Reg.R[tt.args.r1] = before
					c.Bus.WriteByte(c.Reg.R16(int(HL)), before)
					c.Reg.setFlag(flagC)
					op.Handler(c, byte(op.R1), byte(op.R2))
					assert.Equal(t, false, c.Reg.isSet(flagZ))
					assert.Equal(t, false, c.Reg.isSet(flagN))
					assert.Equal(t, false, c.Reg.isSet(flagH))
					assert.Equal(t, false, c.Reg.isSet(flagC))

					want := byte(0b01000001)

					if tt.name != "RL HL" {
						assert.Equal(t, want, c.Reg.R[op.R1])
					} else {
						assert.Equal(t, want, c.Bus.ReadByte(c.Reg.R16(int(HL))))
					}
				})
				t.Run("and carry = 0", func(t *testing.T) {
					c.Reg.R[tt.args.r1] = before
					c.Bus.WriteByte(c.Reg.R16(int(HL)), before)
					c.Reg.clearFlag(flagC)
					op.Handler(c, byte(op.R1), byte(op.R2))
					assert.Equal(t, false, c.Reg.isSet(flagZ))
					assert.Equal(t, false, c.Reg.isSet(flagN))
					assert.Equal(t, false, c.Reg.isSet(flagH))
					assert.Equal(t, false, c.Reg.isSet(flagC))

					want := byte(0b01000000)

					if tt.name != "RL HL" {
						assert.Equal(t, want, c.Reg.R[op.R1])
					} else {
						assert.Equal(t, want, c.Bus.ReadByte(c.Reg.R16(int(HL))))
					}
				})
			})
		})
	}
}

func TestOpeCode_rr(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
		r1     int
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "RR B",
			args: args{0x18, B},
		},
		{
			name: "RR C",
			args: args{0x19, C},
		},
		{
			name: "RR D",
			args: args{0x1A, D},
		},
		{
			name: "RR E",
			args: args{0x1B, E},
		},
		{
			name: "RR H",
			args: args{0x1C, H},
		},
		{
			name: "RR L",
			args: args{0x1D, L},
		},
		{
			name: "RR HL",
			args: args{0x1E, HL},
		},
		{
			name: "RR A",
			args: args{0x1F, A},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := cbOpCodes[tt.args.opcode]

			assert.Equal(t, uint8(tt.args.r1), op.R1)

			t.Run("when bit0 = 1", func(t *testing.T) {
				before := byte(0b10010001)
				t.Run("and carry = 1", func(t *testing.T) {
					c.Reg.R[tt.args.r1] = before
					c.Bus.WriteByte(c.Reg.R16(int(HL)), before)
					c.Reg.setFlag(flagC)
					op.Handler(c, byte(op.R1), byte(op.R2))
					assert.Equal(t, false, c.Reg.isSet(flagZ))
					assert.Equal(t, false, c.Reg.isSet(flagN))
					assert.Equal(t, false, c.Reg.isSet(flagH))
					assert.Equal(t, true, c.Reg.isSet(flagC))

					want := byte(0b11001000)

					if tt.name != "RR HL" {
						assert.Equal(t, want, c.Reg.R[op.R1])
					} else {
						assert.Equal(t, want, c.Bus.ReadByte(c.Reg.R16(int(HL))))
					}
				})
				t.Run("and carry = 0", func(t *testing.T) {
					c.Reg.R[tt.args.r1] = before
					c.Bus.WriteByte(c.Reg.R16(int(HL)), before)
					c.Reg.clearFlag(flagC)
					op.Handler(c, byte(op.R1), byte(op.R2))
					assert.Equal(t, false, c.Reg.isSet(flagZ))
					assert.Equal(t, false, c.Reg.isSet(flagN))
					assert.Equal(t, false, c.Reg.isSet(flagH))
					assert.Equal(t, true, c.Reg.isSet(flagC))

					want := byte(0b01001000)

					if tt.name != "RR HL" {
						assert.Equal(t, want, c.Reg.R[op.R1])
					} else {
						assert.Equal(t, want, c.Bus.ReadByte(c.Reg.R16(int(HL))))
					}
				})
			})
			t.Run("when bit0 = 0 and other = 0", func(t *testing.T) {
				before := byte(0b00000000)
				t.Run("and carry = 1", func(t *testing.T) {
					c.Reg.R[tt.args.r1] = before
					c.Bus.WriteByte(c.Reg.R16(int(HL)), before)
					c.Reg.setFlag(flagC)
					op.Handler(c, byte(op.R1), byte(op.R2))
					assert.Equal(t, false, c.Reg.isSet(flagZ))
					assert.Equal(t, false, c.Reg.isSet(flagN))
					assert.Equal(t, false, c.Reg.isSet(flagH))
					assert.Equal(t, false, c.Reg.isSet(flagC))

					want := byte(0b10000000)

					if tt.name != "RR HL" {
						assert.Equal(t, want, c.Reg.R[op.R1])
					} else {
						assert.Equal(t, want, c.Bus.ReadByte(c.Reg.R16(int(HL))))
					}
				})
				t.Run("and carry = 0", func(t *testing.T) {
					c.Reg.R[tt.args.r1] = before
					c.Bus.WriteByte(c.Reg.R16(int(HL)), before)
					c.Reg.clearFlag(flagC)
					op.Handler(c, byte(op.R1), byte(op.R2))
					assert.Equal(t, true, c.Reg.isSet(flagZ))
					assert.Equal(t, false, c.Reg.isSet(flagN))
					assert.Equal(t, false, c.Reg.isSet(flagH))
					assert.Equal(t, false, c.Reg.isSet(flagC))

					want := byte(0b00000000)

					if tt.name != "RR HL" {
						assert.Equal(t, want, c.Reg.R[op.R1])
					} else {
						assert.Equal(t, want, c.Bus.ReadByte(c.Reg.R16(int(HL))))
					}
				})
			})
			t.Run("when bit0 = 0 and other != 0", func(t *testing.T) {
				before := byte(0b00100000)
				t.Run("and carry = 1", func(t *testing.T) {
					c.Reg.R[tt.args.r1] = before
					c.Bus.WriteByte(c.Reg.R16(int(HL)), before)
					c.Reg.setFlag(flagC)
					op.Handler(c, byte(op.R1), byte(op.R2))
					assert.Equal(t, false, c.Reg.isSet(flagZ))
					assert.Equal(t, false, c.Reg.isSet(flagN))
					assert.Equal(t, false, c.Reg.isSet(flagH))
					assert.Equal(t, false, c.Reg.isSet(flagC))

					want := byte(0b10010000)

					if tt.name != "RR HL" {
						assert.Equal(t, want, c.Reg.R[op.R1])
					} else {
						assert.Equal(t, want, c.Bus.ReadByte(c.Reg.R16(int(HL))))
					}
				})
				t.Run("and carry = 0", func(t *testing.T) {
					c.Reg.R[tt.args.r1] = before
					c.Bus.WriteByte(c.Reg.R16(int(HL)), before)
					c.Reg.clearFlag(flagC)
					op.Handler(c, byte(op.R1), byte(op.R2))
					assert.Equal(t, false, c.Reg.isSet(flagZ))
					assert.Equal(t, false, c.Reg.isSet(flagN))
					assert.Equal(t, false, c.Reg.isSet(flagH))
					assert.Equal(t, false, c.Reg.isSet(flagC))

					want := byte(0b00010000)

					if tt.name != "RR HL" {
						assert.Equal(t, want, c.Reg.R[op.R1])
					} else {
						assert.Equal(t, want, c.Bus.ReadByte(c.Reg.R16(int(HL))))
					}
				})
			})
		})
	}
}

func TestOpeCode_sla(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
		r1     int
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "RL B",
			args: args{0x20, B},
		},
		{
			name: "SLA C",
			args: args{0x21, C},
		},
		{
			name: "SLA D",
			args: args{0x22, D},
		},
		{
			name: "SLA E",
			args: args{0x23, E},
		},
		{
			name: "SLA H",
			args: args{0x24, H},
		},
		{
			name: "SLA L",
			args: args{0x25, L},
		},
		{
			name: "SLA HL",
			args: args{0x26, HL},
		},
		{
			name: "SLA A",
			args: args{0x27, A},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := cbOpCodes[tt.args.opcode]

			assert.Equal(t, uint8(tt.args.r1), op.R1)

			t.Run("when bit7 = 1", func(t *testing.T) {
				before := byte(0b10010000)
				c.Reg.R[tt.args.r1] = before
				c.Bus.WriteByte(c.Reg.R16(int(HL)), before)
				op.Handler(c, byte(op.R1), byte(op.R2))
				assert.Equal(t, false, c.Reg.isSet(flagZ))
				assert.Equal(t, false, c.Reg.isSet(flagN))
				assert.Equal(t, false, c.Reg.isSet(flagH))
				assert.Equal(t, true, c.Reg.isSet(flagC))

				want := byte(0b00100000)

				if tt.name != "SLA HL" {
					assert.Equal(t, want, c.Reg.R[op.R1])
				} else {
					assert.Equal(t, want, c.Bus.ReadByte(c.Reg.R16(int(HL))))
				}
			})
			t.Run("when bit7 = 0 and other = 0", func(t *testing.T) {
				before := byte(0b00000000)
				c.Reg.R[tt.args.r1] = before
				c.Bus.WriteByte(c.Reg.R16(int(HL)), before)
				op.Handler(c, byte(op.R1), byte(op.R2))
				assert.Equal(t, true, c.Reg.isSet(flagZ))
				assert.Equal(t, false, c.Reg.isSet(flagN))
				assert.Equal(t, false, c.Reg.isSet(flagH))
				assert.Equal(t, false, c.Reg.isSet(flagC))

				want := byte(0b00000000)

				if tt.name != "SLA HL" {
					assert.Equal(t, want, c.Reg.R[op.R1])
				} else {
					assert.Equal(t, want, c.Bus.ReadByte(c.Reg.R16(int(HL))))
				}
			})
			t.Run("when bit7 = 0 and other != 0", func(t *testing.T) {
				before := byte(0b00100000)
				c.Reg.R[tt.args.r1] = before
				c.Bus.WriteByte(c.Reg.R16(int(HL)), before)
				op.Handler(c, byte(op.R1), byte(op.R2))
				assert.Equal(t, false, c.Reg.isSet(flagZ))
				assert.Equal(t, false, c.Reg.isSet(flagN))
				assert.Equal(t, false, c.Reg.isSet(flagH))
				assert.Equal(t, false, c.Reg.isSet(flagC))

				want := byte(0b01000000)

				if tt.name != "SLA HL" {
					assert.Equal(t, want, c.Reg.R[op.R1])
				} else {
					assert.Equal(t, want, c.Bus.ReadByte(c.Reg.R16(int(HL))))
				}
			})
		})
	}
}

func TestOpeCode_sra(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
		r1     int
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "SRA B",
			args: args{0x28, B},
		},
		{
			name: "SRA C",
			args: args{0x29, C},
		},
		{
			name: "SRA D",
			args: args{0x2A, D},
		},
		{
			name: "SRA E",
			args: args{0x2B, E},
		},
		{
			name: "SRA H",
			args: args{0x2C, H},
		},
		{
			name: "SRA L",
			args: args{0x2D, L},
		},
		{
			name: "SRA HL",
			args: args{0x2E, HL},
		},
		{
			name: "SRA A",
			args: args{0x2F, A},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := cbOpCodes[tt.args.opcode]

			assert.Equal(t, uint8(tt.args.r1), op.R1)

			t.Run("when bit0 = 1", func(t *testing.T) {
				before := byte(0b10010001)
				c.Reg.R[tt.args.r1] = before
				c.Bus.WriteByte(c.Reg.R16(int(HL)), before)
				op.Handler(c, byte(op.R1), byte(op.R2))
				assert.Equal(t, false, c.Reg.isSet(flagZ))
				assert.Equal(t, false, c.Reg.isSet(flagN))
				assert.Equal(t, false, c.Reg.isSet(flagH))
				assert.Equal(t, true, c.Reg.isSet(flagC))

				want := byte(0b01001000)

				if tt.name != "SRA HL" {
					assert.Equal(t, want, c.Reg.R[op.R1])
				} else {
					assert.Equal(t, want, c.Bus.ReadByte(c.Reg.R16(int(HL))))
				}
			})
			t.Run("when bit0 = 0 and other = 0", func(t *testing.T) {
				before := byte(0b00000000)
				c.Reg.R[tt.args.r1] = before
				c.Bus.WriteByte(c.Reg.R16(int(HL)), before)
				op.Handler(c, byte(op.R1), byte(op.R2))
				assert.Equal(t, true, c.Reg.isSet(flagZ))
				assert.Equal(t, false, c.Reg.isSet(flagN))
				assert.Equal(t, false, c.Reg.isSet(flagH))
				assert.Equal(t, false, c.Reg.isSet(flagC))

				want := byte(0b00000000)

				if tt.name != "SRA HL" {
					assert.Equal(t, want, c.Reg.R[op.R1])
				} else {
					assert.Equal(t, want, c.Bus.ReadByte(c.Reg.R16(int(HL))))
				}
			})
			t.Run("when bit0 = 0 and other != 0", func(t *testing.T) {
				before := byte(0b00100000)
				c.Reg.R[tt.args.r1] = before
				c.Bus.WriteByte(c.Reg.R16(int(HL)), before)
				op.Handler(c, byte(op.R1), byte(op.R2))
				assert.Equal(t, false, c.Reg.isSet(flagZ))
				assert.Equal(t, false, c.Reg.isSet(flagN))
				assert.Equal(t, false, c.Reg.isSet(flagH))
				assert.Equal(t, false, c.Reg.isSet(flagC))

				want := byte(0b00010000)

				if tt.name != "SRA HL" {
					assert.Equal(t, want, c.Reg.R[op.R1])
				} else {
					assert.Equal(t, want, c.Bus.ReadByte(c.Reg.R16(int(HL))))
				}
			})
		})
	}
}

func TestOpeCode_swap(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
		r1     int
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "SWAP B",
			args: args{0x30, B},
		},
		{
			name: "SWAP C",
			args: args{0x31, C},
		},
		{
			name: "SWAP D",
			args: args{0x32, D},
		},
		{
			name: "SWAP E",
			args: args{0x33, E},
		},
		{
			name: "SWAP H",
			args: args{0x34, H},
		},
		{
			name: "SWAP L",
			args: args{0x35, L},
		},
		{
			name: "SWAP HL",
			args: args{0x36, HL},
		},
		{
			name: "SWAP A",
			args: args{0x37, A},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := cbOpCodes[tt.args.opcode]

			assert.Equal(t, uint8(tt.args.r1), op.R1)

			t.Run("when != 0 ", func(t *testing.T) {
				before := byte(0b10010100)
				c.Reg.R[tt.args.r1] = before
				c.Bus.WriteByte(c.Reg.R16(int(HL)), before)
				op.Handler(c, byte(op.R1), byte(op.R2))
				assert.Equal(t, false, c.Reg.isSet(flagZ))
				assert.Equal(t, false, c.Reg.isSet(flagN))
				assert.Equal(t, false, c.Reg.isSet(flagH))
				assert.Equal(t, false, c.Reg.isSet(flagC))

				want := byte(0b01001001)

				if tt.name != "SWAP HL" {
					assert.Equal(t, want, c.Reg.R[op.R1])
				} else {
					assert.Equal(t, want, c.Bus.ReadByte(c.Reg.R16(int(HL))))
				}
			})
			t.Run("when = 0", func(t *testing.T) {
				before := byte(0b00000000)
				c.Reg.R[tt.args.r1] = before
				c.Bus.WriteByte(c.Reg.R16(int(HL)), before)
				op.Handler(c, byte(op.R1), byte(op.R2))
				assert.Equal(t, true, c.Reg.isSet(flagZ))
				assert.Equal(t, false, c.Reg.isSet(flagN))
				assert.Equal(t, false, c.Reg.isSet(flagH))
				assert.Equal(t, false, c.Reg.isSet(flagC))

				want := byte(0b00000000)

				if tt.name != "SWAP HL" {
					assert.Equal(t, want, c.Reg.R[op.R1])
				} else {
					assert.Equal(t, want, c.Bus.ReadByte(c.Reg.R16(int(HL))))
				}
			})
		})
	}
}

func TestOpeCode_srl(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
		r1     int
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "SRL B",
			args: args{0x38, B},
		},
		{
			name: "SRL C",
			args: args{0x39, C},
		},
		{
			name: "SRL D",
			args: args{0x3A, D},
		},
		{
			name: "SRL E",
			args: args{0x3B, E},
		},
		{
			name: "SRL H",
			args: args{0x3C, H},
		},
		{
			name: "SRL L",
			args: args{0x3D, L},
		},
		{
			name: "SRL HL",
			args: args{0x3E, HL},
		},
		{
			name: "SRL A",
			args: args{0x3F, A},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := cbOpCodes[tt.args.opcode]

			assert.Equal(t, uint8(tt.args.r1), op.R1)

			t.Run("when bit0 = 1", func(t *testing.T) {
				before := byte(0b10010001)
				c.Reg.R[tt.args.r1] = before
				c.Bus.WriteByte(c.Reg.R16(int(HL)), before)
				op.Handler(c, byte(op.R1), byte(op.R2))
				assert.Equal(t, false, c.Reg.isSet(flagZ))
				assert.Equal(t, false, c.Reg.isSet(flagN))
				assert.Equal(t, false, c.Reg.isSet(flagH))
				assert.Equal(t, true, c.Reg.isSet(flagC))

				want := byte(0b01001000)

				if tt.name != "SRL HL" {
					assert.Equal(t, want, c.Reg.R[op.R1])
				} else {
					assert.Equal(t, want, c.Bus.ReadByte(c.Reg.R16(int(HL))))
				}
			})
			t.Run("when bit0 = 0 and other = 0", func(t *testing.T) {
				before := byte(0b00000000)
				c.Reg.R[tt.args.r1] = before
				c.Bus.WriteByte(c.Reg.R16(int(HL)), before)
				op.Handler(c, byte(op.R1), byte(op.R2))
				assert.Equal(t, true, c.Reg.isSet(flagZ))
				assert.Equal(t, false, c.Reg.isSet(flagN))
				assert.Equal(t, false, c.Reg.isSet(flagH))
				assert.Equal(t, false, c.Reg.isSet(flagC))

				want := byte(0b00000000)

				if tt.name != "SRL HL" {
					assert.Equal(t, want, c.Reg.R[op.R1])
				} else {
					assert.Equal(t, want, c.Bus.ReadByte(c.Reg.R16(int(HL))))
				}
			})
			t.Run("when bit0 = 0 and other != 0", func(t *testing.T) {
				before := byte(0b00100000)
				c.Reg.R[tt.args.r1] = before
				c.Bus.WriteByte(c.Reg.R16(int(HL)), before)
				op.Handler(c, byte(op.R1), byte(op.R2))
				assert.Equal(t, false, c.Reg.isSet(flagZ))
				assert.Equal(t, false, c.Reg.isSet(flagN))
				assert.Equal(t, false, c.Reg.isSet(flagH))
				assert.Equal(t, false, c.Reg.isSet(flagC))

				want := byte(0b00010000)

				if tt.name != "SRL HL" {
					assert.Equal(t, want, c.Reg.R[op.R1])
				} else {
					assert.Equal(t, want, c.Bus.ReadByte(c.Reg.R16(int(HL))))
				}
			})
		})
	}
}
