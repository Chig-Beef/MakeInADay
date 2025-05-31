package main

type AlienPack struct {
	pack [PACK_WIDTH][PACK_HEIGHT]Alien
	dx   float64
}

func (ap *AlienPack) empty() bool {
	for r := range PACK_HEIGHT {
		for c := range PACK_WIDTH {
			if !ap.pack[c][r].dead {
				return false
			}
		}
	}
	return true
}

func (ap *AlienPack) update(g *Game) {
	flip := false

	// Move all the aliens
	for c := range len(ap.pack) {
		for r := range len(ap.pack[c]) {
			ap.pack[c][r].x += ap.dx
			if ap.pack[c][r].update(g) {
				flip = true
			}
		}
	}

	if flip {
		ap.flip()
	}
}

func (ap *AlienPack) flip() {
	// Speed up
	ap.dx *= 1.01

	// Opposite direction
	ap.dx *= -1

	// Move down a bit
	for c := range len(ap.pack) {
		for r := range len(ap.pack[c]) {
			ap.pack[c][r].y += ENTITY_SIZE / 2
		}
	}
}

func (ap *AlienPack) draw(g *Game) {
	for c := range len(ap.pack) {
		for r := range len(ap.pack[c]) {
			ap.pack[c][r].draw(g)
		}
	}
}

func (ap *AlienPack) newPack(g *Game) {
	// Clear the pack
	ap.pack = [PACK_WIDTH][PACK_HEIGHT]Alien{}

	r := 0

	for c := range PACK_WIDTH {
		ap.pack[c][r] = Alien{
			x:           float64((c + 1) * (ENTITY_SIZE + 1)),
			y:           float64((r + 1) * (ENTITY_SIZE + 1)),
			img:         g.vis.GetImage("alien3"),
			shootChance: 0.0015,
			points:      300,
		}
	}

	r = 1

	for c := range PACK_WIDTH {
		ap.pack[c][r] = Alien{
			x:           float64((c + 1) * (ENTITY_SIZE + 1)),
			y:           float64((r + 1) * (ENTITY_SIZE + 1)),
			img:         g.vis.GetImage("alien2"),
			shootChance: 0.001,
			points:      200,
		}
	}

	r = 2

	for c := range PACK_WIDTH {
		ap.pack[c][r] = Alien{
			x:           float64((c + 1) * (ENTITY_SIZE + 1)),
			y:           float64((r + 1) * (ENTITY_SIZE + 1)),
			img:         g.vis.GetImage("alien1"),
			shootChance: 0.0005,
			points:      100,
		}
	}
}
