package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g := newGame()

	ebiten.SetWindowSize(SCREEN_WIDTH, SCREEN_HEIGHT)
	ebiten.SetWindowTitle("Sudoku")
	ebiten.SetTPS(FPS)

	fmt.Println("Starting")

	err := ebiten.RunGame(&g)
	if err != nil && err != ebiten.Termination {
		panic(err)
	}

	fmt.Println("Ended")
}
