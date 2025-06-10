package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) Draw(screen *ebiten.Image) {
	g.vis.InitFrame(screen)
	g.grid.draw(g)
	g.page.Draw(screen, g.vis)
}
