package scenes

import (
	"mygame/assets"
	"mygame/controls"
	"mygame/game"
	"mygame/gameinput"
	"mygame/viewport"

	graphics "github.com/quasilyte/ebitengine-graphics"
	"github.com/quasilyte/gmath"
	"github.com/quasilyte/gscene"
)

type TestController struct {
	back gscene.Controller
	cm   *gameinput.CameraManager
}

func NewTestController(back gscene.Controller) *TestController {
	return &TestController{}
}

func (c *TestController) Init(ctx gscene.InitContext) {
	worldSize := gmath.Vec{
		X: 2000,
		Y: 1500,
	}
	cameraRect := gmath.Rect{
		Max: game.G.WindowSize,
	}
	layers := []graphics.SceneLayerDrawer{
		0: graphics.NewLayer(),
		1: graphics.NewLayer(),
		2: graphics.NewLayer(),
		3: graphics.NewStaticLayer(), // UI layer
	}
	cam := viewport.NewCamera(viewport.CameraConfig{
		Scene:     ctx.Scene,
		Rect:      cameraRect,
		WorldSize: worldSize,
	})
	game.G.Camera = cam
	ctx.SetDrawer(viewport.NewDrawerWithLayers(cam, layers))

	{
		c.cm = gameinput.NewCameraManager(gameinput.CameraManagerConfig{
			Camera:      cam,
			Input:       game.G.Input,
			EdgeScroll:  true,
			ArrowScroll: true,
			ScrollSpeed: 2,
		})
	}

	{
		spr := game.G.NewSprite(assets.ImageUIPanel)
		spr.Pos.Offset = gmath.Vec{X: 96, Y: 96}
		ctx.Scene.AddGraphics(spr, 1)
	}
}

func (c *TestController) Update(delta float64) {
	c.cm.HandleInput(delta)

	if game.G.Input.ActionIsJustPressed(controls.ActionBack) {
		if c.back != nil {
			game.ChangeScene(c.back)
		}
	}
}
