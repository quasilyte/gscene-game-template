package assets

import resource "github.com/quasilyte/ebitengine-resource"

func registerImageResources(loader *resource.Loader) {
	resources := map[resource.ImageID]resource.ImageInfo{
		ImageShaderNoise: {Path: "$image/shader/noise.png"},

		ImageUIButtonIdle:     {Path: "$image/gui/button_idle.png"},
		ImageUIButtonPressed:  {Path: "$image/gui/button_pressed.png"},
		ImageUIButtonHover:    {Path: "$image/gui/button_hover.png"},
		ImageUIButtonDisabled: {Path: "$image/gui/button_disabled.png"},
		ImageUIPanel:          {Path: "$image/gui/panel.png"},
	}

	for id, info := range resources {
		loader.ImageRegistry.Set(id, info)
		loader.LoadImage(id)
	}
}

const (
	ImageNone resource.ImageID = iota

	ImageShaderNoise

	ImageUIButtonIdle
	ImageUIButtonPressed
	ImageUIButtonHover
	ImageUIButtonDisabled
	ImageUIPanel

	_imageLastID
)

func GetLastImageID() resource.ImageID {
	return _imageLastID
}
