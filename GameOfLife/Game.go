package main

import (
	"GameOfLife/stagui"
	"GameOfLife/vvis"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	vis *vvis.VisualHandler

	mousePos [2]float64

	page *stagui.Page

	grid Grid

	autoStepping bool

	offsetX int
	offsetY int
}

func (g *Game) Update() error {
	var pressed string

	g.getMousePos()

	pressed, _, _ = g.page.Update(g.mousePos)
	switch pressed {
	case "pause":
		g.pause()

	case "play":
		g.play()

	case "reset":
		g.reset()

	case "step":
		g.step()

	default:
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			g.leftClick()
		}

		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
			g.rightClick()
		}
	}

	g.cameraMove()

	if g.autoStepping {
		g.step()
	}

	return nil
}

func (g *Game) cameraMove() {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.offsetY--
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.offsetY++
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.offsetX--
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.offsetX++
	}
}

func (g *Game) pause() {
	g.autoStepping = false
}

func (g *Game) play() {
	g.autoStepping = true
}

func (g *Game) step() {
	g.grid.step()
}

func (g *Game) reset() {
	g.grid = Grid{}
}

func (g *Game) leftClick() {
	// Get actual position
	x := int(g.mousePos[0]) + g.offsetX
	y := int(g.mousePos[1]) + g.offsetY

	// Get row and col
	row := y / TILE_SIZE
	col := x / TILE_SIZE

	// Place the tile
	g.grid.add(Tile{row, col, true})
}

func (g *Game) rightClick() {
	// Get actual position
	x := int(g.mousePos[0]) + g.offsetX
	y := int(g.mousePos[1]) + g.offsetY

	// Get row and col
	row := y / TILE_SIZE
	col := x / TILE_SIZE

	// Place the tile
	g.grid.remove(row, col)
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
