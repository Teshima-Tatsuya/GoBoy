package cpu

import (
	"testing"

	"github.com/Teshima-Tatsuya/GoBoy/pkg/types"
)

func setupRegister() *Register {
	r := &Register{}
	r.reset()

	return r
}

func TestRegister_R16(t *testing.T) {
	r := setupRegister()

	type args struct {
		flag  int
		lower types.Addr
		upper types.Addr
		value types.Addr
	}
	tests := []struct {
		name string
		args args
		want types.Addr
	}{
		{
			name: "AF",
			args: args{AF, F, A, 0x1234},
			want: 0x1234,
		},
		{
			name: "BC",
			args: args{BC, C, B, 0x1234},
			want: 0x1234,
		},
		{
			name: "DE",
			args: args{DE, E, D, 0x1234},
			want: 0x1234,
		},
		{
			name: "HL",
			args: args{HL, L, H, 0x1234},
			want: 0x1234,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r.setR16(tt.args.flag, tt.args.value)

			got := r.R16(int(tt.args.flag))
			lower := r.R[tt.args.lower]
			upper := r.R[tt.args.upper]
			if got != tt.want || lower != 0x34 || upper != 0x12 {
				t.Errorf("got = %v, want %v, lower 0x%2d, upper 0x%2d", got, tt.want, lower, upper)
			}
		})
	}
}

func TestRegister_R16_SP_PC(t *testing.T) {
	r := setupRegister()

	type args struct {
		flag  int
		value types.Addr
	}
	tests := []struct {
		name string
		args args
		want types.Addr
	}{
		{
			name: "SP",
			args: args{SP, 0x1234},
			want: 0x1234,
		},
		{
			name: "PC",
			args: args{PC, 0x1234},
			want: 0x1234,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r.setR16(tt.args.flag, tt.args.value)

			got := r.R16(int(tt.args.flag))
			if got != tt.want {
				t.Errorf("got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRegister_R16WithIncDec(t *testing.T) {
	r := setupRegister()

	type args struct {
		flag types.Addr
	}
	tests := []struct {
		name  string
		args  args
		want1 types.Addr
		want2 types.Addr
	}{
		{
			name:  "HLD",
			args:  args{HLD},
			want1: 0x1234,
			want2: 0x1233,
		},
		{
			name:  "HLI",
			args:  args{HLI},
			want1: 0x1234,
			want2: 0x1235,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r.setR16(HL, 0x1234)
			got := r.R16(int(tt.args.flag))
			got2 := r.R16(HL)
			if got != tt.want1 || got2 != tt.want2 {
				t.Errorf("got1 = 0x%4x, want1 0x%4x, got2 = 0x%4x, want2 0x%4x", got, tt.want1, got2, tt.want2)
			}
		})
	}
}

func TestRegister_Flags(t *testing.T) {
	r := setupRegister()

	type args struct {
		flag  int
		value bool
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "flagZ",
			args: args{flagZ, true},
		},
		{
			name: "flagN",
			args: args{flagN, true},
		},
		{
			name: "flagH",
			args: args{flagH, true},
		},
		{
			name: "flagC",
			args: args{flagC, true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got, want bool
			r.setF(tt.args.flag, true)
			got = r.isSet(tt.args.flag)
			want = true
			if got != want {
				t.Errorf("name = %s, actual = %v, want %v", tt.name, got, want)
			}
			r.setF(tt.args.flag, false)
			got = r.isSet(tt.args.flag)
			want = false
			if got != want {
				t.Errorf("name = %s, actual = %v, want %v", tt.name, got, want)
			}
			r.setF(tt.args.flag, true)
			got = r.isSet(tt.args.flag)
			want = true
			if got != want {
				t.Errorf("name = %s, actual = %v, want %v", tt.name, got, want)
			}
		})
	}
}
