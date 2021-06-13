package common

import (
	"github.com/veandco/go-sdl2/sdl"
	"go-sdl-games/utils"
)

type Element struct {
	Position    Vector
	Rotation    float64
	Active      bool
	Tag         string
	Collisions  []Circle
	Components  []Component
	Destructors []utils.Destructor
}

func (e *Element) Draw(renderer *sdl.Renderer) {
	for _, component := range e.Components {
		err := component.OnDraw(renderer)
		utils.HandleError(err)
	}
}

func (e *Element) Update() {
	for _, component := range e.Components {
		err := component.OnUpdate()
		utils.HandleError(err)
	}
}

func (e *Element) Collision(other *Element) error {
	for _, component := range e.Components {
		err := component.OnCollision(other)
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *Element) AddComponent(component Component) {
	if !e.HasComponent(component) {
		e.Components = append(e.Components, component)
	}
}

func (e *Element) HasComponent(component Component) bool {
	return e.GetBy(component) != nil
}

func (e *Element) GetBy(other Component) Component {
	for _, component := range e.Components {
		if isEqual(component, other) {
			return component
		}
	}
	return nil
}

func (e *Element) AddDestructor(dest utils.Destructor) {
	if dest != nil {
		e.Destructors = append(e.Destructors, dest)
	}
}

func (e *Element) AddCollision(collision Circle) {
	e.Collisions = append(e.Collisions, collision)
}

func (e *Element) Destroy() error {
	for _, d := range e.Destructors {
		err := d.Destroy()
		if err != nil {
			return err
		}
	}
	return nil
}
