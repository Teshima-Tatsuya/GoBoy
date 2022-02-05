package gb

import (
	"image/jpeg"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/Teshima-Tatsuya/GoBoy/pkg/debug"
)

func setup(romData []byte) *GB {
	return NewGB(romData)
}

func TestGB_testroms(t *testing.T) {
	tests := []struct {
		name  string
		file  string
		frame int
	}{
		{"blargg/cpu_instrs", "cpu_instrs.gb", 3200},
		{"blargg/instr_timing", "instr_timing.gb", 100},
		{"blargg/interrupt_time", "interrupt_time.gb", 100},
		{"helloworld", "hello.gb", 100},
		{"mooneye-gb/acceptance/bits", "mem_oam.gb", 100},
		{"mooneye-gb/acceptance/bits", "reg_f.gb", 100},
		{"mooneye-gb/acceptance/bits", "unused_hwio-GS.gb", 100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				err := recover()
				if err != nil {
					debug.Info("%s", err)
				}
			}()

			testPath := "../../test/rom/"
			outPath := "../../test/actual/"
			romData, err := ioutil.ReadFile(testPath + tt.name + "/" + tt.file)
			if err != nil {
				panic(err)
			}
			gb := setup(romData)

			for i := 0; i < tt.frame; i++ {
				gb.Step()
			}
			screen, _ := gb.Display()

			if err := os.MkdirAll(outPath+tt.name, 0777); err != nil {
				panic(err)
			}

			file, _ := os.Create(outPath + tt.name + "/" + strings.Replace(tt.file, ".gb", "", -1) + ".jpg")
			defer file.Close()

			if err := jpeg.Encode(file, screen, &jpeg.Options{100}); err != nil {
				panic(err)
			}
		})
	}
}
