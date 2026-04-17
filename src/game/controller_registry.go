package game

import "github.com/quasilyte/gscene"

type ControllerRegistry struct {
	MainMenuController func() gscene.Controller
}
