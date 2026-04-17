package gui

import (
	"fmt"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
	input "github.com/quasilyte/ebitengine-input"
	resource "github.com/quasilyte/ebitengine-resource"
	sound "github.com/quasilyte/ebitengine-sound"
	"github.com/quasilyte/gmath"
	"github.com/quasilyte/gscene"
)

type AnyWidget = widget.PreferredSizeLocateableWidget

type Builder struct {
	Ready bool

	button *buttonDefaults
	panel  *panelDefaults

	loader *resource.Loader
	audio  *sound.System

	object     *SceneObject
	uiRoot     *ebitenui.UI
	input      *input.Handler
	ScreenSize gmath.Vec

	inputHandled bool
	windowStack  []*window
}

type Context struct {
	Loader *resource.Loader

	Audio *sound.System

	Input *input.Handler
}

func NewBuilder(ctx Context) *Builder {
	b := &Builder{
		loader: ctx.Loader,
		audio:  ctx.Audio,
		input:  ctx.Input,
	}
	return b
}

func (b *Builder) GetTopWindow() *window {
	if len(b.windowStack) > 0 {
		return b.windowStack[len(b.windowStack)-1]
	}
	return nil
}

func (b *Builder) Reset() {
	b.Ready = false
	b.object = nil
	b.windowStack = b.windowStack[:0]
	fmt.Println("reset")
}

func (b *Builder) Init() {
	b.loadPanelWidget()
	b.loadButtonWidget()
}

func (b *Builder) Update(delta float64) {
	b.inputHandled = false

	if b.object != nil {
		if !b.Ready {
			b.Ready = b.object.drawn
		}
	}

	if w := b.GetTopWindow(); w != nil {
		w.Update()
	}
}

func (b *Builder) BuildAt(scene *gscene.Scene, root *widget.Container, layer int) *ebitenui.UI {
	anchor := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)
	anchor.AddChild(root)

	b.object = newSceneObject(anchor)
	scene.AddGraphics(b.object, layer)
	scene.AddObject(b.object)

	b.uiRoot = b.object.ui

	return b.object.ui
}

func (b *Builder) Build(scene *gscene.Scene, root *widget.Container) *ebitenui.UI {
	return b.BuildAt(scene, root, 0)
}
