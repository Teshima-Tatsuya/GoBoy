package cpu

import (
	"testing"

	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/bus"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/cartridge"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/ram"
	"github.com/stretchr/testify/assert"
)

func setupCPU() *CPU {
	romData := make([]byte, 0x8000)
	cart := cartridge.New(romData)
	vram := ram.New(0x2000)
	wram := ram.New(0x2000)
	hram := ram.New(0x0080)
	bus := bus.New(cart, vram, wram, hram)

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
			ldrr(c, byte(tt.args.r1), byte(tt.args.r2), []byte{})

			assert.Equal(t, op.R1, tt.args.r1)
			assert.Equal(t, op.R2, tt.args.r2)
			assert.Equal(t, c.Reg.R[tt.args.r1], want)
			assert.Equal(t, c.Reg.R[tt.args.r2], want)
		})
	}
}