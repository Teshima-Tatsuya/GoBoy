package gb

import (
	"image"
	"time"

	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/apu"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/bus"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/cartridge"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/cpu"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/gpu"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/interrupt"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/io"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/memory"
)

type GB struct {
	Cartridge *cartridge.Cartridge

	cpu *cpu.CPU
	gpu *gpu.GPU
	apu *apu.APU

	timer *io.Timer

	currentCycle uint
}

func NewGB(romData []byte) *GB {
	cart := cartridge.New(romData)

	vram := memory.NewRAM(0x2000)
	wram := memory.NewRAM(0x2000)
	wram2 := memory.NewRAM(0x2000)
	hram := memory.NewRAM(0x0080)
	gpu := gpu.New()
	apu := apu.NewAPU()
	timer := io.NewTimer()
	irq := interrupt.NewIRQ()
	io := io.NewIO(io.NewPad(), io.NewSerial(), timer)
	bus := bus.New(cart, vram, wram, wram2, hram, apu, gpu, irq, io)

	cpu := cpu.New(bus, irq)
	timer.SetRequestIRQ(irq.Request)
	gpu.Init(bus, irq.Request)

	gb := &GB{
		Cartridge:    cart,
		cpu:          cpu,
		gpu:          gpu,
		apu:          apu,
		timer:        timer,
		currentCycle: 0,
	}

	return gb
}

func (gb *GB) Step() {
	time.Sleep(16 * time.Millisecond)
	for {

		var cycle uint
		if gb.gpu.IsDmaStarted() {
			gb.gpu.TransferOAM()
			cycle = 162
		} else {
			cycle = gb.cpu.Step()
		}
		gb.gpu.Step(cycle * 4)

		gb.currentCycle += cycle * 4

		gb.timer.Tick(cycle)

		if gb.currentCycle >= 70224 {
			gb.currentCycle -= 70224
			return
		}
	}
}

func (gb *GB) Display() (*image.RGBA, *image.RGBA) {
	return gb.gpu.Display()
}
