package cpu

import (
	"io/ioutil"
	"log"
	"os"
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
	io := io.NewIO(io.NewPad(), io.NewSerial(), io.NewTimer(), io.NewIRQ(), gpu.New(), 0x2000)
	bus := bus.New(cart, vram, wram, wram2, hram, io)

	return bus
}

func testrom(t assert.TestingT, file string, passstr string) {
	bus := setup(file)
	cpu := New(bus)

	var str string

	str = ""

	for {
		if cpu.Bus.ReadByte(0xff02) == byte(0x81) {
			d := cpu.Bus.ReadByte(0xff01)
			str += string(d)
			cpu.Bus.WriteByte(0xff02, byte(0x00))
		}

		if strings.Contains(str, "Failed") ||
			strings.Contains(str, "Passed") {
			break
		}
		cycle := cpu.Step()
		if cpu.Bus.IO.Timer.Tick(cycle) {
			cpu.IRQ.Request(io.TimerFlag)
		}
	}

	assert.Equal(t, passstr, str)
}

func TestCPU_Blargg_cpu_instrs(t *testing.T) {
	file := "../../../test/blargg/cpu_instrs/cpu_instrs.gb"
	passstr := "cpu_instrs\n\n01:ok  02:ok  03:ok  04:ok  05:ok  06:ok  07:ok  08:ok  09:ok  10:ok  11:ok  \n\nPassed"

	testrom(t, file, passstr)

}

func Test01(t *testing.T) {
	file := "../../../test/blargg/cpu_instrs/individual/01-special.gb"
	passstr := "01-special\n\n\nPassed"

	testrom(t, file, passstr)

}

func Test02(t *testing.T) {
	file := "../../../test/blargg/cpu_instrs/individual/02-interrupts.gb"
	passstr := "02-interrupts\n\n\nPassed"

	testrom(t, file, passstr)
}

func Test03(t *testing.T) {
	file := "../../../test/blargg/cpu_instrs/individual/03-op sp,hl.gb"
	passstr := "03-op sp,hl\n\n\nPassed"

	testrom(t, file, passstr)
}

func Test04(t *testing.T) {
	file := "../../../test/blargg/cpu_instrs/individual/04-op r,imm.gb"
	passstr := "04-op r,imm\n\n\nPassed"

	testrom(t, file, passstr)
}

func Test05(t *testing.T) {
	file := "../../../test/blargg/cpu_instrs/individual/05-op rp.gb"
	passstr := "05-op rp\n\n\nPassed"

	testrom(t, file, passstr)
}

func Test06(t *testing.T) {
	file := "../../../test/blargg/cpu_instrs/individual/06-ld r,r.gb"
	passstr := "06-ld r,r\n\n\nPassed"

	testrom(t, file, passstr)
}

func Test07(t *testing.T) {
	file := "../../../test/blargg/cpu_instrs/individual/07-jr,jp,call,ret,rst.gb"
	passstr := "07-jr,jp,call,ret,rst\n\n\nPassed"

	testrom(t, file, passstr)
}

func Test08(t *testing.T) {
	file := "../../../test/blargg/cpu_instrs/individual/08-misc instrs.gb"
	passstr := "08-misc instrs\n\n\nPassed"

	testrom(t, file, passstr)
}

func Test09(t *testing.T) {
	file := "../../../test/blargg/cpu_instrs/individual/09-op r,r.gb"
	passstr := "09-op r,r\n\n\nPassed"

	testrom(t, file, passstr)
}

func Test10(t *testing.T) {
	file := "../../../test/blargg/cpu_instrs/individual/10-bit ops.gb"
	passstr := "10-bit ops\n\n\nPassed"

	testrom(t, file, passstr)
}

func Test11(t *testing.T) {
	file := "../../../test/blargg/cpu_instrs/individual/11-op a,(hl).gb"
	passstr := "11-op a,(hl)\n\n\nPassed"

	testrom(t, file, passstr)
}

func TestTiming(t *testing.T) {
	file := "../../../test/blargg/instr_timing/instr_timing.gb"
	passstr := "instr_timing\n\n\nPassed"

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
