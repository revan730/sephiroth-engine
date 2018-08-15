package ui

import (
	"fmt"
	"image/color"

	"github.com/faiface/pixel/pixelgl"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/text"
)

type Label struct {
	text    *text.Text
	rawText string
}

func NewLabel(text string, orig pixel.Vec, size float64, font string) *Label {
	label := &Label{
		text:    NewText(font, size, orig),
		rawText: text,
	}
	label.SetText(label.rawText)
	return label
}

func (l *Label) SetText(text string) {
	l.text.Clear()
	fmt.Fprintln(l.text, text)
}

func (l *Label) SetColor(color color.RGBA) {
	l.text.Color = color
	l.SetText(l.rawText)
}

func (l *Label) Draw(win *pixelgl.Window, im pixel.Matrix) {
	l.text.Draw(win, im)
}
