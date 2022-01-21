package emulator

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/stretchr/testify/assert"
)

func testrom(t assert.TestingT, file string, passstr string) {
	romData, err := ioutil.ReadFile(file)

	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}

	e := New(romData)

	if err := ebiten.RunGame(e); err != nil {
		log.Fatal(err)
	}
}

func TestCPU_Blargg_cpu_instrs(t *testing.T) {
	file := "../../test/blargg/cpu_instrs/cpu_instrs.gb"
	passstr := "cpu_instrs\n\n01:ok  02:ok  03:ok  04:ok  05:ok  06:ok  07:ok  08:ok  09:ok  10:ok  11:ok  \n\nPassed all tests\n"

	testrom(t, file, passstr)

}
