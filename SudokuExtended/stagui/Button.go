package stagui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Button struct {
	X, Y, W, H float64

	// Used to identify the button
	Name string

	// Whether the button can be used
	Disabled bool

	// Good for toggle buttons
	Pressed bool

	// Text to display on the button
	Text      string
	FontSize  float64
	TextColor color.Color

	// Color of image of button
	BgColor color.Color
	BgImg   *ebiten.Image

	// The color if pressed
	PressedColor color.Color

	// The color if disabled
	DisabledColor color.Color
}

func (b *Button) Draw(screen *ebiten.Image, vh VisualHandler) {
	// Draw bg
	if b.BgImg == nil {
		b.drawAsSolidColor(vh)
	} else {
		vh.DrawImage(b.BgImg, float64(b.X), float64(b.Y), float64(b.W), float64(b.H), &ebiten.DrawImageOptions{})
	}

	// Draw text
	if b.Text != "" {
		op := text.DrawOptions{}
		op.PrimaryAlign = text.AlignCenter
		op.ColorScale.ScaleWithColor(b.TextColor)
		vh.DrawText(b.Text, b.FontSize, float64(b.X+b.W/2), float64(b.Y), vh.GetFont("button"), &op)
	}
}

func (b *Button) drawAsSolidColor(vh VisualHandler) {
	if b.Disabled {
		vh.DrawRect(b.X, b.Y, b.W, b.H, b.DisabledColor)
		return
	}

	if b.Pressed {
		vh.DrawRect(b.X, b.Y, b.W, b.H, b.PressedColor)
	} else {
		vh.DrawRect(b.X, b.Y, b.W, b.H, b.BgColor)
	}
}

func (b Button) CheckClick(x, y float64) bool {
	return !b.Disabled &&
		b.X <= x && x <= b.X+b.W &&
		b.Y <= y && y <= b.Y+b.H
}
