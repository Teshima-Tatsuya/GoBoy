package video

type LCDS struct {
	Data byte
}

func NewLCDS(data byte) *LCDS {
	return &LCDS{
		Data: data,
	}
}
