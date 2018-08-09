package sephiroth

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Game struct {
	config GameConfig
	scenes *SceneStack
}

// Run is the main game function, responsible for initialization,
// event processing and drawing
func (g *Game) Run() {
	winCfg := pixelgl.WindowConfig{
		Title:  g.config.WindowTitle,
		Bounds: pixel.R(0, 0, g.config.WindowWidth, g.config.WindowHeight),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(winCfg)
	if err != nil {
		panic(err)
	}
	g.scenes.PushScene(g.config.StartScene, g.config.StartScenePayload)
	for !win.Closed() {
		// Event loop, drawing
		g.scenes.Update()
		g.scenes.Draw()
		win.Update()
	}
}
