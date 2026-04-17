package scenes

import (
	"mygame/assets"
	"mygame/controls"
	"mygame/game"
	"mygame/gui"
	"strings"

	"github.com/quasilyte/gscene"
)

// <<edit>> Fill credits with stuff you want to share.

type CreditsController struct {
	scene *gscene.Scene
}

func NewCreditsController() *CreditsController {
	return &CreditsController{}
}

func (c *CreditsController) Init(ctx gscene.InitContext) {
	c.scene = ctx.Scene
	c.initUI()
}

func (c *CreditsController) Update(delta float64) {
	if game.G.Input.ActionIsJustPressed(controls.ActionBack) {
		c.goBack()
	}
}

func (c *CreditsController) goBack() {
	game.ChangeScene(NewMainMenuController())
}

func (c *CreditsController) initUI() {
	topRows := game.G.UI.NewTopLevelGridRows()

	panel := game.G.UI.NewPanel(gui.PanelConfig{
		MinWidth:  240,
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

	creditLines := []string{
		"A game by username",
		"",
		"Made with Ebitengine",
	}

	panelRows.AddChild(game.G.UI.NewText(gui.TextConfig{
		Text: strings.Join(creditLines, "\n"),
	}))

	panelRows.AddChild(game.G.UI.NewText(gui.TextConfig{}))

	panelRows.AddChild(game.G.UI.NewButton(gui.ButtonConfig{
		Text: "Back",
		OnClick: func() {
			c.goBack()
		},
	}))

	game.G.UI.Build(c.scene, topRows)
}
