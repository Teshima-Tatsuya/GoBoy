package gb

import (
	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/bus"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/cartridge"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/cpu"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/gpu"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/io"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/memory"
)

type GB struct {
	Cartridge *cartridge.Cartridge

	// memory
	vRAM *memory.RAM
	wRAM *memory.RAM
	hRAM *memory.RAM

	cpu *cpu.CPU
	gpu *gpu.GPU
}

func NewGB(romData []byte) *GB {
	cart := cartridge.New(romData)

	vram := memory.NewRAM(0x2000)
	wram := memory.NewRAM(0x2000)
	wram2 := memory.NewRAM(0x2000)
	hram := memory.NewRAM(0x0080)
	io := io.NewIO(io.NewPad(), io.NewSerial(), io.NewTimer(), io.NewIRQ(), gpu.New(), 0x2000)
	bus := bus.New(cart, vram, wram, wram2, hram, io)

	cpu := cpu.New(bus)
	cpu.Bus.IO.Timer.SetRequestIRQ(cpu.IRQ.Request)

	gb := &GB{
		Cartridge: cartridge.New(romData),
		vRAM:      memory.NewRAM(0x2000),
		wRAM:      memory.NewRAM(0x2000),
		hRAM:      memory.NewRAM(0x0080),

		cpu: cpu,
		gpu: gpu.New(),
	}

	return gb
}

func (gb *GB) Step() {
	cycle := gb.cpu.Step()
	gb.gpu.Step(cycle)
}
