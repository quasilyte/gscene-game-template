package main

import (
	"mygame/game"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	runner := &gameRunner{}
	bindFlags(&runner.cliArgs)
	runner.Init()

	game.G.SceneManager.ChangeScene(runner.FirstScene())

	if err := ebiten.RunGame(runner); err != nil {
		panic(err)
	}
}
