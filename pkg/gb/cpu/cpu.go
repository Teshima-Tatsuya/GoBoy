package cpu

import (
	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/bus"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/io"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/types"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/util"
)

type CPU struct {
	Reg  Register
	Bus  *bus.Bus
	IRQ  *io.IRQ
	Halt bool
}

func New(bus *bus.Bus) *CPU {
	c := &CPU{
		Bus:  bus,
		IRQ:  bus.IO.IRQ,
		Halt: false,
	}

	c.Reg.reset()

	return c
}

func (c *CPU) Step() int {
	if c.Halt {
		if c.IRQ.Has() {
			c.Halt = false
		}

		return 1
	}

	if c.interrupt() {
		return 1
	}
	opcode := c.fetch()

	var op *OpCode

	if opcode == 0xCB {
		opcode = c.fetch()
		op = cbOpCodes[opcode]
	} else {
		op = opCodes[opcode]
	}

	// log.Info(fmt.Sprintf("PC 0x%04X data 0x%02x%02x", c.Reg.PC-1, c.Bus.ReadByte(c.Reg.PC), c.Bus.ReadByte(c.Reg.PC+1)))
	// log.Info(fmt.Sprintf(" %s", op.Mnemonic))
	op.Handler(c, op.R1, op.R2)
	return 0
}

func (c *CPU) fetch() byte {
	d := c.Bus.ReadByte(c.Reg.PC)
	c.Reg.PC++
	return d
}

func (c *CPU) fetch16() types.Addr {
	lower := c.fetch()
	upper := c.fetch()

	return util.Byte2Addr(upper, lower)
}

func (c *CPU) push(buf byte) {
	c.Reg.SP--
	c.Bus.WriteByte(c.Reg.SP, buf)
}

// push PC
func (c *CPU) pushPC() {
	c.push(util.ExtractUpper(c.Reg.PC))
	c.push(util.ExtractLower(c.Reg.PC))
}

func (c *CPU) pop() byte {
	d := c.Bus.ReadByte(c.Reg.SP)
	c.Reg.SP++
	return d
}

func (c *CPU) popPC() {
	lower := c.pop()
	upper := c.pop()

	c.Reg.PC = util.Byte2Addr(upper, lower)
}

func (c *CPU) getRegMem(r16 int) byte {
	r := c.Reg.R16(r16)
	d := c.Bus.ReadByte(r)

	return d
}

func (c *CPU) interrupt() bool {
	if !c.IRQ.Enabled() || !c.IRQ.Has() {
		return false
	}

	addr := c.IRQ.InterruptAddr()
	c.pushPC()
	c.Reg.PC = addr
	c.IRQ.Disable()

	return true
}
