package gpu

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
	if s.LY >= SCREEN_HEIGHT {
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
