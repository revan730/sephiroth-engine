package game

type GameConfig struct {
	WindowTitle       string
	WindowHeight      float64
	WindowWidth       float64
	StartScene        GameScene
	StartScenePayload interface{}
}
