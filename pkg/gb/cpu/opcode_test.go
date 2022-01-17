package cpu

import (
	"strings"
	"testing"

	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/bus"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/cartridge"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/gpu"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/io"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/memory"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/types"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/util"
	"github.com/stretchr/testify/assert"
)

func setupCPU() *CPU {
	romData := make([]byte, 0x8000)
	cart := cartridge.New(romData)
	vram := memory.NewRAM(0x2000)
	wram := memory.NewRAM(0x2000)
	wram2 := memory.NewRAM(0x2000)
	hram := memory.NewRAM(0x0080)
	io := io.NewIO(io.NewPad(), io.NewSerial(), io.NewTimer(), io.NewIRQ(), gpu.New(), 0x2000)
	bus := bus.New(cart, vram, wram, wram2, hram, io)

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

func TestOpCode_nop(t *testing.T) {}

// test 0x40-0x6F (except 0xX6, 0xXE)
func TestOpCode_ldrr(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
		r1     int
		r2     int
	}

	tests := []struct {
		name string
		args args
	}{
		{name: "LD B, B", args: args{0x40, B, B}},
		{name: "LD B, C", args: args{0x41, B, C}},
		{name: "LD B, D", args: args{0x42, B, D}},
		{name: "LD B, E", args: args{0x43, B, E}},
		{name: "LD B, H", args: args{0x44, B, H}},
		{name: "LD B, L", args: args{0x45, B, L}},
		{name: "LD B, A", args: args{0x47, B, A}},
		{name: "LD C, B", args: args{0x48, C, B}},
		{name: "LD C, C", args: args{0x49, C, C}},
		{name: "LD C, D", args: args{0x4A, C, D}},
		{name: "LD C, E", args: args{0x4B, C, E}},
		{name: "LD C, H", args: args{0x4C, C, H}},
		{name: "LD C, L", args: args{0x4D, C, L}},
		{name: "LD C, A", args: args{0x4F, C, A}},
		{name: "LD D, B", args: args{0x50, D, B}},
		{name: "LD D, C", args: args{0x51, D, C}},
		{name: "LD D, D", args: args{0x52, D, D}},
		{name: "LD D, E", args: args{0x53, D, E}},
		{name: "LD D, H", args: args{0x54, D, H}},
		{name: "LD D, L", args: args{0x55, D, L}},
		{name: "LD D, A", args: args{0x57, D, A}},
		{name: "LD E, B", args: args{0x58, E, B}},
		{name: "LD E, C", args: args{0x59, E, C}},
		{name: "LD E, D", args: args{0x5A, E, D}},
		{name: "LD E, E", args: args{0x5B, E, E}},
		{name: "LD E, H", args: args{0x5C, E, H}},
		{name: "LD E, L", args: args{0x5D, E, L}},
		{name: "LD E, A", args: args{0x5F, E, A}},
		{name: "LD H, B", args: args{0x60, H, B}},
		{name: "LD H, C", args: args{0x61, H, C}},
		{name: "LD H, D", args: args{0x62, H, D}},
		{name: "LD H, E", args: args{0x63, H, E}},
		{name: "LD H, H", args: args{0x64, H, H}},
		{name: "LD H, L", args: args{0x65, H, L}},
		{name: "LD H, A", args: args{0x67, H, A}},
		{name: "LD L, B", args: args{0x68, L, B}},
		{name: "LD L, C", args: args{0x69, L, C}},
		{name: "LD L, D", args: args{0x6A, L, D}},
		{name: "LD L, E", args: args{0x6B, L, E}},
		{name: "LD L, H", args: args{0x6C, L, H}},
		{name: "LD L, L", args: args{0x6D, L, L}},
		{name: "LD L, A", args: args{0x6F, L, A}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := opCodes[tt.args.opcode]
			want := byte(0x12)
			c.Reg.R[tt.args.r2] = want
			op.Handler(c, op.R1, op.R2)

			assert.Equal(t, tt.args.r1, op.R1)
			assert.Equal(t, tt.args.r2, op.R2)
			assert.Equal(t, want, c.Reg.R[op.R1])
			assert.Equal(t, want, c.Reg.R[op.R2])
		})
	}
}

func TestOpCode_ldrm(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
		r1     int
		r2     int
	}

	tests := []struct {
		name string
		args args
	}{
		{name: "LD A,(C)", args: args{0xF2, A, C}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := opCodes[tt.args.opcode]
			addr := types.Addr(0xFF03)
			want := byte(0x12)
			c.Bus.WriteByte(addr, want)
			op.Handler(c, op.R1, op.R2)

			assert.Equal(t, tt.args.r1, op.R1)
			assert.Equal(t, tt.args.r2, op.R2)
			assert.Equal(t, want, c.Reg.R[A])
		})
	}
}

func TestOpCode_ldrm16(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
		r1     int
		r2     int
	}

	tests := []struct {
		name string
		args args
	}{
		{name: "LD A, (BC)", args: args{0x0A, A, BC}},
		{name: "LD A, (DE)", args: args{0x1A, A, DE}},
		{name: "LD A, (HL+)", args: args{0x2A, A, HLI}},
		{name: "LD A, (HL-)", args: args{0x3A, A, HLD}},
		{name: "LD B, (HL)", args: args{0x46, B, HL}},
		{name: "LD C, (HL)", args: args{0x4E, C, HL}},
		{name: "LD D, (HL)", args: args{0x56, D, HL}},
		{name: "LD E, (HL)", args: args{0x5E, E, HL}},
		{name: "LD H, (HL)", args: args{0x66, H, HL}},
		{name: "LD L, (HL)", args: args{0x6E, L, HL}},
		{name: "LD A, (HL)", args: args{0x7E, A, HL}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := opCodes[tt.args.opcode]
			want := byte(0x20)
			c.Bus.WriteByte(c.Reg.R16(tt.args.r2), want)
			if op.R2 == HLI {
				c.Reg.setHL(c.Reg.HL() - 1)
			} else if op.R2 == HLD {
				c.Reg.setHL(c.Reg.HL() + 1)
			}
			op.Handler(c, op.R1, op.R2)

			assert.Equal(t, tt.args.r1, op.R1)
			assert.Equal(t, tt.args.r2, op.R2)
			assert.Equal(t, want, c.Reg.R[op.R1])
		})
	}
}

func TestOpCode_ldrd(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
		r1     int
	}

	tests := []struct {
		name string
		args args
	}{
		{name: "LD B, d8", args: args{0x06, B}},
		{name: "LD C, d8", args: args{0x0E, C}},
		{name: "LD D, d8", args: args{0x16, D}},
		{name: "LD E, d8", args: args{0x1E, E}},
		{name: "LD H, d8", args: args{0x26, H}},
		{name: "LD L, d8", args: args{0x2E, L}},
		{name: "LD A, d8", args: args{0x3E, A}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := opCodes[tt.args.opcode]
			want := byte(0x20)
			c.Bus.WriteByte(c.Reg.PC, want)
			op.Handler(c, op.R1, op.R2)

			assert.Equal(t, tt.args.r1, op.R1)
			assert.Equal(t, want, c.Reg.R[op.R1])
		})
	}
}

func TestOpCode_ldra(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
		r1     int
	}

	tests := []struct {
		name string
		args args
	}{
		{name: "LDH A,(a8)", args: args{0xF0, A}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := opCodes[tt.args.opcode]
			addr := types.Addr(0xFF12)
			a8 := byte(0x12)
			want := byte(0x34)

			c.Bus.WriteByte(c.Reg.PC, a8)

			// $FF12 = 0x34
			c.Bus.WriteByte(addr, want)

			// A = 0x34
			op.Handler(c, op.R1, op.R2)

			assert.Equal(t, tt.args.r1, op.R1)
			assert.Equal(t, want, c.Reg.R[op.R1])
		})
	}
}

func TestOpCode_ldra16(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
		r1     int
	}

	tests := []struct {
		name string
		args args
	}{
		{name: "LD A,(a16)", args: args{0xFA, A}},
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

			op.Handler(c, op.R1, op.R2)

			assert.Equal(t, tt.args.r1, op.R1)
			assert.Equal(t, want, c.Reg.R[A])
		})
	}
}

func TestOpCode_ldr16d16(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
		r1     int
	}

	tests := []struct {
		name string
		args args
	}{
		{name: "LD BC, d16", args: args{0x01, BC}},
		{name: "LD DE, d16", args: args{0x11, DE}},
		{name: "LD HL, d16", args: args{0x21, HL}},
		{name: "LD SP, d16", args: args{0x31, SP}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := opCodes[tt.args.opcode]
			want := byte(0x20)
			c.Bus.WriteByte(c.Reg.PC, want)
			c.Bus.WriteByte(c.Reg.PC+1, want+1)
			op.Handler(c, op.R1, op.R2)

			assert.Equal(t, tt.args.r1, op.R1)
			assert.Equal(t, util.Byte2Addr(want+1, want), c.Reg.R16(op.R1))
		})
	}
}

func TestOpCode_ldr16r16(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
		r1     int
		r2     int
	}

	tests := []struct {
		name string
		args args
	}{
		{name: "LD SP,HL", args: args{0xF9, SP, HL}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := opCodes[tt.args.opcode]
			want := c.Reg.R16(tt.args.r2)
			op.Handler(c, op.R1, op.R2)

			assert.Equal(t, tt.args.r1, op.R1)
			assert.Equal(t, tt.args.r2, op.R2)
			assert.Equal(t, want, c.Reg.R16(op.R1))
		})
	}
}

// TODO fix test
func TestOpCode_ldr16r16d(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
		r1     int
		r2     int
	}

	tests := []struct {
		name string
		args args
	}{
		{name: "LD HL,SP+r8", args: args{0xF8, HL, SP}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := opCodes[tt.args.opcode]

			assert.Equal(t, tt.args.r1, op.R1)
			assert.Equal(t, tt.args.r2, op.R2)

			t.Run("when no carry", func(t *testing.T) {
				v2 := 0x1E00
				d := byte(0xE1)
				want := types.Addr(0x1DE1)
				c.Reg.setR16(tt.args.r2, types.Addr(v2))
				c.Bus.WriteByte(c.Reg.PC, d)
				op.Handler(c, op.R1, op.R2)

				assert.Equal(t, want, c.Reg.R16(tt.args.r1))
				assert.Equal(t, false, c.Reg.isSet(flagZ))
				assert.Equal(t, false, c.Reg.isSet(flagN))
				assert.Equal(t, false, c.Reg.isSet(flagH))
				assert.Equal(t, false, c.Reg.isSet(flagC))
			})
			t.Run("when half carry", func(t *testing.T) {
				v2 := 0x1F20
				d := byte(0xE1)
				want := types.Addr(0x2001)
				c.Reg.setR16(tt.args.r2, types.Addr(v2))
				c.Bus.WriteByte(c.Reg.PC, d)
				op.Handler(c, op.R1, op.R2)

				assert.Equal(t, want, c.Reg.R16(tt.args.r1))
				assert.Equal(t, false, c.Reg.isSet(flagZ))
				assert.Equal(t, false, c.Reg.isSet(flagN))
				assert.Equal(t, true, c.Reg.isSet(flagH))
				assert.Equal(t, false, c.Reg.isSet(flagC))
			})
			t.Run("when carry", func(t *testing.T) {
				v2 := 0xFF20
				d := byte(0xE1)
				want := types.Addr(0x0001)
				c.Reg.setR16(tt.args.r2, types.Addr(v2))
				c.Bus.WriteByte(c.Reg.PC, d)
				op.Handler(c, op.R1, op.R2)

				assert.Equal(t, want, c.Reg.R16(tt.args.r1))
				assert.Equal(t, false, c.Reg.isSet(flagZ))
				assert.Equal(t, false, c.Reg.isSet(flagN))
				assert.Equal(t, true, c.Reg.isSet(flagH))
				assert.Equal(t, true, c.Reg.isSet(flagC))
			})
		})
	}
}

func TestOpCode_ldmr(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
		r1     int
		r2     int
	}

	tests := []struct {
		name string
		args args
	}{
		{name: "LD (C), A", args: args{0xE2, C, A}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := opCodes[tt.args.opcode]

			addr := util.Byte2Addr(0xFF, c.Reg.R[tt.args.r1])
			want := c.Reg.R[tt.args.r2]
			op.Handler(c, op.R1, op.R2)

			assert.Equal(t, tt.args.r1, op.R1)
			assert.Equal(t, tt.args.r2, op.R2)
			assert.Equal(t, want, c.Bus.ReadByte(addr))
		})
	}
}

func TestOpCode_ldm16r(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
		r1     int
		r2     int
	}

	tests := []struct {
		name string
		args args
	}{
		{name: "LD (BC), A", args: args{0x02, BC, A}},
		{name: "LD (DE), A", args: args{0x12, DE, A}},
		{name: "LD (HL+), A", args: args{0x22, HLI, A}},
		{name: "LD (HL-), A", args: args{0x32, HLD, A}},
		{name: "LD (HL), B", args: args{0x70, HL, B}},
		{name: "LD (HL), C", args: args{0x71, HL, C}},
		{name: "LD (HL), D", args: args{0x72, HL, D}},
		{name: "LD (HL), E", args: args{0x73, HL, E}},
		{name: "LD (HL), H", args: args{0x74, HL, H}},
		{name: "LD (HL), L", args: args{0x75, HL, L}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := opCodes[tt.args.opcode]
			want := c.Reg.R[tt.args.r2]
			op.Handler(c, op.R1, op.R2)

			assert.Equal(t, tt.args.r1, op.R1)
			assert.Equal(t, tt.args.r2, op.R2)
			if op.R1 == HLI {
				c.Reg.setHL(c.Reg.HL() - 1)
			} else if op.R1 == HLD {
				c.Reg.setHL(c.Reg.HL() + 1)
			}
			assert.Equal(t, want, c.Bus.ReadByte(c.Reg.R16(tt.args.r1)))
			assert.Equal(t, want, c.Reg.R[tt.args.r2])
		})
	}
}

func TestOpCode_ldm16d(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
		r1     int
	}

	tests := []struct {
		name string
		args args
	}{
		{name: "LD (HL),d8", args: args{0x36, HL}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := opCodes[tt.args.opcode]

			want := byte(0x12)
			c.Bus.WriteByte(c.Reg.PC, 0x12)

			op.Handler(c, op.R1, op.R2)

			assert.Equal(t, tt.args.r1, op.R1)
			assert.Equal(t, want, c.Bus.ReadByte(c.Reg.R16(tt.args.r1)))
		})
	}
}

func TestOpCode_ldar(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
		r1     int
	}

	tests := []struct {
		name string
		args args
	}{
		{name: "LDH (a8),A", args: args{0xE0, A}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := opCodes[tt.args.opcode]

			addr := byte(0x12)
			want := byte(0x01)
			c.Bus.WriteByte(c.Reg.PC, addr)

			op.Handler(c, op.R1, op.R2)

			assert.Equal(t, op.R2, tt.args.r1)
			assert.Equal(t, want, c.Bus.ReadByte(util.Byte2Addr(0xFF, addr)))
		})
	}
}

