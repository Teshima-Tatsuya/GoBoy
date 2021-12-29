package cpu

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/bus"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/cartridge"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/ram"
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
	bus := bus.New(cart, vram, wram, wram2, hram)

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
