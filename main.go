package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/Teshima-Tatsuya/GoBoy/pkg/goboy"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	romData, err := ioutil.ReadFile(os.Args[1])

	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}

	gb := goboy.New(romData)

	if err := ebiten.RunGame(gb); err != nil {
		log.Fatal(err)
	}
}