func TestOpCode_lda16r(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
		r1     int
	}

	tests := []struct {
		name string
		args args
	}{
		{name: "LD (a16),A", args: args{0xEA, A}},
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

			op.Handler(c, op.R1, op.R2)

			assert.Equal(t, op.R2, tt.args.r1)
			assert.Equal(t, want, c.Bus.ReadByte(addr))
		})
	}
}

func TestOpCode_lda16r16(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
		r1     int
	}

	tests := []struct {
		name string
		args args
	}{
		{name: "LD (a16),SP", args: args{0x08, SP}},
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

			op.Handler(c, op.R1, op.R2)

			assert.Equal(t, op.R2, tt.args.r1)
			assert.Equal(t, util.ExtractLower(want), c.Bus.ReadByte(addr))
			assert.Equal(t, util.ExtractUpper(want), c.Bus.ReadByte(addr+1))
		})
	}
}

// -arithmetic-

func TestOpCode_incr(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
		r1     int
	}

	tests := []struct {
		name string
		args args
	}{
		{name: "INC B", args: args{0x04, B}},
		{name: "INC C", args: args{0x0C, C}},
		{name: "INC D", args: args{0x14, D}},
		{name: "INC E", args: args{0x1C, E}},
		{name: "INC H", args: args{0x24, H}},
		{name: "INC L", args: args{0x2C, L}},
		{name: "INC A", args: args{0x3C, A}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := opCodes[tt.args.opcode]

			t.Run("when no carry", func(t *testing.T) {
				c.Reg.R[tt.args.r1] = byte(0x10)
				want := byte(0x11)
				op.Handler(c, op.R1, op.R2)
				assert.Equal(t, tt.args.r1, op.R1)
				assert.Equal(t, want, c.Reg.R[op.R1])
				assert.Equal(t, false, c.Reg.isSet(flagZ))
				assert.Equal(t, false, c.Reg.isSet(flagN))
				assert.Equal(t, false, c.Reg.isSet(flagH))
			})
			t.Run("when harf carry", func(t *testing.T) {
				c.Reg.R[tt.args.r1] = byte(0x1F)
				want := byte(0x20)
				op.Handler(c, op.R1, op.R2)
				assert.Equal(t, tt.args.r1, op.R1)
				assert.Equal(t, want, c.Reg.R[op.R1])
				assert.Equal(t, false, c.Reg.isSet(flagZ))
				assert.Equal(t, false, c.Reg.isSet(flagN))
				assert.Equal(t, true, c.Reg.isSet(flagH))
			})
			t.Run("when zero", func(t *testing.T) {
				c.Reg.R[tt.args.r1] = byte(0xFF)
				want := byte(0x00)
				op.Handler(c, op.R1, op.R2)
				assert.Equal(t, tt.args.r1, op.R1)
				assert.Equal(t, want, c.Reg.R[op.R1])
				assert.Equal(t, true, c.Reg.isSet(flagZ))
				assert.Equal(t, false, c.Reg.isSet(flagN))
				assert.Equal(t, true, c.Reg.isSet(flagH))
			})
		})
	}
}

func TestOpCode_incr16(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
		r1     int
	}

	tests := []struct {
		name string
		args args
	}{
		{name: "INC BC", args: args{0x03, BC}},
		{name: "INC DE", args: args{0x13, DE}},
		{name: "INC HL", args: args{0x23, HL}},
		{name: "INC SP", args: args{0x33, SP}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := opCodes[tt.args.opcode]

			c.Reg.setR16(int(op.R1), 0x1234)
			want := types.Addr(0x1235)

			op.Handler(c, op.R1, op.R2)

			assert.Equal(t, tt.args.r1, op.R1)
			assert.Equal(t, want, c.Reg.R16(int(op.R1)))
		})
	}
}

func TestOpCode_incm16(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
		r1     int
	}

	tests := []struct {
		name string
		args args
	}{
		{name: "INC (HL)", args: args{0x34, HL}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := opCodes[tt.args.opcode]

			var want, act byte

			// zero
			c.Bus.WriteByte(c.Reg.R16(tt.args.r1), byte(0xFF))
			want = byte(0x00)

			op.Handler(c, op.R1, op.R2)

			act = c.Bus.ReadByte(c.Reg.R16(int(op.R1)))

			assert.Equal(t, tt.args.r1, op.R1)
			assert.Equal(t, want, act)
			assert.Equal(t, true, c.Reg.isSet(flagZ))
			assert.Equal(t, false, c.Reg.isSet(flagN))
			assert.Equal(t, true, c.Reg.isSet(flagH))

			// not zero
			c.Bus.WriteByte(c.Reg.R16(tt.args.r1), byte(0x00))
			want = byte(0x01)

			op.Handler(c, op.R1, op.R2)

			act = c.Bus.ReadByte(c.Reg.R16(int(op.R1)))

			assert.Equal(t, tt.args.r1, op.R1)
			assert.Equal(t, want, act)
			assert.Equal(t, false, c.Reg.isSet(flagZ))
			assert.Equal(t, false, c.Reg.isSet(flagN))
			assert.Equal(t, false, c.Reg.isSet(flagH))
		})
	}
}

func TestOpCode_dec(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
		r1     int
	}

	tests := []struct {
		name string
		args args
	}{
		{name: "DEC B", args: args{0x05, B}},
		{name: "DEC C", args: args{0x0D, C}},
		{name: "DEC D", args: args{0x15, D}},
		{name: "DEC E", args: args{0x1D, E}},
		{name: "DEC H", args: args{0x25, H}},
		{name: "DEC L", args: args{0x2D, L}},
		{name: "DEC (HL)", args: args{0x35, HL}},
		{name: "DEC A", args: args{0x3D, A}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := opCodes[tt.args.opcode]

			t.Run("when no carry", func(t *testing.T) {
				if strings.Contains(op.Mnemonic, "HL") {
					c.Bus.WriteByte(c.Reg.R16(HL), byte(0x11))
				} else {
					c.Reg.R[tt.args.r1] = byte(0x11)
				}
				want := byte(0x10)
				op.Handler(c, op.R1, op.R2)
				assert.Equal(t, tt.args.r1, op.R1)
				if strings.Contains(op.Mnemonic, "HL") {
					assert.Equal(t, want, c.Bus.ReadByte(c.Reg.R16(HL)))
				} else {
					assert.Equal(t, want, c.Reg.R[op.R1])
				}
				assert.Equal(t, false, c.Reg.isSet(flagZ))
				assert.Equal(t, true, c.Reg.isSet(flagN))
				assert.Equal(t, false, c.Reg.isSet(flagH))
			})
			t.Run("when harf carry", func(t *testing.T) {
				if strings.Contains(op.Mnemonic, "HL") {
					c.Bus.WriteByte(c.Reg.R16(HL), byte(0x10))
				} else {
					c.Reg.R[tt.args.r1] = byte(0x10)
				}
				want := byte(0x0F)
				op.Handler(c, op.R1, op.R2)
				assert.Equal(t, tt.args.r1, op.R1)
				if strings.Contains(op.Mnemonic, "HL") {
					assert.Equal(t, want, c.Bus.ReadByte(c.Reg.R16(HL)))
				} else {
					assert.Equal(t, want, c.Reg.R[op.R1])
				}
				assert.Equal(t, false, c.Reg.isSet(flagZ))
				assert.Equal(t, true, c.Reg.isSet(flagN))
				assert.Equal(t, true, c.Reg.isSet(flagH))
			})
			t.Run("when zero", func(t *testing.T) {
				if strings.Contains(op.Mnemonic, "HL") {
					c.Bus.WriteByte(c.Reg.R16(HL), byte(0x01))
				} else {
					c.Reg.R[tt.args.r1] = byte(0x01)
				}
				want := byte(0x00)
				op.Handler(c, op.R1, op.R2)
				assert.Equal(t, tt.args.r1, op.R1)
				if strings.Contains(op.Mnemonic, "HL") {
					assert.Equal(t, want, c.Bus.ReadByte(c.Reg.R16(HL)))
				} else {
					assert.Equal(t, want, c.Reg.R[op.R1])
				}
				assert.Equal(t, true, c.Reg.isSet(flagZ))
				assert.Equal(t, true, c.Reg.isSet(flagN))
				assert.Equal(t, false, c.Reg.isSet(flagH))
			})
		})
	}
}

func TestOpCode_decr16(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
		r1     int
	}

	tests := []struct {
		name string
		args args
	}{
		{name: "DEC BC", args: args{0x0B, BC}},
		{name: "DEC DE", args: args{0x1B, DE}},
		{name: "DEC HL", args: args{0x2B, HL}},
		{name: "DEC SP", args: args{0x3B, SP}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := opCodes[tt.args.opcode]

			c.Reg.setR16(int(op.R1), 0x1234)
			want := types.Addr(0x1233)

			op.Handler(c, op.R1, op.R2)

			assert.Equal(t, tt.args.r1, op.R1)
			assert.Equal(t, want, c.Reg.R16(int(op.R1)))
		})
	}
}

func TestOpCode_add(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
		r1     int
		r2     int
	}

	tests := []struct {
		name string
		args args
	}{
		{name: "ADD A, B", args: args{0x80, A, B}},
		{name: "ADD A, C", args: args{0x81, A, C}},
		{name: "ADD A, D", args: args{0x82, A, D}},
		{name: "ADD A, E", args: args{0x83, A, E}},
		{name: "ADD A, H", args: args{0x84, A, H}},
		{name: "ADD A, L", args: args{0x85, A, L}},
		{name: "ADD A, HL", args: args{0x86, A, HL}},
		{name: "ADD A, A", args: args{0x87, A, A}},
		{name: "ADD A, d8", args: args{0xC6, A, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := opCodes[tt.args.opcode]

			assert.Equal(t, tt.args.r1, op.R1)
			assert.Equal(t, tt.args.r2, op.R2)
			if strings.Contains(op.Mnemonic, "A, A") {
				t.Skip()
			}

			t.Run("when no carry", func(t *testing.T) {
				c.Reg.R[tt.args.r1] = 0xE1
				setVal := byte(0x0E)
				if strings.Contains(op.Mnemonic, "HL") {
					c.Bus.WriteByte(c.Reg.R16(HL), setVal)
				} else if strings.Contains(op.Mnemonic, "d8") {
					c.Bus.WriteByte(c.Reg.PC, setVal)
				} else {
					c.Reg.R[tt.args.r2] = setVal
				}
				op.Handler(c, op.R1, op.R2)
				assert.Equal(t, byte(0xEF), c.Reg.R[A])
				assert.Equal(t, false, c.Reg.isSet(flagZ))
				assert.Equal(t, false, c.Reg.isSet(flagN))
				assert.Equal(t, false, c.Reg.isSet(flagH))
				assert.Equal(t, false, c.Reg.isSet(flagC))
			})
			t.Run("when Harf carry", func(t *testing.T) {
				c.Reg.R[tt.args.r1] = 0xE1
				setVal := byte(0x0F)
				if strings.Contains(op.Mnemonic, "HL") {
					c.Bus.WriteByte(c.Reg.R16(HL), setVal)
				} else if strings.Contains(op.Mnemonic, "d8") {
					c.Bus.WriteByte(c.Reg.PC, setVal)
				} else {
					c.Reg.R[tt.args.r2] = setVal
				}
				op.Handler(c, op.R1, op.R2)
				assert.Equal(t, byte(0xF0), c.Reg.R[A])
				assert.Equal(t, false, c.Reg.isSet(flagZ))
				assert.Equal(t, false, c.Reg.isSet(flagN))
				assert.Equal(t, true, c.Reg.isSet(flagH))
				assert.Equal(t, false, c.Reg.isSet(flagC))
			})
			t.Run("when carry and zero", func(t *testing.T) {
				c.Reg.R[tt.args.r1] = 0xE1
				setVal := byte(0x1F)
				if strings.Contains(op.Mnemonic, "HL") {
					c.Bus.WriteByte(c.Reg.R16(HL), setVal)
				} else if strings.Contains(op.Mnemonic, "d8") {
					c.Bus.WriteByte(c.Reg.PC, setVal)
				} else {
					c.Reg.R[tt.args.r2] = setVal
				}
				op.Handler(c, op.R1, op.R2)
				assert.Equal(t, byte(0x00), c.Reg.R[A])
				assert.Equal(t, true, c.Reg.isSet(flagZ))
				assert.Equal(t, false, c.Reg.isSet(flagN))
				assert.Equal(t, true, c.Reg.isSet(flagH))
				assert.Equal(t, true, c.Reg.isSet(flagC))
			})
		})
	}
}

func TestOpCode_addr16(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
		r1     int
		r2     int
	}

	tests := []struct {
		name string
		args args
	}{
		{name: "ADD HL,BC", args: args{0x09, HL, BC}},
		{name: "ADD HL,DE", args: args{0x19, HL, DE}},
		{name: "ADD HL,HL", args: args{0x29, HL, HL}},
		{name: "ADD HL,SP", args: args{0x39, HL, SP}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := opCodes[tt.args.opcode]

			assert.Equal(t, tt.args.r1, op.R1)
			assert.Equal(t, tt.args.r2, op.R2)

			if strings.Contains(op.Mnemonic, "HL,HL") {
				t.Skip()
			}

			t.Run("when no carry", func(t *testing.T) {
				c.Reg.setR16(tt.args.r1, 0x00E1)
				c.Reg.setR16(tt.args.r2, 0x000E)
				op.Handler(c, op.R1, op.R2)
				assert.Equal(t, types.Addr(0x00EF), c.Reg.R16(int(op.R1)))
				assert.Equal(t, false, c.Reg.isSet(flagN))
				assert.Equal(t, false, c.Reg.isSet(flagH))
				assert.Equal(t, false, c.Reg.isSet(flagC))
			})
			t.Run("when Harf carry", func(t *testing.T) {
				c.Reg.setR16(tt.args.r1, 0x0FF1)
				c.Reg.setR16(tt.args.r2, 0x000F)
				op.Handler(c, op.R1, op.R2)
				assert.Equal(t, types.Addr(0x1000), c.Reg.R16(int(op.R1)))
				assert.Equal(t, false, c.Reg.isSet(flagN))
				assert.Equal(t, true, c.Reg.isSet(flagH))
				assert.Equal(t, false, c.Reg.isSet(flagC))
			})
			t.Run("when carry and zero", func(t *testing.T) {
				c.Reg.setR16(tt.args.r1, 0xFFF1)
				c.Reg.setR16(tt.args.r2, 0x000F)
				op.Handler(c, op.R1, op.R2)
				assert.Equal(t, types.Addr(0x0000), c.Reg.R16(int(op.R1)))
				assert.Equal(t, false, c.Reg.isSet(flagN))
				assert.Equal(t, true, c.Reg.isSet(flagH))
				assert.Equal(t, true, c.Reg.isSet(flagC))
			})
		})
	}
}

