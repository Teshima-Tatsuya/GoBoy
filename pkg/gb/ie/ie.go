package ie

type IE struct {
	buf byte
}

func New() *IE {
	buf := byte(0x00)

	return &IE{
		buf: buf,
	}

}

func (r *IE) Read() byte {
	return r.buf
}

func (r *IE) Write(value byte) {
	r.buf = value
}
