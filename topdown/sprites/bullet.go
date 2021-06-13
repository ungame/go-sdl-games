package sprites

import (
	"github.com/veandco/go-sdl2/sdl"
	"go-sdl-games/topdown/common"
	"go-sdl-games/topdown/mover"
)

const (
	BulletTexture = "\\topdown\\assets\\player\\player_bullet.bmp"
	BulletTag     = "bullet"
	BulletSpeed   = 15
	BulletRadius  = 8
	MaxBullets    = 30
)

func NewBullet(renderer *sdl.Renderer) *common.Element {
	var bullet common.Element
	bullet.Tag = BulletTag
	bullet.Position = common.Vector{}

	sprite := common.NewSprite(&bullet, renderer, BulletTexture)
	bullet.AddComponent(sprite)
	bullet.AddDestructor(sprite.Tex)

	bulletMover := mover.NewBulletMover(&bullet, BulletSpeed)
	bullet.AddComponent(bulletMover)

	collision := common.Circle{
		Center: bullet.Position,
		Radius: BulletRadius,
	}

	bullet.AddCollision(collision)

	return &bullet
}

var Bullets []*common.Element

func InitBullets(renderer *sdl.Renderer) {
	for i := 0; i < MaxBullets; i++ {
		Bullets = append(Bullets, NewBullet(renderer))
	}
}

func NextBullet() (*common.Element, bool) {
	for _, bullet := range Bullets {
		if !bullet.Active {
			return bullet, true
		}
	}
	return nil, false
}

func ResetBullets() {
	for index := range Bullets {
		Bullets[index].Active = false
	}
}