// TODO Fix test
func TestOpCode_addr16d(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
		r1     int
		r2     int
	}

	tests := []struct {
		name string
		args args
	}{
		{name: "ADD SP,r8", args: args{0xE8, SP, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := opCodes[tt.args.opcode]

			assert.Equal(t, tt.args.r1, op.R1)

			t.Run("when no carry", func(t *testing.T) {
				c.Reg.setR16(tt.args.r1, 0x01E1)
				c.Bus.WriteByte(c.Reg.PC, 0x0D)
				op.Handler(c, op.R1, op.R2)
				assert.Equal(t, types.Addr(0x01EE), c.Reg.R16(int(op.R1)))
				assert.Equal(t, false, c.Reg.isSet(flagN))
				assert.Equal(t, false, c.Reg.isSet(flagH))
				assert.Equal(t, false, c.Reg.isSet(flagC))
				assert.Equal(t, false, c.Reg.isSet(flagZ))
			})
			t.Run("when Harf carry", func(t *testing.T) {
				c.Reg.setR16(tt.args.r1, 0x0FD1)
				c.Bus.WriteByte(c.Reg.PC, 0x0F)
				op.Handler(c, op.R1, op.R2)
				assert.Equal(t, types.Addr(0x0FE0), c.Reg.R16(int(op.R1)))
				assert.Equal(t, false, c.Reg.isSet(flagN))
				assert.Equal(t, true, c.Reg.isSet(flagH))
				assert.Equal(t, false, c.Reg.isSet(flagC))
				assert.Equal(t, false, c.Reg.isSet(flagZ))
			})
			t.Run("when carry", func(t *testing.T) {
				c.Reg.setR16(tt.args.r1, 0x0DF1)
				c.Bus.WriteByte(c.Reg.PC, 0x0F)
				op.Handler(c, op.R1, op.R2)
				assert.Equal(t, types.Addr(0x0E00), c.Reg.R16(int(op.R1)))
				assert.Equal(t, false, c.Reg.isSet(flagN))
				assert.Equal(t, true, c.Reg.isSet(flagH))
				assert.Equal(t, true, c.Reg.isSet(flagC))
				assert.Equal(t, false, c.Reg.isSet(flagZ))
			})
			t.Run("when negative no carry", func(t *testing.T) {
				c.Reg.setR16(tt.args.r1, 0x0F81)
				c.Bus.WriteByte(c.Reg.PC, 0x80)
				op.Handler(c, op.R1, op.R2)
				assert.Equal(t, types.Addr(0x0F01), c.Reg.R16(int(op.R1)))
				assert.Equal(t, false, c.Reg.isSet(flagN))
				assert.Equal(t, false, c.Reg.isSet(flagH))
				assert.Equal(t, false, c.Reg.isSet(flagC))
				assert.Equal(t, false, c.Reg.isSet(flagZ))
			})
			t.Run("when negative half carry", func(t *testing.T) {
				c.Reg.setR16(tt.args.r1, 0x0FD1) //4049
				c.Bus.WriteByte(c.Reg.PC, 0x8F)  // -113
				op.Handler(c, op.R1, op.R2)
				assert.Equal(t, types.Addr(0x0F60), c.Reg.R16(int(op.R1))) // 3936
				assert.Equal(t, false, c.Reg.isSet(flagN))
				assert.Equal(t, true, c.Reg.isSet(flagH))
				assert.Equal(t, false, c.Reg.isSet(flagC))
				assert.Equal(t, false, c.Reg.isSet(flagZ))
			})
			t.Run("when negative carry", func(t *testing.T) {
				c.Reg.setR16(tt.args.r1, 0x0FD1) //4049
				c.Bus.WriteByte(c.Reg.PC, 0xE0)  // -32
				op.Handler(c, op.R1, op.R2)
				assert.Equal(t, types.Addr(0x0FB1), c.Reg.R16(int(op.R1))) // 4017
				assert.Equal(t, false, c.Reg.isSet(flagN))
				assert.Equal(t, false, c.Reg.isSet(flagH))
				assert.Equal(t, true, c.Reg.isSet(flagC))
				assert.Equal(t, false, c.Reg.isSet(flagZ))
			})
		})
	}
}

func TestOpCode_adc(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
		r1     int
		r2     int
	}

	tests := []struct {
		name string
		args args
	}{
		{name: "ADC A, B", args: args{0x88, A, B}},
		{name: "ADC A, C", args: args{0x89, A, C}},
		{name: "ADC A, D", args: args{0x8A, A, D}},
		{name: "ADC A, E", args: args{0x8B, A, E}},
		{name: "ADC A, H", args: args{0x8C, A, H}},
		{name: "ADC A, L", args: args{0x8D, A, L}},
		{name: "ADC A,(HL)", args: args{0x8E, A, HL}},
		{name: "ADC A, A", args: args{0x8F, A, A}},
		{name: "ADC A,d8", args: args{0xCE, A, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := opCodes[tt.args.opcode]

			assert.Equal(t, tt.args.r1, op.R1)
			assert.Equal(t, tt.args.r2, op.R2)

			if strings.Contains(op.Mnemonic, "ADC A, A") {
				t.Skip()
			}
			t.Run("when no carry", func(t *testing.T) {
				t.Run("and carry = 1", func(t *testing.T) {
					c.Reg.R[tt.args.r1] = 0xE1

					setval := byte(0x0D)
					if strings.Contains(op.Mnemonic, "HL") {
						c.Bus.WriteByte(c.Reg.R16(HL), setval)
					} else if strings.Contains(op.Mnemonic, "d8") {
						c.Bus.WriteByte(c.Reg.PC, setval)
					} else {
						c.Reg.R[tt.args.r2] = setval
					}
					c.Reg.setF(flagC, true)

					op.Handler(c, op.R1, op.R2)
					assert.Equal(t, byte(0xEF), c.Reg.R[A])
					assert.Equal(t, false, c.Reg.isSet(flagZ))
					assert.Equal(t, false, c.Reg.isSet(flagN))
					assert.Equal(t, false, c.Reg.isSet(flagH))
					assert.Equal(t, false, c.Reg.isSet(flagC))
				})
				t.Run("and carry = 0", func(t *testing.T) {
					c.Reg.R[tt.args.r1] = 0xE1
					setval := byte(0x0E)
					if strings.Contains(op.Mnemonic, "HL") {
						c.Bus.WriteByte(c.Reg.R16(HL), setval)
					} else if strings.Contains(op.Mnemonic, "d8") {
						c.Bus.WriteByte(c.Reg.PC, setval)
					} else {
						c.Reg.R[tt.args.r2] = setval
					}
					c.Reg.setF(flagC, false)

					op.Handler(c, op.R1, op.R2)
					assert.Equal(t, byte(0xEF), c.Reg.R[A])
					assert.Equal(t, false, c.Reg.isSet(flagZ))
					assert.Equal(t, false, c.Reg.isSet(flagN))
					assert.Equal(t, false, c.Reg.isSet(flagH))
					assert.Equal(t, false, c.Reg.isSet(flagC))
				})
			})
			t.Run("when Harf carry", func(t *testing.T) {
				t.Run("and carry = 1", func(t *testing.T) {
					c.Reg.R[tt.args.r1] = 0xE1
					setval := byte(0x0F)
					if strings.Contains(op.Mnemonic, "HL") {
						c.Bus.WriteByte(c.Reg.R16(HL), setval)
					} else if strings.Contains(op.Mnemonic, "d8") {
						c.Bus.WriteByte(c.Reg.PC, setval)
					} else {
						c.Reg.R[tt.args.r2] = setval
					}
					c.Reg.setF(flagC, true)

					op.Handler(c, op.R1, op.R2)
					assert.Equal(t, byte(0xF1), c.Reg.R[A])
					assert.Equal(t, false, c.Reg.isSet(flagZ))
					assert.Equal(t, false, c.Reg.isSet(flagN))
					assert.Equal(t, true, c.Reg.isSet(flagH))
					assert.Equal(t, false, c.Reg.isSet(flagC))
				})
				t.Run("and carry = 0", func(t *testing.T) {
					c.Reg.R[tt.args.r1] = 0xE1
					setval := byte(0x0F)
					if strings.Contains(op.Mnemonic, "HL") {
						c.Bus.WriteByte(c.Reg.R16(HL), setval)
					} else if strings.Contains(op.Mnemonic, "d8") {
						c.Bus.WriteByte(c.Reg.PC, setval)
					} else {
						c.Reg.R[tt.args.r2] = setval
					}
					c.Reg.setF(flagC, false)

					op.Handler(c, op.R1, op.R2)
					assert.Equal(t, byte(0xF0), c.Reg.R[A])
					assert.Equal(t, false, c.Reg.isSet(flagZ))
					assert.Equal(t, false, c.Reg.isSet(flagN))
					assert.Equal(t, true, c.Reg.isSet(flagH))
					assert.Equal(t, false, c.Reg.isSet(flagC))
				})
			})
			t.Run("when carry and zero", func(t *testing.T) {
				t.Run("and carry = 1", func(t *testing.T) {
					c.Reg.R[tt.args.r1] = 0xF1
					setval := byte(0x0E)
					if strings.Contains(op.Mnemonic, "HL") {
						c.Bus.WriteByte(c.Reg.R16(HL), setval)
					} else if strings.Contains(op.Mnemonic, "d8") {
						c.Bus.WriteByte(c.Reg.PC, setval)
					} else {
						c.Reg.R[tt.args.r2] = setval
					}
					c.Reg.setF(flagC, true)

					op.Handler(c, op.R1, op.R2)
					assert.Equal(t, byte(0x00), c.Reg.R[A])
					assert.Equal(t, true, c.Reg.isSet(flagZ))
					assert.Equal(t, false, c.Reg.isSet(flagN))
					assert.Equal(t, true, c.Reg.isSet(flagH))
					assert.Equal(t, true, c.Reg.isSet(flagC))
				})
				t.Run("when and carry = 0", func(t *testing.T) {
					c.Reg.R[tt.args.r1] = 0xE1
					setval := byte(0x1F)
					if strings.Contains(op.Mnemonic, "HL") {
						c.Bus.WriteByte(c.Reg.R16(HL), setval)
					} else if strings.Contains(op.Mnemonic, "d8") {
						c.Bus.WriteByte(c.Reg.PC, setval)
					} else {
						c.Reg.R[tt.args.r2] = setval
					}
					c.Reg.setF(flagC, false)

					op.Handler(c, op.R1, op.R2)
					assert.Equal(t, byte(0x00), c.Reg.R[A])
					assert.Equal(t, true, c.Reg.isSet(flagZ))
					assert.Equal(t, false, c.Reg.isSet(flagN))
					assert.Equal(t, true, c.Reg.isSet(flagH))
					assert.Equal(t, true, c.Reg.isSet(flagC))
				})
			})
		})
	}
}

func TestOpCode_sbc(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
		r1     int
		r2     int
	}

	tests := []struct {
		name string
		args args
	}{
		{name: "SBC A, B", args: args{0x98, A, B}},
		{name: "SBC A, C", args: args{0x99, A, C}},
		{name: "SBC A, D", args: args{0x9A, A, D}},
		{name: "SBC A, E", args: args{0x9B, A, E}},
		{name: "SBC A, H", args: args{0x9C, A, H}},
		{name: "SBC A, L", args: args{0x9D, A, L}},
		{name: "SBC A,(HL)", args: args{0x9E, A, HL}},
		{name: "SBC A, A", args: args{0x9F, A, A}},
		{name: "SBC A,d8", args: args{0xDE, A, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := opCodes[tt.args.opcode]

			assert.Equal(t, tt.args.r1, op.R1)
			assert.Equal(t, tt.args.r2, op.R2)

			if strings.Contains(op.Mnemonic, "SBC A, A") {
				t.Skip()
			}
			t.Run("when no carry", func(t *testing.T) {
				t.Run("and carry = 1", func(t *testing.T) {
					c.Reg.R[tt.args.r1] = 0xE2

					setval := byte(0x01)
					if strings.Contains(op.Mnemonic, "HL") {
						c.Bus.WriteByte(c.Reg.R16(HL), setval)
					} else if strings.Contains(op.Mnemonic, "d8") {
						c.Bus.WriteByte(c.Reg.PC, setval)
					} else {
						c.Reg.R[tt.args.r2] = setval
					}
					c.Reg.setF(flagC, true)

					op.Handler(c, op.R1, op.R2)
					assert.Equal(t, byte(0xE0), c.Reg.R[A])
					assert.Equal(t, false, c.Reg.isSet(flagZ))
					assert.Equal(t, true, c.Reg.isSet(flagN))
					assert.Equal(t, false, c.Reg.isSet(flagH))
					assert.Equal(t, false, c.Reg.isSet(flagC))
				})
				t.Run("and carry = 0", func(t *testing.T) {
					c.Reg.R[tt.args.r1] = 0xE2
					setval := byte(0x01)
					if strings.Contains(op.Mnemonic, "HL") {
						c.Bus.WriteByte(c.Reg.R16(HL), setval)
					} else if strings.Contains(op.Mnemonic, "d8") {
						c.Bus.WriteByte(c.Reg.PC, setval)
					} else {
						c.Reg.R[tt.args.r2] = setval
					}
					c.Reg.setF(flagC, false)

					op.Handler(c, op.R1, op.R2)
					assert.Equal(t, byte(0xE1), c.Reg.R[A])
					assert.Equal(t, false, c.Reg.isSet(flagZ))
					assert.Equal(t, true, c.Reg.isSet(flagN))
					assert.Equal(t, false, c.Reg.isSet(flagH))
					assert.Equal(t, false, c.Reg.isSet(flagC))
				})
			})
			t.Run("when Harf carry", func(t *testing.T) {
				t.Run("and carry = 1", func(t *testing.T) {
					c.Reg.R[tt.args.r1] = 0xE1
					setval := byte(0x0F)
					if strings.Contains(op.Mnemonic, "HL") {
						c.Bus.WriteByte(c.Reg.R16(HL), setval)
					} else if strings.Contains(op.Mnemonic, "d8") {
						c.Bus.WriteByte(c.Reg.PC, setval)
					} else {
						c.Reg.R[tt.args.r2] = setval
					}
					c.Reg.setF(flagC, true)

					op.Handler(c, op.R1, op.R2)
					assert.Equal(t, byte(0xD1), c.Reg.R[A])
					assert.Equal(t, false, c.Reg.isSet(flagZ))
					assert.Equal(t, true, c.Reg.isSet(flagN))
					assert.Equal(t, true, c.Reg.isSet(flagH))
					assert.Equal(t, false, c.Reg.isSet(flagC))
				})
				t.Run("and carry = 0", func(t *testing.T) {
					c.Reg.R[tt.args.r1] = 0xE1
					setval := byte(0x0F)
					if strings.Contains(op.Mnemonic, "HL") {
						c.Bus.WriteByte(c.Reg.R16(HL), setval)
					} else if strings.Contains(op.Mnemonic, "d8") {
						c.Bus.WriteByte(c.Reg.PC, setval)
					} else {
						c.Reg.R[tt.args.r2] = setval
					}
					c.Reg.setF(flagC, false)

					op.Handler(c, op.R1, op.R2)
					assert.Equal(t, byte(0xD2), c.Reg.R[A])
					assert.Equal(t, false, c.Reg.isSet(flagZ))
					assert.Equal(t, true, c.Reg.isSet(flagN))
					assert.Equal(t, true, c.Reg.isSet(flagH))
					assert.Equal(t, false, c.Reg.isSet(flagC))
				})
			})
			t.Run("when carry and zero", func(t *testing.T) {
				t.Run("and carry = 1", func(t *testing.T) {
					c.Reg.R[tt.args.r1] = 0xF1
					setval := byte(0xF0)
					if strings.Contains(op.Mnemonic, "HL") {
						c.Bus.WriteByte(c.Reg.R16(HL), setval)
					} else if strings.Contains(op.Mnemonic, "d8") {
						c.Bus.WriteByte(c.Reg.PC, setval)
					} else {
						c.Reg.R[tt.args.r2] = setval
					}
					c.Reg.setF(flagC, true)

					op.Handler(c, op.R1, op.R2)
					assert.Equal(t, byte(0x00), c.Reg.R[A])
					assert.Equal(t, true, c.Reg.isSet(flagZ))
					assert.Equal(t, true, c.Reg.isSet(flagN))
					assert.Equal(t, false, c.Reg.isSet(flagH))
					assert.Equal(t, false, c.Reg.isSet(flagC))
				})
				t.Run("when and carry = 0", func(t *testing.T) {
					c.Reg.R[tt.args.r1] = 0xE1
					setval := byte(0xE2)
					if strings.Contains(op.Mnemonic, "HL") {
						c.Bus.WriteByte(c.Reg.R16(HL), setval)
					} else if strings.Contains(op.Mnemonic, "d8") {
						c.Bus.WriteByte(c.Reg.PC, setval)
					} else {
						c.Reg.R[tt.args.r2] = setval
					}
					c.Reg.setF(flagC, false)

					op.Handler(c, op.R1, op.R2)
					assert.Equal(t, byte(0xFF), c.Reg.R[A])
					assert.Equal(t, false, c.Reg.isSet(flagZ))
					assert.Equal(t, true, c.Reg.isSet(flagN))
					assert.Equal(t, true, c.Reg.isSet(flagH))
					assert.Equal(t, true, c.Reg.isSet(flagC))
				})
			})
		})
	}
}

