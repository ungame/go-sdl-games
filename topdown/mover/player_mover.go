package mover

import (
	"github.com/veandco/go-sdl2/sdl"
	"go-sdl-games/topdown/common"
	"go-sdl-games/topdown/global"
)

type PlayerMover struct {
	Container *common.Element
	Speed     float64
	Sprite    *common.Sprite
}

func NewPlayerMover(element *common.Element, speed float64) *PlayerMover {
	var playerMover PlayerMover
	playerMover.Container = element
	playerMover.Speed = speed
	sprite := element.GetBy(&common.Sprite{})
	playerMover.Sprite, _ = sprite.(*common.Sprite)
	return &playerMover
}

func (m *PlayerMover) OnDraw(_ *sdl.Renderer) error {
	return nil
}

func (m *PlayerMover) OnUpdate() error {
	keys := sdl.GetKeyboardState()
	container := m.Container

	if keys[sdl.SCANCODE_LEFT] == 1 {
		if container.Position.X-(m.Sprite.Width/2) > 0 {
			container.Position.X -= m.Speed * global.GetDeltaTime()
		}
	} else if keys[sdl.SCANCODE_RIGHT] == 1 {
		if container.Position.X+(m.Sprite.Width/2) < global.MAX_WIDTH {
			container.Position.X += m.Speed * global.GetDeltaTime()
		}
	} else if keys[sdl.SCANCODE_UP] == 1 {
		if container.Position.Y-(m.Sprite.Height/2) > 0 {
			container.Position.Y -= m.Speed * global.GetDeltaTime()
		}
	} else if keys[sdl.SCANCODE_DOWN] == 1 {
		if container.Position.Y+(m.Sprite.Height/2) < global.MAX_HEIGHT {
			container.Position.Y += m.Speed * global.GetDeltaTime()
		}
	}

	return nil
}

func (m *PlayerMover) OnCollision(_ *common.Element) error {
	return nil
}
