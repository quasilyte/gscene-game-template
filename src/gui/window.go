package gui

import (
	"image"
	"image/color"

	"mygame/controls"

	"github.com/ebitenui/ebitenui"
	eimage "github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/quasilyte/gmath"
	"github.com/quasilyte/gsignal"
	"github.com/quasilyte/gslices"
)

type window struct {
	b         *Builder
	ui        *ebitenui.UI
	w         *widget.Window
	bg        *widget.Window
	width     int
	height    int
	seq       int
	closeFunc func(submit bool)
	pad       image.Point

	autoclose      bool
	open           bool
	closeOnUnfocus bool
	closeOnEsc     bool
	clickPressed   bool
	isTooltip      bool

	onHandleInput func() bool

	EventOpened gsignal.Event[gsignal.Void]
	EventClosed gsignal.Event[gsignal.Void]
}

func newWindow(b *Builder, uiRoot *ebitenui.UI) *window {
	return &window{
		b:  b,
		ui: uiRoot,
	}
}

func (w *window) Update() {
	if !w.IsOpen() {
		return
	}

	if !w.isTooltip {
		if w.b.windowStack[len(w.b.windowStack)-1] == w {
			if w.handleInput() {
				w.b.inputHandled = true
			}
		}
	}
}

func (w *window) handleInput() bool {
	if w.b.inputHandled {
		return false
	}

	if w.onHandleInput != nil {
		if w.onHandleInput() {
			return true
		}
	}

	if w.closeOnEsc && w.b.input.ActionIsJustPressed(controls.ActionBack) {
		w.Close()
		return true
	}

	if w.clickPressed {
		if !w.closeOnUnfocus {
			return false
		}

		hasClick := w.b.input.ActionIsJustReleased(controls.ActionRightClick) ||
			w.b.input.ActionIsJustReleased(controls.ActionClick)
		if hasClick {
			w.clickPressed = false
			pos := w.b.input.CursorPos()
			rect := w.w.GetContainer().GetWidget().Rect
			if !pos.ToStd().In(rect) {
				w.Close()
				return true
			}
		}
	} else {
		if w.b.input.ActionIsJustPressed(controls.ActionClick) || w.b.input.ActionIsJustPressed(controls.ActionRightClick) {
			w.clickPressed = true
		}
	}

	return false
}

func (w *window) IsOpen() bool {
	return w.open
}

func (w *window) doClose(submit bool) {
	if !w.IsOpen() {
		return
	}

	w.closeFunc(submit)
	w.w = nil
	w.seq++
}

func (w *window) Close() {
	w.doClose(false)
}

func (w *window) SetPos(pos gmath.Vec) {
	w.w.SetLocation(w.calcLocation(pos))
}

func (w *window) calcLocation(pos gmath.Vec) image.Rectangle {
	pad := w.pad
	rect := image.Rectangle{
		Min: pos.ToStd(),
		Max: pos.ToStd().Add(image.Pt(w.width, w.height)),
	}
	if rect.Max.X+pad.X > int(w.b.ScreenSize.X) {
		if rect.Max.X+pad.X-w.width/2 > int(w.b.ScreenSize.X) {
			rect.Min.X -= w.width + pad.X
			rect.Max.X -= w.width + pad.X
		} else {
			rect.Min.X -= w.width/2 + pad.X
			rect.Max.X -= w.width/2 + pad.X
		}
	}
	if rect.Max.Y+pad.Y >= int(w.b.ScreenSize.Y) {
		overflow := rect.Max.Y - int(w.b.ScreenSize.Y)
		rect.Min.Y -= overflow + pad.Y
		rect.Max.Y -= overflow + pad.Y
	}
	return rect
}

func (w *window) EnsureClosed() {
	if w.IsOpen() {
		w.Close()
	}
}

type WindowOpenConfig struct {
	Content     AnyWidget
	Blur        bool
	Width       int
	Height      int
	Pos         gmath.Vec
	OnClosed    func(submit bool)
	HandleInput func() bool

	HideCloseButton       bool
	DisableCloseOnUnfocus bool
	DisableCloseOnEsc     bool
}

