package cpu

import (
	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/irq"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/interfaces/bus"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/types"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/util"
)

type CPU struct {
	Reg Register
	Bus bus.IO
	IRQ *irq.IRQ
}

func New(bus bus.IO, irq *irq.IRQ) *CPU {
	c := &CPU{
		Bus: bus,
		IRQ: irq,
	}

	c.Reg.reset()

	return c
}

func (c *CPU) Step() int {
	opcode := c.fetch()

	var op *OpCode

	if opcode == 0xCB {
		opcode = c.fetch()
		op = cbOpCodes[opcode]
	} else {
		op = opCodes[opcode]
	}

	// log.Info(fmt.Sprintf("0x%02X:\"%s\" PC 0x%04X, SP 0x%04X, 0x%02X 0x%02X", op.Code, op.Mnemonic, c.Reg.PC-1, c.Reg.SP, c.Bus.ReadByte(c.Reg.PC), c.Bus.ReadByte(c.Reg.PC+1)))
	// if strings.Contains(op.Mnemonic, "(BC") {
	// 	log.Info(fmt.Sprintf(" (BC) = %02X", c.getRegMem(BC)))
	// } else if strings.Contains(op.Mnemonic, "(DE)") {
	// 	log.Info(fmt.Sprintf(" (DE) = %02X", c.getRegMem(DE)))
	// } else if strings.Contains(op.Mnemonic, "(HL") {
	// 	log.Info(fmt.Sprintf(" (HL) = %02X", c.getRegMem(HL)))
	// } else if strings.Contains(op.Mnemonic, "(AF") {
	// 	log.Info(fmt.Sprintf(" (AF) = %02X", c.getRegMem(AF)))
	// } else if strings.Contains(op.Mnemonic, "(a16") {
	// 	lower := c.Bus.ReadByte(c.Reg.PC)
	// 	upper := c.Bus.ReadByte(c.Reg.PC + 1)
	// 	log.Info(fmt.Sprintf(" (a16) = %02X", c.Bus.ReadByte(util.Byte2Addr(upper, lower))))
	// } else if strings.Contains(op.Mnemonic, "(a8") {
	// 	lower := c.Bus.ReadByte(c.Reg.PC)
	// 	upper := byte(0xFF)
	// 	log.Info(fmt.Sprintf(" (a8) = %02X", c.Bus.ReadByte(util.Byte2Addr(upper, lower))))
	// }

	// c.Reg.Dump()
	op.Handler(c, op.R1, op.R2)
	// if strings.Contains(op.Mnemonic, "(BC") {
	// 	log.Info(fmt.Sprintf(" (BC) = %02X", c.getRegMem(BC)))
	// } else if strings.Contains(op.Mnemonic, "(DE)") {
	// 	log.Info(fmt.Sprintf(" (DE) = %02X", c.getRegMem(DE)))
	// } else if strings.Contains(op.Mnemonic, "(HL+") {
	// 	log.Info(fmt.Sprintf(" (HL) = %02X", c.Bus.ReadByte(c.Reg.R16(HL)-1)))
	// } else if strings.Contains(op.Mnemonic, "(HL-") {
	// 	log.Info(fmt.Sprintf(" (HL) = %02X", c.Bus.ReadByte(c.Reg.R16(HL)+1)))
	// } else if strings.Contains(op.Mnemonic, "(HL") {
	// 	log.Info(fmt.Sprintf(" (HL) = %02X", c.getRegMem(HL)))
	// } else if strings.Contains(op.Mnemonic, "(AF") {
	// 	log.Info(fmt.Sprintf(" (AF) = %02X", c.getRegMem(AF)))
	// } else if strings.Contains(op.Mnemonic, "(a16") {
	// 	lower := c.Bus.ReadByte(c.Reg.PC - 2)
	// 	upper := c.Bus.ReadByte(c.Reg.PC - 1)
	// 	log.Info(fmt.Sprintf(" (a16) = %02X", c.Bus.ReadByte(util.Byte2Addr(upper, lower))))
	// } else if strings.Contains(op.Mnemonic, "(a8") {
	// 	lower := c.Bus.ReadByte(c.Reg.PC - 1)
	// 	upper := byte(0xFF)
	// 	log.Info(fmt.Sprintf(" (a8) = %02X", c.Bus.ReadByte(util.Byte2Addr(upper, lower))))
	// }

	// fmt.Println()

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
