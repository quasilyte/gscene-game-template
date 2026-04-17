package gui

import "github.com/ebitenui/ebitenui/widget"

func (b *Builder) NewTopLevelGridRows() *widget.Container {
	return widget.NewContainer(
		widget.ContainerOpts.Layout(
			widget.NewGridLayout(
				widget.GridLayoutOpts.Columns(1),
				widget.GridLayoutOpts.Spacing(8, 8),
			),
		),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
				VerticalPosition:   widget.AnchorLayoutPositionCenter,
			}),
		),
	)
}

func (b *Builder) NewPanelGridRows(minWidth int) *widget.Container {
	return widget.NewContainer(
		widget.ContainerOpts.Layout(
			widget.NewGridLayout(
				widget.GridLayoutOpts.Columns(1),
				widget.GridLayoutOpts.Spacing(8, 8),
				widget.GridLayoutOpts.DefaultStretch(true, false),
			),
		),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
				VerticalPosition:   widget.AnchorLayoutPositionStart,
			}),
			widget.WidgetOpts.MinSize(minWidth, 0),
		),
	)
}

func (b *Builder) NewTextGridRows(minWidth int) *widget.Container {
	return widget.NewContainer(
		widget.ContainerOpts.Layout(
			widget.NewGridLayout(
				widget.GridLayoutOpts.Columns(1),
				widget.GridLayoutOpts.Spacing(8, 2),
				widget.GridLayoutOpts.DefaultStretch(true, false),
			),
		),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
				VerticalPosition:   widget.AnchorLayoutPositionStart,
			}),
			widget.WidgetOpts.MinSize(minWidth, 0),
		),
	)
}

type GridColsConfig struct {
	Cols       int
	MinWidth   int
	ColSpacing int
	ColScale   []bool
}

func (b *Builder) NewGridCols(config GridColsConfig) *widget.Container {
	return widget.NewContainer(
		widget.ContainerOpts.Layout(
			widget.NewGridLayout(
				widget.GridLayoutOpts.Columns(config.Cols),
				widget.GridLayoutOpts.Spacing(config.ColSpacing, 8),
				widget.GridLayoutOpts.Stretch(config.ColScale, nil),
			),
		),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
				VerticalPosition:   widget.AnchorLayoutPositionStart,
				StretchHorizontal:  true,
			}),
			widget.WidgetOpts.MinSize(config.MinWidth, 0),
		),
	)
}