func TestOpCode_sub(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
		r1     int
	}

	tests := []struct {
		name string
		args args
	}{
		{name: "SUB B", args: args{0x90, B}},
		{name: "SUB C", args: args{0x91, C}},
		{name: "SUB D", args: args{0x92, D}},
		{name: "SUB E", args: args{0x93, E}},
		{name: "SUB H", args: args{0x94, H}},
		{name: "SUB L", args: args{0x95, L}},
		{name: "SUB (HL)", args: args{0x96, HL}},
		{name: "SUB A", args: args{0x97, A}},
		{name: "SUB d8", args: args{0xD6, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := opCodes[tt.args.opcode]

			assert.Equal(t, tt.args.r1, op.R1)

			if strings.Contains(op.Mnemonic, "A") {
				t.Skip()
			}
			t.Run("when no carry", func(t *testing.T) {
				c.Reg.R[A] = 0xE1
				if strings.Contains(op.Mnemonic, "HL") {
					c.Reg.setR16(int(HL), 0x0101)
					c.Bus.WriteByte(c.Reg.R16(HL), 0x01)
				} else if strings.Contains(op.Mnemonic, "d8") {
					c.Bus.WriteByte(c.Reg.PC, 0x01)
				} else {
					c.Reg.R[tt.args.r1] = 0x01
				}

				op.Handler(c, op.R1, op.R2)
				assert.Equal(t, byte(0xE0), c.Reg.R[A])
				assert.Equal(t, false, c.Reg.isSet(flagZ))
				assert.Equal(t, true, c.Reg.isSet(flagN))
				assert.Equal(t, false, c.Reg.isSet(flagH))
				assert.Equal(t, false, c.Reg.isSet(flagC))
			})
			t.Run("when Harf carry", func(t *testing.T) {
				c.Reg.R[A] = 0xE1
				if strings.Contains(op.Mnemonic, "HL") {
					c.Reg.setR16(int(HL), 0x0202)
					c.Bus.WriteByte(c.Reg.R16(HL), 0x02)
				} else if strings.Contains(op.Mnemonic, "d8") {
					c.Bus.WriteByte(c.Reg.PC, 0x02)
				} else {
					c.Reg.R[tt.args.r1] = 0x02
				}

				op.Handler(c, op.R1, op.R2)
				assert.Equal(t, byte(0xDF), c.Reg.R[A])
				assert.Equal(t, false, c.Reg.isSet(flagZ))
				assert.Equal(t, true, c.Reg.isSet(flagN))
				assert.Equal(t, true, c.Reg.isSet(flagH))
				assert.Equal(t, false, c.Reg.isSet(flagC))
			})
			t.Run("when zero", func(t *testing.T) {
				c.Reg.R[A] = 0xE1
				if strings.Contains(op.Mnemonic, "HL") {
					c.Reg.setR16(int(HL), 0xE1E1)
					c.Bus.WriteByte(c.Reg.R16(HL), 0xE1)
				} else if strings.Contains(op.Mnemonic, "d8") {
					c.Bus.WriteByte(c.Reg.PC, 0xE1)
				} else {
					c.Reg.R[tt.args.r1] = 0xE1
				}

				op.Handler(c, op.R1, op.R2)
				assert.Equal(t, byte(0x00), c.Reg.R[A])
				assert.Equal(t, true, c.Reg.isSet(flagZ))
				assert.Equal(t, true, c.Reg.isSet(flagN))
				assert.Equal(t, false, c.Reg.isSet(flagH))
				assert.Equal(t, false, c.Reg.isSet(flagC))
			})
			t.Run("when carry", func(t *testing.T) {
				c.Reg.R[A] = 0xE1
				if strings.Contains(op.Mnemonic, "HL") {
					c.Reg.setR16(int(HL), 0xE2E2)
					c.Bus.WriteByte(c.Reg.R16(HL), 0xE2)
				} else if strings.Contains(op.Mnemonic, "d8") {
					c.Bus.WriteByte(c.Reg.PC, 0xE2)
				} else {
					c.Reg.R[tt.args.r1] = 0xE2
				}

				op.Handler(c, op.R1, op.R2)
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
func TestOpCode_and(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
		r1     int
	}

	tests := []struct {
		name string
		args args
	}{
		{name: "AND B", args: args{0xA0, B}},
		{name: "AND C", args: args{0xA1, C}},
		{name: "AND D", args: args{0xA2, D}},
		{name: "AND E", args: args{0xA3, E}},
		{name: "AND H", args: args{0xA4, H}},
		{name: "AND L", args: args{0xA5, L}},
		{name: "AND (HL)", args: args{0xA6, HL}},
		{name: "AND A", args: args{0xA7, A}},
		{name: "AND d8", args: args{0xE6, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := opCodes[tt.args.opcode]
			assert.Equal(t, tt.args.r1, op.R1)

			if strings.Contains(op.Mnemonic, "AND A") {
				t.Skip()
			}
			t.Run("when oposite", func(t *testing.T) {
				c.Reg.R[A] = 0b11110000
				val := byte(0b00001111)
				if strings.Contains(op.Mnemonic, "HL") {
					c.Bus.WriteByte(c.Reg.R16(int(HL)), val)
				} else if strings.Contains(op.Mnemonic, "d8") {
					c.Bus.WriteByte(c.Reg.PC, val)
				} else {
					c.Reg.R[tt.args.r1] = val
				}
				op.Handler(c, op.R1, op.R2)
				assert.Equal(t, byte(0b00000000), c.Reg.R[A])
			})
			t.Run("when equal", func(t *testing.T) {
				c.Reg.R[A] = 0b11110000
				val := byte(0b11110000)
				if strings.Contains(op.Mnemonic, "HL") {
					c.Bus.WriteByte(c.Reg.R16(int(HL)), val)
				} else if strings.Contains(op.Mnemonic, "d8") {
					c.Bus.WriteByte(c.Reg.PC, val)
				} else {
					c.Reg.R[tt.args.r1] = val
				}

				op.Handler(c, op.R1, op.R2)
				assert.Equal(t, byte(0b11110000), c.Reg.R[A])
			})
			t.Run("when other", func(t *testing.T) {
				c.Reg.R[A] = 0b11110000
				val := byte(0b10100000)
				if strings.Contains(op.Mnemonic, "HL") {
					c.Bus.WriteByte(c.Reg.R16(int(HL)), val)
				} else if strings.Contains(op.Mnemonic, "d8") {
					c.Bus.WriteByte(c.Reg.PC, val)
				} else {
					c.Reg.R[tt.args.r1] = val
				}

				op.Handler(c, op.R1, op.R2)
				assert.Equal(t, byte(0b10100000), c.Reg.R[A])
			})
		})
	}
}

// test orr orHL ord8
func TestOpCode_or(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
		r1     int
	}

	tests := []struct {
		name string
		args args
	}{
		{name: "OR B", args: args{0xB0, B}},
		{name: "OR C", args: args{0xB1, C}},
		{name: "OR D", args: args{0xB2, D}},
		{name: "OR E", args: args{0xB3, E}},
		{name: "OR H", args: args{0xB4, H}},
		{name: "OR L", args: args{0xB5, L}},
		{name: "OR (HL)", args: args{0xB6, HL}},
		{name: "OR A", args: args{0xB7, A}},
		{name: "OR d8", args: args{0xF6, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := opCodes[tt.args.opcode]
			assert.Equal(t, tt.args.r1, op.R1)

			if strings.Contains(op.Mnemonic, "OR A") {
				t.Skip()
			}
			t.Run("when oposite", func(t *testing.T) {
				c.Reg.R[A] = 0b11110000
				val := byte(0b00001111)
				if strings.Contains(op.Mnemonic, "HL") {
					c.Bus.WriteByte(c.Reg.R16(int(HL)), val)
				} else if strings.Contains(op.Mnemonic, "d8") {
					c.Bus.WriteByte(c.Reg.PC, val)
				} else {
					c.Reg.R[tt.args.r1] = val
				}

				op.Handler(c, op.R1, op.R2)
				assert.Equal(t, byte(0b11111111), c.Reg.R[A])
				assert.Equal(t, false, c.Reg.isSet(flagZ))
				assert.Equal(t, false, c.Reg.isSet(flagN))
				assert.Equal(t, false, c.Reg.isSet(flagH))
				assert.Equal(t, false, c.Reg.isSet(flagC))
			})
			t.Run("when equal", func(t *testing.T) {
				c.Reg.R[A] = 0b11110000
				val := byte(0b11110000)
				if strings.Contains(op.Mnemonic, "HL") {
					c.Bus.WriteByte(c.Reg.R16(int(HL)), val)
				} else if strings.Contains(op.Mnemonic, "d8") {
					c.Bus.WriteByte(c.Reg.PC, val)
				} else {
					c.Reg.R[tt.args.r1] = val
				}

				op.Handler(c, op.R1, op.R2)
				assert.Equal(t, byte(0b11110000), c.Reg.R[A])
				assert.Equal(t, false, c.Reg.isSet(flagZ))
				assert.Equal(t, false, c.Reg.isSet(flagN))
				assert.Equal(t, false, c.Reg.isSet(flagH))
				assert.Equal(t, false, c.Reg.isSet(flagC))
			})
			t.Run("when other", func(t *testing.T) {
				c.Reg.R[A] = 0b11110000
				val := byte(0b00000101)
				if strings.Contains(op.Mnemonic, "HL") {
					c.Bus.WriteByte(c.Reg.R16(int(HL)), val)
				} else if strings.Contains(op.Mnemonic, "d8") {
					c.Bus.WriteByte(c.Reg.PC, val)
				} else {
					c.Reg.R[tt.args.r1] = val
				}

				op.Handler(c, op.R1, op.R2)
				assert.Equal(t, byte(0b11110101), c.Reg.R[A])
				assert.Equal(t, false, c.Reg.isSet(flagZ))
				assert.Equal(t, false, c.Reg.isSet(flagN))
				assert.Equal(t, false, c.Reg.isSet(flagH))
				assert.Equal(t, false, c.Reg.isSet(flagC))
			})
			t.Run("when zero", func(t *testing.T) {
				c.Reg.R[A] = 0b00000000
				val := byte(0b00000000)
				if strings.Contains(op.Mnemonic, "HL") {
					c.Bus.WriteByte(c.Reg.R16(int(HL)), val)
				} else if strings.Contains(op.Mnemonic, "d8") {
					c.Bus.WriteByte(c.Reg.PC, val)
				} else {
					c.Reg.R[tt.args.r1] = val
				}

				op.Handler(c, op.R1, op.R2)
				assert.Equal(t, byte(0b00000000), c.Reg.R[A])
				assert.Equal(t, true, c.Reg.isSet(flagZ))
				assert.Equal(t, false, c.Reg.isSet(flagN))
				assert.Equal(t, false, c.Reg.isSet(flagH))
				assert.Equal(t, false, c.Reg.isSet(flagC))
			})
		})
	}
}

// test xorr xorHL xord8
func TestOpCode_xor(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
		r1     int
	}

	tests := []struct {
		name string
		args args
	}{
		{name: "XOR B", args: args{0xA8, B}},
		{name: "XOR C", args: args{0xA9, C}},
		{name: "XOR D", args: args{0xAA, D}},
		{name: "XOR E", args: args{0xAB, E}},
		{name: "XOR H", args: args{0xAC, H}},
		{name: "XOR L", args: args{0xAD, L}},
		{name: "XOR (HL)", args: args{0xAE, HL}},
		{name: "XOR A", args: args{0xAF, A}},
		{name: "XOR d8", args: args{0xEE, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := opCodes[tt.args.opcode]
			assert.Equal(t, tt.args.r1, op.R1)

			if strings.Contains(op.Mnemonic, "XOR A") {
				t.Skip()
			}
			t.Run("when oposite", func(t *testing.T) {
				c.Reg.R[A] = 0b11110000
				val := byte(0b00001111)
				if strings.Contains(op.Mnemonic, "HL") {
					c.Bus.WriteByte(c.Reg.R16(int(HL)), val)
				} else if strings.Contains(op.Mnemonic, "d8") {
					c.Bus.WriteByte(c.Reg.PC, val)
				} else {
					c.Reg.R[tt.args.r1] = val
				}

				op.Handler(c, op.R1, op.R2)
				assert.Equal(t, byte(0b11111111), c.Reg.R[A])
				assert.Equal(t, false, c.Reg.isSet(flagZ))
				assert.Equal(t, false, c.Reg.isSet(flagN))
				assert.Equal(t, false, c.Reg.isSet(flagH))
				assert.Equal(t, false, c.Reg.isSet(flagC))
			})
			t.Run("when equal", func(t *testing.T) {
				c.Reg.R[A] = 0b11110000
				val := byte(0b11110000)
				if strings.Contains(op.Mnemonic, "HL") {
					c.Bus.WriteByte(c.Reg.R16(int(HL)), val)
				} else if strings.Contains(op.Mnemonic, "d8") {
					c.Bus.WriteByte(c.Reg.PC, val)
				} else {
					c.Reg.R[tt.args.r1] = val
				}

				op.Handler(c, op.R1, op.R2)
				assert.Equal(t, byte(0b00000000), c.Reg.R[A])
				assert.Equal(t, true, c.Reg.isSet(flagZ))
				assert.Equal(t, false, c.Reg.isSet(flagN))
				assert.Equal(t, false, c.Reg.isSet(flagH))
				assert.Equal(t, false, c.Reg.isSet(flagC))
			})
			t.Run("when other", func(t *testing.T) {
				c.Reg.R[A] = 0b11110000
				val := byte(0b01010101)
				if strings.Contains(op.Mnemonic, "HL") {
					c.Bus.WriteByte(c.Reg.R16(int(HL)), val)
				} else if strings.Contains(op.Mnemonic, "d8") {
					c.Bus.WriteByte(c.Reg.PC, val)
				} else {
					c.Reg.R[tt.args.r1] = val
				}

				op.Handler(c, op.R1, op.R2)
				assert.Equal(t, byte(0b10100101), c.Reg.R[A])
				assert.Equal(t, false, c.Reg.isSet(flagZ))
				assert.Equal(t, false, c.Reg.isSet(flagN))
				assert.Equal(t, false, c.Reg.isSet(flagH))
				assert.Equal(t, false, c.Reg.isSet(flagC))
			})
		})
	}
}

func TestOpCode_cp(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
		r1     int
	}

	tests := []struct {
		name string
		args args
	}{
		{name: "CP B", args: args{0xB8, B}},
		{name: "CP C", args: args{0xB9, C}},
		{name: "CP D", args: args{0xBA, D}},
		{name: "CP E", args: args{0xBB, E}},
		{name: "CP H", args: args{0xBC, H}},
		{name: "CP L", args: args{0xBD, L}},
		{name: "CP (HL)", args: args{0xBE, HL}},
		{name: "CP A", args: args{0xBF, A}},
		{name: "CP d8", args: args{0xFE, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := opCodes[tt.args.opcode]
			// prepare
			c.Reg.R[A] = 0x12

			assert.Equal(t, tt.args.r1, op.R1)

			if strings.Contains(op.Mnemonic, "CP A") {
				t.Skip()
			}

			t.Run("when equal", func(t *testing.T) {
				val := byte(0x12)
				if strings.Contains(op.Mnemonic, "HL") {
					c.Reg.setR16(int(HL), util.Byte2Addr(val, val))
					c.Bus.WriteByte(c.Reg.R16(int(HL)), val)
				} else if strings.Contains(op.Mnemonic, "d8") {
					c.Bus.WriteByte(c.Reg.PC, val)
				} else {
					c.Reg.R[tt.args.r1] = val
				}

				op.Handler(c, op.R1, op.R2)
				assert.Equal(t, true, c.Reg.isSet(flagZ))
				assert.Equal(t, true, c.Reg.isSet(flagN))
				assert.Equal(t, false, c.Reg.isSet(flagH))
				assert.Equal(t, false, c.Reg.isSet(flagC))
			})
			t.Run("when greater than A", func(t *testing.T) {
				val := byte(0x13)
				if strings.Contains(op.Mnemonic, "HL") {
					c.Reg.setR16(int(HL), util.Byte2Addr(val, val))
					c.Bus.WriteByte(c.Reg.R16(int(HL)), val)
				} else if strings.Contains(op.Mnemonic, "d8") {
					c.Bus.WriteByte(c.Reg.PC, val)
				} else {
					c.Reg.R[tt.args.r1] = val
				}

				op.Handler(c, op.R1, op.R2)
				assert.Equal(t, false, c.Reg.isSet(flagZ))
				assert.Equal(t, true, c.Reg.isSet(flagN))
				assert.Equal(t, true, c.Reg.isSet(flagH))
				assert.Equal(t, true, c.Reg.isSet(flagC))
			})
			t.Run("when less than A with borrow", func(t *testing.T) {
				val := byte(0x03)
				if strings.Contains(op.Mnemonic, "HL") {
					c.Reg.setR16(int(HL), util.Byte2Addr(val, val))
					c.Bus.WriteByte(c.Reg.R16(int(HL)), val)
				} else if strings.Contains(op.Mnemonic, "d8") {
					c.Bus.WriteByte(c.Reg.PC, val)
				} else {
					c.Reg.R[tt.args.r1] = val
				}

				op.Handler(c, op.R1, op.R2)
				assert.Equal(t, false, c.Reg.isSet(flagZ))
				assert.Equal(t, true, c.Reg.isSet(flagN))
				assert.Equal(t, true, c.Reg.isSet(flagH))
				assert.Equal(t, false, c.Reg.isSet(flagC))
			})
			t.Run("when less than A no borrow", func(t *testing.T) {
				val := byte(0x02)
				if strings.Contains(op.Mnemonic, "HL") {
					c.Reg.setR16(int(HL), util.Byte2Addr(val, val))
					c.Bus.WriteByte(c.Reg.R16(int(HL)), val)
				} else if strings.Contains(op.Mnemonic, "d8") {
					c.Bus.WriteByte(c.Reg.PC, val)
				} else {
					c.Reg.R[tt.args.r1] = val
				}

				op.Handler(c, op.R1, op.R2)
				assert.Equal(t, false, c.Reg.isSet(flagZ))
				assert.Equal(t, true, c.Reg.isSet(flagN))
				assert.Equal(t, false, c.Reg.isSet(flagH))
				assert.Equal(t, false, c.Reg.isSet(flagC))
			})
		})
	}
}

// -----ret----
func TestOpCode_ret(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
		flag   int
	}

	tests := []struct {
		name string
		args args
	}{
		{name: "RET", args: args{0xC9, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := opCodes[tt.args.opcode]

			assert.Equal(t, tt.args.flag, op.R1)

			c.Reg.SP = 0xFFFC

			// lower
			c.Bus.WriteByte(c.Reg.SP, 0x34)
			// upper
			c.Bus.WriteByte(c.Reg.SP+1, 0x12)

			op.Handler(c, op.R1, op.R2)

			assert.Equal(t, types.Addr(0xFFFE), c.Reg.SP)
			assert.Equal(t, types.Addr(0x1234), c.Reg.PC)
		})
	}
}

