package vvis

import (
	"bytes"
	"image"
	"image/color"

	"errors"

	"MineSweeper/stagerror"

	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type VisualHandler struct {
	actualWidth, actualHeight   float64
	virtualWidth, virtualHeight float64

	images map[string]*ebiten.Image
	fonts  map[string]*text.GoTextFaceSource

	screen *ebiten.Image

	isRelease bool
}

const ASPECT_RATIO = 2

func (vh *VisualHandler) ChangeRes(screenWidth int) {
	vh.actualWidth = float64(screenWidth)
	vh.actualHeight = vh.actualWidth / ASPECT_RATIO
}

func (vh *VisualHandler) Init(screenWidth, virtualWidth int, isRelease bool) {
	vh.actualWidth = float64(screenWidth)
	vh.actualHeight = vh.actualWidth / ASPECT_RATIO
	vh.virtualWidth = float64(virtualWidth)
	vh.virtualHeight = vh.virtualWidth / ASPECT_RATIO
	vh.isRelease = isRelease
	vh.images = map[string]*ebiten.Image{}
	vh.fonts = map[string]*text.GoTextFaceSource{}
}

func (vh *VisualHandler) InitFrame(screen *ebiten.Image) {
	vh.screen = screen
}

func (vh *VisualHandler) Translate(d float64) float64 {
	return d * vh.actualWidth / vh.virtualWidth
}

func (vh *VisualHandler) Untranslate(d float64) float64 {
	return d * vh.virtualWidth / vh.actualWidth
}

func (vh *VisualHandler) DrawRect(x, y, w, h float64, clr color.Color) {
	nx := vh.Translate(x)
	ny := vh.Translate(y)
	nw := vh.Translate(w)
	nh := vh.Translate(h)

	vector.DrawFilledRect(
		vh.screen,
		float32(nx),
		float32(ny),
		float32(nw),
		float32(nh),
		clr,
		false,
	)
}

func (vh *VisualHandler) DrawCircle(x, y, r float64, clr color.Color) {
	nx := vh.Translate(x)
	ny := vh.Translate(y)
	nr := vh.Translate(r)

	vector.DrawFilledCircle(
		vh.screen,
		float32(nx),
		float32(ny),
		float32(nr),
		clr,
		true,
	)
}

func (vh *VisualHandler) DrawLine(x1, y1, x2, y2, w float64, clr color.Color) {
	nx1 := vh.Translate(x1)
	ny1 := vh.Translate(y1)
	nx2 := vh.Translate(x2)
	ny2 := vh.Translate(y2)
	nw := vh.Translate(w)

	vector.StrokeLine(
		vh.screen,
		float32(nx1),
		float32(ny1),
		float32(nx2),
		float32(ny2),
		float32(nw),
		clr,
		true,
	)
}

func (vh *VisualHandler) DrawImage(img *ebiten.Image, x, y, w, h float64, op *ebiten.DrawImageOptions) {
	nx := vh.Translate(x)
	ny := vh.Translate(y)
	nw := vh.Translate(w)
	nh := vh.Translate(h)

	ogW, ogH := img.Bounds().Dx(), img.Bounds().Dy()
	op.GeoM.Scale(nw/float64(ogW), nh/float64(ogH))
	op.GeoM.Translate(nx, ny)
	vh.screen.DrawImage(img, op)
}

func (vh *VisualHandler) DrawImageWithRot(img *ebiten.Image, x, y, w, h, rot float64, op *ebiten.DrawImageOptions) {
	nx := vh.Translate(x)
	ny := vh.Translate(y)
	nw := vh.Translate(w)
	nh := vh.Translate(h)

	ogW, ogH := img.Bounds().Dx(), img.Bounds().Dy()

	op.GeoM.Translate(-float64(ogW/2), -float64(ogH/2))
	op.GeoM.Rotate(rot)
	op.GeoM.Translate(float64(ogW/2), float64(ogH/2))

	op.GeoM.Scale(nw/float64(ogW), nh/float64(ogH))
	op.GeoM.Translate(nx, ny)
	vh.screen.DrawImage(img, op)
}

func (vh *VisualHandler) InitImages() {
	vh.setIcon()
}

func (vh *VisualHandler) GetImage(name string) *ebiten.Image {
	i, ok := vh.images[name]
	if ok {
		return i
	}

	stagerror.SaveToLog(errors.New("Couldn't find this image: "+name), vh.isRelease)

	// Will hard error if default doesn't exist
	return vh.images["missing"]
}

func (vh *VisualHandler) setIcon() {
	ebiten.SetWindowIcon([]image.Image{vh.GetImage("icon")})
}

func (vh *VisualHandler) RegisterImage(name string, img *ebiten.Image) {
	vh.images[name] = img
}

func (vh *VisualHandler) GetFont(name string) *text.GoTextFaceSource {
	f, ok := vh.fonts[name]
	if ok {
		return f
	}

	stagerror.SaveToLog(errors.New("Couldn't find this font: "+name), vh.isRelease)

	// Will hard error if default doesn't exist
	return vh.fonts["default"]
}

func (vh *VisualHandler) DrawText(str string, size, x, y float64, drawFont *text.GoTextFaceSource, op *text.DrawOptions) {
	nx := vh.Translate(x)
	ny := vh.Translate(y)
	nsize := vh.Translate(size)

	op.GeoM.Translate(nx, ny)
	text.Draw(
		vh.screen,
		str,
		&text.GoTextFace{
			Source: drawFont,
			Size:   nsize,
		},
		op,
	)
}

func (vh *VisualHandler) LoadFonts() {
	vh.fonts = make(map[string]*text.GoTextFaceSource)

	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if err != nil {
		// Panic because it's vital
		stagerror.SaveToLog(errors.New("Couldn't load fonts"), vh.isRelease)
	}

	vh.fonts["default"] = s
	vh.fonts["button"] = s
	vh.fonts["textBox"] = s
}

func (vh *VisualHandler) LoadImage(name string, source string) {
	img, _, err := ebitenutil.NewImageFromFile(source)
	if err != nil {
		stagerror.SaveToLog(errors.New("Couldn't load image: "+source), vh.isRelease)
		return
	}
	vh.images[name] = img
}

func (vh *VisualHandler) Thinline() float64 {
	return vh.virtualWidth / vh.actualWidth
}
