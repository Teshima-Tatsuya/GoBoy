package joypad

import (
	"github.com/Teshima-Tatsuya/GoBoy/pkg/gb/pad"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func Press() pad.Button {
	for key, button := range keyMap {
		if inpututil.IsKeyJustPressed(key) {
			return button
		}
	}
	return 0x00
}

func Release() pad.Button {
	for key, button := range keyMap {
		if inpututil.IsKeyJustReleased(key) {
			return button
		}
	}
	return 0x00
}

var keyMap = map[ebiten.Key]pad.Button{
	ebiten.KeyZ:         pad.A,
	ebiten.KeyX:         pad.B,
	ebiten.KeyBackspace: pad.Select,
	ebiten.KeyEnter:     pad.Start,
	ebiten.KeyRight:     pad.Right,
	ebiten.KeyLeft:      pad.Left,
	ebiten.KeyUp:        pad.Up,
	ebiten.KeyDown:      pad.Down,
}
