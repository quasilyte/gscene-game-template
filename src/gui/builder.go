package gui

import (
	"mygame/dat"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
	resource "github.com/quasilyte/ebitengine-resource"
	sound "github.com/quasilyte/ebitengine-sound"
	"github.com/quasilyte/gscene"
)

type AnyWidget = widget.PreferredSizeLocateableWidget

type Builder struct {
	Ready bool

	button *buttonDefaults
	panel  *panelDefaults

	content *dat.Content
	loader  *resource.Loader
	audio   *sound.System

	object *SceneObject
}

type Context struct {
	Loader *resource.Loader

	Content *dat.Content

	Audio *sound.System
}

func NewBuilder(ctx Context) *Builder {
	b := &Builder{
		loader:  ctx.Loader,
		audio:   ctx.Audio,
		content: ctx.Content,
	}
	return b
}

func (b *Builder) Reset() {
	b.Ready = false
	b.object = nil
}

func (b *Builder) Init() {
	b.loadPanelWidget()
	b.loadButtonWidget()
}

func (b *Builder) Update(delta float64) {
	if b.object != nil {
		b.object.Update(delta)
		if !b.Ready {
			b.Ready = b.object.drawn
		}
	}
}

func (b *Builder) BuildAt(scene *gscene.Scene, root *widget.Container, layer int) *ebitenui.UI {
	anchor := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)
	anchor.AddChild(root)

	b.object = newSceneObject(anchor)
	scene.AddGraphics(b.object, layer)

	return b.object.ui
}

func (b *Builder) Build(scene *gscene.Scene, root *widget.Container) *ebitenui.UI {
	return b.BuildAt(scene, root, 0)
}
