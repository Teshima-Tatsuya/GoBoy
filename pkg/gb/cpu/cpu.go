package cpu

import (
	"fmt"
	"os"

	"github.com/Teshima-Tatsuya/GoBoy/pkg/interfaces/bus"
	"github.com/apex/log"
)

type CPU struct {
	PC  uint16
	SP  uint16
	Reg Registers
	Bus bus.IO
}

func New(bus bus.IO) *CPU {
	return &CPU{
		PC: 0x100, // Gameboy Start Addr
		Reg: Registers{
			A: 0x01,
			B: 0x00,
			C: 0x00,
			D: 0xFF,
			E: 0x56,
			F: 0x80,
			H: 0x00,
			L: 0x0D,
		},
		Bus: bus,
	}
}

func (c *CPU) Step() int {
	opcode := c.Bus.ReadByte(c.PC)

	switch opcode {
	case 0x00:
		c.PC++
		return 1
	default:
		log.Info(fmt.Sprintf("Unsupported opcode 0x%02x at 0x%04x\n\n\n", opcode, c.PC))
		os.Exit(-1)
		c.PC++
	}

	return 0
}
