package cartridge

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBank_Size(t *testing.T) {
	type args struct {
		size byte
		want uint16
	}

	tests := []struct {
		name string
		args args
	}{
		{name: "ROM_32KB", args: args{0x00, 2}},
		{name: "ROM_64KB", args: args{0x01, 4}},
		{name: "ROM_128KB", args: args{0x02, 8}},
		{name: "ROM_256KB", args: args{0x03, 16}},
		{name: "ROM_512KB", args: args{0x04, 32}},
		{name: "ROM_1024KB", args: args{0x05, 64}},
		{name: "ROM_2048KB", args: args{0x06, 128}},
		{name: "ROM_4096KB", args: args{0x07, 256}},
		{name: "ROM_8192KB", args: args{0x08, 512}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			size := getRomSize(tt.args.size)
			bank := NewBank(size)
			assert.Equal(t, tt.args.want, bank.Size)
		})
	}
}
