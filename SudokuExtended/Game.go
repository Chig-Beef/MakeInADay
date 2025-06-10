package main

import (
	"strconv"
	"math/rand"
	"Sudoku/stagui"
	"Sudoku/vvis"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	vis *vvis.VisualHandler

	mousePos [2]float64

	page *stagui.Page

	grid [GRID_SIZE][GRID_SIZE]Tile

	autoSolving bool
	robotSolving bool
	checkedAllPoss bool
	autoRow int
	autoCol int

	curNumber int
}

func (g *Game) Update() error {
	var pressed string

	g.getMousePos()

	if g.autoSolving {
		g.auto()
	}

	pressed, _, _ = g.page.Update(g.mousePos)
	switch pressed {
	case "reset":
		g.reset()

	case "auto":
		g.clearBoard(0)
		g.autoSolving = true
		g.robotSolving = false
		g.checkedAllPoss = false

	case "0":
		g.changeNumber(0)

	case "1":
		g.changeNumber(1)

	case "2":
		g.changeNumber(2)

	case "3":
		g.changeNumber(3)

	case "4":
		g.changeNumber(4)

	case "5":
		g.changeNumber(5)

	case "6":
		g.changeNumber(6)

	case "7":
		g.changeNumber(7)

	case "8":
		g.changeNumber(8)

	case "9":
		g.changeNumber(9)
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
		g.leftClick()
	}

	return nil
}

func (g *Game) clearBoard(n int) {
	for r := range GRID_SIZE {
		for c := range GRID_SIZE {
			if g.grid[r][c].locked {
				continue
			}
			g.grid[r][c].num = n
		}
	}
}

func (g *Game) changeNumber(n int) {
	g.curNumber = n
	g.page.Text[0].Text = "CurCode: " + strconv.Itoa(n)
}

func (g *Game) checkWinState() bool {
	// 0 Check
	for r := range GRID_SIZE {
		for c := range GRID_SIZE {
			if g.grid[r][c].num == 0 {
				return false
			}
		}
	}

	// Col Check
	for c := range GRID_SIZE {
		var check uint16 = 0
		for r := range GRID_SIZE {
			check |= uint16(1 << (g.grid[r][c].num-1))
		}

		if check != 511 {
			return false
		}
	}

	// Row Check
	for r := range GRID_SIZE {
		var check uint16 = 0
		for c := range GRID_SIZE {
			check |= uint16(1 << (g.grid[r][c].num-1))
		}

		if check != 511 {
			return false
		}
	}

	// Sub-Grid Check
	for or := range SUB_GRID_SIZE {
		for oc := range SUB_GRID_SIZE {
			var check uint16 = 0
			for ir := range SUB_GRID_SIZE {
				for ic := range SUB_GRID_SIZE {
					check |= uint16(1 << (g.grid[or*SUB_GRID_SIZE+ir][oc*SUB_GRID_SIZE+ic].num-1))
				}
			}

			if check != 511 {
				return false
			}
		}
	}

	return true
}

func (g *Game) reset() {
	i := rand.Intn(len(premadeGrids))

	for r := range GRID_SIZE {
		for c := range GRID_SIZE {
			n := premadeGrids[i][r][c]
			g.grid[r][c] = Tile{num:n, locked:n != 0}
		}
	}

	g.autoSolving = false
}

func (g *Game) leftClick() {
	if g.autoSolving {
		return
	}

	x := int(g.mousePos[0] / TILE_SIZE)
	y := int(g.mousePos[1] / TILE_SIZE)

	if x < 0 || x >= GRID_SIZE || y < 0 || y >= GRID_SIZE {
		return
	}

	if g.grid[y][x].locked {
		return
	}

	g.grid[y][x].num = g.curNumber
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
