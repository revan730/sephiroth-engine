package main

import (
	"fmt"

	"github.com/faiface/pixel/pixelgl"
	sephiroth "github.com/revan730/sephiroth-engine"
)

type TestScene struct {
	textPrinted    bool
	strFromPayload string
}

func (t *TestScene) Update() {
	if !t.textPrinted {
		fmt.Println("Hello from test scene!")
		t.textPrinted = true
	}
}

func (t *TestScene) Draw() {
	return
}

func (t *TestScene) OnStart(payload interface{}) {
	fmt.Println("Got " + payload.(string) + " from payload")
	t.strFromPayload = payload.(string)
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
	cfg := sephiroth.GameConfig{
		WindowTitle:       "Sephiroth Engine",
		WindowHeight:      720,
		WindowWidth:       1280,
		StartScene:        &TestScene{},
		StartScenePayload: "much string",
	}
	game := sephiroth.NewGame(cfg)
	pixelgl.Run(game.Run)
}
