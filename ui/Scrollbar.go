package ui

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

type Scrollbar struct {
	scrollPos   int
	barColor    pixel.RGBA
	seekerColor pixel.RGBA
	scrollRect  pixel.Rect
	seekerRect  pixel.Rect
}

func NewScrollbar() *Scrollbar {
	return &Scrollbar{}
}

func (s *Scrollbar) Draw(tg pixel.Target, im pixel.Matrix) {
	imd := imdraw.New(nil)
	imd.Color = s.barColor
	imd.Push(s.scrollRect.Min, s.scrollRect.Max)
	imd.Rectangle(0)
	imd.Color = s.seekerColor
	imd.Push(s.seekerRect.Min, s.seekerRect.Max)
	imd.Rectangle(0)
	imd.Draw(tg)
}

func (s *Scrollbar) Center() pixel.Vec {
	return s.scrollRect.Center()
}
