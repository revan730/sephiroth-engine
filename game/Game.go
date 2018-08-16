package game

import (
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Game struct {
	Config GameConfig
	Scenes *SceneStack
}

// Run is the main game function, responsible for initialization,
// event processing and drawing
func (g *Game) Run() {
	winCfg := pixelgl.WindowConfig{
		Title:  g.Config.WindowTitle,
		Bounds: pixel.R(0, 0, g.Config.WindowWidth, g.Config.WindowHeight),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(winCfg)
	if err != nil {
		panic(err)
	}
	g.Config.StartScene.SetWindow(win)
	g.Scenes.PushScene(g.Config.StartScene, g.Config.StartScenePayload)
	fps := time.Tick(time.Second / 120) // TODO: Make configurable
	for !win.Closed() {
		// Event loop, drawing
		g.Scenes.Update()
		g.Scenes.Draw()
		win.Update()
		<-fps
	}
}
