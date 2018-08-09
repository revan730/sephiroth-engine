package sephiroth

// NewGame returns main game function which must be passed to pixelgl.Run from main
// game function
func NewGame(cfg GameConfig) *Game {
	return &Game{
		config: cfg,
		scenes: NewSceneStack(),
	}
}
