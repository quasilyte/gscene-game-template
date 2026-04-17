package main

import (
	"mygame/assets"
	"mygame/controls"
	"mygame/game"
	"mygame/gui"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	graphics "github.com/quasilyte/ebitengine-graphics"
	input "github.com/quasilyte/ebitengine-input"
	resource "github.com/quasilyte/ebitengine-resource"
	sound "github.com/quasilyte/ebitengine-sound"
	"github.com/quasilyte/gscene"
	// "github.com/quasilyte/ld58-game/src/assets"
	// "github.com/quasilyte/ld58-game/src/controls"
	// "github.com/quasilyte/ld58-game/src/game"
	// "github.com/quasilyte/ld58-game/src/gui"
)

// This file contains default things that rarely
// need to change.

type gameRunner struct {
	inputSystem input.System

	cliArgs cliArgs
}

func (g *gameRunner) Update() error {
	const delta = 1.0 / 60.0

	g.inputSystem.Update()

	if game.G.Camera != nil {
		game.G.Camera.Update(delta)
	}

	game.G.SceneManager.Update()
	game.G.UI.Update(delta)
	game.G.Audio.Update()
	game.G.Playlist.Update(delta)
	return nil
}

func (g *gameRunner) Draw(screen *ebiten.Image) {
	game.G.SceneManager.Draw(screen)
}

func (g *gameRunner) Layout(_, _ int) (int, int) {
	panic("should never happen")
}

func (g *gameRunner) LayoutF(_, _ float64) (float64, float64) {
	return game.G.WindowSize.X, game.G.WindowSize.Y
}

func (g *gameRunner) Init() {
	game.G = &game.GlobalContext{
		SoundVolume: 3,
		MusicVolume: 3,

		Session: &game.Session{},
	}
	game.G.SceneManager = gscene.NewManager()

	sampleRate := 48000
	audioContext := audio.NewContext(sampleRate)
	game.G.Loader = resource.NewLoader(audioContext)
	game.G.Loader.OpenAssetFunc = assets.MakeOpenAssetFunc(g.cliArgs.gameDataFolder)

	game.G.Audio.Init(audioContext, game.G.Loader)
	game.G.Playlist = sound.NewPlaylist(&game.G.Audio)

	game.G.Rand.SetSeed(time.Now().UnixNano())

	g.inputSystem.Init(input.SystemConfig{
		DevicesEnabled: input.AnyDevice,
	})
	game.G.Input = g.inputSystem.NewHandler(0, controls.DefaultKeymap())

	graphics.CompileShaders()

	assets.RegisterResources(game.G.Loader)

	game.G.UI = gui.NewBuilder(gui.Context{
		Loader: game.G.Loader,
		Audio:  &game.G.Audio,
	})
	game.G.UI.Init()

	g.InitContextCustom()
}
