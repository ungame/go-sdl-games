package sprites

import (
	"github.com/veandco/go-sdl2/sdl"
	"go-sdl-games/loader"
	"go-sdl-games/topdown/common"
	"time"
)

const (
	ANIMATE_IDLE    = "idle"
	ANIMATE_DESTROY = "destroy"
)

type Animator struct {
	Container       *common.Element
	Sequences       map[string]*Sequence
	Current         string
	LastFrameChange time.Time
	Finished        bool
}

func NewAnimator(element *common.Element, sequences map[string]*Sequence, defaultSequence string) *Animator {
	var animator Animator
	animator.Container = element
	animator.Sequences = sequences
	animator.Current = defaultSequence
	animator.LastFrameChange = time.Now()
	return &animator
}

func (a *Animator) OnDraw(renderer *sdl.Renderer) error {
	texture := a.Sequences[a.Current].GetTexture()
	return common.DrawTexture(
		texture,
		a.Container.Position,
		a.Container.Rotation,
		renderer,
	)
}

func (a *Animator) OnUpdate() error {
	sequence := a.Sequences[a.Current]

	frameInterval := float64(time.Second) / sequence.SampleRate

	if time.Since(a.LastFrameChange) >= time.Duration(frameInterval) {
		a.Finished = !sequence.Next()
		a.LastFrameChange = time.Now()
	}

	return nil
}

func (a *Animator) OnCollision(_ *common.Element) error {
	return nil
}

func (a *Animator) ChangeSequence(name string) {
	a.Current = name
	a.LastFrameChange = time.Now()
}

type Sequence struct {
	Textures   []*sdl.Texture
	Frame      int
	SampleRate float64
	Loop       bool
}

func NewSequence(dirPath string, sampleRate float64, loop bool, renderer *sdl.Renderer) *Sequence {
	var sequence Sequence
	sequence.Textures = loader.Textures(dirPath, renderer)
	sequence.SampleRate = sampleRate
	sequence.Loop = loop
	return &sequence
}

func (s *Sequence) GetTexture() *sdl.Texture {
	return s.Textures[s.Frame]
}

func (s *Sequence) Next() bool {
	if s.IsLastFrame() {
		if s.Loop {
			s.Frame = 0
		} else {
			return false
		}
	} else {
		s.Frame++
	}
	return true
}

func (s *Sequence) IsLastFrame() bool {
	return s.Frame == len(s.Textures)-1
}
