package stagui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Slider struct {
	X float64
	Y float64
	W float64
	H float64

	Value float64

	Name    string // Used to identify the slider
	Pressed bool

	LineColor   color.Color
	SliderColor color.Color
}

func (s *Slider) Draw(vh VisualHandler) {
	// Bar
	vh.DrawRect(s.X, s.Y+s.H/4, s.W, s.H/2, s.LineColor)

	// Slidy thing
	vh.DrawRect(s.X+s.Value*(s.W-8), s.Y, 8, s.H, s.SliderColor)
}

func (s *Slider) CheckCollide(x, y float64) bool {
	return s.X <= x && x <= s.X+s.W &&
		s.Y <= y && y <= s.Y+s.H
}

func (s *Slider) Update(curMousePos [2]float64) bool {
	x := curMousePos[0]
	y := curMousePos[1]

	if !ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
		return false
	}

	if !s.CheckCollide(x, y) {
		return false
	}

	dx := x - s.X

	s.Value = float64(dx / (s.W - 8))
	if s.Value > 1 {
		s.Value = 1
	}

	// Changed
	return true
}
