package sprites

import (
	"github.com/veandco/go-sdl2/sdl"
	"go-sdl-games/topdown/common"
	"math"
	"time"
)

type Shooter struct {
	Container   *common.Element
	Cooldown    time.Duration
	LastShot    time.Time
	RightWeapon common.Vector
	LeftWeapon  common.Vector
}

func NewShooter(element *common.Element, cooldown time.Duration, rightWeapon, leftWeapon common.Vector) *Shooter {
	var shooter Shooter
	shooter.Container = element
	shooter.Cooldown = cooldown
	shooter.RightWeapon = rightWeapon
	shooter.LeftWeapon = leftWeapon
	return &shooter
}

func (s *Shooter) OnDraw(_ *sdl.Renderer) error {
	return nil
}

func (s *Shooter) OnCollision(_ *common.Element) error {
	return nil
}

func (s *Shooter) OnUpdate() error {
	keys := sdl.GetKeyboardState()

	if keys[sdl.SCANCODE_SPACE] == 1 {
		if s.CanShoot() {
			s.Shoot(s.Container.Position.X+s.RightWeapon.X, s.Container.Position.Y-s.RightWeapon.Y)
			s.Shoot(s.Container.Position.X-s.LeftWeapon.X, s.Container.Position.Y-s.LeftWeapon.Y)
			s.LastShot = time.Now()
		}
	}

	if keys[sdl.SCANCODE_R] == 1 {
		ResetBullets()
	}

	return nil
}

func (s *Shooter) CanShoot() bool {
	return time.Since(s.LastShot) >= s.Cooldown
}

func (s *Shooter) Shoot(x, y float64) {
	bullet, exists := NextBullet()
	if exists {
		bullet.Active = true
		bullet.Position.X = x
		bullet.Position.Y = y
		bullet.Rotation = 270 * (math.Pi / 180)
	}
}
