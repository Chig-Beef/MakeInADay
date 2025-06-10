package main

import "slices"

type Grid struct {
	tiles []Tile
}

func (g *Grid) step() {
	lastFrame := g.clone()
	g.bulk()
	g.applyRules(lastFrame)
	g.strip()
}

func (g *Grid) applyRules(lastFrame Grid) {
	for i := range g.tiles {
		g.tiles[i].step(lastFrame)
	}
}

func (g *Grid) clone() Grid {
	newGrid := Grid{}

	// Allocate the whole slice
	newGrid.tiles = make([]Tile, len(g.tiles))

	// Copy over each tile
	for i, t := range g.tiles {
		newGrid.tiles[i] = t
	}

	return newGrid
}

func (g *Grid) bulk() {
	// Only use current length, don't create new slice
	curLength := len(g.tiles)

	for i := range curLength {

		// Paranoia, shouldn't be possible because of strip
		if g.tiles[i].on {
			surTiles := g.tiles[i].getSurrounds()
			for n := range 8 {
				g.add(surTiles[n])
			}
		}
	}
}

func (g *Grid) isOn(row, col int) bool {
	for _, o := range g.tiles {
		if o.row == row && o.col == col {
			return o.on
		}
	}
	return false
}

func (g *Grid) remove(row, col int) {
	for i, o := range g.tiles {
		if o.row == row && o.col == col {
			g.tiles = slices.Delete(g.tiles, i, i+1)
			return
		}
	}
}

func (g *Grid) add(t Tile) {
	for _, o := range g.tiles {
		if o.row == t.row && o.col == t.col {
			return
		}
	}

	g.tiles = append(g.tiles, t)
}

func (g *Grid) strip() {
	// Will give a little extra space,
	// but will be used later when bulking
	newTiles := make([]Tile, 0, len(g.tiles))

	for _, o := range g.tiles {
		if o.on {
			newTiles = append(newTiles, o)
		}
	}

	g.tiles = newTiles
}

func (g *Grid) draw(game *Game) {
	for _, t := range g.tiles {
		t.draw(game)
	}
}
