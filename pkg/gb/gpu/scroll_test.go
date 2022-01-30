package gpu

import (
	"testing"

	"github.com/Teshima-Tatsuya/GoBoy/pkg/types"
	"github.com/stretchr/testify/assert"
)

func TestScroll(t *testing.T) {

	type args struct {
		value byte
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "Scroll"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Run("isVBlankPeriod", func(t *testing.T) {
				s := NewScroll()
				s.LY = 143
				assert.Equal(t, false, s.isVBlankPeriod())
				s.LY = 144
				assert.Equal(t, true, s.isVBlankPeriod())
				s.LY = 153
				assert.Equal(t, true, s.isVBlankPeriod())
				s.LY = 154
				assert.Equal(t, false, s.isVBlankPeriod())
			})
			t.Run("isHBlankPeriod", func(t *testing.T) {
				s := NewScroll()
				s.LY = 0
				assert.Equal(t, true, s.isHBlankPeriod())
				s.LY = 143
				assert.Equal(t, true, s.isHBlankPeriod())
				s.LY = 144
				assert.Equal(t, false, s.isHBlankPeriod())
			})
			t.Run("isVBlankStart", func(t *testing.T) {
				s := NewScroll()
				s.LY = 143
				assert.Equal(t, false, s.isVBlankStart())
				s.LY = 144
				assert.Equal(t, true, s.isVBlankStart())
			})
		})
	}
}

func TestScroll_ReadWrite(t *testing.T) {

	type args struct {
		addr types.Addr
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "SCY", args: args{addr: 0x42}},
		{name: "SCX", args: args{addr: 0x43}},
		{name: "LY", args: args{addr: 0x44}},
		{name: "LYC", args: args{addr: 0x45}},
		{name: "WY", args: args{addr: 0x4A}},
		{name: "WX", args: args{addr: 0x4B}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewScroll()
			s.Write(tt.args.addr, 0x12)
			switch tt.name {
			case "SCY":
				assert.Equal(t, byte(0x12), s.SCY)
			case "SCX":
				assert.Equal(t, byte(0x12), s.SCX)
			case "LY":
				assert.Equal(t, byte(0x12), s.LY)
			case "LYC":
				assert.Equal(t, byte(0x12), s.LYC)
			case "WY":
				assert.Equal(t, byte(0x12), s.WY)
			case "WX":
				assert.Equal(t, byte(0x12), s.WX)
			}
			assert.Equal(t, byte(0x12), s.Read(tt.args.addr))
		})
	}
}
