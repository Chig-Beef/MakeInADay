package main

func (g *Game) loadImages() {
	g.vis.LoadImage("missing", "assets/images/missing.png")
	g.vis.LoadImage("icon", "assets/images/icon.png")
	g.vis.LoadImage("player", "assets/images/player.png")
	g.vis.LoadImage("alien1", "assets/images/alien1.png")
	g.vis.LoadImage("alien2", "assets/images/alien2.png")
	g.vis.LoadImage("alien3", "assets/images/alien3.png")
	g.vis.LoadImage("tower", "assets/images/tower.png")
}
