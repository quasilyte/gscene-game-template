package assets

import (
	"embed"
	"io"
	"path/filepath"
	"strings"

	resource "github.com/quasilyte/ebitengine-resource"
)

//go:embed all:_data
var gameAssets embed.FS

const (
	SoundGroupEffect uint = iota
	SoundGroupMusic
)

func MakeOpenAssetFunc(gamedataFolder string) func(path string) io.ReadCloser {
	return func(path string) io.ReadCloser {
		if strings.HasPrefix(path, "$") {
			f, err := gameAssets.Open("_data/" + path[len("$"):])
			if err != nil {
				panic(err)
			}
			return f
		}

		f, err := openfile(filepath.Join(gamedataFolder, path))
		if err != nil {
			panic(err)
		}
		return f
	}
}

func RegisterResources(l *resource.Loader) {
	registerImageResources(l)
	registerAudioResources(l)
	registerShaderResources(l)
}

func VolumeMultiplier(level int) float64 {
	switch level {
	case 1:
		return 0.01
	case 2:
		return 0.15
	case 3:
		return 0.45
	case 4:
		return 0.8
	case 5:
		return 1.0
	default:
		return 0
	}
}
