package main

func (g *Game) loadImages() {
	g.vis.LoadImage("missing", "assets/images/missing.png")
	g.vis.LoadImage("icon", "assets/images/icon.png")
}
