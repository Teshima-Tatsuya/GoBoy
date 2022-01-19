package gb

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testrom(t assert.TestingT, file string, passstr string) {
	romData, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}

	gb := NewGB(romData)

	var str string

	str = ""

	for i := 0; i < 40000000; i++ {
		if gb.cpu.Bus.ReadByte(0xff02) == byte(0x81) {
			d := gb.cpu.Bus.ReadByte(0xff01)
			str += string(d)
			gb.cpu.Bus.WriteByte(0xff02, byte(0x00))
		}

		gb.cpu.Step()
	}

	assert.Equal(t, passstr, str)
}

func TestCPU_Blargg_cpu_instrs(t *testing.T) {
	file := "../../test/blargg/cpu_instrs/cpu_instrs.gb"
	passstr := "cpu_instrs\n\n01:ok  02:ok  03:ok  04:ok  05:ok  06:ok  07:ok  08:ok  09:ok  10:ok  11:ok  \n\nPassed all tests\n"

	testrom(t, file, passstr)

}
