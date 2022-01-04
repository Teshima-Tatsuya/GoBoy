package bus

import (
	"testing"

	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/cartridge"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/ie"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/io"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/ram"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/types"
	"github.com/stretchr/testify/assert"
)

func setup() *Bus {
	romData := make([]byte, 0x8000)
	cart := cartridge.New(romData)
	vram := ram.New(0x2000)
	wram := ram.New(0x2000)
	wram2 := ram.New(0x2000)
	hram := ram.New(0x0080)
	io := io.New(0x0080)
	ie := ie.New()
	bus := New(cart, vram, wram, wram2, hram, io, ie)

	return bus
}

func TestBus_CartReadWrite(t *testing.T) {
	b := setup()
	type args struct {
		start types.Addr
		end   types.Addr
	}

	tests := []struct {
		name string
		args args
	}{
		{name: "CART", args: args{0x0000, 0x7FFF}},
		{name: "VRAM", args: args{0x8000, 0x9FFF}},
		{name: "WRAM", args: args{0xC000, 0xCFFF}},
		{name: "WRAM2", args: args{0xD000, 0xDFFF}},
		{name: "ERAM", args: args{0xE000, 0xFDFF}},
		{name: "IO", args: args{0xFF00, 0xFF7F}},
		{name: "HRAM", args: args{0xFF80, 0xFFFE}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s, e byte
			// Read
			s = b.ReadByte(tt.args.start)
			e = b.ReadByte(tt.args.end)
			assert.Equal(t, byte(0x00), s)
			assert.Equal(t, byte(0x00), e)

			// Write
			b.WriteByte(tt.args.start, 0x01)
			b.WriteByte(tt.args.end, 0x02)
			s = b.ReadByte(tt.args.start)
			e = b.ReadByte(tt.args.end)
			assert.Equal(t, byte(0x01), s)
			assert.Equal(t, byte(0x02), e)
		})
	}

}
