package arucomarker

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
)

const (
	slideSize  = 150
	markSize   = 5
	borderSize = 1
)

func drawBox(img *image.RGBA, x, y, size int, color color.Color) {
	for dy := 0; dy < size; dy++ {
		for dx := 0; dx < size; dx++ {
			img.Set(x+dx, y+dy, color)
		}
	}
}

// DrawMarker 绘制标签
func DrawMarker() error {
	img := image.NewRGBA(image.Rect(0, 0, slideSize, slideSize))
	draw.Draw(img, img.Bounds(), &image.Uniform{color.Black}, image.ZP, draw.Src)

	slideCount := markSize + 2*borderSize
	boxSize := slideSize / slideCount
	fmt.Println(slideSize, markSize, borderSize, slideCount, boxSize)
	index := 0
	for dy := borderSize; dy < slideCount-borderSize; dy++ {
		for dx := borderSize; dx < slideCount-borderSize; dx++ {
			boxColor := color.Black
			if index%2 == 0 {
				boxColor = color.White
			}
			drawBox(img, dx*boxSize, dy*boxSize, boxSize, boxColor)
			index++
		}
	}

	f, _ := os.Create("out.png")
	defer f.Close()
	err := png.Encode(f, img)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
