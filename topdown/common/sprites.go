package common

import (
	"github.com/veandco/go-sdl2/sdl"
	"go-sdl-games/loader"
	"go-sdl-games/utils"
)

type Sprite struct {
	Container     *Element
	Tex           *sdl.Texture
	Width, Height float64
}

func NewSprite(container *Element, renderer *sdl.Renderer, fileName string) *Sprite {
	texture := loader.TextureFromBMP(fileName, renderer)
	_, _, width, height, err := texture.Query()
	utils.HandleError(err)
	var sprite Sprite
	sprite.Container = container
	sprite.Tex = texture
	sprite.Width = float64(width)
	sprite.Height = float64(height)
	return &sprite
}

func (s *Sprite) OnDraw(renderer *sdl.Renderer) error {
	return DrawTexture(s.Tex, s.Container.Position, s.Container.Rotation, renderer)
}

func (s *Sprite) OnUpdate() error {
	return nil
}

func (s *Sprite) OnCollision(_ *Element) error {
	return nil
}
