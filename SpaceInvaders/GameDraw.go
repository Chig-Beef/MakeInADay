package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) Draw(screen *ebiten.Image) {
	g.vis.InitFrame(screen)

	switch g.phase {
	case PLAY:
		g.play.Draw(screen, g.vis)
		g.player.draw(g)
		g.pack.draw(g)

		for _, t := range g.towers {
			t.draw(g)
		}

		for _, b := range g.bullets {
			b.draw(g)
		}

	case LOSE:
		g.lose.Draw(screen, g.vis)
	}
}
