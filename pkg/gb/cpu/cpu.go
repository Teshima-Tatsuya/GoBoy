package cpu

import (
	"fmt"

	"github.com/Teshima-Tatsuya/GoBoy/pkg/interfaces/bus"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/types"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/util"
	"github.com/apex/log"
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
		log.Info("PREFIX CB")
		opcode = c.fetch()
		op = cbOpCodes[opcode]
	} else {
		op = opCodes[opcode]
	}

	log.Info(fmt.Sprintf("0x%02X:\"%s\" PC 0x%04X, SP 0x%04X, 0x%02X 0x%02X", op.Code, op.Mnemonic, c.Reg.PC-1, c.Reg.SP, c.Bus.ReadByte(c.Reg.PC), c.Bus.ReadByte(c.Reg.PC+1)))
	//	log.Info(fmt.Sprintf("    (BC):0x%02X (DE):0x%02X (HL):0x%02X (AF):0x%02X", c.Bus.ReadByte(c.Reg.R16(BC)), c.Bus.ReadByte(c.Reg.R16(DE)), c.Bus.ReadByte(c.Reg.R16(HL)), c.Bus.ReadByte(c.Reg.R16(AF))))
	c.Reg.Dump()
	fmt.Println()
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

	c.Reg.PC = (types.Addr(upper) << 8) | types.Addr(lower)
}
