package stagui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type StaticText struct {
	Text string

	X float64
	Y float64

	Color color.Color
	Font  *text.GoTextFaceSource
	Size  float64
	Align text.Align
}

func (st *StaticText) Draw(screen *ebiten.Image, vh VisualHandler) {
	nx := vh.Translate(st.X)
	ny := vh.Translate(st.Y)
	nsize := vh.Translate(st.Size)

	op := text.DrawOptions{}
	op.PrimaryAlign = st.Align
	op.GeoM.Translate(nx, ny)
	text.Draw(
		screen,
		st.Text,
		&text.GoTextFace{
			Source: st.Font,
			Size:   nsize,
		},
		&op,
	)
}
