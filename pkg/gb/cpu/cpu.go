package cpu

import (
	"fmt"

	"github.com/Teshima-Tatsuya/GoBoy/pkg/interfaces/bus"
	"github.com/apex/log"
)

type CPU struct {
	// Program Counter
	PC uint16
	// Stack Pointer
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
	opcode := c.fetch()

	var op *OpCode
	if opcode == 0xCB {
		// Prefix Ope
		c.fetch()
		op = opCodes[opcode]
	} else {
		op = opCodes[opcode]
	}

	opereands := c.fetchOperands(op.Size)

	log.Info(fmt.Sprintf("opcode 0x%02x at 0x%04x\n\n\n", opcode, c.PC))
	op.Handler(c, opereands)

	return 0
}

func (c *CPU) fetch() byte {
	d := c.Bus.ReadByte(c.PC)
	c.PC++
	return d
}

func (c *CPU) fetchOperands(size uint8) []byte {
	ops := []byte{}
	for i := 0; i < int(size); i++ {
		ops = append(ops, c.fetch())
	}

	return ops
}
