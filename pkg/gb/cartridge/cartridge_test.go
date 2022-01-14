package cartridge

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupCartridge() *Cartridge {
	romData := make([]byte, 8000)
	return New(romData)
}
func TestCartridge_getRomSize(t *testing.T) {
	type args struct {
		size byte
		want int
	}

	tests := []struct {
		name string
		args args
	}{
		{name: "ROM_32KB", args: args{0x00, 32 * 1024}},
		{name: "ROM_64KB", args: args{0x01, 64 * 1024}},
		{name: "ROM_128KB", args: args{0x02, 128 * 1024}},
		{name: "ROM_256KB", args: args{0x03, 256 * 1024}},
		{name: "ROM_512KB", args: args{0x04, 512 * 1024}},
		{name: "ROM_1024KB", args: args{0x05, 1024 * 1024}},
		{name: "ROM_2048KB", args: args{0x06, 2048 * 1024}},
		{name: "ROM_4096KB", args: args{0x07, 4096 * 1024}},
		{name: "ROM_8192KB", args: args{0x08, 8192 * 1024}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			size := getRomSize(tt.args.size)
			assert.Equal(t, tt.args.want, size)
		})
	}
}
