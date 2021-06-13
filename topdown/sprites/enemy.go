package sprites

import (
	"github.com/veandco/go-sdl2/sdl"
	"go-sdl-games/topdown/common"
	"go-sdl-games/topdown/global"
)

const (
	EnemyIdleTextureDir    = "\\topdown\\assets\\enemies\\idle"
	EnemyDestroyTextureDir = "\\topdown\\assets\\enemies\\destroy"
	EnemyIdleSampleRate    = 5
	EnemyDestroySampleRate = 15
	EnemyTag               = "enemy"
	EnemyRotation          = 180
	EnemyRadius            = 38
	DefaultEnemySize       = 105
)

func NewEnemy(renderer *sdl.Renderer, position common.Vector) *common.Element {
	var enemy common.Element
	enemy.Position = position
	enemy.Rotation = EnemyRotation
	enemy.Tag = EnemyTag

	sequences := make(map[string]*Sequence, 2)

	idle := NewSequence(EnemyIdleTextureDir, EnemyIdleSampleRate, true, renderer)
	sequences[ANIMATE_IDLE] = idle
	for index := range idle.Textures {
		enemy.AddDestructor(idle.Textures[index])
	}

	destroy := NewSequence(EnemyDestroyTextureDir, EnemyDestroySampleRate, false, renderer)
	sequences[ANIMATE_DESTROY] = destroy
	for index := range destroy.Textures {
		enemy.AddDestructor(destroy.Textures[index])
	}

	animator := NewAnimator(&enemy, sequences, ANIMATE_IDLE)
	enemy.AddComponent(animator)

	vulnerableToBullets := NewVulnerableToBullets(&enemy)
	enemy.AddComponent(vulnerableToBullets)

	enemy.Active = true

	collision := common.Circle{
		Center: enemy.Position,
		Radius: EnemyRadius,
	}
	enemy.AddCollision(collision)

	return &enemy
}

var Enemies []*common.Element

func InitEnemies(renderer *sdl.Renderer) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {

			x := (float64(i)/5)*global.MAX_WIDTH + float64(DefaultEnemySize/2)
			y := float64(j*DefaultEnemySize) + float64(DefaultEnemySize/2)

			enemy := NewEnemy(renderer, common.Vector{X: x, Y: y})
			Enemies = append(Enemies, enemy)
		}
	}
}
