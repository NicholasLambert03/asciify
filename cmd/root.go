/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
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

func validateArguments(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		fmt.Println("Warning: No filepath provided, defaulted to showing help message.")
		cmd.Help()
		os.Exit(1)
	}
	if len(args) > 1 {
		return fmt.Errorf("too many arguments provided, expected 1 but got %d", len(args))
	}
	// Check if the file exists
	if _, err := os.Stat(args[0]); os.IsNotExist(err) {
		return fmt.Errorf("file does not exist: %s", args[0])
	}
	return nil

	//Filetype checked in the image.Decode function for
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "asciify [filepath]",
	Short: "A simple CLI tool to convert images to ASCII art",
	Long:  `Asciify is a command-line tool that converts .png .jpeg or .gif into ASCII art.`,
	Args:  validateArguments,
	Run: func(cmd *cobra.Command, args []string) {
		widthFlag, _ := cmd.Flags().GetInt("width") // Get the width from the command line flag
		//Variable Declarations
		var filepath string = args[0] // Get the file path from the command line argument
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
			fmt.Println("Error opening file:", err)
			os.Exit(1)

		}
		defer imageFile.Close() // defer runs when main returns

		// Decode PNG
		img, _, err = image.Decode(imageFile)
		if err != nil {
			fmt.Println("Error decoding image:", err)
			os.Exit(1)

		}
		if widthFlag > 0 { // If the width flag is set, resize the image
			img = resize(img, widthFlag)
		}

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
	rootCmd.Flags().IntP("width", "w", 0, "Custom width of the output ASCII art in characters")
}
