package main

import (
	"SpaceInvaders/stagui"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var buttonColor = color.RGBA{64, 64, 64, 255}

func (g *Game) initUI() {
	g.play = &stagui.Page{
		Title:   "",
		BgColor: color.Black,
		BgDraw:  true,
		Text: []*stagui.StaticText{
			{
				Text:  "",
				X:     1,
				Y:     1,
				Color: color.White,
				Font:  g.vis.GetFont("default"),
				Size:  3,
			},
		},
	}

	g.lose = &stagui.Page{
		Title:   "You Lost",
		BgColor: color.Black,
		BgDraw:  true,
		Buttons: []*stagui.Button{
			{
				X:         43,
				Y:         25,
				W:         14,
				H:         7,
				Name:      "reset",
				BgColor:   buttonColor,
				Text:      "reset",
				TextColor: color.White,
				FontSize:  3,
			},
		},
		Text: []*stagui.StaticText{
			{
				Text:  "",
				X:     50,
				Y:     15,
				Color: color.White,
				Font:  g.vis.GetFont("default"),
				Size:  3,
				Align: text.AlignCenter,
			},
		},
	}
}
