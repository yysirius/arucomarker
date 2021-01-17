package arucomarker

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
)

// MarkerInfo 标签信息
type MarkerInfo struct {
	SlideSize  int
	MarkerSize int
	BorderSize int
	Name       string
	Data       []uint8
}

func drawBox(img *image.RGBA, x, y, size int, color color.Color) {
	for dy := 0; dy < size; dy++ {
		for dx := 0; dx < size; dx++ {
			img.Set(x+dx, y+dy, color)
		}
	}
}

// DrawMarker 绘制标签
func DrawMarker(info MarkerInfo) error {
	img := image.NewRGBA(image.Rect(0, 0, info.SlideSize, info.SlideSize))
	draw.Draw(img, img.Bounds(), &image.Uniform{color.Black}, image.ZP, draw.Src)

	slideCount := info.MarkerSize + 2*info.BorderSize
	boxSize := info.SlideSize / slideCount
	if boxSize == 0 {
		return errors.New("Slide size is too limit")
	}
	fmt.Printf("Draw marker size %dX%d width %d border %d\n", info.MarkerSize, info.MarkerSize, info.SlideSize, info.BorderSize)
	byteIndex := 0
	currentByte := info.Data[byteIndex]
	bitIndex := 7
	for dy := info.BorderSize; dy < slideCount-info.BorderSize; dy++ {
		for dx := info.BorderSize; dx < slideCount-info.BorderSize; dx++ {
			//fmt.Println(currentByte, byteIndex, bitIndex)
			if (currentByte & (0x1 << bitIndex)) != 0 {
				drawBox(img, dx*boxSize, dy*boxSize, boxSize, color.White)
				//fmt.Println("set 1")
			}
			if bitIndex == 0 {
				byteIndex++
				if byteIndex >= len(info.Data) {
					break
				}
				currentByte = info.Data[byteIndex]
				bitIndex = 7
			} else {
				bitIndex--
			}
		}
	}

	f, _ := os.Create(fmt.Sprintf("%s.png", info.Name))
	defer f.Close()
	err := png.Encode(f, img)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
