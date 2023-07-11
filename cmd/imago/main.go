package main

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	"strings"

	"github.com/metal3d/imago/operations"
	"github.com/spf13/cobra"
)

var (
	version = "dev" // set by Makefile
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "imago",
		Short: "A simple image processing tool",
		Long:  `A simple image processing tool built in Go`,
	}

	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version number of application",
		Long:  `Print the version number of application`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(version)
		},
	}

	rootCmd.AddCommand(
		versionCmd,
		dropShadowCommand(),
		resizeCommand(),
	)
	rootCmd.Execute()

}

func dropShadowCommand() *cobra.Command {

	var (
		sigma   *float64
		opacity *float64
		size    *int
		offset  *int
	)

	dropShadowCmd := &cobra.Command{
		Use:   "dropshadow [flags] input output",
		Short: "Add a drop shadow to an image, input and output image can be files or - for stdin/stdout",
		Long:  `Add a drop shadow to an image`,
		RunE: func(cmd *cobra.Command, args []string) error {

			if len(args) != 2 {
				return fmt.Errorf("dropshadow requires two arguments: input and output images")
			}

			return dropShadow(args[0], args[1], &operations.DropShadowOptions{
				Sigma:   *sigma,
				Opacity: *opacity,
				Size:    *size,
				Offset:  image.Pt(*offset, *offset),
			})
		},
	}
	sigma = dropShadowCmd.Flags().Float64P("sigma", "s", 25, "Blur sigma (standard deviation, strength of the blur)")
	opacity = dropShadowCmd.Flags().Float64P("opacity", "o", 0.8, "opacity of the shadow")
	size = dropShadowCmd.Flags().IntP("size", "z", 0, "size to extend the image (in pixels)")
	offset = dropShadowCmd.Flags().IntP("offset", "f", 0, "offset of the shadow (in pixels)")

	return dropShadowCmd
}

func resizeCommand() *cobra.Command {
	var (
		percent    *float64
		dimensions *string
	)

	resizeCmd := &cobra.Command{
		Use:   "resize [flags] input output",
		Short: "Resize an image",
		Long:  `Resize an image`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 2 {
				return fmt.Errorf("resize requires two arguments: input and output images")
			}
			var width, height int

			if *dimensions != "" {
				var err error
				width, height, err = parseDimensions(*dimensions)
				if err != nil {
					return err
				}
			}

			return resize(args[0], args[1], &operations.ResizeOptions{
				Percent: *percent,
				Width:   width,
				Height:  height,
			})
		},
	}

	percent = resizeCmd.Flags().Float64P(
		"percent", "p", 0,
		"percent to resize the image. If set, width and height are ignored",
	)
	dimensions = resizeCmd.Flags().StringP(
		"dimensions", "d", "0x0",
		`dimensions to resize the image (e.g. 100x200). If one value is set to 0, the value is calculated to respect the aspect-ratio.`,
	)
	return resizeCmd
}

func parseDimensions(dimensions string) (width, height int, err error) {
	dimensions = strings.TrimSpace(dimensions)
	dimensions = strings.ToLower(dimensions)
	if dimensions == "" {
		return 0, 0, nil
	}

	_, err = fmt.Sscanf(dimensions, "%dx%d", &width, &height)
	if err == nil {
		return
	}
	_, err = fmt.Sscanf(dimensions, "%dx", &width)
	if err == nil {
		return
	}

	_, err = fmt.Sscanf(dimensions, "x%d", &height)
	if err == nil {
		return
	}
	err = fmt.Errorf("invalid dimensions: %s", dimensions)
	return

}
