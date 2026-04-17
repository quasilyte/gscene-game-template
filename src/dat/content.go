package dat

import (
	"fmt"
	"mygame/assets"
	"mygame/datscan"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	graphics "github.com/quasilyte/ebitengine-graphics"
	resource "github.com/quasilyte/ebitengine-resource"
)

type Content struct {
	Images  map[string]resource.ImageID
	Audio   map[string]resource.AudioID
	Shaders map[string]*ebiten.Shader
}

func ScanContent(path string, l *resource.Loader) *Content {
	content := &Content{
		Images:  make(map[string]resource.ImageID),
		Audio:   make(map[string]resource.AudioID),
		Shaders: make(map[string]*ebiten.Shader),
	}

	filesByType := make(map[datscan.FileKind][]datscan.File, 64)
	err := datscan.Walk(datscan.WalkConfig{
		GameDataPath: path,
		Error: func(f *datscan.File, err error) {
			fmt.Printf("SCAN ERROR %s\n", err.Error())
		},
		Visit: func(f *datscan.File) error {
			filesByType[f.Kind] = append(filesByType[f.Kind], *f)
			return nil
		},
	})
	if err != nil {
		panic(fmt.Errorf("%q: %w", path, err))
	}

	imageIDSeq := assets.GetLastImageID()
	for _, f := range filesByType[datscan.FileImage] {
		id := imageIDSeq
		imageIDSeq++
		content.Images[f.Path] = id
		l.ImageRegistry.Set(id, resource.ImageInfo{
			Path: f.Path,
		})
	}

	audioIDSeq := assets.GetLastAudioID()
	for _, f := range filesByType[datscan.FileSound] {
		id := audioIDSeq
		audioIDSeq++
		content.Audio[f.Path] = id
		l.AudioRegistry.Set(id, resource.AudioInfo{
			Path:   f.Path,
			Volume: f.Arg,
			Group:  assets.SoundGroupEffect,
		})
	}
	for _, f := range filesByType[datscan.FileMusicOGG] {
		id := audioIDSeq
		audioIDSeq++
		content.Audio[f.Path] = id
		l.AudioRegistry.Set(id, resource.AudioInfo{
			Path:   f.Path,
			Volume: f.Arg,
			Group:  assets.SoundGroupMusic,
		})
	}

	graphics.CompileShaders()
	for _, f := range filesByType[datscan.FileShaderSource] {
		src, err := os.ReadFile(f.AbsPath)
		if err != nil {
			panic(fmt.Errorf("%q: %w", f.Path, err))
		}
		compiled, err := ebiten.NewShader(src)
		if err != nil {
			panic(fmt.Errorf("%q: %w", f.Path, err))
		}
		content.Shaders[f.Path] = compiled
	}

	return content
}
