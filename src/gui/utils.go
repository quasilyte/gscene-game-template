package gui

import (
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	resource "github.com/quasilyte/ebitengine-resource"
)

func loadNineSliced(l *resource.Loader, id resource.ImageID, offsetX, offsetY int) *image.NineSlice {
	i := l.LoadImage(id).Data
	return nineSliceImage(i, offsetX, offsetY)
}

func nineSliceImage(i *ebiten.Image, offsetX, offsetY int) *image.NineSlice {
	size := i.Bounds().Size()
	w := size.X
	h := size.Y
	return image.NewNineSlice(i,
		[3]int{offsetX, w - 2*offsetX, offsetX},
		[3]int{offsetY, h - 2*offsetY, offsetY},
	)
}

func (b *Builder) loadNineSlicedImage(id resource.ImageID, offsetX, offsetY int) *image.NineSlice {
	i := b.loader.LoadImage(id).Data
	return nineSliceImage(i, offsetX, offsetY)
}

func (b *Builder) loadImage(id resource.ImageID) *ebiten.Image {
	return b.loader.LoadImage(id).Data
}

func (b *Builder) AddTooltip(w *widget.Widget, fn func() AnyWidget) {
	needRefresh := true // Need to create it first, so start with true
	w.CursorExitEvent.AddHandler(func(args interface{}) {
		needRefresh = true
	})

	tooltipPanel := b.NewPanel(PanelConfig{})
	var content AnyWidget
	tt := widget.NewToolTip(
		widget.ToolTipOpts.Content(tooltipPanel),
		widget.ToolTipOpts.ToolTipUpdater(func(c widget.Containerer) {
			if needRefresh {
				needRefresh = false
				content = fn()
				c.RemoveChildren() // Make it re-entrant
				c.AddChild(content)
			}
		}),
	)
	w.ToolTips = append(w.ToolTips, tt)
}
