package assets

import (
	resource "github.com/quasilyte/ebitengine-resource"
)

func registerShaderResources(loader *resource.Loader) {
	resources := map[resource.ShaderID]resource.ShaderInfo{
		ShaderMelt: {Path: "$shader/melt.shader.go"},
		ShaderCRT:  {Path: "$shader/crt.shader.go"},
	}

	for id, res := range resources {
		loader.ShaderRegistry.Set(id, res)
		loader.LoadShader(id)
	}
}

const (
	ShaderNone resource.ShaderID = iota

	ShaderMelt
	ShaderCRT
)