type windowOpenConfig struct {
	Width          int
	Height         int
	Blur           bool
	CloseButton    bool
	CloseOnUnfocus bool
	CloseOnEsc     bool
	IsTooltip      bool
	BlockLower     bool
	HandleInput    func() bool
	OnClose        func(submit bool)
	OnUpdate       func()

	Contents AnyWidget
}

func (b *Builder) OpenWindow(config WindowOpenConfig) *window {
	winPanel := b.NewPanel(PanelConfig{})
	winPanel.AddChild(config.Content)

	w := newWindow(b, b.uiRoot)

	width, height := winPanel.PreferredSize()
	if config.Width != 0 {
		width = config.Width
	}
	if config.Height != 0 {
		height = config.Height
	}

	pos := config.Pos
	if pos.IsZero() {
		pos = gmath.Vec{
			X: (b.ScreenSize.X * 0.5) - (float64(width) * 0.5),
			Y: (b.ScreenSize.Y * 0.5) - (float64(height) * 0.5),
		}
	}
	w.Open(winPanel, pos, windowOpenConfig{
		Width:          width,
		Height:         height,
		Blur:           config.Blur,
		CloseButton:    !config.HideCloseButton,
		CloseOnUnfocus: !config.DisableCloseOnUnfocus,
		CloseOnEsc:     !config.DisableCloseOnEsc,
		OnClose:        config.OnClosed,
		HandleInput:    config.HandleInput,
		BlockLower:     true,
	})

	return w
}

func (w *window) Open(contents *widget.Container, pos gmath.Vec, config windowOpenConfig) {
	windowContents := contents
	if config.CloseButton {
		stackedContents := widget.NewContainer(
			widget.ContainerOpts.Layout(widget.NewStackedLayout()),
		)
		stackedContents.AddChild(contents)
		{
			buttonAnchor := widget.NewContainer(
				widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
			)
			stackedContents.AddChild(buttonAnchor)
			buttonAnchor.AddChild(w.b.NewTinyButton(TinyButtonConfig{
				Text: "x",
				OnClick: func() {
					w.EnsureClosed()
				},
				LayoutData: widget.AnchorLayoutData{
					HorizontalPosition: widget.AnchorLayoutPositionEnd,
				},
			}))
		}
		windowContents = stackedContents
	}

	w.width = config.Width
	w.height = config.Height
	w.w = widget.NewWindow(
		widget.WindowOpts.DisableRelayering(true),
		widget.WindowOpts.BlockLower(config.BlockLower),
		widget.WindowOpts.Contents(windowContents),
		widget.WindowOpts.Location(w.calcLocation(pos)),
	)

	closeBgFunc := func() {}
	if config.Blur {
		bg := eimage.NewNineSliceColor(color.NRGBA{0, 0, 0, 100})
		bgContainer := widget.NewContainer(
			widget.ContainerOpts.BackgroundImage(bg),
			widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
		)
		w.bg = widget.NewWindow(
			widget.WindowOpts.DisableRelayering(true),
			widget.WindowOpts.Contents(bgContainer),
			widget.WindowOpts.BlockLower(false),
			widget.WindowOpts.Location(image.Rect(0, 0, int(w.b.ScreenSize.X), int(w.b.ScreenSize.Y))),
		)
		closeBgFunc = w.ui.AddWindow(w.bg)
	} else {
		w.bg = nil
	}

	w.clickPressed = false
	closeFunc := w.ui.AddWindow(w.w)
	w.open = true
	w.isTooltip = config.IsTooltip
	w.onHandleInput = config.HandleInput
	w.closeOnEsc = config.CloseOnEsc
	w.closeOnUnfocus = config.CloseOnUnfocus
	if !config.IsTooltip {
		if top := w.b.GetTopWindow(); top != nil {
			top.clickPressed = false
		}
		w.b.windowStack = append(w.b.windowStack, w)
	}
	w.closeFunc = func(submit bool) {
		w.open = false
		if !config.IsTooltip {
			gslices.Delete(&w.b.windowStack, w)
		}
		w.EventClosed.Emit(gsignal.Void{})

		closeBgFunc()
		closeFunc()
		if config.OnClose != nil {
			config.OnClose(submit)
		}
	}
	w.w.SetCloseFunction(func() {
		w.closeFunc(false)
	})
	w.EventOpened.Emit(gsignal.Void{})
}
