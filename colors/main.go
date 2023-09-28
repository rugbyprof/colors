package main

import (
	"fmt"
	"image"
	_ "image/png" // import this package to decode PNGs
	"os"
)

type PixelInfo struct {
	X, Y    int
	R, G, B uint8
}

func getPixelInfo(filePath string) ([]PixelInfo, error) {
	var pixelArray []PixelInfo

	reader, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("Error: %v", err)
	}
	defer reader.Close()

	img, _, err := image.Decode(reader)
	if err != nil {
		return nil, fmt.Errorf("Error: %v", err)
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			color := img.At(x, y)
			r, g, b, _ := color.RGBA()
			pixel := PixelInfo{
				X: x,
				Y: y,
				R: uint8(r >> 8),
				G: uint8(g >> 8),
				B: uint8(b >> 8),
			}
			pixelArray = append(pixelArray, pixel)
		}
	}
	return pixelArray, nil
}

func main() {
	pixels, err := getPixelInfo("colors.png")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, pixel := range pixels {
		fmt.Printf("Pixel at (%d, %d) - R: %d, G: %d, B: %d\n", pixel.X, pixel.Y, pixel.R, pixel.G, pixel.B)
	}
}
