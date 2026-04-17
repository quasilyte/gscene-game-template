package game

import (
	"mygame/assets"
	"mygame/gui"
	"mygame/viewport"

	input "github.com/quasilyte/ebitengine-input"
	resource "github.com/quasilyte/ebitengine-resource"
	sound "github.com/quasilyte/ebitengine-sound"
	"github.com/quasilyte/gmath"
	"github.com/quasilyte/gscene"
	// "github.com/quasilyte/ld58-game/src/assets"
	// "github.com/quasilyte/ld58-game/src/gui"
	// "github.com/quasilyte/ld58-game/src/viewport"
)

// <<edit>> the context needs to hold everything
// your game needs to access globally.
// Add fields as necessary.

var G *GlobalContext

type GlobalContext struct {
	SceneManager *gscene.Manager

	Playlist    *sound.Playlist
	Audio       sound.System
	SoundVolume int
	MusicVolume int

	WindowSize gmath.Vec

	Session *Session

	Input *input.Handler

	Camera *viewport.Camera

	Loader *resource.Loader

	Rand gmath.Rand

	UI *gui.Builder

	Controllers ControllerRegistry
}

func ChangeScene(c gscene.Controller) {
	G.Camera = nil
	G.UI.Reset()

	G.SceneManager.ChangeScene(c)
}

func (ctx *GlobalContext) PlaySound(id resource.AudioID) {
	resourceID := id
	numSamples := 1
	// numSamples := assets.NumSamples(id)
	if numSamples > 0 {
		resourceID += resource.AudioID(ctx.Rand.IntRange(0, numSamples-1))
	}
	ctx.Audio.PlaySound(resourceID)
}

func (ctx *GlobalContext) AdjustVolumeLevels() {
	ctx.Audio.SetGroupVolume(assets.SoundGroupMusic,
		assets.VolumeMultiplier(ctx.MusicVolume))
	ctx.Audio.SetGroupVolume(assets.SoundGroupEffect,
		assets.VolumeMultiplier(ctx.SoundVolume))
}

func (ctx *GlobalContext) AdjustMusic() {
	ctx.Playlist.SetPaused(true)
	ctx.Playlist.SetPaused(false)
}
