package main

import (
	"fmt"
	"graphics_test/pkg/graphics"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

func main() {
	fmt.Println("Starting graphics test ... ")

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Graphics test")

	game := graphics.NewGame(&graphics.Config{
		Screen: struct {
			Width  int
			Height int
		}{screenWidth, screenHeight},
	})

	game.Init()

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}

}
