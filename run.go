package sephiroth

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

// NewGame returns main game function which must be passed to pixelgl.Run from main
// game function
func NewGame(*cfg GameConfig) func() {

	return func() {
		winCfg := pixelgl.WindowConfig{
			Title: cfg.WindowTitle,
			Bounds: pixel.R(0,0, cfg.WindowWidth, cfg.WindowHeight),
		}
		win, err := pixelgl.NewWindow(winCfg)
		if err != nil {
			panic(err)
		}
	}
}