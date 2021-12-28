package cpu

import (
	"github.com/Teshima-Tatsuya/GoBoy/pkg/interfaces/bus"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/types"
)

type CPU struct {
	Reg Register
	Bus bus.IO
}

func New(bus bus.IO) *CPU {
	c := &CPU{
		Bus: bus,
	}

	c.Reg.reset()

	return c
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

	op.Handler(c, op.R1, op.R2)

	return 0
}

func (c *CPU) fetch() byte {
	d := c.Bus.ReadByte(c.Reg.PC)
	c.Reg.PC++
	return d
}

func (c *CPU) fetch16() types.Addr {
	lower := uint16(c.Bus.ReadByte(c.Reg.PC))
	upper := uint16(c.Bus.ReadByte(c.Reg.PC + 1))
	c.Reg.PC += 2

	return types.Addr((upper << 8) | lower)
}

func (c *CPU) pop() byte {
	d := c.Bus.ReadByte(c.Reg.SP)
	c.Reg.SP++
	return d
}

func (c *CPU) pop2PC() {
	lower := c.pop()
	upper := c.pop()

	c.Reg.PC = (types.Addr(upper) << 8) | types.Addr(lower)
}
