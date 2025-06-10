package main

import (
	"GameOfLife/stagui"
	"image/color"
)

var buttonColor = color.RGBA{64, 64, 64, 255}

func (g *Game) initUI() {
	g.page = &stagui.Page{
		Title:   "",
		BgColor: color.Black,
		BgDraw:  false,
		Buttons: []*stagui.Button{
			{
				X:         85,
				Y:         1,
				W:         14,
				H:         5,
				Name:      "reset",
				BgColor:   buttonColor,
				Text:      "reset",
				TextColor: color.White,
				FontSize:  3,
			},
			{
				X:         70,
				Y:         1,
				W:         14,
				H:         5,
				Name:      "pause",
				BgColor:   buttonColor,
				Text:      "pause",
				TextColor: color.White,
				FontSize:  3,
			},
			{
				X:         55,
				Y:         1,
				W:         14,
				H:         5,
				Name:      "play",
				BgColor:   buttonColor,
				Text:      "play",
				TextColor: color.White,
				FontSize:  3,
			},
			{
				X:         40,
				Y:         1,
				W:         14,
				H:         5,
				Name:      "step",
				BgColor:   buttonColor,
				Text:      "step",
				TextColor: color.White,
				FontSize:  3,
			},
		},
		Text: []*stagui.StaticText{
		},
	}
}
