package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
)

func main() {

	//Load image

	imageFile, err := os.Open("/home/nicholas/Pictures/hq_cat.png") // := is shorthand for regular variable declaration
	if err != nil {
		log.Fatal(err)
	}
	defer imageFile.Close() // defer runs when main returns

	// Decode PNG
	image, _, err := image.Decode(imageFile)
	if err != nil {
		log.Fatal(err)
	}

	levels := []string{"*", "%", "D", "J", "@"}

	for y := image.Bounds().Min.Y; y < image.Bounds().Max.Y; y++ {
		for x := image.Bounds().Min.X; x < image.Bounds().Max.X; x++ {
			c := color.GrayModel.Convert(image.At(x, y)).(color.Gray)
			level := c.Y / 51
			if level == 5 {
				level--
			}
			fmt.Print(levels[level])
		}
	}

}
