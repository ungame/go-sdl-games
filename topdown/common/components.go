package common

import (
	"github.com/veandco/go-sdl2/sdl"
	"reflect"
)

type Component interface {
	OnUpdate() error
	OnDraw(renderer *sdl.Renderer) error
	OnCollision(other *Element) error
}

func isEqual(c1, c2 Component) bool {
	return reflect.TypeOf(c1) == reflect.TypeOf(c2)
}
