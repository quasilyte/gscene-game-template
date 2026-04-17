package controls

import (
	input "github.com/quasilyte/ebitengine-input"
)

// <<edit>> Adjust for your game controls and actions.
// The defaults here demonstrate how to do it.

const (
	ActionUnknown input.Action = iota

	ActionPanUp
	ActionPanDown
	ActionPanLeft
	ActionPanRight

	ActionClick
	ActionRightClick

	ActionConfirm
	ActionRestart

	ActionBack
)

func DefaultKeymap() input.Keymap {
	return input.Keymap{
		ActionPanUp:    {input.KeyUp, input.KeyGamepadUp},
		ActionPanDown:  {input.KeyDown, input.KeyGamepadDown},
		ActionPanLeft:  {input.KeyLeft, input.KeyGamepadLeft},
		ActionPanRight: {input.KeyRight, input.KeyGamepadRight},

		ActionClick:      {input.KeyMouseLeft, input.KeyGamepadA},
		ActionRightClick: {input.KeyMouseRight, input.KeyGamepadB},

		ActionConfirm: {input.KeyEnter, input.KeySpace},
		ActionRestart: {input.KeyR},

		ActionBack: {input.KeyEscape, input.KeyGamepadBack},
	}
}
