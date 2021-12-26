package cpu

import "os"

type OpCode struct {
	Code    byte
	Size    uint8
	Cycles  uint8
	Handler func(*CPU, []byte)
}

var opCodes = []*OpCode{
	{0x00, 0, 1, nop},
	{0x01, 0, 1, notimplementd},
	{0x02, 0, 1, notimplementd},
	{0x03, 0, 1, notimplementd},
	{0x04, 0, 1, notimplementd},
	{0x05, 0, 1, notimplementd},
	{0x06, 0, 1, notimplementd},
	{0x07, 0, 1, notimplementd},
	{0x08, 0, 1, notimplementd},
	{0x09, 0, 1, notimplementd},
	{0x0A, 0, 1, notimplementd},
	{0x0B, 0, 1, notimplementd},
	{0x0C, 0, 1, notimplementd},
	{0x0D, 0, 1, notimplementd},
	{0x0E, 0, 1, notimplementd},
	{0x0F, 0, 1, notimplementd},
	{0x10, 0, 1, notimplementd},
	{0x11, 0, 1, notimplementd},
	{0x12, 0, 1, notimplementd},
	{0x13, 0, 1, notimplementd},
	{0x14, 0, 1, notimplementd},
	{0x15, 0, 1, notimplementd},
	{0x16, 0, 1, notimplementd},
	{0x17, 0, 1, notimplementd},
	{0x18, 0, 1, notimplementd},
	{0x19, 0, 1, notimplementd},
	{0x1A, 0, 1, notimplementd},
	{0x1B, 0, 1, notimplementd},
	{0x1C, 0, 1, notimplementd},
	{0x1D, 0, 1, notimplementd},
	{0x1E, 0, 1, notimplementd},
	{0x1F, 0, 1, notimplementd},
	{0x20, 0, 1, notimplementd},
	{0x21, 0, 1, notimplementd},
	{0x22, 0, 1, notimplementd},
	{0x23, 0, 1, notimplementd},
	{0x24, 0, 1, notimplementd},
	{0x25, 0, 1, notimplementd},
	{0x26, 0, 1, notimplementd},
	{0x27, 0, 1, notimplementd},
	{0x28, 0, 1, notimplementd},
	{0x29, 0, 1, notimplementd},
	{0x2A, 0, 1, notimplementd},
	{0x2B, 0, 1, notimplementd},
	{0x2C, 0, 1, notimplementd},
	{0x2D, 0, 1, notimplementd},
	{0x2E, 0, 1, notimplementd},
	{0x2F, 0, 1, notimplementd},
	{0x30, 0, 1, notimplementd},
	{0x31, 0, 1, notimplementd},
	{0x32, 0, 1, notimplementd},
	{0x33, 0, 1, notimplementd},
	{0x34, 0, 1, notimplementd},
	{0x35, 0, 1, notimplementd},
	{0x36, 0, 1, notimplementd},
	{0x37, 0, 1, notimplementd},
	{0x38, 0, 1, notimplementd},
	{0x39, 0, 1, notimplementd},
	{0x3A, 0, 1, notimplementd},
	{0x3B, 0, 1, notimplementd},
	{0x3C, 0, 1, notimplementd},
	{0x3D, 0, 1, notimplementd},
	{0x3E, 0, 1, notimplementd},
	{0x3F, 0, 1, notimplementd},
	{0x40, 0, 1, notimplementd},
	{0x41, 0, 1, notimplementd},
	{0x42, 0, 1, notimplementd},
	{0x43, 0, 1, notimplementd},
	{0x44, 0, 1, notimplementd},
	{0x45, 0, 1, notimplementd},
	{0x46, 0, 1, notimplementd},
	{0x47, 0, 1, notimplementd},
	{0x48, 0, 1, notimplementd},
	{0x49, 0, 1, notimplementd},
	{0x4A, 0, 1, notimplementd},
	{0x4B, 0, 1, notimplementd},
	{0x4C, 0, 1, notimplementd},
	{0x4D, 0, 1, notimplementd},
	{0x4E, 0, 1, notimplementd},
	{0x4F, 0, 1, notimplementd},
	{0x50, 0, 1, notimplementd},
	{0x51, 0, 1, notimplementd},
	{0x52, 0, 1, notimplementd},
	{0x53, 0, 1, notimplementd},
	{0x54, 0, 1, notimplementd},
	{0x55, 0, 1, notimplementd},
	{0x56, 0, 1, notimplementd},
	{0x57, 0, 1, notimplementd},
	{0x58, 0, 1, notimplementd},
	{0x59, 0, 1, notimplementd},
	{0x5A, 0, 1, notimplementd},
	{0x5B, 0, 1, notimplementd},
	{0x5C, 0, 1, notimplementd},
	{0x5D, 0, 1, notimplementd},
	{0x5E, 0, 1, notimplementd},
	{0x5F, 0, 1, notimplementd},
	{0x60, 0, 1, notimplementd},
	{0x61, 0, 1, notimplementd},
	{0x62, 0, 1, notimplementd},
	{0x63, 0, 1, notimplementd},
	{0x64, 0, 1, notimplementd},
	{0x65, 0, 1, notimplementd},
	{0x66, 0, 1, notimplementd},
	{0x67, 0, 1, notimplementd},
	{0x68, 0, 1, notimplementd},
	{0x69, 0, 1, notimplementd},
	{0x6A, 0, 1, notimplementd},
	{0x6B, 0, 1, notimplementd},
	{0x6C, 0, 1, notimplementd},
	{0x6D, 0, 1, notimplementd},
	{0x6E, 0, 1, notimplementd},
	{0x6F, 0, 1, notimplementd},
	{0x70, 0, 1, notimplementd},
	{0x71, 0, 1, notimplementd},
	{0x72, 0, 1, notimplementd},
	{0x73, 0, 1, notimplementd},
	{0x74, 0, 1, notimplementd},
	{0x75, 0, 1, notimplementd},
	{0x76, 0, 1, notimplementd},
	{0x77, 0, 1, notimplementd},
	{0x78, 0, 1, notimplementd},
	{0x79, 0, 1, notimplementd},
	{0x7A, 0, 1, notimplementd},
	{0x7B, 0, 1, notimplementd},
	{0x7C, 0, 1, notimplementd},
	{0x7D, 0, 1, notimplementd},
	{0x7E, 0, 1, notimplementd},
	{0x7F, 0, 1, notimplementd},
	{0x80, 0, 1, notimplementd},
	{0x81, 0, 1, notimplementd},
	{0x82, 0, 1, notimplementd},
	{0x83, 0, 1, notimplementd},
	{0x84, 0, 1, notimplementd},
	{0x85, 0, 1, notimplementd},
	{0x86, 0, 1, notimplementd},
	{0x87, 0, 1, notimplementd},
	{0x88, 0, 1, notimplementd},
	{0x89, 0, 1, notimplementd},
	{0x8A, 0, 1, notimplementd},
	{0x8B, 0, 1, notimplementd},
	{0x8C, 0, 1, notimplementd},
	{0x8D, 0, 1, notimplementd},
	{0x8E, 0, 1, notimplementd},
	{0x8F, 0, 1, notimplementd},
	{0x90, 0, 1, notimplementd},
	{0x91, 0, 1, notimplementd},
	{0x92, 0, 1, notimplementd},
	{0x93, 0, 1, notimplementd},
	{0x94, 0, 1, notimplementd},
	{0x95, 0, 1, notimplementd},
	{0x96, 0, 1, notimplementd},
	{0x97, 0, 1, notimplementd},
	{0x98, 0, 1, notimplementd},
	{0x99, 0, 1, notimplementd},
	{0x9A, 0, 1, notimplementd},
	{0x9B, 0, 1, notimplementd},
	{0x9C, 0, 1, notimplementd},
	{0x9D, 0, 1, notimplementd},
	{0x9E, 0, 1, notimplementd},
	{0x9F, 0, 1, notimplementd},
	{0xA0, 0, 1, notimplementd},
	{0xA1, 0, 1, notimplementd},
	{0xA2, 0, 1, notimplementd},
	{0xA3, 0, 1, notimplementd},
	{0xA4, 0, 1, notimplementd},
	{0xA5, 0, 1, notimplementd},
	{0xA6, 0, 1, notimplementd},
	{0xA7, 0, 1, notimplementd},
	{0xA8, 0, 1, notimplementd},
	{0xA9, 0, 1, notimplementd},
	{0xAA, 0, 1, notimplementd},
	{0xAB, 0, 1, notimplementd},
	{0xAC, 0, 1, notimplementd},
	{0xAD, 0, 1, notimplementd},
	{0xAE, 0, 1, notimplementd},
	{0xAF, 0, 1, notimplementd},
	{0xB0, 0, 1, notimplementd},
	{0xB1, 0, 1, notimplementd},
	{0xB2, 0, 1, notimplementd},
	{0xB3, 0, 1, notimplementd},
	{0xB4, 0, 1, notimplementd},
	{0xB5, 0, 1, notimplementd},
	{0xB6, 0, 1, notimplementd},
	{0xB7, 0, 1, notimplementd},
	{0xB8, 0, 1, notimplementd},
	{0xB9, 0, 1, notimplementd},
	{0xBA, 0, 1, notimplementd},
	{0xBB, 0, 1, notimplementd},
	{0xBC, 0, 1, notimplementd},
	{0xBD, 0, 1, notimplementd},
	{0xBE, 0, 1, notimplementd},
	{0xBF, 0, 1, notimplementd},
	{0xC0, 0, 1, notimplementd},
	{0xC1, 0, 1, notimplementd},
	{0xC2, 0, 1, notimplementd},
	{0xC3, 0, 1, notimplementd},
	{0xC4, 0, 1, notimplementd},
	{0xC5, 0, 1, notimplementd},
	{0xC6, 0, 1, notimplementd},
	{0xC7, 0, 1, notimplementd},
	{0xC8, 0, 1, notimplementd},
	{0xC9, 0, 1, notimplementd},
	{0xCA, 0, 1, notimplementd},
	{0xCB, 0, 1, notimplementd},
	{0xCC, 0, 1, notimplementd},
	{0xCD, 0, 1, notimplementd},
	{0xCE, 0, 1, notimplementd},
	{0xCF, 0, 1, notimplementd},
	{0xD0, 0, 1, notimplementd},
	{0xD1, 0, 1, notimplementd},
	{0xD2, 0, 1, notimplementd},
	{0xD3, 0, 1, notimplementd},
	{0xD4, 0, 1, notimplementd},
	{0xD5, 0, 1, notimplementd},
	{0xD6, 0, 1, notimplementd},
	{0xD7, 0, 1, notimplementd},
	{0xD8, 0, 1, notimplementd},
	{0xD9, 0, 1, notimplementd},
	{0xDA, 0, 1, notimplementd},
	{0xDB, 0, 1, notimplementd},
	{0xDC, 0, 1, notimplementd},
	{0xDD, 0, 1, notimplementd},
	{0xDE, 0, 1, notimplementd},
	{0xDF, 0, 1, notimplementd},
	{0xE0, 0, 1, notimplementd},
	{0xE1, 0, 1, notimplementd},
	{0xE2, 0, 1, notimplementd},
	{0xE3, 0, 1, notimplementd},
	{0xE4, 0, 1, notimplementd},
	{0xE5, 0, 1, notimplementd},
	{0xE6, 0, 1, notimplementd},
	{0xE7, 0, 1, notimplementd},
	{0xE8, 0, 1, notimplementd},
	{0xE9, 0, 1, notimplementd},
	{0xEA, 0, 1, notimplementd},
	{0xEB, 0, 1, notimplementd},
	{0xEC, 0, 1, notimplementd},
	{0xED, 0, 1, notimplementd},
	{0xEE, 0, 1, notimplementd},
	{0xEF, 0, 1, notimplementd},
	{0xF0, 0, 1, notimplementd},
	{0xF1, 0, 1, notimplementd},
	{0xF2, 0, 1, notimplementd},
	{0xF3, 0, 1, notimplementd},
	{0xF4, 0, 1, notimplementd},
	{0xF5, 0, 1, notimplementd},
	{0xF6, 0, 1, notimplementd},
	{0xF7, 0, 1, notimplementd},
	{0xF8, 0, 1, notimplementd},
	{0xF9, 0, 1, notimplementd},
	{0xFA, 0, 1, notimplementd},
	{0xFB, 0, 1, notimplementd},
	{0xFC, 0, 1, notimplementd},
	{0xFD, 0, 1, notimplementd},
	{0xFE, 0, 1, notimplementd},
	{0xFF, 0, 1, notimplementd},
}

func nop(c *CPU, ope []byte) {}

func notimplementd(c *CPU, ope []byte) {
	os.Exit(-1)
}
