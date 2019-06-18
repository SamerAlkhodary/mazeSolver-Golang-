package maze

import (
	"algorithms/greedy/mazeSolver/model"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func init() {
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
}

func ParseMaze(path string) (image.Image, model.Dimensions) {
	imageFile, err := os.Open(path)
	if err != nil {
		log.Printf("Error happend while reading file : %v", err)
		os.Exit(0)
	}
	img, dim, err := getPixelArray(imageFile)
	if err != nil {
		log.Printf("Error happend while getting pixelArray : %v", err)
	}
	return img, dim

}
func getPixelArray(imageFIle *os.File) (image.Image, model.Dimensions, error) {
	imgConfig, _, er := image.DecodeConfig(imageFIle)

	if er != nil {
		log.Printf("ErGetPowerArrayror happend while getting Configs : %v", er)

	}
	dim := model.Dimensions{Height: imgConfig.Height, Width: imgConfig.Width}
	imageFIle.Seek(0, 0)
	img, _, err := image.Decode(imageFIle)
	return img, dim, err
}
func CreateImageResult(img image.Image, path string, item *model.Item) {
	image := printImage(img, item)
	fg, err := os.Create(path)
	defer fg.Close()
	if err != nil {
		panic(err)
	}
	err = png.Encode(fg, image)
	if err != nil {
		panic(err)
	}

}
func printImage(img image.Image, item *model.Item) image.Image {
	bounds := img.Bounds()
	target := image.NewNRGBA(image.Rect(0, 0, bounds.Dx()-1, bounds.Dy()))

	for i := bounds.Dy(); i >= 0; i-- {

		for j := bounds.Dx(); j >= 0; j-- {

			target.Set(j, i, img.At(j, i))

		}
	}
	lineDrawer(target, item)
	return target
}
func lineDrawer(image *image.NRGBA, item *model.Item) {
	current := item
	previous := item.GetPreviousItem()
	for previous != nil {
		x1, y1 := current.GetNode().GetCoord()
		x2, y2 := previous.GetNode().GetCoord()
		if x1 == x2 {

			delta, factor := delta(y1, y2)

			for l := 0; l <= (delta); l++ {
				image.Set(y2+(factor*l), (x1 - cellSize), color.RGBA{R: 255, G: 0, B: 0, A: 255})
			}

		} else if y1 == y2 {
			delta, factor := delta(x1, x2)
			for l := 0; l < (delta); l++ {
				image.Set(y1, x2-cellSize+(l*factor), color.RGBA{R: 255, G: 0, B: 0, A: 255})
			}

		}
		current = previous
		previous = previous.GetPreviousItem()

	}

}
func delta(a, b int) (int, int) {
	if a < b {
		return b - a, -1
	}
	return a - b, 1

}
