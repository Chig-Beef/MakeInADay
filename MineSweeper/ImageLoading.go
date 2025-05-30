package main

func (g *Game) loadImages() {
	g.vis.LoadImage("missing", "assets/images/missing.png")
	g.vis.LoadImage("icon", "assets/images/icon.png")
	g.vis.LoadImage("mine", "assets/images/mine.png")
	g.vis.LoadImage("face", "assets/images/face.png")
	g.vis.LoadImage("flag", "assets/images/flag.png")
}
