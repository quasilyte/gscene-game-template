package assets

import (
	resource "github.com/quasilyte/ebitengine-resource"

	_ "image/png"
)

const (
	AudioNone resource.AudioID = iota

	AudioError
	AudioClick

	_audioLastID
)

func registerAudioResources(loader *resource.Loader) {
	audioResources := map[resource.AudioID]resource.AudioInfo{
		AudioError: {Path: "audio/sfx/error.wav", Volume: -0.1, Group: SoundGroupEffect},
		AudioClick: {Path: "audio/sfx/button_click_soft.wav", Group: SoundGroupEffect},
	}

	for id, res := range audioResources {
		loader.AudioRegistry.Set(id, res)
		loader.LoadAudio(id)
	}
}

func GetLastAudioID() resource.AudioID {
	return _audioLastID
}
