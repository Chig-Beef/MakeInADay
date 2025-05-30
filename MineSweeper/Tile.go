package main

import (
	"image/color"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Tile struct {
	r, c                              int
	surrounds                         int
	mine, flagged, revealed, exploded bool
}

func (t *Tile) flag(g *Game) {
	// Can only flag hidden tiles
	if t.revealed {
		return
	}

	// Taking back a flag
	if t.flagged {
		t.flagged = false
		g.flags++
		g.updateFlagCounter()
		return
	}

	// Trying to place a flag
	if !t.flagged {
		// No more flags
		if g.flags == 0 {
			return
		}

		// Place a flag
		t.flagged = true
		g.flags--
		g.updateFlagCounter()
		return
	}
}

func (t *Tile) reveal(g *Game) {
	// Can't double reveal
	if t.revealed {
		return
	}

	// Don't accidentally reveal mines
	if t.flagged {
		return
	}

	t.revealed = true

	if t.mine {
		g.phase = LOSE
		t.exploded = true
		return
	}

	g.checkWin()

	if t.surrounds == 0 {
		t.revealNeighbours(g)
	}
}

func (t *Tile) revealNeighbours(g *Game) {
	for nr := t.r - 1; nr <= t.r+1; nr++ {
		for nc := t.c - 1; nc <= t.c+1; nc++ {
			// Each surrounding position

			// Don't check self
			if nr == t.r && nc == t.c {
				continue
			}

			// Out of bounds check
			if nr < 0 || nc < 0 || nr >= GRID_SIZE || nc >= GRID_SIZE {
				continue
			}

			g.grid[nr][nc].reveal(g)
		}
	}
}

func (t *Tile) draw(g *Game) {
	x := float64(t.c) * TILE_SIZE
	y := float64(t.r) * TILE_SIZE

	if t.flagged {
		g.vis.DrawRect(x, y, TILE_SIZE, TILE_SIZE, color.RGBA{128, 128, 128, 255})
		g.vis.DrawImage(g.vis.GetImage("flag"), x, y, TILE_SIZE, TILE_SIZE, &ebiten.DrawImageOptions{})
		return
	}

	if t.exploded {
		g.vis.DrawRect(x, y, TILE_SIZE, TILE_SIZE, color.RGBA{255, 0, 0, 255})
		g.vis.DrawImage(g.vis.GetImage("mine"), x, y, TILE_SIZE, TILE_SIZE, &ebiten.DrawImageOptions{})
		return
	}

	if t.revealed {
		g.vis.DrawRect(x, y, TILE_SIZE, TILE_SIZE, color.RGBA{64, 64, 64, 255})
		if t.surrounds != 0 {
			g.vis.DrawText(strconv.Itoa(t.surrounds), TILE_SIZE*0.8, x, y, g.vis.GetFont("default"), &text.DrawOptions{})
		}
		return
	}

	if t.mine && g.phase == LOSE {
		g.vis.DrawRect(x, y, TILE_SIZE, TILE_SIZE, color.RGBA{64, 64, 64, 255})
		g.vis.DrawImage(g.vis.GetImage("mine"), x, y, TILE_SIZE, TILE_SIZE, &ebiten.DrawImageOptions{})
		return
	}

	g.vis.DrawRect(x, y, TILE_SIZE, TILE_SIZE, color.RGBA{128, 128, 128, 255})
}