func TestOpCode_retf(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
		flag   int
	}

	tests := []struct {
		name string
		args args
	}{
		{name: "RET NZ", args: args{0xC0, flagZ}},
		{name: "RET Z", args: args{0xC8, flagZ}},
		{name: "RET NC", args: args{0xD0, flagC}},
		{name: "RET Z", args: args{0xD8, flagC}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := opCodes[tt.args.opcode]

			assert.Equal(t, tt.args.flag, op.R1)

			t.Run("when flag = 0", func(t *testing.T) {
				c.Reg.PC = 0x5678
				c.Reg.SP = 0xFFFC
				c.Reg.setF(tt.args.flag, false)
				// lower
				c.Bus.WriteByte(c.Reg.SP, 0x34)
				// upper
				c.Bus.WriteByte(c.Reg.SP+1, 0x12)
				op.Handler(c, op.R1, op.R2)
				if strings.Contains(op.Mnemonic, "N") {
					assert.Equal(t, types.Addr(0xFFFE), c.Reg.SP)
					assert.Equal(t, types.Addr(0x1234), c.Reg.PC)
				} else {
					assert.Equal(t, types.Addr(0xFFFC), c.Reg.SP)
					assert.Equal(t, types.Addr(0x5678), c.Reg.PC)
				}
			})
			t.Run("when flag = 1", func(t *testing.T) {
				c.Reg.PC = 0x5678
				c.Reg.SP = 0xFFFC
				c.Reg.setF(tt.args.flag, true)
				// lower
				c.Bus.WriteByte(c.Reg.SP, 0x34)
				// upper
				c.Bus.WriteByte(c.Reg.SP+1, 0x12)
				op.Handler(c, op.R1, op.R2)
				if strings.Contains(op.Mnemonic, "N") {
					assert.Equal(t, types.Addr(0xFFFC), c.Reg.SP)
					assert.Equal(t, types.Addr(0x5678), c.Reg.PC)
				} else {
					assert.Equal(t, types.Addr(0xFFFE), c.Reg.SP)
					assert.Equal(t, types.Addr(0x1234), c.Reg.PC)
				}
			})
		})
	}
}

// -----jp-----
func TestOpCode_jpa16(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
		addr   types.Addr
	}

	tests := []struct {
		name string
		args args
	}{
		{name: "JP a16", args: args{0xC3, 0x1234}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := opCodes[tt.args.opcode]
			want := tt.args.addr
			c.Bus.WriteByte(c.Reg.PC, 0x34)
			c.Bus.WriteByte(c.Reg.PC+1, 0x12)
			op.Handler(c, op.R1, op.R2)

			assert.Equal(t, want, c.Reg.PC)
		})
	}
}

func TestOpCode_jpfa16(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
		flag   int
	}

	tests := []struct {
		name string
		args args
	}{
		{name: "JP Z", args: args{0xCA, flagZ}},
		{name: "JP C", args: args{0xDA, flagC}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := opCodes[tt.args.opcode]
			t.Run("when flag = 0", func(t *testing.T) {
				c.Reg.PC = 0x0100
				want := types.Addr(0x0100)
				c.Bus.WriteByte(c.Reg.PC, 0x34)
				c.Bus.WriteByte(c.Reg.PC+1, 0x12)

				c.Reg.setF(tt.args.flag, false)
				op.Handler(c, op.R1, op.R2)

				assert.Equal(t, want, c.Reg.PC)
			})
			t.Run("when flag = 1", func(t *testing.T) {
				c.Reg.PC = 0x0100
				want := types.Addr(0x1234)
				c.Bus.WriteByte(c.Reg.PC, 0x34)
				c.Bus.WriteByte(c.Reg.PC+1, 0x12)

				c.Reg.setF(tt.args.flag, true)
				op.Handler(c, op.R1, op.R2)

				assert.Equal(t, want, c.Reg.PC)
			})
		})
	}
}

func TestOpCode_jpnfa16(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
		flag   int
	}

	tests := []struct {
		name string
		args args
	}{
		{name: "JP NZ", args: args{0xC2, flagZ}},
		{name: "JP NC", args: args{0xD2, flagC}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := opCodes[tt.args.opcode]
			t.Run("when flag = 0", func(t *testing.T) {
				c.Reg.PC = 0x0100
				want := types.Addr(0x1234)
				c.Bus.WriteByte(c.Reg.PC, 0x34)
				c.Bus.WriteByte(c.Reg.PC+1, 0x12)

				c.Reg.setF(tt.args.flag, false)
				op.Handler(c, op.R1, op.R2)

				assert.Equal(t, want, c.Reg.PC)
			})
			t.Run("when flag = 1", func(t *testing.T) {
				c.Reg.PC = 0x0100
				want := types.Addr(0x0100)
				c.Bus.WriteByte(c.Reg.PC, 0x34)
				c.Bus.WriteByte(c.Reg.PC+1, 0x12)

				c.Reg.setF(tt.args.flag, true)
				op.Handler(c, op.R1, op.R2)

				assert.Equal(t, want, c.Reg.PC)
			})
		})
	}
}

func TestOpCode_jpm16(t *testing.T) {
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
		{name: "JP (HL)", args: args{0xE9, HL, 0x1234}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := opCodes[tt.args.opcode]
			want := tt.args.addr
			c.Reg.setR16(tt.args.r16, 0x1234)

			op.Handler(c, op.R1, op.R2)

			assert.Equal(t, want, c.Reg.PC)
		})
	}
}

// -----jr-----

func TestOpCode_jrr8(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
	}

	tests := []struct {
		name string
		args args
	}{
		{name: "JR r8", args: args{0x18}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := opCodes[tt.args.opcode]
			t.Run("when value is positive", func(t *testing.T) {
				c.Reg.PC = 0x0100
				c.Bus.WriteByte(c.Reg.PC, 0x10)
				op.Handler(c, op.R1, op.R2)

				assert.Equal(t, types.Addr(0x0111), c.Reg.PC)
			})
			t.Run("when value is negative", func(t *testing.T) {
				c.Reg.PC = 0x0100
				c.Bus.WriteByte(c.Reg.PC, 0xFE) // -2
				op.Handler(c, op.R1, op.R2)

				assert.Equal(t, types.Addr(0x00FF), c.Reg.PC)
			})
		})
	}
}

func TestOpCode_jrfr8(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
		flag   int
	}

	tests := []struct {
		name string
		args args
	}{
		{name: "JR Z", args: args{0x28, flagZ}},
		{name: "JR C", args: args{0x38, flagC}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := opCodes[tt.args.opcode]
			t.Run("when flag = 1", func(t *testing.T) {
				c.Reg.setF(tt.args.flag, true)
				t.Run("when value is positive", func(t *testing.T) {
					c.Reg.PC = 0x0100
					c.Bus.WriteByte(c.Reg.PC, 0x10)
					op.Handler(c, op.R1, op.R2)

					assert.Equal(t, types.Addr(0x0111), c.Reg.PC)
				})
				t.Run("when value is negative", func(t *testing.T) {
					c.Reg.PC = 0x0100
					c.Bus.WriteByte(c.Reg.PC, 0xFE) // -2
					op.Handler(c, op.R1, op.R2)

					assert.Equal(t, types.Addr(0x00FF), c.Reg.PC)
				})
			})
			t.Run("when flag = 0", func(t *testing.T) {
				c.Reg.setF(tt.args.flag, false)
				c.Reg.PC = 0x0100
				c.Bus.WriteByte(c.Reg.PC, 0x10)
				op.Handler(c, op.R1, op.R2)

				assert.Equal(t, types.Addr(0x0101), c.Reg.PC)
			})
		})
	}
}

func TestOpCode_jrnfr8(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
		flag   int
	}

	tests := []struct {
		name string
		args args
	}{
		{name: "JR NZ", args: args{0x20, flagZ}},
		{name: "JR NC", args: args{0x30, flagC}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := opCodes[tt.args.opcode]
			t.Run("when flag = 0", func(t *testing.T) {
				c.Reg.setF(tt.args.flag, false)
				t.Run("when value is positive", func(t *testing.T) {
					c.Reg.PC = 0x0100
					c.Bus.WriteByte(c.Reg.PC, 0x10)
					op.Handler(c, op.R1, op.R2)

					assert.Equal(t, types.Addr(0x0111), c.Reg.PC)
				})
				t.Run("when value is negative", func(t *testing.T) {
					c.Reg.PC = 0x0100
					c.Bus.WriteByte(c.Reg.PC, 0xFE) // -2
					op.Handler(c, op.R1, op.R2)

					assert.Equal(t, types.Addr(0x00FF), c.Reg.PC)
				})
			})
			t.Run("when flag =1", func(t *testing.T) {
				c.Reg.setF(tt.args.flag, true)
				c.Reg.PC = 0x0100
				c.Bus.WriteByte(c.Reg.PC, 0x10)
				op.Handler(c, op.R1, op.R2)

				assert.Equal(t, types.Addr(0x0101), c.Reg.PC)
			})
		})
	}
}

