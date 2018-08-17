package ui

import (
	"strconv"

	"github.com/faiface/pixel/pixelgl"

	"github.com/faiface/pixel"
)

type Menu struct {
	labels       []*Label
	canvas       *pixelgl.Canvas
	scrollPos    pixel.Vec
	scrollRegion pixel.Matrix
	activePos    int
	isActive     bool
}

func NewMenu() *Menu {
	menu := &Menu{
		labels:       make([]*Label, 10),
		canvas:       pixelgl.NewCanvas(pixel.R(-100, -100, 100, 100)),
		scrollPos:    pixel.ZV,
		scrollRegion: pixel.Matrix{},
		activePos:    0,
		isActive:     true,
	}

	// TODO: Only here to debug this shit, move to func params
	for i, _ := range menu.labels {
		menu.labels[i] = NewLabel(strconv.Itoa(i), pixel.V(0, 50.0*float64(i)), 0, "")
	}
	return menu
}

func (m *Menu) Draw(win *pixelgl.Window, im pixel.Matrix) {
	for _, l := range m.labels {
		l.Draw(m.canvas, im)
	}
	m.canvas.Draw(win, im)
}

func (m *Menu) Bounds() pixel.Rect {
	return m.canvas.Bounds()
}
