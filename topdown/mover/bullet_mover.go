package mover

import (
	"github.com/veandco/go-sdl2/sdl"
	"go-sdl-games/topdown/common"
	"go-sdl-games/topdown/global"
	"math"
)

type BulletMover struct {
	Container *common.Element
	Speed     float64
}

func NewBulletMover(element *common.Element, speed float64) *BulletMover {
	var bulletMover BulletMover
	bulletMover.Container = element
	bulletMover.Speed = speed
	return &bulletMover
}

func (m *BulletMover) OnDraw(_ *sdl.Renderer) error {
	return nil
}

func (m *BulletMover) OnUpdate() error {
	container := m.Container
	container.Position.X += m.ComputeTrajectoryFromX()
	container.Position.Y += m.ComputeTrajectoryFromY()

	m.UpdateCollisions(container.Position)

	return nil
}

func (m *BulletMover) ComputeTrajectoryFromX() float64 {
	return m.Speed * math.Cos(m.Container.Rotation) * global.GetDeltaTime()
}

func (m *BulletMover) ComputeTrajectoryFromY() float64 {
	return m.Speed * math.Sin(m.Container.Rotation) * global.GetDeltaTime()
}

func (m *BulletMover) UpdateCollisions(position common.Vector) {
	for index := range m.Container.Collisions {
		m.Container.Collisions[index].Center = position
	}
}

func (m *BulletMover) OnCollision(_ *common.Element) error {
	m.Container.Active = false
	return nil
}
