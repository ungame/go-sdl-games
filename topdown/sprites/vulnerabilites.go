package sprites

import (
	"github.com/veandco/go-sdl2/sdl"
	"go-sdl-games/topdown/common"
)

type VulnerableToBullets struct {
	Container *common.Element
	Animator  *Animator
}

func NewVulnerableToBullets(element *common.Element) *VulnerableToBullets {
	var vtb VulnerableToBullets
	vtb.Container = element
	animator := element.GetBy(&Animator{})
	vtb.Animator, _ = animator.(*Animator)
	return &vtb
}

func (vtb *VulnerableToBullets) OnDraw(_ *sdl.Renderer) error {
	return nil
}

func (vtb *VulnerableToBullets) OnUpdate() error {
	switch vtb.Animator.Current {
	case ANIMATE_DESTROY:
		if vtb.Animator.Finished {
			vtb.Container.Active = false
		}
	}
	return nil
}

func (vtb *VulnerableToBullets) OnCollision(other *common.Element) error {
	if other.Tag == BulletTag {
		vtb.Animator.ChangeSequence(ANIMATE_DESTROY)
	}
	return nil
}
