package gb

import (
	"github.com/DisgoOrg/log"
	"github.com/Teshima-Tatsuya/GoBoy/pkg/cartridge"
)

type GB struct {
	Cartridge cartridge.Cartridge
}

func NewGB(romData []byte) {
	c := cartridge.New(romData)
	log.Info(c.Title)

}
