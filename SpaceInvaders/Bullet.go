package main

import "image/color"

type Bullet struct {
	x, y float64
	dy   float64
}

func (b *Bullet) update(g *Game) bool {
	b.y += b.dy
	return b.outOfBounds() || b.hitTarget(g)
}

func (b *Bullet) hitTarget(g *Game) bool {
	// Player collision
	if g.player.hit(b) {
		return true
	}

	// Tower collision
	for i := range g.towers {
		if g.towers[i].hit(b, g) {
			return true
		}
	}

	// Enemy collision
	for c := range len(g.pack.pack) {
		for r := range len(g.pack.pack[c]) {
			a := &g.pack.pack[c][r]
			if a.hit(b, g) {
				return true
			}
		}
	}

	return false
}

func (b *Bullet) outOfBounds() bool {
	return b.y < -BULLET_HEIGHT || b.y > 100
}

func (b *Bullet) draw(g *Game) {
	g.vis.DrawRect(b.x, b.y, BULLET_WIDTH, BULLET_HEIGHT, color.RGBA{0, 255, 0, 255})
}
