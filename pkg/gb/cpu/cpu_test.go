package cpu

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/bus"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/cartridge"
)

func setup() *bus.Bus {
	file := "../../../test/blargg-gb-tests/cpu_instrs/individual/01-special.gb"
	romData, err := ioutil.ReadFile(file)

	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}

	cart := cartridge.New(romData)
	bus := bus.New(cart)

	return bus
}

func Test(t *testing.T) {
	bus := setup()
	cpu := New(bus)

	for {
		cpu.Step()
	}
}
