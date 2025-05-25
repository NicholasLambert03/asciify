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

// resize function resizes an image to a specified width while maintaining aspect ratio
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

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "asciify [filepath]",
	Short: "A simple CLI tool to convert images to ASCII art",
	Long:  `Asciify is a command-line tool that converts images into ASCII art.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		//Variable Declarations
		var filepath string = args[0] // Get the file path from the command line argument
		if _, err := os.Stat(filepath); os.IsNotExist(err) {
			log.Fatalf("File does not exist: %s", filepath)
		}
		var imageFile *os.File
		var err error
		var img image.Image
		var cd_map = [69]string{" ", ".", "'", "`", "^", "\"", ":", ";", "I", "l",
			"!", "i", ">", "<", "~", "+", "_", "-", "?", "]",
			"[", "}", "{", "1", ")", "(", "|", "\\", "/", "t",
			"f", "j", "r", "x", "n", "u", "v", "c", "z", "X",
			"Y", "U", "J", "C", "L", "Q", "0", "O", "Z", "m",
			"w", "q", "p", "d", "b", "k", "h", "a", "o", "*",
			"#", "M", "W", "&", "8", "%", "B", "@", "$"}

		//Load image

		imageFile, err = os.Open(filepath) // := is shorthand for regular variable declaration
		if err != nil {
			log.Fatalln(err)
		}
		defer imageFile.Close() // defer runs when main returns

		// Decode PNG
		img, _, err = image.Decode(imageFile)
		if err != nil {
			log.Fatal(err)

		}

		//img = resize(img, 20)

		for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
			for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
				c := color.GrayModel.Convert(img.At(x, y)).(color.Gray) // Converts image to greyscale for intensity map and then gets grey color component
				level := int(c.Y) * (len(cd_map) - 1) / 255             // Maps the intensity to a level in the cd_map array
				if level > len(cd_map) {
					level--
				}
				fmt.Print(cd_map[level])
			}
			fmt.Print("\n")
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.asciify.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
