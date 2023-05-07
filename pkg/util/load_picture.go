package util

import (
	"github.com/faiface/pixel"
	"image"
	"os"
)

func LoadPicture(path string) pixel.Picture {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}
	if err = file.Close(); err != nil {
		panic(err)
	}
	return pixel.PictureDataFromImage(img)
}
