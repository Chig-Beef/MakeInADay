package main

import "image/color"

var tileOnColor = color.RGBA{255, 255, 0, 255}

type Tile struct {
	row, col int
	on bool
}

func (t *Tile) step(lastGrid Grid) {
	n := t.countSurrounds(lastGrid)

	// Die (alone)
	if n < 2 {
		t.on = false
		return
	}

	// Die (popular)
	if n > 3 {
		t.on = false
		return
	}

	// Exist
	if n == 2 {
		return
	}

	// Born
	if n == 3 {
		t.on = true
		return
	}
}

func (t *Tile) countSurrounds(lastGrid Grid) int {
	var n int

	if lastGrid.isOn(t.row-1, t.col-1) {
		n++
	}

	if lastGrid.isOn(t.row-1, t.col) {
		n++
	}

	if lastGrid.isOn(t.row-1, t.col+1) {
		n++
	}

	if lastGrid.isOn(t.row, t.col-1) {
		n++
	}

	if lastGrid.isOn(t.row, t.col+1) {
		n++
	}

	if lastGrid.isOn(t.row+1, t.col-1) {
		n++
	}

	if lastGrid.isOn(t.row+1, t.col) {
		n++
	}

	if lastGrid.isOn(t.row+1, t.col+1) {
		n++
	}

	return n
}

func (t *Tile) getSurrounds() [8]Tile {
	return [8]Tile{
		{t.row-1, t.col-1, false},
		{t.row-1, t.col, false},
		{t.row-1, t.col+1, false},
		{t.row, t.col-1, false},
		{t.row, t.col+1, false},
		{t.row+1, t.col-1, false},
		{t.row+1, t.col, false},
		{t.row+1, t.col+1, false},
	}
}

func (t *Tile) draw(g *Game) {
	clr := tileOnColor

	g.vis.DrawRect(
		float64(t.col*TILE_SIZE-g.offsetX),
		float64(t.row*TILE_SIZE-g.offsetY),
		TILE_SIZE,
		TILE_SIZE,
		clr,
	)
}
