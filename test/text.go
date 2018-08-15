package main

import (
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	sephiroth "github.com/revan730/sephiroth-engine"
	"github.com/revan730/sephiroth-engine/game"
	"github.com/revan730/sephiroth-engine/ui"
	"golang.org/x/image/colornames"
)

type TestScene struct {
	textPrinted bool
	text        *text.Text
	label       *ui.Label
	window      *pixelgl.Window
}

func (t *TestScene) Update() {
	if !t.textPrinted {
		fmt.Println("Hello from test scene!")
		t.textPrinted = true
	}
}

func (t *TestScene) Draw() {
	t.text.Draw(t.window, pixel.IM.Scaled(t.text.Orig, 2).Moved(t.window.Bounds().Center().Sub(t.text.Bounds().Center())))
	t.label.Draw(t.window, pixel.IM.Scaled(t.text.Orig, 2).Moved(t.window.Bounds().Center().Sub(t.text.Bounds().Center())))
}

func (t *TestScene) OnStart(payload interface{}) {
	t.text = ui.NewText("", 0, pixel.V(0, 0))
	t.label = ui.NewLabel("I'm a label", pixel.V(0, -50), 0, "")
	fmt.Fprintln(t.text, "Hello World!")
	t.text.Color = colornames.Green
	fmt.Fprintln(t.text, "Multiple colors")
	t.label.SetColor(colornames.Blueviolet)
	t.label.SetText("I'am a label, you can change my text and color")
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
