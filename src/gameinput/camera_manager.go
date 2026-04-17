package gameinput

import (
	"mygame/controls"
	"mygame/viewport"

	input "github.com/quasilyte/ebitengine-input"
	"github.com/quasilyte/gmath"
)

type CameraManager struct {
	camera *viewport.Camera
	input  *input.Handler

	scrollSpeed float64
	edgeScroll  bool
	arrowScroll bool

	cameraPanStartPos gmath.Vec
	cameraPanDragPos  gmath.Vec
}

type CameraManagerConfig struct {
	Camera      *viewport.Camera
	Input       *input.Handler
	ScrollSpeed int
	EdgeScroll  bool
	ArrowScroll bool
}

func NewCameraManager(config CameraManagerConfig) *CameraManager {
	return &CameraManager{
		camera:      config.Camera,
		input:       config.Input,
		edgeScroll:  config.EdgeScroll,
		arrowScroll: config.ArrowScroll,
		scrollSpeed: []float64{
			128,
			384,
			512,
			640,
			896,
		}[config.ScrollSpeed],
	}
}

func (m *CameraManager) HandleInput(delta float64) bool {
	h := m.input
	cameraPanSpeed := m.scrollSpeed * delta
	cameraPanBoundary := 8.0
	cam := m.camera

	changed := false

	var cameraPan gmath.Vec
	if m.arrowScroll {
		if h.ActionIsPressed(controls.ActionPanRight) {
			cameraPan.X += cameraPanSpeed
		}
		if h.ActionIsPressed(controls.ActionPanDown) {
			cameraPan.Y += cameraPanSpeed
		}
		if h.ActionIsPressed(controls.ActionPanLeft) {
			cameraPan.X -= cameraPanSpeed
		}
		if h.ActionIsPressed(controls.ActionPanUp) {
			cameraPan.Y -= cameraPanSpeed
		}
	}

	if cameraPan.IsZero() {
		if info, ok := h.JustPressedActionInfo(controls.ActionPanWheel); ok {
			m.cameraPanDragPos = cam.GetOffset()
			m.cameraPanStartPos = info.Pos
		} else if info, ok := h.PressedActionInfo(controls.ActionPanWheel); ok {
			posDelta := m.cameraPanStartPos.Sub(info.Pos)
			newPos := m.cameraPanDragPos.Add(posDelta)
			cam.SetOffset(newPos)
			changed = true
		}
	}

	if cameraPan.IsZero() && cameraPanBoundary != 0 && m.edgeScroll {
		// Mouse cursor can pan the camera too.
		cursor := h.CursorPos()
		windowSize := cam.GetViewportRect()
		if windowSize.Contains(cursor) {
			if cursor.X >= windowSize.Width()-cameraPanBoundary {
				cameraPan.X += cameraPanSpeed
			}
			if cursor.Y >= windowSize.Height()-cameraPanBoundary {
				cameraPan.Y += cameraPanSpeed
			}
			if cursor.X < cameraPanBoundary {
				cameraPan.X -= cameraPanSpeed
			}
			if cursor.Y < cameraPanBoundary {
				cameraPan.Y -= cameraPanSpeed
			}
		}
	}

	if !cameraPan.IsZero() {
		changed = true
		cam.Pan(cameraPan)
	}

	return changed
}
