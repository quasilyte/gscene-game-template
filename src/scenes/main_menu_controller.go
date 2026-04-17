package scenes

import (
	"mygame/assets"
	"mygame/dat"
	"mygame/game"
	"mygame/gui"
	"os"
	"runtime"

	"github.com/quasilyte/gscene"
)

type MainMenuController struct {
	scene *gscene.Scene
}

func NewMainMenuController() *MainMenuController {
	return &MainMenuController{}
}

func (c *MainMenuController) Init(ctx gscene.InitContext) {
	c.scene = ctx.Scene
	c.initUI()

	game.G.AdjustMusic()
}

func (c *MainMenuController) Update(delta float64) {

}

func (c *MainMenuController) initUI() {
	topRows := game.G.UI.NewTopLevelGridRows()

	panel := game.G.UI.NewPanel(gui.PanelConfig{
		MinWidth:  240,
		MinHeight: 100,
	})
	topRows.AddChild(panel)

	panelRows := game.G.UI.NewPanelGridRows(120)
	panel.AddChild(panelRows)

	panelRows.AddChild(game.G.UI.NewText(gui.TextConfig{
		Text: dat.GameTitle,
		Font: assets.Font2,
	}))

	panelRows.AddChild(game.G.UI.NewText(gui.TextConfig{}))

	panelRows.AddChild(game.G.UI.NewButton(gui.ButtonConfig{
		Text: "Play",
		OnClick: func() {
		},
	}))

	{
		b := game.G.UI.NewButton(gui.ButtonConfig{
			Text: "Settings",
			OnClick: func() {
				back := NewMainMenuController()
				game.ChangeScene(NewSettingsController(back))
			},
		})
		panelRows.AddChild(b)
	}

	{
		b := game.G.UI.NewButton(gui.ButtonConfig{
			Text: "Credits",
			OnClick: func() {
				game.ChangeScene(NewCreditsController())
			},
		})
		panelRows.AddChild(b)
	}

	if runtime.GOARCH != "wasm" {
		panelRows.AddChild(game.G.UI.NewText(gui.TextConfig{}))

		panelRows.AddChild(game.G.UI.NewButton(gui.ButtonConfig{
			Text: "Exit",
			OnClick: func() {
				os.Exit(0)
			},
		}))
	}

	game.G.UI.Build(c.scene, topRows)
}
