package main

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type Alien struct {
	x, y        float64
	img         *ebiten.Image
	shootChance float64
	points      int
	dead        bool
}

func (a *Alien) hit(b *Bullet, g *Game) bool {
	if !rectCollide(b.x, b.y, BULLET_WIDTH, BULLET_HEIGHT, a.x, a.y, ENTITY_SIZE, ENTITY_SIZE) {
		return false
	}

	if a.dead {
		return false
	}

	if b.dy != -1 {
		return false
	}

	a.dead = true
	g.score += a.points

	return true
}

func (a *Alien) shoot(g *Game) {
	if rand.Float64() >= a.shootChance {
		return
	}

	g.bullets = append(g.bullets, Bullet{
		x:  a.x + ENTITY_SIZE/2 - BULLET_WIDTH/2,
		y:  a.y + ENTITY_SIZE + 1,
		dy: 0.5,
	})
}

func (a *Alien) update(g *Game) bool {
	if a.dead {
		return false
	}

	a.shoot(g)

	// Possibly flip direction
	if a.x <= 1 {
		return true
	}

	if a.x >= 99-ENTITY_SIZE {
		return true
	}

	return false
}

func (a *Alien) draw(g *Game) {
	if a.dead {
		return
	}

	g.vis.DrawImage(a.img, a.x, a.y, ENTITY_SIZE, ENTITY_SIZE, &ebiten.DrawImageOptions{})
}
