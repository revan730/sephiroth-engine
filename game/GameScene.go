package game

import "github.com/faiface/pixel/pixelgl"

// GameScene is an entity responsible for
// some part of game (ex. main menu, splash screen, world map etc.)
type GameScene interface {
	Update()
	Draw()
	OnStart(payload interface{})
	OnPause()
	SetWindow(window *pixelgl.Window)
	OnResume()
	OnDestroy()
}
