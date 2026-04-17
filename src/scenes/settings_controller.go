package scenes

import (
	"mygame/assets"
	"mygame/controls"
	"mygame/game"
	"mygame/gui"

	"github.com/quasilyte/gscene"
	"github.com/quasilyte/gsignal"
)

type SettingsController struct {
	scene *gscene.Scene
	back  gscene.Controller
}

func NewSettingsController(back gscene.Controller) *SettingsController {
	return &SettingsController{back: back}
}

func (c *SettingsController) Init(ctx gscene.InitContext) {
	c.scene = ctx.Scene
	c.initUI()
}

func (c *SettingsController) Update(delta float64) {
	if game.G.Input.ActionIsJustPressed(controls.ActionBack) {
		c.goBack()
	}
}

func (c *SettingsController) goBack() {
	game.ChangeScene(c.back)
}

func (c *SettingsController) initUI() {
	topRows := game.G.UI.NewTopLevelGridRows()

	panel := game.G.UI.NewPanel(gui.PanelConfig{
		MinHeight: 100,
	})
	topRows.AddChild(panel)

	panelRows := game.G.UI.NewPanelGridRows(120)
	panel.AddChild(panelRows)

	panelRows.AddChild(game.G.UI.NewText(gui.TextConfig{
		Text: "Credits",
		Font: assets.Font2,
	}))

	panelRows.AddChild(game.G.UI.NewText(gui.TextConfig{}))

	volumeLabels := []string{
		"off",
		"very low",
		"low",
		"normal",
		"loud",
		"very loud",
	}

	{
		b := game.G.UI.NewSettingButton(gui.SettingsButtonConfig{
			Label:       "Sound volume",
			ValueLabels: volumeLabels,
			Value:       &game.G.SoundVolume,
			MinWidth:    240,
		})
		b.EventChanged.Connect(nil, func(gsignal.Void) {
			game.G.AdjustVolumeLevels()
			game.G.PlaySound(assets.AudioClick)
		})
		panelRows.AddChild(b.Container)
	}

	{
		b := game.G.UI.NewSettingButton(gui.SettingsButtonConfig{
			Label:       "Music volume",
			ValueLabels: volumeLabels,
			Value:       &game.G.MusicVolume,
			MinWidth:    240,
		})
		b.EventChanged.Connect(nil, func(gsignal.Void) {
			game.G.AdjustVolumeLevels()
			game.G.AdjustMusic()
		})
		panelRows.AddChild(b.Container)
	}

	panelRows.AddChild(game.G.UI.NewText(gui.TextConfig{}))

	panelRows.AddChild(game.G.UI.NewButton(gui.ButtonConfig{
		Text: "Back",
		OnClick: func() {
			c.goBack()
		},
	}))

	game.G.UI.Build(c.scene, topRows)
}
