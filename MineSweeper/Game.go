package main

import (
	"MineSweeper/stagui"
	"MineSweeper/vvis"
	"math/rand"
	"slices"
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	vis *vvis.VisualHandler

	mousePos [2]float64
	phase    Phase

	page *stagui.Page

	grid       [GRID_SIZE][GRID_SIZE]Tile
	flags      int
	firstClick bool

	startTime time.Time
}

func (g *Game) Update() error {
	var pressed string

	g.getMousePos()

	switch g.phase {
	case PLAY:
		g.page.Text[1].Text = strconv.Itoa(int(time.Now().Sub(g.startTime).Seconds()))

		pressed, _, _ = g.page.Update(g.mousePos)
		switch pressed {
		case "reset":
			g.reset()
		}

		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
			g.leftClick()
		}

		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton2) {
			g.rightClick()
		}

	case WIN:
		pressed, _, _ = g.page.Update(g.mousePos)
		switch pressed {
		case "reset":
			g.reset()
		}

	case LOSE:
		pressed, _, _ = g.page.Update(g.mousePos)
		switch pressed {
		case "reset":
			g.reset()
		}
	}

	return nil
}

func (g *Game) leftClick() {
	x := int(g.mousePos[0] / TILE_SIZE)
	y := int(g.mousePos[1] / TILE_SIZE)

	if x < 0 || x >= GRID_SIZE || y < 0 || y >= GRID_SIZE {
		return
	}

	if g.firstClick {
		g.firstClick = false
		g.placeMines(x, y)
		g.countMines()
	}

	g.grid[y][x].reveal(g)
}

func (g *Game) countMines() {
	for r := range GRID_SIZE {
		for c := range GRID_SIZE {
			// Each tile

			for nr := r - 1; nr <= r+1; nr++ {
				for nc := c - 1; nc <= c+1; nc++ {
					// Each surrounding position

					// Don't check self
					if nr == r && nc == c {
						continue
					}

					// Out of bounds check
					if nr < 0 || nc < 0 || nr >= GRID_SIZE || nc >= GRID_SIZE {
						continue
					}

					if g.grid[nr][nc].mine {
						g.grid[r][c].surrounds++
					}
				}
			}
		}
	}
}

func (g *Game) placeMines(x, y int) {
	// List of tiles to avoid
	avoids := []int{y*GRID_SIZE + x}

	for range START_MINES {
		inlinePos := rand.Intn(GRID_SIZE*GRID_SIZE - len(avoids))

		// Avoid positions
		for _, avoidPos := range avoids {
			if inlinePos >= avoidPos {
				inlinePos++
			}
		}

		// Convert to a 2D position
		mineX := inlinePos % GRID_SIZE
		mineY := inlinePos / GRID_SIZE

		// Turn it into a mine
		g.grid[mineY][mineX].mine = true

		// Add this new position to the list of tiles to avoid (sorted)
		insertPos := 0
		for insertPos = 0; insertPos < len(avoids); insertPos++ {
			if avoids[insertPos] > inlinePos {
				break
			}
		}
		avoids = slices.Insert(avoids, insertPos, mineY*GRID_SIZE+mineX)
	}
}

func (g *Game) rightClick() {
	x := int(g.mousePos[0] / TILE_SIZE)
	y := int(g.mousePos[1] / TILE_SIZE)

	if x < 0 || x >= GRID_SIZE || y < 0 || y >= GRID_SIZE {
		return
	}

	g.grid[y][x].flag(g)
}

func (g *Game) updateFlagCounter() {
	g.page.Text[0].Text = strconv.Itoa(g.flags)
}

func (g *Game) reset() {
	g.phase = PLAY
	g.flags = START_MINES
	g.firstClick = true

	g.grid = [GRID_SIZE][GRID_SIZE]Tile{}
	for r := range GRID_SIZE {
		for c := range GRID_SIZE {
			g.grid[r][c].r = r
			g.grid[r][c].c = c
		}
	}

	g.startTime = time.Now()
}

func (g *Game) checkWin() {
	// Check for a non-mined tile without a mine
	for r := range GRID_SIZE {
		for c := range GRID_SIZE {
			if !g.grid[r][c].revealed && !g.grid[r][c].mine {
				return
			}
		}
	}

	// We won!
	g.phase = WIN
}

func (g *Game) getMousePos() (float64, float64) {
	x, y := ebiten.CursorPosition()
	g.mousePos = [2]float64{g.vis.Untranslate(float64(x)), g.vis.Untranslate(float64(y))}
	return g.mousePos[0], g.mousePos[1]
}

func (g *Game) Layout(int, int) (int, int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}

func newGame() Game {
	g := Game{}

	g.vis = &vvis.VisualHandler{}
	g.vis.Init(SCREEN_WIDTH, 100, IS_RELEASE)

	g.loadImages()

	g.vis.InitImages()
	g.vis.LoadFonts()

	g.initUI()

	g.reset()

	return g
}
