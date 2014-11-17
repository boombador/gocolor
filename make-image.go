package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"
)

func main() {
	fmt.Println("beginning test")
	dx := 640
	dy := 480
	img := image.NewRGBA(image.Rect(0, 0, dx, dy))
	x := 0
	y := 0
	for y < dy {
		i := (x + y) % 255
		img.Set(x, y, color.RGBA{uint8(i), uint8(i), uint8(i), 255})
		x += 1
		if x >= dx {
			x = 0
			y += 1
		}
	}

	file, err := os.Create("output.jpeg")
	if err != nil {
		log.Fatal(err)
	}

	err = jpeg.Encode(file, img, &jpeg.Options{100})
	if err != nil {
		log.Fatal(err)
	}
}
