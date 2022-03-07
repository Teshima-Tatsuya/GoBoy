package cpu

import (
	"testing"

	"github.com/Teshima-Tatsuya/GoBoy/pkg/types"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/util"
	"github.com/Teshima-Tatsuya/GoBoy/test/mock"
	"github.com/stretchr/testify/assert"
)

func setupCPU() *CPU {
	irq := io.NewIRQ()
	bus := &mock.MockBus{}

	return New(bus, irq)
}

func TestCPU_fetch(t *testing.T) {
	c := setupCPU()
	assert := assert.New(t)

	want := byte(0x12)
	c.Bus.WriteByte(c.Reg.PC, want)

	d := c.fetch()
	assert.Equal(want, d)
	assert.Equal(types.Addr(0x101), c.Reg.PC)
}

func TestCPU_fetch16(t *testing.T) {
	c := setupCPU()
	assert := assert.New(t)

	lower := byte(0x12)
	upper := byte(0x13)
	want := types.Addr(0x1312)
	c.Bus.WriteByte(c.Reg.PC, lower)
	c.Bus.WriteByte(c.Reg.PC+1, upper)

	d := c.fetch16()
	assert.Equal(want, d)
	assert.Equal(types.Addr(0x102), c.Reg.PC)
}

func TestCPU_pushpop(t *testing.T) {
	c := setupCPU()
	assert := assert.New(t)

	want := byte(0x01)

	before_sp := c.Reg.SP
	c.push(want)
	assert.Equal(before_sp-1, c.Reg.SP)
	assert.Equal(want, c.Bus.ReadByte(c.Reg.SP))

	d := c.pop()
	assert.Equal(before_sp, c.Reg.SP)
	assert.Equal(want, d)
}

func TestCPU_pushpopPC(t *testing.T) {
	c := setupCPU()
	assert := assert.New(t)

	want := types.Addr(0x1234)
	c.Reg.PC = want

	before_sp := c.Reg.SP
	c.pushPC()
	assert.Equal(util.ExtractLower(want), c.Bus.ReadByte(before_sp-2))
	assert.Equal(util.ExtractUpper(want), c.Bus.ReadByte(before_sp-1))

	c.popPC()
	assert.Equal(want, c.Reg.PC)
	assert.Equal(before_sp, c.Reg.SP)
}
