package gpu

import "github.com/Teshima-Tatsuya/GoBoy/pkg/types"

type Scroll struct {
	// FF42, FF43
	SCY, SCX byte
	/*
	  FF44
	   LY indicates the current horizontal line
	   LY can hold any value from 0 to 153,
	   with values from 144 to 153 indicating the VBlank period.
	    if LY >= 154 LY = 0
	*/
	LY byte

	/*
	  FF45
	   The Game Boy permanently compares the value of the LYC and LY registers.
	   When both values are identical, the “LYC=LY” flag in the STAT register is
	   set, and (if enabled) a STAT interrupt is requested.
	*/
	LYC byte
	// FF4A, FF4B
	WX, WY byte
}

func NewScroll() *Scroll {
	return &Scroll{}
}

func (s *Scroll) isVBlankPeriod() bool {
	if SCREEN_HEIGHT <= s.LY && s.LY <= 153 {
		return true
	}

	return false
}

func (s *Scroll) isHBlankPeriod() bool {
	if s.LY < SCREEN_HEIGHT {
		return true
	}

	return false
}

func (s *Scroll) isVBlankStart() bool {
	return s.LY == SCREEN_HEIGHT
}

func (s *Scroll) Read(addr types.Addr) byte {
	switch addr {
	case SCYAddr:
		return s.SCY
	case SCXAddr:
		return s.SCX
	case LYAddr:
		return s.LY
	case LYCAddr:
		return s.LYC
	case WXAddr:
		return s.WX
	case WYAddr:
		return s.WY
	default:
		panic("Scroll Read")
	}
}

func (s *Scroll) Write(addr types.Addr, value byte) {
	switch addr {
	case SCYAddr:
		s.SCY = value
	case SCXAddr:
		s.SCX = value
	case LYAddr:
		s.LY = value
	case LYCAddr:
		s.LYC = value
	case WXAddr:
		s.WX = value
	case WYAddr:
		s.WY = value
	default:
		panic("Scroll Write")
	}
}
