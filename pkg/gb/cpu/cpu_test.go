package cpu

import (
	"io/ioutil"
	"log"
	"os"
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

func setup(file string) *bus.Bus {
	log.Printf("load file %s", file)
	romData, err := ioutil.ReadFile(file)

	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}

	cart := cartridge.New(romData)
	vram := ram.New(0x2000)
	wram := ram.New(0x2000)
	wram2 := ram.New(0x2000)
	hram := ram.New(0x0080)
	io := io.New(0x0080)
	ie := ie.New()
	bus := bus.New(cart, vram, wram, wram2, hram, io, ie)

	return bus
}

func Test01(t *testing.T) {
	file := "../../../test/blargg/cpu_instrs/individual/01-special.gb"
	bus := setup(file)
	cpu := New(bus)

	for {
		cpu.Step()
	}
}

func Test06(t *testing.T) {
	file := "../../../test/blargg/cpu_instrs/individual/06-ld r,r.gb"
	bus := setup(file)
	cpu := New(bus)

	for {
		cpu.Step()
	}
}

func TestCPU_fetch(t *testing.T) {
	file := "../../../test/blargg/cpu_instrs/individual/06-ld r,r.gb"
	bus := setup(file)
	c := New(bus)
	assert := assert.New(t)

	want := byte(0x12)
	c.Bus.WriteByte(c.Reg.PC, want)

	d := c.fetch()
	assert.Equal(want, d)
	assert.Equal(types.Addr(0x101), c.Reg.PC)
}

func TestCPU_fetch16(t *testing.T) {
	file := "../../../test/blargg/cpu_instrs/individual/06-ld r,r.gb"
	bus := setup(file)
	c := New(bus)
	assert := assert.New(t)

	lower := byte(0x12)
	upper := byte(0x13)
	want := types.Addr(0x1312)
	c.Bus.WriteByte(c.Reg.PC, lower)
	c.Bus.WriteByte(c.Reg.PC+1, upper)

	d := c.fetch16()
	assert.Equal(want, d)
	assert.Equal(types.Addr(0x102), c.Reg.PC)
}

func TestCPU_pushpop(t *testing.T) {
	file := "../../../test/blargg/cpu_instrs/individual/06-ld r,r.gb"
	bus := setup(file)
	c := New(bus)
	assert := assert.New(t)

	want := byte(0x01)

	before_sp := c.Reg.SP
	c.push(want)
	assert.Equal(before_sp-1, c.Reg.SP)
	assert.Equal(want, c.Bus.ReadByte(c.Reg.SP))

	d := c.pop()
	assert.Equal(before_sp, c.Reg.SP)
	assert.Equal(want, d)
}

func TestCPU_pushpopPC(t *testing.T) {
	file := "../../../test/blargg/cpu_instrs/individual/06-ld r,r.gb"
	bus := setup(file)
	c := New(bus)
	assert := assert.New(t)

	want := types.Addr(0x1234)
	c.Reg.PC = want

	c.pushPC()
	assert.Equal(util.ExtractLower(want), c.Bus.ReadByte(c.Reg.SP))
	assert.Equal(util.ExtractUpper(want), c.Bus.ReadByte(c.Reg.SP+1))

	c.popPC()
	assert.Equal(want, c.Reg.PC)
}
