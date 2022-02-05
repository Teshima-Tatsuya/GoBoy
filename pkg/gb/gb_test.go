package gb

import (
	"image/jpeg"
	"io/ioutil"
	"os"
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
		{"blargg/cpu_instrs", "cpu_instrs", 3200},
		{"blargg/instr_timing", "instr_timing", 100},
		{"blargg/interrupt_time", "interrupt_time", 100},
		{"helloworld", "hello", 100},
		{"mooneye-gb/acceptance/bits", "mem_oam", 100},
		{"mooneye-gb/acceptance/bits", "reg_f", 100},
		{"mooneye-gb/acceptance/bits", "unused_hwio-GS", 100},
		{"mooneye-gb/acceptance/instr", "daa", 100},
		{"mooneye-gb/acceptance/interrupts", "ie_push", 100},
		{"mooneye-gb/acceptance/oam_dma", "basic", 100},
		{"mooneye-gb/acceptance/oam_dma", "reg_read", 100},
		{"mooneye-gb/acceptance/oam_dma", "sources-GS", 100},
		{"mooneye-gb/acceptance/ppu", "hblank_ly_scx_timing-GS", 100},
		{"mooneye-gb/acceptance/ppu", "intr_1_2_timing", 100},
		{"mooneye-gb/acceptance/ppu", "intr_2_0_timing", 100},
		{"mooneye-gb/acceptance/ppu", "intr_2_mode0_timing_sprites", 100},
		{"mooneye-gb/acceptance/ppu", "intr_2_mode0_timing", 100},
		{"mooneye-gb/acceptance/ppu", "intr_2_mode3_timing", 100},
		{"mooneye-gb/acceptance/ppu", "intr_2_oam_ok_timing", 100},
		{"mooneye-gb/acceptance/ppu", "stat_irq_blocking", 100},
		{"mooneye-gb/acceptance/ppu", "stat_lyc_onoff", 100},
		{"mooneye-gb/acceptance/ppu", "vblank_stat_inft-GS", 100},
		{"mooneye-gb/acceptance/serial", "boot_sclk_align-dmgABCmgb", 100},
		{"mooneye-gb/acceptance/timer", "div_write", 100},
		{"mooneye-gb/acceptance/timer", "rapid_toggle", 100},
		{"mooneye-gb/acceptance/timer", "tim00_div_trigger", 100},
		{"mooneye-gb/acceptance/timer", "tim00", 100},
		{"mooneye-gb/acceptance/timer", "tim01_div_trigger", 100},
		{"mooneye-gb/acceptance/timer", "tim01", 100},
		{"mooneye-gb/acceptance/timer", "tim10_div_trigger", 100},
		{"mooneye-gb/acceptance/timer", "tim10", 100},
		{"mooneye-gb/acceptance/timer", "tim11_div_trigger", 100},
		{"mooneye-gb/acceptance/timer", "tim11", 100},
		{"mooneye-gb/acceptance/timer", "tima_reload", 100},
		{"mooneye-gb/acceptance/timer", "tima_write_reloading", 100},
		{"mooneye-gb/acceptance/timer", "tma_write_reloading", 100},
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
			romData, err := ioutil.ReadFile(testPath + tt.name + "/" + tt.file + ".gb")
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

			file, _ := os.Create(outPath + tt.name + "/" + tt.file + ".jpg")
			defer file.Close()

			if err := jpeg.Encode(file, screen, &jpeg.Options{100}); err != nil {
				panic(err)
			}
		})
	}
}
