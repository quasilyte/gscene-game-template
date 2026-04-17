package gui

import (
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
)

type panelDefaults struct {
	image   *image.NineSlice
	padding widget.Insets
}

func (b *Builder) loadPanelWidget() {
	idle := b.loadNineSlicedImage("image/gui/panel.png", 5, 5)
	b.panel = &panelDefaults{
		image: idle,
		padding: widget.Insets{
			Left:   8,
			Right:  8,
			Top:    8,
			Bottom: 12,
		},
	}
}

type PanelConfig struct {
	LayoutData any

	MinWidth  int
	MinHeight int
}

func (b *Builder) NewPanel(config PanelConfig) *widget.Container {
	defaults := b.panel

	padding := &defaults.padding

	var ld any
	if config.LayoutData != nil {
		ld = config.LayoutData
	} else {
		ld = widget.AnchorLayoutData{
			HorizontalPosition: widget.AnchorLayoutPositionCenter,
			VerticalPosition:   widget.AnchorLayoutPositionCenter,
		}
	}

	opts := []widget.ContainerOpt{
		widget.ContainerOpts.BackgroundImage(defaults.image),
		widget.ContainerOpts.Layout(widget.NewAnchorLayout(
			widget.AnchorLayoutOpts.Padding(padding),
		)),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(ld),
			widget.WidgetOpts.MinSize(config.MinWidth, config.MinHeight),
		),
	}

	panel := widget.NewContainer(opts...)

	return panel
}
