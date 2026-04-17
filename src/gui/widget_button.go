package gui

import (
	"strings"

	"mygame/assets"
	"mygame/styles"

	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/quasilyte/gmath"
	"github.com/quasilyte/gsignal"
)

type buttonDefaults struct {
	image      *widget.ButtonImage
	padding    widget.Insets
	textColors *widget.ButtonTextColor
}

func (b *Builder) loadButtonWidget() {
	disabled := b.loadNineSlicedImage(assets.ImageUIButtonDisabled, 5, 5)
	idle := b.loadNineSlicedImage(assets.ImageUIButtonIdle, 5, 5)
	hover := b.loadNineSlicedImage(assets.ImageUIButtonHover, 5, 5)
	pressed := b.loadNineSlicedImage(assets.ImageUIButtonPressed, 5, 5)
	buttonPadding := widget.Insets{
		Left:   8,
		Right:  8,
		Top:    6,
		Bottom: 8,
	}
	b.button = &buttonDefaults{
		image: &widget.ButtonImage{
			Idle:     idle,
			Hover:    hover,
			Pressed:  pressed,
			Disabled: disabled,
		},
		padding: buttonPadding,
		textColors: &widget.ButtonTextColor{
			Idle:     styles.NormalTextColor.Color(),
			Disabled: styles.DisabledTextColor.Color(),
		},
	}
}

type TinyButtonConfig struct {
	Text       string
	OnClick    func()
	LayoutData any
}

func (b *Builder) NewTinyButton(config TinyButtonConfig) *widget.Button {
	return b.NewButton(ButtonConfig{
		Text:       config.Text,
		OnClick:    config.OnClick,
		LayoutData: config.LayoutData,

		MinWidth:  24,
		MinHeight: 24,
		Font:      assets.FontTiny,
	})
}

type ButtonConfig struct {
	Text       string
	OnClick    func()
	Tooltip    string
	MinWidth   int
	MinHeight  int
	Font       *text.Face
	LayoutData any
}

func (b *Builder) NewButton(config ButtonConfig) *widget.Button {
	ff := config.Font
	if ff == nil {
		ff = assets.Font1
	}

	defaults := b.button

	colors := b.button.textColors
	padding := &defaults.padding
	options := []widget.ButtonOpt{
		widget.ButtonOpts.Image(defaults.image),
	}

	if config.LayoutData != nil {
		options = append(options, widget.ButtonOpts.WidgetOpts(widget.WidgetOpts.LayoutData(config.LayoutData)))
	}

	if config.Text != "" {
		options = append(options,
			widget.ButtonOpts.Text(config.Text, ff, colors),
			widget.ButtonOpts.TextPadding(padding))
	}

	if strings.Contains(config.Text, "[color=") {
		options = append(options, widget.ButtonOpts.TextProcessBBCode(true))
	}

	options = append(options, widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
		if config.OnClick != nil {
			config.OnClick()
		}
	}))

	if config.MinWidth != 0 || config.MinHeight != 0 {
		options = append(options, widget.ButtonOpts.WidgetOpts(widget.WidgetOpts.MinSize(config.MinWidth, config.MinHeight)))
	}

	buttonWidget := widget.NewButton(options...)

	if config.Tooltip != "" {
		b.AddTooltip(buttonWidget.GetWidget(), func() AnyWidget {
			return b.NewText(TextConfig{
				Text: config.Tooltip,
				Font: assets.FontTiny,
			})
		})
	}

	return buttonWidget
}

type IconButtonConfig struct {
	Icon       *ebiten.Image
	OnClick    func()
	Tooltip    string
	MinWidth   int
	MinHeight  int
	Font       *text.GoXFace
	LayoutData any
}

type IconButton struct {
	Button *widget.Button
	Root   *widget.Container
}

