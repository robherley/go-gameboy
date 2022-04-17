package ui

import (
	"errors"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	WindowTitle  = "go-gameboy"
	WindowWidth  = 600
	WindowHeight = 600
)

func NewSDLWindow() error {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return err
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow(WindowTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		WindowWidth, WindowHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		return err
	}
	defer window.Destroy()

	surface, err := window.GetSurface()
	if err != nil {
		return err
	}
	surface.FillRect(nil, 0)

	rect := sdl.Rect{X: 0, Y: 0, W: 200, H: 200}
	surface.FillRect(&rect, 0x724CF9)
	window.UpdateSurface()

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return errors.New("quit")
			}
		}
	}
}
