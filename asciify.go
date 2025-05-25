package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"math"
	"os"

	"golang.org/x/image/draw"
)

func resize(img image.Image, width int) image.Image {
	//Variable Declarations
	var aspectRatio float64
	var height int
	var resized_image draw.Image

	//Calculate new image height based on width
	aspectRatio = (float64)(img.Bounds().Max.Y) / (float64)(img.Bounds().Max.X)
	height = int(math.Round(float64(width) * aspectRatio))

	//Resize image

	resized_image = image.NewRGBA(image.Rect(0, 0, width, height))
	draw.NearestNeighbor.Scale(resized_image, resized_image.Bounds(), img, img.Bounds(), draw.Over, nil)

	return resized_image

}

func main() {

	//Variable Declarations
	var imageFile *os.File
	var err error
	var img image.Image
	var cd_map = [5]string{"*", "%", "D", "J", "@"}

	//Load image

	imageFile, err = os.Open("./testimages/hq_cat.png") // := is shorthand for regular variable declaration
	if err != nil {
		log.Fatal(err)
	}
	defer imageFile.Close() // defer runs when main returns

	// Decode PNG
	img, _, err = image.Decode(imageFile)
	if err != nil {
		log.Fatal(err)

	}

	img = resize(img, 20)

	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			c := color.GrayModel.Convert(img.At(x, y)).(color.Gray) // Converts image to greyscale for intensity map and then gets grey color component
			level := c.Y / 51
			if level == 5 {
				level--
			}
			fmt.Print(cd_map[level])
		}
		fmt.Print("\n")
	}

}