func TestOpCode_rst(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
		addr   byte
	}

	tests := []struct {
		name string
		args args
	}{
		{name: "RST 00H", args: args{0xC7, 0x00}},
		{name: "RST 08H", args: args{0xCF, 0x08}},
		{name: "RST 10H", args: args{0xD7, 0x10}},
		{name: "RST 18H", args: args{0xDF, 0x18}},
		{name: "RST 20H", args: args{0xE7, 0x20}},
		{name: "RST 28H", args: args{0xEF, 0x28}},
		{name: "RST 30H", args: args{0xF7, 0x30}},
		{name: "RST 38H", args: args{0xFF, 0x38}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := opCodes[tt.args.opcode]

			c.Reg.PC = 0x1234

			op.Handler(c, op.R1, op.R2)

			assert.Equal(t, int(tt.args.addr), op.R1)
			assert.Equal(t, byte(0x34), c.Bus.ReadByte(c.Reg.SP))
			assert.Equal(t, byte(0x12), c.Bus.ReadByte(c.Reg.SP+1))
			assert.Equal(t, util.Byte2Addr(0x00, tt.args.addr), c.Reg.PC)
		})
	}
}

// ----push----
func TestOpCode_push(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
		flag   int
	}

	tests := []struct {
		name string
		args args
	}{
		{name: "PUSH BC", args: args{0xC5, BC}},
		{name: "PUSH DE", args: args{0xD5, DE}},
		{name: "PUSH HL", args: args{0xE5, HL}},
		{name: "PUSH AF", args: args{0xF5, AF}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := opCodes[tt.args.opcode]

			upper := util.ExtractUpper(c.Reg.R16(int(op.R1)))
			lower := util.ExtractLower(c.Reg.R16(int(op.R1)))
			before_sp := c.Reg.SP

			op.Handler(c, op.R1, op.R2)

			assert.Equal(t, before_sp-2, c.Reg.SP)
			assert.Equal(t, lower, c.Bus.ReadByte((c.Reg.SP)))
			assert.Equal(t, upper, c.Bus.ReadByte((c.Reg.SP + 1)))
		})
	}
}

func TestOpCode_pop(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
		flag   int
	}

	tests := []struct {
		name string
		args args
	}{
		{name: "POP BC", args: args{0xC1, BC}},
		{name: "POP DE", args: args{0xD1, DE}},
		{name: "POP HL", args: args{0xE1, HL}},
		{name: "POP AF", args: args{0xF1, AF}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := opCodes[tt.args.opcode]

			c.push(0x12) // upper
			c.push(0x34) // lower

			op.Handler(c, op.R1, op.R2)
			if strings.Contains(op.Mnemonic, "AF") {
				assert.Equal(t, types.Addr(0x1230), c.Reg.R16(int(op.R1)))
			} else {
				assert.Equal(t, types.Addr(0x1234), c.Reg.R16(int(op.R1)))
			}
		})
	}
}

func TestOpCode_call(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
		flag   int
	}

	tests := []struct {
		name string
		args args
	}{
		{name: "CALL NZ,a16", args: args{0xC4, flagZ}},
		{name: "CALL Z,a16", args: args{0xCC, flagZ}},
		{name: "CALL a16", args: args{0xCD, 0}},
		{name: "CALL NC,a16", args: args{0xD4, flagC}},
		{name: "CALL C,a16", args: args{0xDC, flagC}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := opCodes[tt.args.opcode]

			t.Run("when flag = 0", func(t *testing.T) {
				c.Reg.PC = 0x5678
				c.Reg.SP = 0xFFFC
				c.Reg.setF(tt.args.flag, false)
				// lower
				c.Bus.WriteByte(c.Reg.PC, 0x34)
				// upper
				c.Bus.WriteByte(c.Reg.PC+1, 0x12)
				op.Handler(c, op.R1, op.R2)
				if strings.Contains(op.Mnemonic, "CALL a16") {
					assert.Equal(t, types.Addr(0x1234), c.Reg.PC)
					assert.Equal(t, types.Addr(0xFFFA), c.Reg.SP)
					assert.Equal(t, byte(0x7A), c.Bus.ReadByte(c.Reg.SP))
					assert.Equal(t, byte(0x56), c.Bus.ReadByte(c.Reg.SP+1))
				} else if strings.Contains(op.Mnemonic, "N") {
					assert.Equal(t, types.Addr(0x1234), c.Reg.PC)
					assert.Equal(t, types.Addr(0xFFFA), c.Reg.SP)
					assert.Equal(t, byte(0x7A), c.Bus.ReadByte(c.Reg.SP))
					assert.Equal(t, byte(0x56), c.Bus.ReadByte(c.Reg.SP+1))
				} else {
					assert.Equal(t, types.Addr(0x567A), c.Reg.PC)
					assert.Equal(t, types.Addr(0xFFFC), c.Reg.SP)
				}
			})
			t.Run("when flag = 1", func(t *testing.T) {
				c.Reg.PC = 0x5678
				c.Reg.SP = 0xFFFC
				c.Reg.setF(tt.args.flag, true)
				// lower
				c.Bus.WriteByte(c.Reg.PC, 0x34)
				// upper
				c.Bus.WriteByte(c.Reg.PC+1, 0x12)
				op.Handler(c, op.R1, op.R2)
				if strings.Contains(op.Mnemonic, "CALL a16") {
					assert.Equal(t, types.Addr(0x1234), c.Reg.PC)
					assert.Equal(t, types.Addr(0xFFFA), c.Reg.SP)
					assert.Equal(t, byte(0x7A), c.Bus.ReadByte(c.Reg.SP))
					assert.Equal(t, byte(0x56), c.Bus.ReadByte(c.Reg.SP+1))
				} else if strings.Contains(op.Mnemonic, "N") {
					assert.Equal(t, types.Addr(0x567A), c.Reg.PC)
					assert.Equal(t, types.Addr(0xFFFC), c.Reg.SP)
				} else {
					assert.Equal(t, types.Addr(0x1234), c.Reg.PC)
					assert.Equal(t, types.Addr(0xFFFA), c.Reg.SP)
					assert.Equal(t, byte(0x7A), c.Bus.ReadByte(c.Reg.SP))
					assert.Equal(t, byte(0x56), c.Bus.ReadByte(c.Reg.SP+1))
				}
			})
		})
	}
}

func TestOpCode_daa(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
	}

	tests := []struct {
		name string
		args args
	}{
		{name: "DAA", args: args{0x27}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := opCodes[tt.args.opcode]

			t.Run("when low 4 bit greater than or equal A", func(t *testing.T) {
				c.Reg.R[A] = 0x0A
				c.Reg.setZNHC(false, false, false, false)
				op.Handler(c, op.R1, op.R2)
				assert.Equal(t, byte(0x10), c.Reg.R[A])
				assert.Equal(t, false, c.Reg.isSet(flagZ))
				assert.Equal(t, false, c.Reg.isSet(flagH))
				assert.Equal(t, false, c.Reg.isSet(flagC))
			})
			t.Run("when low 4 bit less than A", func(t *testing.T) {
				c.Reg.R[A] = 0x09
				c.Reg.setZNHC(false, false, false, false)
				t.Run("and Harf carry = 0", func(t *testing.T) {
					op.Handler(c, op.R1, op.R2)
					assert.Equal(t, byte(0x09), c.Reg.R[A])
					assert.Equal(t, false, c.Reg.isSet(flagZ))
					assert.Equal(t, false, c.Reg.isSet(flagH))
					assert.Equal(t, false, c.Reg.isSet(flagC))
				})
				t.Run("and Harf carry = 1", func(t *testing.T) {
					c.Reg.setF(flagH, true)
					op.Handler(c, op.R1, op.R2)
					assert.Equal(t, byte(0x0F), c.Reg.R[A])
					assert.Equal(t, false, c.Reg.isSet(flagZ))
					assert.Equal(t, false, c.Reg.isSet(flagH))
					assert.Equal(t, false, c.Reg.isSet(flagC))
				})
			})
			t.Run("when high 4 bit greater than or equal A", func(t *testing.T) {
				c.Reg.R[A] = 0xA0
				c.Reg.setZNHC(false, false, false, false)
				op.Handler(c, op.R1, op.R2)
				assert.Equal(t, byte(0x00), c.Reg.R[A])
				assert.Equal(t, true, c.Reg.isSet(flagZ))
				assert.Equal(t, false, c.Reg.isSet(flagH))
				assert.Equal(t, true, c.Reg.isSet(flagC))
			})
			t.Run("when high 4 bit less than or equal A", func(t *testing.T) {
				c.Reg.R[A] = 0x90
				c.Reg.setZNHC(false, false, false, false)
				t.Run("and Carry = 0", func(t *testing.T) {
					op.Handler(c, op.R1, op.R2)
					assert.Equal(t, byte(0x90), c.Reg.R[A])
					assert.Equal(t, false, c.Reg.isSet(flagZ))
					assert.Equal(t, false, c.Reg.isSet(flagH))
					assert.Equal(t, false, c.Reg.isSet(flagC))
				})
				t.Run("and Carry = 1", func(t *testing.T) {
					c.Reg.setF(flagC, true)
					op.Handler(c, op.R1, op.R2)
					assert.Equal(t, byte(0xF0), c.Reg.R[A])
					assert.Equal(t, false, c.Reg.isSet(flagZ))
					assert.Equal(t, false, c.Reg.isSet(flagH))
					assert.Equal(t, true, c.Reg.isSet(flagC))
				})
			})
		})
	}
}

// PREFIX CB

