package sephiroth

import "github.com/revan730/sephiroth-engine/game"

// NewGame returns main game function which must be passed to pixelgl.Run from main
// game function
func NewGame(cfg game.GameConfig) *game.Game {
	// TODO: move startscene from gameconfig so it can be passed to game scenes
	return &game.Game{
		Config: cfg,
		Scenes: game.NewSceneStack(),
	}
}
