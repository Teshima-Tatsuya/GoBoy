package bus

import (
	"testing"

	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/cartridge"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/ram"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/types"
	"github.com/stretchr/testify/assert"
)

func setup() *Bus {
	romData := make([]byte, 0x8000)
	cart := cartridge.New(romData)
	vram := ram.New(0x2000)
	wram := ram.New(0x2000)
	hram := ram.New(0x0080)
	bus := New(cart, vram, wram, hram)

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
		{
			name: "Cart",
			args: args{0x0000, 0x7FFF},
		},
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
