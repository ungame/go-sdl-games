package sprites

import (
	"github.com/veandco/go-sdl2/sdl"
	"go-sdl-games/topdown/common"
	"go-sdl-games/topdown/global"
	"go-sdl-games/topdown/mover"
	"go-sdl-games/utils"
	"time"
)

const (
	PlayerTexture      = "\\topdown\\assets\\player\\player.bmp"
	PlayerTag          = "player"
	PlayerSpeed        = 10
	PlayerShotCooldown = time.Millisecond * 250
)

func NewPlayer(renderer *sdl.Renderer) *common.Element {
	var player common.Element
	player.Destructors = []utils.Destructor{}
	player.Position = common.Vector{}
	player.Active = true
	player.Tag = PlayerTag

	sprite := common.NewSprite(&player, renderer, PlayerTexture)
	player.AddComponent(sprite)
	player.AddDestructor(sprite.Tex)

	player.Position.X = global.MAX_WIDTH / 2
	player.Position.Y = global.MAX_HEIGHT - (sprite.Height / 2)

	playerMover := mover.NewPlayerMover(&player, PlayerSpeed)
	player.AddComponent(playerMover)

	rightWeapon := common.Vector{
		X: 25,
		Y: 20,
	}
	leftWeapon := common.Vector{
		X: 25,
		Y: 20,
	}

	shooter := NewShooter(&player, PlayerShotCooldown, rightWeapon, leftWeapon)
	player.AddComponent(shooter)

	return &player
}
