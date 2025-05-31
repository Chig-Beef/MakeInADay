package main

import (
	"SpaceInvaders/stagui"
	"SpaceInvaders/vvis"
	"slices"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	vis *vvis.VisualHandler

	mousePos [2]float64
	phase    Phase

	lose, play *stagui.Page

	score int

	player  Player
	bullets []Bullet
	pack    AlienPack
	towers  [3]Tower
}

func (g *Game) Update() error {
	var pressed string

	g.getMousePos()

	switch g.phase {
	case PLAY:
		g.player.update(g)

		g.pack.update(g)

		for i := len(g.bullets) - 1; i >= 0; i-- {
			if g.bullets[i].update(g) {
				g.bullets = slices.Delete(g.bullets, i, i+1)
			}
		}

		g.checkLoseConditions()
		g.checkWinConditions()

	case LOSE:
		pressed, _, _ = g.lose.Update(g.mousePos)
		switch pressed {
		case "reset":
			g.reset()
		}
	}

	g.updateScore()

	return nil
}

func (g *Game) updateScore() {
	g.play.Text[0].Text = strconv.Itoa(g.score)
	g.lose.Text[0].Text = strconv.Itoa(g.score)
}

func (g *Game) checkWinConditions() {
	if g.pack.empty() {
		g.pack.newPack(g)

		// Speed up
		g.pack.dx *= 1.01
	}
}

func (g *Game) checkLoseConditions() {
	// Player lives check
	if g.player.lives == 0 {
		g.phase = LOSE
		return
	}

	// Tower lives check
	allTowersDead := true
	for i := range g.towers {
		if g.towers[i].health > 0 {
			allTowersDead = false
			break
		}
	}
	if allTowersDead {
		g.phase = LOSE
		return
	}

	// Enemy breached check
	towerY := g.towers[0].y
	for r := range PACK_HEIGHT {
		for c := range PACK_WIDTH {
			e := &g.pack.pack[c][r]

			if e.dead {
				continue
			}

			if e.y+ENTITY_SIZE >= towerY {
				g.phase = LOSE
				return
			}
		}
	}
}

func (g *Game) reset() {
	g.phase = PLAY

	g.player = Player{
		x:                 50,
		y:                 49 - ENTITY_SIZE,
		weaponCooldownMax: FPS,
		img:               g.vis.GetImage("player"),
		lives:             3,
	}

	g.score = 0
	g.pack.newPack(g)
	g.bullets = []Bullet{}
	g.pack.dx = 0.05

	for i := range 3 {
		g.towers[i] = Tower{
			x:      float64(25 * (i + 1)),
			y:      50 - 2*(ENTITY_SIZE+1),
			health: 10,
			img:    g.vis.GetImage("tower"),
		}
	}
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