func TestOpCode_rlc(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
		r1     int
	}

	tests := []struct {
		name string
		args args
	}{
		{name: "RLC B", args: args{0x00, B}},
		{name: "RLC C", args: args{0x01, C}},
		{name: "RLC D", args: args{0x02, D}},
		{name: "RLC E", args: args{0x03, E}},
		{name: "RLC H", args: args{0x04, H}},
		{name: "RLC L", args: args{0x05, L}},
		{name: "RLC HL", args: args{0x06, HL}},
		{name: "RLC A", args: args{0x07, A}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := cbOpCodes[tt.args.opcode]

			assert.Equal(t, tt.args.r1, op.R1)

			t.Run("when bit7 = 1", func(t *testing.T) {
				before := byte(0b10010000)
				if strings.Contains(op.Mnemonic, "HL") {
					c.Bus.WriteByte(c.Reg.R16(int(HL)), before)
				} else {
					c.Reg.R[tt.args.r1] = before
				}

				op.Handler(c, op.R1, op.R2)
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
				if strings.Contains(op.Mnemonic, "HL") {
					c.Bus.WriteByte(c.Reg.R16(int(HL)), before)
				} else {
					c.Reg.R[tt.args.r1] = before
				}

				op.Handler(c, op.R1, op.R2)
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
				if strings.Contains(op.Mnemonic, "HL") {
					c.Bus.WriteByte(c.Reg.R16(int(HL)), before)
				} else {
					c.Reg.R[tt.args.r1] = before
				}

				op.Handler(c, op.R1, op.R2)
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

func TestOpCode_rrc(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
		r1     int
	}

	tests := []struct {
		name string
		args args
	}{
		{name: "RRC B", args: args{0x08, B}},
		{name: "RRC C", args: args{0x09, C}},
		{name: "RRC D", args: args{0x0A, D}},
		{name: "RRC E", args: args{0x0B, E}},
		{name: "RRC H", args: args{0x0C, H}},
		{name: "RRC L", args: args{0x0D, L}},
		{name: "RRC HL", args: args{0x0E, HL}},
		{name: "RRC A", args: args{0x0F, A}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := cbOpCodes[tt.args.opcode]

			assert.Equal(t, tt.args.r1, op.R1)

			t.Run("when bit0 = 1", func(t *testing.T) {
				before := byte(0b10010001)
				if strings.Contains(op.Mnemonic, "HL") {
					c.Bus.WriteByte(c.Reg.R16(int(HL)), before)
				} else {
					c.Reg.R[tt.args.r1] = before
				}

				op.Handler(c, op.R1, op.R2)
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
				if strings.Contains(op.Mnemonic, "HL") {
					c.Bus.WriteByte(c.Reg.R16(int(HL)), before)
				} else {
					c.Reg.R[tt.args.r1] = before
				}

				op.Handler(c, op.R1, op.R2)
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
				if strings.Contains(op.Mnemonic, "HL") {
					c.Bus.WriteByte(c.Reg.R16(int(HL)), before)
				} else {
					c.Reg.R[tt.args.r1] = before
				}

				op.Handler(c, op.R1, op.R2)
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

func TestOpCode_rl(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
		r1     int
	}

	tests := []struct {
		name string
		args args
	}{
		{name: "RL B", args: args{0x10, B}},
		{name: "RL C", args: args{0x11, C}},
		{name: "RL D", args: args{0x12, D}},
		{name: "RL E", args: args{0x13, E}},
		{name: "RL H", args: args{0x14, H}},
		{name: "RL L", args: args{0x15, L}},
		{name: "RL HL", args: args{0x16, HL}},
		{name: "RL A", args: args{0x17, A}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := cbOpCodes[tt.args.opcode]

			assert.Equal(t, tt.args.r1, op.R1)

			t.Run("when bit7 = 1", func(t *testing.T) {
				before := byte(0b10010000)
				t.Run("and carry = 1", func(t *testing.T) {
					if strings.Contains(op.Mnemonic, "HL") {
						c.Bus.WriteByte(c.Reg.R16(int(HL)), before)
					} else {
						c.Reg.R[tt.args.r1] = before
					}

					c.Reg.setF(flagC, true)
					op.Handler(c, op.R1, op.R2)
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
					if strings.Contains(op.Mnemonic, "HL") {
						c.Bus.WriteByte(c.Reg.R16(int(HL)), before)
					} else {
						c.Reg.R[tt.args.r1] = before
					}

					c.Reg.setF(flagC, false)
					op.Handler(c, op.R1, op.R2)
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
					if strings.Contains(op.Mnemonic, "HL") {
						c.Bus.WriteByte(c.Reg.R16(int(HL)), before)
					} else {
						c.Reg.R[tt.args.r1] = before
					}

					c.Reg.setF(flagC, true)
					op.Handler(c, op.R1, op.R2)
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
					if strings.Contains(op.Mnemonic, "HL") {
						c.Bus.WriteByte(c.Reg.R16(int(HL)), before)
					} else {
						c.Reg.R[tt.args.r1] = before
					}

					c.Reg.setF(flagC, false)
					op.Handler(c, op.R1, op.R2)
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
					if strings.Contains(op.Mnemonic, "HL") {
						c.Bus.WriteByte(c.Reg.R16(int(HL)), before)
					} else {
						c.Reg.R[tt.args.r1] = before
					}

					c.Reg.setF(flagC, true)
					op.Handler(c, op.R1, op.R2)
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
					if strings.Contains(op.Mnemonic, "HL") {
						c.Bus.WriteByte(c.Reg.R16(int(HL)), before)
					} else {
						c.Reg.R[tt.args.r1] = before
					}

					c.Reg.setF(flagC, false)
					op.Handler(c, op.R1, op.R2)
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

func TestOpCode_rr(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
		r1     int
	}

	tests := []struct {
		name string
		args args
	}{
		{name: "RR B", args: args{0x18, B}},
		{name: "RR C", args: args{0x19, C}},
		{name: "RR D", args: args{0x1A, D}},
		{name: "RR E", args: args{0x1B, E}},
		{name: "RR H", args: args{0x1C, H}},
		{name: "RR L", args: args{0x1D, L}},
		{name: "RR HL", args: args{0x1E, HL}},
		{name: "RR A", args: args{0x1F, A}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := cbOpCodes[tt.args.opcode]

			assert.Equal(t, tt.args.r1, op.R1)

			t.Run("when bit0 = 1", func(t *testing.T) {
				before := byte(0b10010001)
				t.Run("and carry = 1", func(t *testing.T) {
					if strings.Contains(op.Mnemonic, "HL") {
						c.Bus.WriteByte(c.Reg.R16(int(HL)), before)
					} else {
						c.Reg.R[tt.args.r1] = before
					}

					c.Reg.setF(flagC, true)
					op.Handler(c, op.R1, op.R2)
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
					if strings.Contains(op.Mnemonic, "HL") {
						c.Bus.WriteByte(c.Reg.R16(int(HL)), before)
					} else {
						c.Reg.R[tt.args.r1] = before
					}

					c.Reg.setF(flagC, false)
					op.Handler(c, op.R1, op.R2)
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
					if strings.Contains(op.Mnemonic, "HL") {
						c.Bus.WriteByte(c.Reg.R16(int(HL)), before)
					} else {
						c.Reg.R[tt.args.r1] = before
					}

					c.Reg.setF(flagC, true)
					op.Handler(c, op.R1, op.R2)
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
					if strings.Contains(op.Mnemonic, "HL") {
						c.Bus.WriteByte(c.Reg.R16(int(HL)), before)
					} else {
						c.Reg.R[tt.args.r1] = before
					}

					c.Reg.setF(flagC, false)
					op.Handler(c, op.R1, op.R2)
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
					if strings.Contains(op.Mnemonic, "HL") {
						c.Bus.WriteByte(c.Reg.R16(int(HL)), before)
					} else {
						c.Reg.R[tt.args.r1] = before
					}

					c.Reg.setF(flagC, true)
					op.Handler(c, op.R1, op.R2)
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
					if strings.Contains(op.Mnemonic, "HL") {
						c.Bus.WriteByte(c.Reg.R16(int(HL)), before)
					} else {
						c.Reg.R[tt.args.r1] = before
					}

					c.Reg.setF(flagC, false)
					op.Handler(c, op.R1, op.R2)
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

func TestOpCode_sla(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
		r1     int
	}

	tests := []struct {
		name string
		args args
	}{
		{name: "SLA B", args: args{0x20, B}},
		{name: "SLA C", args: args{0x21, C}},
		{name: "SLA D", args: args{0x22, D}},
		{name: "SLA E", args: args{0x23, E}},
		{name: "SLA H", args: args{0x24, H}},
		{name: "SLA L", args: args{0x25, L}},
		{name: "SLA HL", args: args{0x26, HL}},
		{name: "SLA A", args: args{0x27, A}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := cbOpCodes[tt.args.opcode]

			assert.Equal(t, tt.args.r1, op.R1)

			t.Run("when bit7 = 1", func(t *testing.T) {
				before := byte(0b10010000)
				if strings.Contains(op.Mnemonic, "HL") {
					c.Bus.WriteByte(c.Reg.R16(int(HL)), before)
				} else {
					c.Reg.R[tt.args.r1] = before
				}

				op.Handler(c, op.R1, op.R2)
				assert.Equal(t, false, c.Reg.isSet(flagZ))
				assert.Equal(t, false, c.Reg.isSet(flagN))
				assert.Equal(t, false, c.Reg.isSet(flagH))
				assert.Equal(t, true, c.Reg.isSet(flagC))

				want := byte(0b00100000)

				if !strings.Contains(op.Mnemonic, "HL") {
					assert.Equal(t, want, c.Reg.R[op.R1])
				} else {
					assert.Equal(t, want, c.Bus.ReadByte(c.Reg.R16(int(HL))))
				}
			})
			t.Run("when bit7 = 0 and other = 0", func(t *testing.T) {
				before := byte(0b00000000)
				if strings.Contains(op.Mnemonic, "HL") {
					c.Bus.WriteByte(c.Reg.R16(int(HL)), before)
				} else {
					c.Reg.R[tt.args.r1] = before
				}

				op.Handler(c, op.R1, op.R2)
				assert.Equal(t, true, c.Reg.isSet(flagZ))
				assert.Equal(t, false, c.Reg.isSet(flagN))
				assert.Equal(t, false, c.Reg.isSet(flagH))
				assert.Equal(t, false, c.Reg.isSet(flagC))

				want := byte(0b00000000)

				if !strings.Contains(op.Mnemonic, "HL") {
					assert.Equal(t, want, c.Reg.R[op.R1])
				} else {
					assert.Equal(t, want, c.Bus.ReadByte(c.Reg.R16(int(HL))))
				}
			})
			t.Run("when bit7 = 0 and other != 0", func(t *testing.T) {
				before := byte(0b00100000)
				if strings.Contains(op.Mnemonic, "HL") {
					c.Bus.WriteByte(c.Reg.R16(int(HL)), before)
				} else {
					c.Reg.R[tt.args.r1] = before
				}

				op.Handler(c, op.R1, op.R2)
				assert.Equal(t, false, c.Reg.isSet(flagZ))
				assert.Equal(t, false, c.Reg.isSet(flagN))
				assert.Equal(t, false, c.Reg.isSet(flagH))
				assert.Equal(t, false, c.Reg.isSet(flagC))

				want := byte(0b01000000)

				if !strings.Contains(op.Mnemonic, "HL") {
					assert.Equal(t, want, c.Reg.R[op.R1])
				} else {
					assert.Equal(t, want, c.Bus.ReadByte(c.Reg.R16(int(HL))))
				}
			})
		})
	}
}

func TestOpCode_sra(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
		r1     int
	}

	tests := []struct {
		name string
		args args
	}{
		{name: "SRA B", args: args{0x28, B}},
		{name: "SRA C", args: args{0x29, C}},
		{name: "SRA D", args: args{0x2A, D}},
		{name: "SRA E", args: args{0x2B, E}},
		{name: "SRA H", args: args{0x2C, H}},
		{name: "SRA L", args: args{0x2D, L}},
		{name: "SRA HL", args: args{0x2E, HL}},
		{name: "SRA A", args: args{0x2F, A}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := cbOpCodes[tt.args.opcode]

			assert.Equal(t, tt.args.r1, op.R1)

			t.Run("when bit0 = 1 and bit7 = 1", func(t *testing.T) {
				before := byte(0b10010001)
				if strings.Contains(op.Mnemonic, "HL") {
					c.Bus.WriteByte(c.Reg.R16(int(HL)), before)
				} else {
					c.Reg.R[tt.args.r1] = before
				}

				op.Handler(c, op.R1, op.R2)
				assert.Equal(t, false, c.Reg.isSet(flagZ))
				assert.Equal(t, false, c.Reg.isSet(flagN))
				assert.Equal(t, false, c.Reg.isSet(flagH))
				assert.Equal(t, true, c.Reg.isSet(flagC))

				want := byte(0b11001000)

				if tt.name != "SRA HL" {
					assert.Equal(t, want, c.Reg.R[op.R1])
				} else {
					assert.Equal(t, want, c.Bus.ReadByte(c.Reg.R16(int(HL))))
				}
			})
			t.Run("when bit0 = 0 and other = 0", func(t *testing.T) {
				before := byte(0b00000000)
				if strings.Contains(op.Mnemonic, "HL") {
					c.Bus.WriteByte(c.Reg.R16(int(HL)), before)
				} else {
					c.Reg.R[tt.args.r1] = before
				}

				op.Handler(c, op.R1, op.R2)
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
				if strings.Contains(op.Mnemonic, "HL") {
					c.Bus.WriteByte(c.Reg.R16(int(HL)), before)
				} else {
					c.Reg.R[tt.args.r1] = before
				}

				op.Handler(c, op.R1, op.R2)
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

func TestOpCode_swap(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
		r1     int
	}

	tests := []struct {
		name string
		args args
	}{
		{name: "SWAP B", args: args{0x30, B}},
		{name: "SWAP C", args: args{0x31, C}},
		{name: "SWAP D", args: args{0x32, D}},
		{name: "SWAP E", args: args{0x33, E}},
		{name: "SWAP H", args: args{0x34, H}},
		{name: "SWAP L", args: args{0x35, L}},
		{name: "SWAP HL", args: args{0x36, HL}},
		{name: "SWAP A", args: args{0x37, A}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := cbOpCodes[tt.args.opcode]

			assert.Equal(t, tt.args.r1, op.R1)

			t.Run("when != 0 ", func(t *testing.T) {
				before := byte(0b10010100)
				if strings.Contains(op.Mnemonic, "HL") {
					c.Bus.WriteByte(c.Reg.R16(int(HL)), before)
				} else {
					c.Reg.R[tt.args.r1] = before
				}

				op.Handler(c, op.R1, op.R2)
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
				if strings.Contains(op.Mnemonic, "HL") {
					c.Bus.WriteByte(c.Reg.R16(int(HL)), before)
				} else {
					c.Reg.R[tt.args.r1] = before
				}

				op.Handler(c, op.R1, op.R2)
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

func TestOpCode_srl(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
		r1     int
	}

	tests := []struct {
		name string
		args args
	}{
		{name: "SRL B", args: args{0x38, B}},
		{name: "SRL C", args: args{0x39, C}},
		{name: "SRL D", args: args{0x3A, D}},
		{name: "SRL E", args: args{0x3B, E}},
		{name: "SRL H", args: args{0x3C, H}},
		{name: "SRL L", args: args{0x3D, L}},
		{name: "SRL HL", args: args{0x3E, HL}},
		{name: "SRL A", args: args{0x3F, A}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := cbOpCodes[tt.args.opcode]

			assert.Equal(t, tt.args.r1, op.R1)

			t.Run("when bit0 = 1", func(t *testing.T) {
				before := byte(0b10010001)
				if strings.Contains(op.Mnemonic, "HL") {
					c.Bus.WriteByte(c.Reg.R16(int(HL)), before)
				} else {
					c.Reg.R[tt.args.r1] = before
				}

				op.Handler(c, op.R1, op.R2)
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
				if strings.Contains(op.Mnemonic, "HL") {
					c.Bus.WriteByte(c.Reg.R16(int(HL)), before)
				} else {
					c.Reg.R[tt.args.r1] = before
				}

				op.Handler(c, op.R1, op.R2)
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
				if strings.Contains(op.Mnemonic, "HL") {
					c.Bus.WriteByte(c.Reg.R16(int(HL)), before)
				} else {
					c.Reg.R[tt.args.r1] = before
				}

				op.Handler(c, op.R1, op.R2)
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

func TestOpCode_bit(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
		bit    int
		r1     int
	}

	tests := []struct {
		name string
		args args
	}{
		{name: "BIT 0,B", args: args{0x40, 0, B}},
		{name: "BIT 0,C", args: args{0x41, 0, C}},
		{name: "BIT 0,D", args: args{0x42, 0, D}},
		{name: "BIT 0,E", args: args{0x43, 0, E}},
		{name: "BIT 0,H", args: args{0x44, 0, H}},
		{name: "BIT 0,L", args: args{0x45, 0, L}},
		{name: "BIT 0,(HL)", args: args{0x46, 0, HL}},
		{name: "BIT 0,A", args: args{0x47, 0, A}},
		{name: "BIT 1,B", args: args{0x48, 1, B}},
		{name: "BIT 1,C", args: args{0x49, 1, C}},
		{name: "BIT 1,D", args: args{0x4A, 1, D}},
		{name: "BIT 1,E", args: args{0x4B, 1, E}},
		{name: "BIT 1,H", args: args{0x4C, 1, H}},
		{name: "BIT 1,L", args: args{0x4D, 1, L}},
		{name: "BIT 1,(HL)", args: args{0x4E, 1, HL}},
		{name: "BIT 1,A", args: args{0x4F, 1, A}},
		{name: "BIT 2,B", args: args{0x50, 2, B}},
		{name: "BIT 2,C", args: args{0x51, 2, C}},
		{name: "BIT 2,D", args: args{0x52, 2, D}},
		{name: "BIT 2,E", args: args{0x53, 2, E}},
		{name: "BIT 2,H", args: args{0x54, 2, H}},
		{name: "BIT 2,L", args: args{0x55, 2, L}},
		{name: "BIT 2,(HL)", args: args{0x56, 2, HL}},
		{name: "BIT 2,A", args: args{0x57, 2, A}},
		{name: "BIT 3,B", args: args{0x58, 3, B}},
		{name: "BIT 3,C", args: args{0x59, 3, C}},
		{name: "BIT 3,D", args: args{0x5A, 3, D}},
		{name: "BIT 3,E", args: args{0x5B, 3, E}},
		{name: "BIT 3,H", args: args{0x5C, 3, H}},
		{name: "BIT 3,L", args: args{0x5D, 3, L}},
		{name: "BIT 3,(HL)", args: args{0x5E, 3, HL}},
		{name: "BIT 3,A", args: args{0x5F, 3, A}},
		{name: "BIT 4,B", args: args{0x60, 4, B}},
		{name: "BIT 4,C", args: args{0x61, 4, C}},
		{name: "BIT 4,D", args: args{0x62, 4, D}},
		{name: "BIT 4,E", args: args{0x63, 4, E}},
		{name: "BIT 4,H", args: args{0x64, 4, H}},
		{name: "BIT 4,L", args: args{0x65, 4, L}},
		{name: "BIT 4,(HL)", args: args{0x66, 4, HL}},
		{name: "BIT 4,A", args: args{0x67, 4, A}},
		{name: "BIT 5,B", args: args{0x68, 5, B}},
		{name: "BIT 5,C", args: args{0x69, 5, C}},
		{name: "BIT 5,D", args: args{0x6A, 5, D}},
		{name: "BIT 5,E", args: args{0x6B, 5, E}},
		{name: "BIT 5,H", args: args{0x6C, 5, H}},
		{name: "BIT 5,L", args: args{0x6D, 5, L}},
		{name: "BIT 5,(HL)", args: args{0x6E, 5, HL}},
		{name: "BIT 5,A", args: args{0x6F, 5, A}},
		{name: "BIT 6,B", args: args{0x70, 6, B}},
		{name: "BIT 6,C", args: args{0x71, 6, C}},
		{name: "BIT 6,D", args: args{0x72, 6, D}},
		{name: "BIT 6,E", args: args{0x73, 6, E}},
		{name: "BIT 6,H", args: args{0x74, 6, H}},
		{name: "BIT 6,L", args: args{0x75, 6, L}},
		{name: "BIT 6,(HL)", args: args{0x76, 6, HL}},
		{name: "BIT 6,A", args: args{0x77, 6, A}},
		{name: "BIT 7,B", args: args{0x78, 7, B}},
		{name: "BIT 7,C", args: args{0x79, 7, C}},
		{name: "BIT 7,D", args: args{0x7A, 7, D}},
		{name: "BIT 7,E", args: args{0x7B, 7, E}},
		{name: "BIT 7,H", args: args{0x7C, 7, H}},
		{name: "BIT 7,L", args: args{0x7D, 7, L}},
		{name: "BIT 7,(HL)", args: args{0x7E, 7, HL}},
		{name: "BIT 7,A", args: args{0x7F, 7, A}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := cbOpCodes[tt.args.opcode]

			assert.Equal(t, tt.args.bit, op.R1)
			assert.Equal(t, tt.args.r1, op.R2)

			t.Run("when bitN = 1", func(t *testing.T) {
				before := byte(0b00000000)
				before |= (1 << tt.args.bit)
				if strings.Contains(op.Mnemonic, "HL") {
					c.Bus.WriteByte(c.Reg.R16(int(HL)), before)
				} else {
					c.Reg.R[tt.args.r1] = before
				}

				op.Handler(c, op.R1, op.R2)
				assert.Equal(t, false, c.Reg.isSet(flagZ))
				assert.Equal(t, false, c.Reg.isSet(flagN))
				assert.Equal(t, true, c.Reg.isSet(flagH))
			})
			t.Run("when bitN = 0", func(t *testing.T) {
				before := byte(0b00000000)
				if strings.Contains(op.Mnemonic, "HL") {
					c.Bus.WriteByte(c.Reg.R16(int(HL)), before)
				} else {
					c.Reg.R[tt.args.r1] = before
				}

				op.Handler(c, op.R1, op.R2)
				assert.Equal(t, true, c.Reg.isSet(flagZ))
				assert.Equal(t, false, c.Reg.isSet(flagN))
				assert.Equal(t, true, c.Reg.isSet(flagH))
			})
		})
	}
}

func TestOpCode_res(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
		bit    int
		r1     int
	}

	tests := []struct {
		name string
		args args
	}{
		{name: "RES 0,B", args: args{0x80, 0, B}},
		{name: "RES 0,C", args: args{0x81, 0, C}},
		{name: "RES 0,D", args: args{0x82, 0, D}},
		{name: "RES 0,E", args: args{0x83, 0, E}},
		{name: "RES 0,H", args: args{0x84, 0, H}},
		{name: "RES 0,L", args: args{0x85, 0, L}},
		{name: "RES 0,(HL)", args: args{0x86, 0, HL}},
		{name: "RES 0,A", args: args{0x87, 0, A}},
		{name: "RES 1,B", args: args{0x88, 1, B}},
		{name: "RES 1,C", args: args{0x89, 1, C}},
		{name: "RES 1,D", args: args{0x8A, 1, D}},
		{name: "RES 1,E", args: args{0x8B, 1, E}},
		{name: "RES 1,H", args: args{0x8C, 1, H}},
		{name: "RES 1,L", args: args{0x8D, 1, L}},
		{name: "RES 1,(HL)", args: args{0x8E, 1, HL}},
		{name: "RES 1,A", args: args{0x8F, 1, A}},
		{name: "RES 2,B", args: args{0x90, 2, B}},
		{name: "RES 2,C", args: args{0x91, 2, C}},
		{name: "RES 2,D", args: args{0x92, 2, D}},
		{name: "RES 2,E", args: args{0x93, 2, E}},
		{name: "RES 2,H", args: args{0x94, 2, H}},
		{name: "RES 2,L", args: args{0x95, 2, L}},
		{name: "RES 2,(HL)", args: args{0x96, 2, HL}},
		{name: "RES 2,A", args: args{0x97, 2, A}},
		{name: "RES 3,B", args: args{0x98, 3, B}},
		{name: "RES 3,C", args: args{0x99, 3, C}},
		{name: "RES 3,D", args: args{0x9A, 3, D}},
		{name: "RES 3,E", args: args{0x9B, 3, E}},
		{name: "RES 3,H", args: args{0x9C, 3, H}},
		{name: "RES 3,L", args: args{0x9D, 3, L}},
		{name: "RES 3,(HL)", args: args{0x9E, 3, HL}},
		{name: "RES 3,A", args: args{0x9F, 3, A}},
		{name: "RES 4,B", args: args{0xA0, 4, B}},
		{name: "RES 4,C", args: args{0xA1, 4, C}},
		{name: "RES 4,D", args: args{0xA2, 4, D}},
		{name: "RES 4,E", args: args{0xA3, 4, E}},
		{name: "RES 4,H", args: args{0xA4, 4, H}},
		{name: "RES 4,L", args: args{0xA5, 4, L}},
		{name: "RES 4,(HL)", args: args{0xA6, 4, HL}},
		{name: "RES 4,A", args: args{0xA7, 4, A}},
		{name: "RES 5,B", args: args{0xA8, 5, B}},
		{name: "RES 5,C", args: args{0xA9, 5, C}},
		{name: "RES 5,D", args: args{0xAA, 5, D}},
		{name: "RES 5,E", args: args{0xAB, 5, E}},
		{name: "RES 5,H", args: args{0xAC, 5, H}},
		{name: "RES 5,L", args: args{0xAD, 5, L}},
		{name: "RES 5,(HL)", args: args{0xAE, 5, HL}},
		{name: "RES 5,A", args: args{0xAF, 5, A}},
		{name: "RES 6,B", args: args{0xB0, 6, B}},
		{name: "RES 6,C", args: args{0xB1, 6, C}},
		{name: "RES 6,D", args: args{0xB2, 6, D}},
		{name: "RES 6,E", args: args{0xB3, 6, E}},
		{name: "RES 6,H", args: args{0xB4, 6, H}},
		{name: "RES 6,L", args: args{0xB5, 6, L}},
		{name: "RES 6,(HL)", args: args{0xB6, 6, HL}},
		{name: "RES 6,A", args: args{0xB7, 6, A}},
		{name: "RES 7,B", args: args{0xB8, 7, B}},
		{name: "RES 7,C", args: args{0xB9, 7, C}},
		{name: "RES 7,D", args: args{0xBA, 7, D}},
		{name: "RES 7,E", args: args{0xBB, 7, E}},
		{name: "RES 7,H", args: args{0xBC, 7, H}},
		{name: "RES 7,L", args: args{0xBD, 7, L}},
		{name: "RES 7,(HL)", args: args{0xBE, 7, HL}},
		{name: "RES 7,A", args: args{0xBF, 7, A}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := cbOpCodes[tt.args.opcode]

			assert.Equal(t, tt.args.bit, op.R1)
			assert.Equal(t, tt.args.r1, op.R2)

			t.Run("when bitN = 1", func(t *testing.T) {
				before := byte(0b00000000)
				before |= (1 << tt.args.bit)
				if strings.Contains(op.Mnemonic, "HL") {
					c.Bus.WriteByte(c.Reg.R16(int(HL)), before)
				} else {
					c.Reg.R[tt.args.r1] = before
				}

				op.Handler(c, op.R1, op.R2)
				if !strings.Contains(op.Mnemonic, "HL") {
					assert.Equal(t, byte(0), util.Bit(c.Reg.R[op.R2], int(op.R1)))
				} else {
					assert.Equal(t, byte(0), util.Bit(c.Bus.ReadByte(c.Reg.R16(int(op.R2))), int(op.R1)))
				}
			})
			t.Run("when bitN = 0", func(t *testing.T) {
				before := byte(0b00000000)
				if strings.Contains(op.Mnemonic, "HL") {
					c.Bus.WriteByte(c.Reg.R16(int(HL)), before)
				} else {
					c.Reg.R[tt.args.r1] = before
				}

				op.Handler(c, op.R1, op.R2)
				if !strings.Contains(op.Mnemonic, "HL") {
					assert.Equal(t, byte(0), util.Bit(c.Reg.R[op.R2], int(op.R1)))
				} else {
					assert.Equal(t, byte(0), util.Bit(c.Bus.ReadByte(c.Reg.R16(int(op.R2))), int(op.R1)))
				}
			})
		})
	}
}

func TestOpCode_set(t *testing.T) {
	c := setupCPU()

	type args struct {
		opcode byte
		bit    int
		r1     int
	}

	tests := []struct {
		name string
		args args
	}{
		{name: "SET 0,B", args: args{0xC0, 0, B}},
		{name: "SET 0,C", args: args{0xC1, 0, C}},
		{name: "SET 0,D", args: args{0xC2, 0, D}},
		{name: "SET 0,E", args: args{0xC3, 0, E}},
		{name: "SET 0,H", args: args{0xC4, 0, H}},
		{name: "SET 0,L", args: args{0xC5, 0, L}},
		{name: "SET 0,(HL)", args: args{0xC6, 0, HL}},
		{name: "SET 0,A", args: args{0xC7, 0, A}},
		{name: "SET 1,B", args: args{0xC8, 1, B}},
		{name: "SET 1,C", args: args{0xC9, 1, C}},
		{name: "SET 1,D", args: args{0xCA, 1, D}},
		{name: "SET 1,E", args: args{0xCB, 1, E}},
		{name: "SET 1,H", args: args{0xCC, 1, H}},
		{name: "SET 1,L", args: args{0xCD, 1, L}},
		{name: "SET 1,(HL)", args: args{0xCE, 1, HL}},
		{name: "SET 1,A", args: args{0xCF, 1, A}},
		{name: "SET 2,B", args: args{0xD0, 2, B}},
		{name: "SET 2,C", args: args{0xD1, 2, C}},
		{name: "SET 2,D", args: args{0xD2, 2, D}},
		{name: "SET 2,E", args: args{0xD3, 2, E}},
		{name: "SET 2,H", args: args{0xD4, 2, H}},
		{name: "SET 2,L", args: args{0xD5, 2, L}},
		{name: "SET 2,(HL)", args: args{0xD6, 2, HL}},
		{name: "SET 2,A", args: args{0xD7, 2, A}},
		{name: "SET 3,B", args: args{0xD8, 3, B}},
		{name: "SET 3,C", args: args{0xD9, 3, C}},
		{name: "SET 3,D", args: args{0xDA, 3, D}},
		{name: "SET 3,E", args: args{0xDB, 3, E}},
		{name: "SET 3,H", args: args{0xDC, 3, H}},
		{name: "SET 3,L", args: args{0xDD, 3, L}},
		{name: "SET 3,(HL)", args: args{0xDE, 3, HL}},
		{name: "SET 3,A", args: args{0xDF, 3, A}},
		{name: "SET 4,B", args: args{0xE0, 4, B}},
		{name: "SET 4,C", args: args{0xE1, 4, C}},
		{name: "SET 4,D", args: args{0xE2, 4, D}},
		{name: "SET 4,E", args: args{0xE3, 4, E}},
		{name: "SET 4,H", args: args{0xE4, 4, H}},
		{name: "SET 4,L", args: args{0xE5, 4, L}},
		{name: "SET 4,(HL)", args: args{0xE6, 4, HL}},
		{name: "SET 4,A", args: args{0xE7, 4, A}},
		{name: "SET 5,B", args: args{0xE8, 5, B}},
		{name: "SET 5,C", args: args{0xE9, 5, C}},
		{name: "SET 5,D", args: args{0xEA, 5, D}},
		{name: "SET 5,E", args: args{0xEB, 5, E}},
		{name: "SET 5,H", args: args{0xEC, 5, H}},
		{name: "SET 5,L", args: args{0xED, 5, L}},
		{name: "SET 5,(HL)", args: args{0xEE, 5, HL}},
		{name: "SET 5,A", args: args{0xEF, 5, A}},
		{name: "SET 6,B", args: args{0xF0, 6, B}},
		{name: "SET 6,C", args: args{0xF1, 6, C}},
		{name: "SET 6,D", args: args{0xF2, 6, D}},
		{name: "SET 6,E", args: args{0xF3, 6, E}},
		{name: "SET 6,H", args: args{0xF4, 6, H}},
		{name: "SET 6,L", args: args{0xF5, 6, L}},
		{name: "SET 6,(HL)", args: args{0xF6, 6, HL}},
		{name: "SET 6,A", args: args{0xF7, 6, A}},
		{name: "SET 7,B", args: args{0xF8, 7, B}},
		{name: "SET 7,C", args: args{0xF9, 7, C}},
		{name: "SET 7,D", args: args{0xFA, 7, D}},
		{name: "SET 7,E", args: args{0xFB, 7, E}},
		{name: "SET 7,H", args: args{0xFC, 7, H}},
		{name: "SET 7,L", args: args{0xFD, 7, L}},
		{name: "SET 7,(HL)", args: args{0xFE, 7, HL}},
		{name: "SET 7,A", args: args{0xFF, 7, A}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.regreset()
			op := cbOpCodes[tt.args.opcode]

			assert.Equal(t, tt.args.bit, op.R1)
			assert.Equal(t, tt.args.r1, op.R2)

			t.Run("when bitN = 1", func(t *testing.T) {
				before := byte(0b00000000)
				before |= (1 << tt.args.bit)
				if strings.Contains(op.Mnemonic, "HL") {
					c.Bus.WriteByte(c.Reg.R16(int(HL)), before)
				} else {
					c.Reg.R[tt.args.r1] = before
				}

				op.Handler(c, op.R1, op.R2)
				if !strings.Contains(op.Mnemonic, "HL") {
					assert.Equal(t, byte(1), util.Bit(c.Reg.R[op.R2], int(op.R1)))
				} else {
					assert.Equal(t, byte(1), util.Bit(c.Bus.ReadByte(c.Reg.R16(int(op.R2))), int(op.R1)))
				}
			})
			t.Run("when bitN = 0", func(t *testing.T) {
				before := byte(0b00000000)
				if strings.Contains(op.Mnemonic, "HL") {
					c.Bus.WriteByte(c.Reg.R16(int(HL)), before)
				} else {
					c.Reg.R[tt.args.r1] = before
				}

				op.Handler(c, op.R1, op.R2)
				if !strings.Contains(op.Mnemonic, "HL") {
					assert.Equal(t, byte(1), util.Bit(c.Reg.R[op.R2], int(op.R1)))
				} else {
					assert.Equal(t, byte(1), util.Bit(c.Bus.ReadByte(c.Reg.R16(int(op.R2))), int(op.R1)))
				}
			})
		})
	}
}
