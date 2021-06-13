package loader

import (
	"github.com/veandco/go-sdl2/sdl"
	"go-sdl-games/utils"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

func GetRootDirPath() string {
	p, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return p
}

func FromBMP(fileName string) *sdl.Surface {
	bmp, err := sdl.LoadBMP(fileName)
	utils.HandleError(err)
	return bmp
}

func TextureFromBMP(fileName string, renderer *sdl.Renderer) *sdl.Texture {
	fullFileName := path.Join(GetRootDirPath(), fileName)
	bmp := FromBMP(fullFileName)
	defer bmp.Free()
	tex, err := renderer.CreateTextureFromSurface(bmp)
	utils.HandleError(err)
	return tex
}

func Textures(dirPath string, renderer *sdl.Renderer) []*sdl.Texture {
	fullDirPath := path.Join(GetRootDirPath(), dirPath)

	files, err := ioutil.ReadDir(fullDirPath)
	utils.HandleError(err)

	textures := make([]*sdl.Texture, 0, len(files))
	for _, file := range files {
		fullFileName := path.Join(dirPath, file.Name())
		texture := TextureFromBMP(fullFileName, renderer)
		textures = append(textures, texture)
	}

	return textures
}
