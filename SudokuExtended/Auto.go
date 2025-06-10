package main

import (
	"time"
)

func (g *Game) auto() {
	t1 := time.Now()
	var diff time.Duration = 0

	for diff.Milliseconds() < MS_PER_FRAME {

		switchingToRobot := false

		if g.robotSolving {
			g.bruteForce()
		} else {
			if g.humanSolve() {
			} else {
				switchingToRobot = true
			}
		}

		// Done!
		if g.checkWinState() {
			g.autoSolving = false
			return
		} else if switchingToRobot {
			g.setAllLowestPossible()
			g.robotSolving = true
			g.autoRow = 0
			g.autoCol = 0
		}

		t2 := time.Now()
		diff = t2.Sub(t1)
	}
}

func (g *Game) humanSolve() bool {
	if !g.checkedAllPoss {
		if g.checkAllValidPoss() {
			g.checkedAllPoss = true
		}
		return true
	}
	
	// Now find singles and place them
	if g.placeEnsuredTile() {
		g.checkedAllPoss = false
		return true
	}

	return !(g.autoRow == 0 && g.autoCol == 0)
}

// Returns true if a wrap occurs
func (g *Game) nextTile() bool {
	// Move to next tile
	g.autoMove(g.autoRow, g.autoCol+1)

	// Wrapped?
	return g.autoRow == 0 && g.autoCol == 0 
}

func (g *Game) placeEnsuredTile() bool {
	// Number already given
	if g.grid[g.autoRow][g.autoCol].num != 0 {
		g.nextTile()
		return false
	}

	if g.grid[g.autoRow][g.autoCol].locked {
		g.nextTile()
		return false
	}

	numValids := 0
	lastIndex := 0
	for i := range 9 {
		if g.grid[g.autoRow][g.autoCol].canBe[i] {
			numValids++
			lastIndex = i
		}
	}

	if numValids == 1 {
		g.grid[g.autoRow][g.autoCol].num = lastIndex+1
		g.grid[g.autoRow][g.autoCol].known = true
		g.nextTile()
		return true
	}

	g.nextTile()
	return false
}

func (g *Game) setAllLowestPossible() {
	for r := range GRID_SIZE {
		for c := range GRID_SIZE {
			if g.grid[r][c].locked {
				continue
			}

			if g.grid[r][c].known {
				continue
			}

			// Find lowest posible
			var i int
			for i = 0; i < 9; i++ {
				if g.grid[r][c].canBe[i] {
					break
				}
			}

			// Can't be anything?
			if i == 9 {
				panic("Couldn't find a lowest possible value")
				continue
			}

			// Set to lowest possible
			g.grid[r][c].num = i+1
		}
	}
}

func (g *Game) setLowestPossible(r, c int) {
	if g.grid[r][c].locked {
		return
	}

	if g.grid[r][c].known {
		return
	}

	// Find lowest posible
	var i int
	for i = 0; i < 9; i++ {
		if g.grid[r][c].canBe[i] {
			break
		}
	}

	// Can't be anything?
	if i == 9 {
		return
	}

	// Set to lowest possible
	g.grid[r][c].num = i+1
}

func (g *Game) nextLowestPossible(r, c int) bool {
	n := g.grid[r][c].num

	for i := n; i < 9; i++ {
		if g.grid[r][c].canBe[i] {
			g.grid[r][c].num = i+1
			return false
		}
	}

	// Wrap back around
	for i := 0; i < n; i++ {
		if g.grid[r][c].canBe[i] {
			g.grid[r][c].num = i+1
			return true
		}
	}

	// No possible thingy
	return false
}

func (g *Game) atLastPossible(r, c int) bool {
	n := g.grid[r][c].num-1

	var i int
	for i = 8; i >= 0; i-- {
		if g.grid[r][c].canBe[i] {
			break
		}
	}

	return n == i
}

