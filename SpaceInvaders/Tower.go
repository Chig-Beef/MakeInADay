package main

import "github.com/hajimehoshi/ebiten/v2"

type Tower struct {
	x, y   float64
	health int
	img    *ebiten.Image
}

func (t *Tower) hit(b *Bullet, g *Game) bool {
	if t.health == 0 {
		return false
	}

	if !rectCollide(b.x, b.y, BULLET_WIDTH, BULLET_HEIGHT, t.x, t.y, ENTITY_SIZE, ENTITY_SIZE) {
		return false
	}

	t.health -= 1
	return true
}

func (t *Tower) draw(g *Game) {
	if t.health > 0 {
		g.vis.DrawImage(t.img, t.x, t.y, ENTITY_SIZE, ENTITY_SIZE, &ebiten.DrawImageOptions{})
	}
}
