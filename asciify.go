package main

import (
	"fmt"
	"image/color"
	"image/png"
	"log"
	"os"
)

func main() {

	//Load image

	imageFile, err := os.Open("/home/nicholas/Pictures/sunlight.png") // := is shorthand for regular variable declaration
	if err != nil {
		log.Fatal(err)
	}
	defer imageFile.Close() // defer runs when main returns

	// Decode PNG
	image, err := png.Decode(imageFile)
	if err != nil {
		log.Fatal(err)
	}

	levels := []string{".", ",", "$", "Y", "+", "P", "*", "%", "D", "J", "@"}

	for y := image.Bounds().Min.Y; y < image.Bounds().Max.Y; y++ {
		for x := image.Bounds().Min.X; x < image.Bounds().Max.X; x++ {
			c := color.GrayModel.Convert(image.At(x, y)).(color.Gray)
			level := c.Y / 23
			if level == 11 {
				level--
			}
			fmt.Print(levels[level])
		}
		fmt.Println("\n")
	}

}
