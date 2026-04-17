package gui

import (
	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/quasilyte/gscene"
)

type SceneObject struct {
	drawn bool
	ui    *ebitenui.UI
}

func newSceneObject(root *widget.Container) *SceneObject {
	ui := &ebitenui.UI{
		Container:           root,
		DisableDefaultFocus: true,
	}
	return &SceneObject{
		ui: ui,
	}
}

func (o *SceneObject) IsDisposed() bool { return false }

func (o *SceneObject) Init(scene *gscene.Scene) {}

func (o *SceneObject) Update(delta float64) {
	o.ui.Update()
}

func (o *SceneObject) Draw(dst *ebiten.Image) {
	o.ui.Draw(dst)
	o.drawn = true
}
