package topdown

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"go-sdl-games/topdown/common"
	"go-sdl-games/topdown/global"
	"go-sdl-games/topdown/sprites"
	"go-sdl-games/utils"
	"time"
)

var Elements []*common.Element

func Run() {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	utils.HandleError(err)
	defer sdl.Quit()

	window, err := sdl.CreateWindow(
		"Topdown Game",
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		global.MAX_WIDTH,
		global.MAX_HEIGHT,
		sdl.WINDOW_OPENGL,
	)
	utils.HandleError(err)
	defer utils.HandleDestroy(window)

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	utils.HandleError(err)
	defer utils.HandleDestroy(renderer)

	player := sprites.NewPlayer(renderer)
	onFinish.Append(player)

	sprites.InitBullets(renderer)
	sprites.InitEnemies(renderer)

	for index := range sprites.Bullets {
		onFinish.Append(sprites.Bullets[index])
	}

	for index := range sprites.Enemies {
		onFinish.Append(sprites.Enemies[index])
	}

	Elements = append(Elements, player)
	Elements = append(Elements, sprites.Bullets...)
	Elements = append(Elements, sprites.Enemies...)

	defer onFinish.Destroy()

GameLoop:
	for {

		frameStart := time.Now()

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				fmt.Println("Quit...")
				break GameLoop
			}
		}

		err = renderer.SetDrawColor(255, 255, 255, 1)
		utils.HandleError(err)

		err = renderer.Clear()
		utils.HandleError(err)

		for _, element := range Elements {
			if element.Active {
				element.Draw(renderer)
				element.Update()
			}
		}

		err = common.CheckCollisions(Elements)
		utils.HandleError(err)

		renderer.Present()

		_ = global.SetDeltaTime(frameStart)
		fmt.Printf("%v\r", global.GetDeltaTime())
	}
}

type OnFinish struct {
	destructors []utils.Destructor
}

var onFinish OnFinish

func (f *OnFinish) Append(destructor ...utils.Destructor) {
	if destructor != nil {
		f.destructors = append(f.destructors, destructor...)
	}
}

func (f *OnFinish) Destroy() {
	for index := range f.destructors {
		utils.HandleDestroy(f.destructors[index])
	}
}
