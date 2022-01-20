package gpu

type Scroll struct {
	// FF42, FF43
	SCY, SCX byte
	/*
	  FF44
	   LY indicates the current horizontal line
	   LY can hold any value from 0 to 153,
	   with values from 144 to 153 indicating the VBlank period.
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
