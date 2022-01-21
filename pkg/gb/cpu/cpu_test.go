package cpu

import (
	"io/ioutil"
	"log"
	"os"
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

func setup(file string) *bus.Bus {
	log.Printf("load file %s", file)
	romData, err := ioutil.ReadFile(file)

	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}

	cart := cartridge.New(romData)
	vram := memory.NewRAM(0x2000)
	wram := memory.NewRAM(0x2000)
	wram2 := memory.NewRAM(0x2000)
	hram := memory.NewRAM(0x0080)
	g := gpu.New()
	io := io.NewIO(io.NewPad(), io.NewSerial(), io.NewTimer(), io.NewIRQ(), g, 0x2000)
	bus := bus.New(cart, vram, wram, wram2, hram, io)

	g.Init(bus, io.IRQ.Request)
	return bus
}

func testrom(t assert.TestingT, file string, passstr string) {
	bus := setup(file)
	cpu := New(bus)
	cpu.Bus.IO.Timer.SetRequestIRQ(cpu.IRQ.Request)

	var str string

	str = ""

	for i := 0; i < 40000000; i++ {
		if cpu.Bus.ReadByte(0xff02) == byte(0x81) {
			d := cpu.Bus.ReadByte(0xff01)
			str += string(d)
			cpu.Bus.WriteByte(0xff02, byte(0x00))
		}

		cpu.Step()
	}

	assert.Equal(t, passstr, str)
}

func TestCPU_Blargg_cpu_instrs(t *testing.T) {
	file := "../../../test/blargg/cpu_instrs/cpu_instrs.gb"
	passstr := "cpu_instrs\n\n01:ok  02:ok  03:ok  04:ok  05:ok  06:ok  07:ok  08:ok  09:ok  10:ok  11:ok  \n\nPassed all tests\n"

	testrom(t, file, passstr)

}

func TestTiming(t *testing.T) {
	file := "../../../test/blargg/instr_timing/instr_timing.gb"
	passstr := "instr_timing\n\n\nPassed"

	testrom(t, file, passstr)
}

func TestMooneye_acceptance(t *testing.T) {
	file := "../../../test/mooneye-gb/acceptance/bits/mem_oam.gb"
	passstr := "pass"

	testrom(t, file, passstr)
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

	before_sp := c.Reg.SP
	c.pushPC()
	assert.Equal(util.ExtractLower(want), c.Bus.ReadByte(before_sp-2))
	assert.Equal(util.ExtractUpper(want), c.Bus.ReadByte(before_sp-1))

	c.popPC()
	assert.Equal(want, c.Reg.PC)
	assert.Equal(before_sp, c.Reg.SP)
}
