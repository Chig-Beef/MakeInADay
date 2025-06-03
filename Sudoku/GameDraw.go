package main

import (
	"strconv"
	"image/color"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

func (g *Game) Draw(screen *ebiten.Image) {
	g.vis.InitFrame(screen)

	g.page.Draw(screen, g.vis)

	var lineColor color.Color = color.White
	if g.checkWinState() {
		lineColor = color.RGBA{0, 255, 0, 255}
	}

	// Draw grid lines
	g.vis.DrawLine(0, 0, 0, TILE_SIZE*9, 1, lineColor)
	g.vis.DrawLine(TILE_SIZE, 0, TILE_SIZE, TILE_SIZE*9, DFLT_LINE_WIDTH, lineColor)
	g.vis.DrawLine(TILE_SIZE*2, 0, TILE_SIZE*2, TILE_SIZE*9, DFLT_LINE_WIDTH, lineColor)
	g.vis.DrawLine(TILE_SIZE*3, 0, TILE_SIZE*3, TILE_SIZE*9, 1, lineColor)
	g.vis.DrawLine(TILE_SIZE*4, 0, TILE_SIZE*4, TILE_SIZE*9, DFLT_LINE_WIDTH, lineColor)
	g.vis.DrawLine(TILE_SIZE*5, 0, TILE_SIZE*5, TILE_SIZE*9, DFLT_LINE_WIDTH, lineColor)
	g.vis.DrawLine(TILE_SIZE*6, 0, TILE_SIZE*6, TILE_SIZE*9, 1, lineColor)
	g.vis.DrawLine(TILE_SIZE*7, 0, TILE_SIZE*7, TILE_SIZE*9, DFLT_LINE_WIDTH, lineColor)
	g.vis.DrawLine(TILE_SIZE*8, 0, TILE_SIZE*8, TILE_SIZE*9, DFLT_LINE_WIDTH, lineColor)
	g.vis.DrawLine(TILE_SIZE*9, 0, TILE_SIZE*9, TILE_SIZE*9, 1, lineColor)

	g.vis.DrawLine(0, 0, TILE_SIZE*9, 0, 1, lineColor)
	g.vis.DrawLine(0, TILE_SIZE, TILE_SIZE*9, TILE_SIZE, DFLT_LINE_WIDTH, lineColor)
	g.vis.DrawLine(0, TILE_SIZE*2, TILE_SIZE*9, TILE_SIZE*2, DFLT_LINE_WIDTH, lineColor)
	g.vis.DrawLine(0, TILE_SIZE*3, TILE_SIZE*9, TILE_SIZE*3, 1, lineColor)
	g.vis.DrawLine(0, TILE_SIZE*4, TILE_SIZE*9, TILE_SIZE*4, DFLT_LINE_WIDTH, lineColor)
	g.vis.DrawLine(0, TILE_SIZE*5, TILE_SIZE*9, TILE_SIZE*5, DFLT_LINE_WIDTH, lineColor)
	g.vis.DrawLine(0, TILE_SIZE*6, TILE_SIZE*9, TILE_SIZE*6, 1, lineColor)
	g.vis.DrawLine(0, TILE_SIZE*7, TILE_SIZE*9, TILE_SIZE*7, DFLT_LINE_WIDTH, lineColor)
	g.vis.DrawLine(0, TILE_SIZE*8, TILE_SIZE*9, TILE_SIZE*8, DFLT_LINE_WIDTH, lineColor)
	g.vis.DrawLine(0, TILE_SIZE*9, TILE_SIZE*9, TILE_SIZE*9, 1, lineColor)

	for r := range GRID_SIZE {
		for c := range GRID_SIZE {
			if g.grid[r][c].num == 0 {
				continue
			}

			op := text.DrawOptions{}

			op.PrimaryAlign = text.AlignCenter
			op.SecondaryAlign = text.AlignCenter

			g.vis.DrawText(strconv.Itoa(g.grid[r][c].num),TILE_SIZE, float64(c)*TILE_SIZE+TILE_SIZE/2, float64(r)*TILE_SIZE+TILE_SIZE/2, g.vis.GetFont("default"), &op)
		}
	}
}
