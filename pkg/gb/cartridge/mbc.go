package cartridge

import "github.com/Teshima-Tatsuya/GoBoy/pkg/types"

// Memory Bank Controller
// Which provides BANK switch
// to expand address spaces
type MBC interface {
	Read(addr types.Addr) byte
	Write(addr types.Addr, value byte)
	SwitchROMBank(bank uint16)
	SwitchRAMBank(bank uint8)
}
