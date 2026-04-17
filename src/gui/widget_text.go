package gui

import (
	"image/color"
	"strings"

	"mygame/assets"
	"mygame/styles"

	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type TextConfig struct {
	Text     string
	Font     *text.Face
	Color    color.Color
	MinWidth int
	MaxWidth int

	LayoutData any

	AlignLeft   bool
	ForceBBCode bool
	AlignRight  bool
	AlignTop    bool
}

func (b *Builder) SimpleTooltip(s string) func() AnyWidget {
	return func() AnyWidget {
		return b.NewText(TextConfig{
			Text: s,
		})
	}
}

func (b *Builder) NewText(config TextConfig) *widget.Text {
	var clr color.Color = styles.NormalTextColor.Color()
	if config.Color != nil {
		clr = config.Color
	}

	ff := assets.FontTiny
	if config.Font != nil {
		ff = config.Font
	}

	verticalPos := widget.TextPositionCenter
	if config.AlignTop {
		verticalPos = widget.TextPositionStart
	}

	opts := []widget.TextOpt{
		widget.TextOpts.Text(config.Text, ff, clr),
	}
	if config.LayoutData != nil {
		opts = append(opts, widget.TextOpts.WidgetOpts(widget.WidgetOpts.LayoutData(config.LayoutData)))
	}
	if config.MinWidth != 0 {
		opts = append(opts, widget.TextOpts.WidgetOpts(widget.WidgetOpts.MinSize(config.MinWidth, 0)))
	}
	if config.MaxWidth != 0 {
		opts = append(opts, widget.TextOpts.MaxWidth(float64(config.MaxWidth)))
	}
	switch {
	case config.AlignLeft:
		opts = append(opts, widget.TextOpts.Position(widget.TextPositionStart, verticalPos))
	case config.AlignRight:
		opts = append(opts, widget.TextOpts.Position(widget.TextPositionEnd, verticalPos))
	default:
		opts = append(opts, widget.TextOpts.Position(widget.TextPositionCenter, verticalPos))
	}
	if config.ForceBBCode || strings.Contains(config.Text, "[color=") {
		opts = append(opts, widget.TextOpts.ProcessBBCode(true))
	}
	return widget.NewText(opts...)
}
