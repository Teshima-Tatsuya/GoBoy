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

func test(name, filename string, frame int) {
	testPath := "../../test/rom/"
	outPath := "../../test/actual/"
	romData, err := ioutil.ReadFile(testPath + name + "/" + filename + ".gb")
	if err != nil {
		panic(err)
	}
	gb := setup(romData)

	for i := 0; i < frame; i++ {
		gb.Step()
	}
	screen, _ := gb.Display()

	if err := os.MkdirAll(outPath+name, 0777); err != nil {
		panic(err)
	}

	file, _ := os.Create(outPath + name + "/" + filename + ".jpg")
	defer file.Close()

	if err := jpeg.Encode(file, screen, &jpeg.Options{100}); err != nil {
		panic(err)
	}
}

func TestGB_test_blargg(t *testing.T) {
	tests := []struct {
		name  string
		file  string
		frame int
	}{
		{"blargg/cpu_instrs", "cpu_instrs", 3200},
		{"blargg/instr_timing", "instr_timing", 100},
		{"blargg/interrupt_time", "interrupt_time", 100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				err := recover()
				if err != nil {
					debug.Info("%s", err)
				}
			}()

			test(tt.name, tt.file, tt.frame)
		})
	}
}
func TestGB_test_hello(t *testing.T) {
	tests := []struct {
		name  string
		file  string
		frame int
	}{
		{"helloworld", "hello", 100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				err := recover()
				if err != nil {
					debug.Info("%s", err)
				}
			}()

			test(tt.name, tt.file, tt.frame)
		})
	}
}

func TestGB_test_mooneye(t *testing.T) {
	tests := []struct {
		name  string
		file  string
		frame int
	}{
		{"mooneye-gb/acceptance/bits", "mem_oam", 10},
		{"mooneye-gb/acceptance/bits", "reg_f", 10},
		{"mooneye-gb/acceptance/bits", "unused_hwio-GS", 10},
		{"mooneye-gb/acceptance/instr", "daa", 100},
		{"mooneye-gb/acceptance/interrupts", "ie_push", 100},
		{"mooneye-gb/acceptance/oam_dma", "basic", 100},
		{"mooneye-gb/acceptance/oam_dma", "reg_read", 100},
		{"mooneye-gb/acceptance/oam_dma", "sources-GS", 100},
		{"mooneye-gb/acceptance/ppu", "hblank_ly_scx_timing-GS", 100},
		{"mooneye-gb/acceptance/ppu", "intr_1_2_timing-GS", 100},
		{"mooneye-gb/acceptance/ppu", "intr_2_0_timing", 100},
		{"mooneye-gb/acceptance/ppu", "intr_2_mode0_timing_sprites", 100},
		{"mooneye-gb/acceptance/ppu", "intr_2_mode0_timing", 100},
		{"mooneye-gb/acceptance/ppu", "intr_2_mode3_timing", 100},
		{"mooneye-gb/acceptance/ppu", "intr_2_oam_ok_timing", 100},
		{"mooneye-gb/acceptance/ppu", "stat_irq_blocking", 100},
		{"mooneye-gb/acceptance/ppu", "stat_lyc_onoff", 100},
		{"mooneye-gb/acceptance/ppu", "vblank_stat_intr-GS", 100},
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
		{"mooneye-gb/acceptance", "add_sp_e_timing", 100},
		{"mooneye-gb/acceptance", "boot_div-dmg0", 100},
		{"mooneye-gb/acceptance", "boot_div-dmgABCmgb", 100},
		{"mooneye-gb/acceptance", "boot_div-S", 100},
		{"mooneye-gb/acceptance", "boot_div2-S", 100},
		{"mooneye-gb/acceptance", "boot_hwio-dmg0", 100},
		{"mooneye-gb/acceptance", "boot_hwio-dmgABCmgb", 100},
		{"mooneye-gb/acceptance", "boot_hwio-S", 100},
		{"mooneye-gb/acceptance", "boot_regs-dmg0", 100},
		{"mooneye-gb/acceptance", "boot_regs-dmgABC", 100},
		{"mooneye-gb/acceptance", "boot_regs-mgb", 100},
		{"mooneye-gb/acceptance", "boot_regs-sgb", 100},
		{"mooneye-gb/acceptance", "boot_regs-sgb2", 100},
		{"mooneye-gb/acceptance", "call_cc_timing", 100},
		{"mooneye-gb/acceptance", "call_cc_timing2", 100},
		{"mooneye-gb/acceptance", "call_timing", 100},
		{"mooneye-gb/acceptance", "call_cc_timing2", 100},
		{"mooneye-gb/acceptance", "di_timing-GS", 100},
		{"mooneye-gb/acceptance", "div_timing", 100},
		{"mooneye-gb/acceptance", "ei_sequence", 100},
		{"mooneye-gb/acceptance", "ei_timing", 100},
		{"mooneye-gb/acceptance", "halt_ime0_ei", 100},
		{"mooneye-gb/acceptance", "halt_ime0_nointr_timing", 100},
		{"mooneye-gb/acceptance", "halt_ime1_timing", 100},
		{"mooneye-gb/acceptance", "halt_ime1_timing2-GS", 100},
		{"mooneye-gb/acceptance", "if_ie_registers", 100},
		{"mooneye-gb/acceptance", "intr_timing", 100},
		{"mooneye-gb/acceptance", "jp_cc_timing", 100},
		{"mooneye-gb/acceptance", "jp_timing", 100},
		{"mooneye-gb/acceptance", "ld_hl_sp_e_timing", 100},
		{"mooneye-gb/acceptance", "ld_hl_sp_e_timing", 100},
		{"mooneye-gb/acceptance", "oam_dma_restart", 100},
		{"mooneye-gb/acceptance", "oam_dma_start", 100},
		{"mooneye-gb/acceptance", "oam_dma_timing", 100},
		{"mooneye-gb/acceptance", "pop_timing", 100},
		{"mooneye-gb/acceptance", "push_timing", 100},
		{"mooneye-gb/acceptance", "rapid_di_ei", 100},
		{"mooneye-gb/acceptance", "ret_cc_timing", 100},
		{"mooneye-gb/acceptance", "ret_timing", 100},
		{"mooneye-gb/acceptance", "reti_intr_timing", 100},
		{"mooneye-gb/acceptance", "reti_timing", 100},
		{"mooneye-gb/acceptance", "rst_timing", 20},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				err := recover()
				if err != nil {
					debug.Info("%s", err)
				}
			}()

			test(tt.name, tt.file, tt.frame)
		})
	}
}

func TestGB_test_temp(t *testing.T) {
	tests := []struct {
		name  string
		file  string
		frame int
	}{
		{"mooneye-gb/acceptance/ppu", "intr_2_0_timing", 2000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test(tt.name, tt.file, tt.frame)
		})
	}
}
