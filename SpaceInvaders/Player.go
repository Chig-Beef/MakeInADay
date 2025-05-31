package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	x, y float64
	img  *ebiten.Image

	weaponCooldown    int
	weaponCooldownMax int
	lives             int
}

func (p *Player) hit(b *Bullet) bool {
	if !rectCollide(b.x, b.y, BULLET_WIDTH, BULLET_HEIGHT, p.x, p.y, ENTITY_SIZE, ENTITY_SIZE) {
		return false
	}

	p.lives--
	return true
}

func (p *Player) shoot(g *Game) {
	p.weaponCooldown = p.weaponCooldownMax
	g.bullets = append(g.bullets, Bullet{
		x:  p.x + ENTITY_SIZE/2 - BULLET_WIDTH/2,
		y:  p.y - BULLET_HEIGHT - 1,
		dy: -1,
	})
}

func (p *Player) update(g *Game) {
	if p.weaponCooldown > 0 {
		p.weaponCooldown--
	} else {
		if ebiten.IsKeyPressed(ebiten.KeySpace) {
			p.shoot(g)
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		p.x--
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		p.x++
	}

	if p.x < 0 {
		p.x = 0
	}

	if p.x > 100-ENTITY_SIZE {
		p.x = 100 - ENTITY_SIZE
	}
}

func (p *Player) draw(g *Game) {
	g.vis.DrawImage(p.img, p.x, p.y, ENTITY_SIZE, ENTITY_SIZE, &ebiten.DrawImageOptions{})

	for i := range p.lives {
		g.vis.DrawImage(p.img, float64(100-(ENTITY_SIZE+1)*(i+1)), 1, ENTITY_SIZE, ENTITY_SIZE, &ebiten.DrawImageOptions{})
	}
}
