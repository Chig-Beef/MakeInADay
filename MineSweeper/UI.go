package main

import (
	"MineSweeper/stagui"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var buttonColor = color.RGBA{64, 64, 64, 255}

func (g *Game) initUI() {
	g.page = &stagui.Page{
		Title:   "",
		BgColor: color.Black,
		BgDraw:  true,
		Buttons: []*stagui.Button{
			{
				X:       92.5,
				Y:       0.5,
				W:       7,
				H:       7,
				Name:    "reset",
				BgColor: buttonColor,
				BgImg:   g.vis.GetImage("face"),
			},
		},
		Text: []*stagui.StaticText{
			{
				Text:  "10",
				X:     75,
				Y:     10,
				Color: color.White,
				Font:  g.vis.GetFont("default"),
				Size:  5,
				Align: text.AlignCenter,
			},
			{
				Text:  "10",
				X:     75,
				Y:     16,
				Color: color.White,
				Font:  g.vis.GetFont("default"),
				Size:  5,
				Align: text.AlignCenter,
			},
		},
	}
}
