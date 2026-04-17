package assets

import resource "github.com/quasilyte/ebitengine-resource"

func registerImageResources(loader *resource.Loader) {
	resources := map[resource.ImageID]resource.ImageInfo{
		ImageShaderNoise: {Path: "$image/shader/noise.png"},
	}

	for id, info := range resources {
		loader.ImageRegistry.Set(id, info)
		loader.LoadImage(id)
	}
}

const (
	ImageNone resource.ImageID = iota

	ImageShaderNoise

	_imageLastID
)

func GetLastImageID() resource.ImageID {
	return _imageLastID
}
