package main

import (
	"fmt"

	"github.com/faiface/pixel/pixelgl"
	sephiroth "github.com/revan730/sephiroth-engine"
	"github.com/revan730/sephiroth-engine/game"
	"github.com/revan730/sephiroth-engine/ui"
)

type TestScene struct {
	textPrinted    bool
	text text.Text
	window *pixelgl.Window
}

func (t *TestScene) Update() {
	if !t.textPrinted {
		fmt.Println("Hello from test scene!")
		t.textPrinted = true
	}
}

func (t *TestScene) Draw() {
	t.text.Draw(t.window)
}

func (t *TestScene) OnStart(payload interface{}) {
	t.text := ui.NewText("", 0, pixel.V(100, 500))
	fmt.Printf("Test", t.text)
}

func (t *TestScene) SetWindow(window *pixelgl.Window) {
	t.window = window
}

func (t *TestScene) OnDestroy() {
	fmt.Println("Test scene destoying")
}

func (t *TestScene) OnPause() {
	return
}

func (t *TestScene) OnResume() {
	return
}

func main() {
	cfg := game.GameConfig{
		WindowTitle:       "Sephiroth Engine",
		WindowHeight:      720,
		WindowWidth:       1280,
		StartScene:        &TestScene{},
		StartScenePayload: "much string",
	}
	game := sephiroth.NewGame(cfg)
	pixelgl.Run(game.Run)
}
