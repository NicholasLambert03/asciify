/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"math"
	"os"

	"github.com/spf13/cobra"
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

// convertCmd represents the convert command
var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
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
	},
}

func init() {
	rootCmd.AddCommand(convertCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// convertCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// convertCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