func (g *Game) checkAllValidPoss() bool {
	// Number already given
	if g.grid[g.autoRow][g.autoCol].num != 0 {
		return g.nextTile()
	}

	if g.grid[g.autoRow][g.autoCol].locked {
		return g.nextTile()
	}

	g.grid[g.autoRow][g.autoCol].canBe = [9]bool{true, true, true, true, true, true, true, true, true}

	for r := range GRID_SIZE {
		if r == g.autoRow {
			continue
		}

		if g.grid[r][g.autoCol].num == 0 {
			continue
		}

		g.grid[g.autoRow][g.autoCol].canBe[g.grid[r][g.autoCol].num-1] = false
	}

	for c := range GRID_SIZE {
		if c == g.autoCol {
			continue
		}

		if g.grid[g.autoRow][c].num == 0 {
			continue
		}

		g.grid[g.autoRow][g.autoCol].canBe[g.grid[g.autoRow][c].num-1] = false
	}

	startRow := g.autoRow-g.autoRow%3
	startCol := g.autoCol-g.autoCol%3
	
	for subRow := range 3 {
		for subCol := range 3 {
			r := startRow+subRow
			c := startCol+subCol

			if c == g.autoCol {
				continue
			}

			if r == g.autoRow {
				continue
			}

			if g.grid[r][c].num == 0 {
				continue
			}

			g.grid[g.autoRow][g.autoCol].canBe[g.grid[r][c].num-1] = false
		}
	}

	return g.nextTile()
}

func (g *Game) bruteForce() {
	// Invalid move
	if g.autoRow < 0 || g.autoRow >= GRID_SIZE {
		g.autoSolving = false
		return
	}

	if g.grid[g.autoRow][g.autoCol].locked {
		g.autoMove(g.autoRow, g.autoCol+1)
		return
	}

	if g.grid[g.autoRow][g.autoCol].known {
		g.autoMove(g.autoRow, g.autoCol+1)
		return
	}

	// // This tile is done
	// if g.grid[g.autoRow][g.autoCol].num == 9 {
	// 	g.grid[g.autoRow][g.autoCol].num = 1
	//
	// 	// Move to the next tile
	// 	g.autoMove(g.autoRow, g.autoCol+1)
	// 	return
	// }

	if g.atLastPossible(g.autoRow, g.autoCol) {
		// Set to lowest possible
		g.setLowestPossible(g.autoRow, g.autoCol)

		// Move to next tile
		g.autoMove(g.autoRow, g.autoCol+1)
		return
	}

	g.nextLowestPossible(g.autoRow, g.autoCol)

	// for {
	// 	g.grid[g.autoRow][g.autoCol].num++
	//
	// 	// Done the tile
	// 	if g.grid[g.autoRow][g.autoCol].num == 9 {
	// 		break
	// 	}
	//
	// 	// Valid number
	// 	valid := true
	// 	if g.countRowMatches(g.autoRow, g.autoCol, g.grid[g.autoRow][g.autoCol].num) > 0 {
	// 		valid = false
	// 	}
	// 	if g.countColMatches(g.autoRow, g.autoCol, g.grid[g.autoRow][g.autoCol].num) > 0 {
	// 		valid = false
	// 	}
	// 	if valid {
	// 		break
	// 	}
	// }

	g.autoMove(0, 0)
}

func (g *Game) countRowMatches(r, c, n int) int {
	count := 0

	for c = c+1; c < GRID_SIZE; c++ {
		if g.grid[r][c].num == n {
			count++
		}
	}

	return count
}

func (g *Game) countColMatches(r, c, n int) int {
	count := 0

	for r = r+1; r < GRID_SIZE; r++ {
		if g.grid[r][c].num == n {
			count++
		}
	}

	return count
}

func (g *Game) autoMove(r, c int) {
	// Wrap
	if c < 0 {
		r--
		c = GRID_SIZE-1
	}

	if c >= GRID_SIZE {
		r++
		c = 0
	}

	// Out of bounds
	if r < 0 || r >= GRID_SIZE {
		g.autoRow = 0
		g.autoCol = 0
		return
	}

	// Correct (no change needed
	g.autoRow = r
	g.autoCol = c
}