func (b *Builder) NewIconButton(config IconButtonConfig) IconButton {
	button := b.NewButton(ButtonConfig{
		OnClick:    config.OnClick,
		Tooltip:    config.Tooltip,
		MinWidth:   config.MinWidth,
		MinHeight:  config.MinHeight,
		LayoutData: config.LayoutData,
	})

	combo := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewStackedLayout()))
	combo.GetWidget().LayoutData = config.LayoutData

	iconWidget := widget.NewGraphic(widget.GraphicOpts.Image(config.Icon))

	combo.AddChild(button)
	combo.AddChild(iconWidget)

	return IconButton{
		Button: button,
		Root:   combo,
	}
}

type SettingsButton struct {
	Container *widget.Container

	minValue    int
	valueLabels []string
	value       *int
	valueWidget *widget.Text
	wrapAround  bool

	EventChanged gsignal.Event[gsignal.Void]
}

type SettingsButtonConfig struct {
	Label        string
	ValueLabels  []string
	Value        *int
	MinValue     int
	ButtonLabels [2]string
	Tooltip      string
	MinWidth     int
	WrapAround   bool
}

func (b *Builder) NewSettingButton(config SettingsButtonConfig) *SettingsButton {
	result := &SettingsButton{
		valueLabels: config.ValueLabels,
		value:       config.Value,
		wrapAround:  config.WrapAround,
		minValue:    config.MinValue,
	}

	columns := b.NewGridCols(GridColsConfig{
		Cols:       3,
		ColSpacing: 4,
		MinWidth:   config.MinWidth,
		ColScale:   []bool{true, false, false},
	})

	labelPanel := b.NewPanel(PanelConfig{})
	innerCols := b.NewGridCols(GridColsConfig{
		Cols:       2,
		ColSpacing: 4,
		ColScale:   []bool{true, false},
	})
	labelPanel.AddChild(innerCols)
	innerCols.AddChild(b.NewText(TextConfig{
		Text:      config.Label,
		Font:      assets.FontTiny,
		AlignLeft: true,
		LayoutData: widget.AnchorLayoutData{
			VerticalPosition: widget.AnchorLayoutPositionCenter,
		},
	}))
	valueDisplay := b.NewText(TextConfig{
		Text:       "?",
		Font:       assets.FontTiny,
		Color:      styles.ImportantTextColor.Color(),
		AlignRight: true,
	})
	innerCols.AddChild(valueDisplay)
	result.valueWidget = valueDisplay

	minusLabel := config.ButtonLabels[0]
	if minusLabel == "" {
		minusLabel = "-"
	}
	minusButton := b.NewTinyButton(TinyButtonConfig{
		Text: minusLabel,
		OnClick: func() {
			if result.changeValue(-1) {
				result.updateValueLabel()
			}
		},
	})

	plusLabel := config.ButtonLabels[1]
	if plusLabel == "" {
		plusLabel = "+"
	}
	plusButton := b.NewTinyButton(TinyButtonConfig{
		Text: plusLabel,
		OnClick: func() {
			if result.changeValue(+1) {
				result.updateValueLabel()
			}
		},
	})

	columns.AddChild(labelPanel)
	columns.AddChild(minusButton)
	columns.AddChild(plusButton)

	result.Container = columns

	result.updateValueLabel()

	return result
}

func (b *SettingsButton) updateValueLabel() {
	b.valueWidget.Label = b.valueLabels[*b.value-b.minValue]
}

func (b *SettingsButton) Reload() {
	b.updateValueLabel()
}

func (b *SettingsButton) changeValue(delta int) bool {
	var newValue int
	if b.wrapAround {
		newValue = *b.value + delta
		switch {
		case newValue < b.minValue:
			newValue = len(b.valueLabels) - 1 + b.minValue
		case newValue >= len(b.valueLabels)+b.minValue:
			newValue = b.minValue
		}
	} else {
		newValue = gmath.Clamp(*b.value+delta, b.minValue, len(b.valueLabels)-1+b.minValue)
	}

	changed := newValue != *b.value
	*b.value = newValue

	if changed {
		b.EventChanged.Emit(gsignal.Void{})
	}
	return changed
}
