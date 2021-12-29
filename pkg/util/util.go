package util

import "github.com/Teshima-Tatsuya/GoBoy/pkg/types"

func Byte2Addr(upper byte, lower byte) types.Addr {
	return types.Addr((uint16(upper) << 8) | uint16(lower))
}

func ExtractLower(addr types.Addr) byte {
	return byte(addr & 0x00ff)
}

func ExtractUpper(addr types.Addr) byte {
	return byte(addr >> 8)
}
