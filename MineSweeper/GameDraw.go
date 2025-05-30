package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) Draw(screen *ebiten.Image) {
	g.vis.InitFrame(screen)

	switch g.phase {
	case PLAY:
		g.page.Draw(screen, g.vis)
	case WIN:
		g.page.Draw(screen, g.vis)
	case LOSE:
		g.page.Draw(screen, g.vis)
	}

	for r := range GRID_SIZE {
		for c := range GRID_SIZE {
			g.grid[r][c].draw(g)
		}
	}
}
