package main

import (
	"mygame/assets"
	"mygame/dat"
	"mygame/game"
	"mygame/scenes"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/quasilyte/gmath"
	"github.com/quasilyte/gscene"
)

// This file will contain game-specific code.
// <<edit>> configure the way your game is initialized.

func (g *gameRunner) FirstScene() gscene.Controller {
	// This will define the game's entry point in terms of scenes.
	return scenes.NewMainMenuController()
}

func (g *gameRunner) InitContextCustom() {
	game.G.WindowSize = gmath.Vec{
		X: 1920 / 2,
		Y: 1080 / 2,
	}

	ebiten.SetWindowTitle(dat.GameTitle)
	ebiten.SetFullscreen(true)

	game.G.Playlist.Add(assets.AudioExampleMusic)

	g.initControllerRegistry()
}

func (g *gameRunner) initControllerRegistry() {
	// The controllers registry is needed when a scene A needs to go to
	// scene B, but it can't import the package containing B due to
	// a circular imports error.
	// You would use a registry to get around that - any scene package
	// should have game.G access, and to required controllers through it.
	//
	// Note: the main menu controller is added here as an example.
	// Do not add new controllers to the registry unless you have
	// a circular import problem to solve.

	game.G.Controllers.MainMenuController = func() gscene.Controller {
		return scenes.NewMainMenuController()
	}
}
