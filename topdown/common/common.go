package common

import (
	"github.com/veandco/go-sdl2/sdl"
	"go-sdl-games/utils"
)

func DrawTexture(texture *sdl.Texture, position Vector, rotation float64, renderer *sdl.Renderer) error {
	width, height, err := GetDimensions(texture)
	utils.HandleError(err)

	x := position.X - float64(width)/2.0
	y := position.Y - float64(height)/2.0

	src := sdl.Rect{
		X: 0,
		Y: 0,
		W: width,
		H: height,
	}

	dst := sdl.Rect{
		X: int32(x),
		Y: int32(y),
		W: width,
		H: height,
	}

	point := sdl.Point{
		X: width / 2,
		Y: height / 2,
	}

	return renderer.CopyEx(
		texture,
		&src,
		&dst,
		rotation,
		&point,
		sdl.FLIP_NONE,
	)
}

func GetDimensions(texture *sdl.Texture) (width int32, height int32, err error) {
	_, _, width, height, err = texture.Query()
	return
}
