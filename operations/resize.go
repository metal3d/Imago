package operations

import (
	"image"

	"github.com/disintegration/imaging"
)

type ResizeOptions struct {
	Width   int
	Height  int
	Percent float64
}

func NewResizeOptions() *ResizeOptions {
	return &ResizeOptions{
		Width:   0,
		Height:  0,
		Percent: 0,
	}
}

func Resize(img image.Image, options *ResizeOptions) (image.Image, error) {
	if options == nil {
		options = NewResizeOptions()
	}
	bounds := img.Bounds()

	// if percent is set, calculate the width and height
	if options.Percent > 0 {
		options.Width = int(float64(bounds.Dx())*options.Percent) / 100
		options.Height = int(float64(bounds.Dy())*options.Percent) / 100
	}

	// if width or height is 0, calculate the missing value
	if options.Width == 0 {
		options.Width = int(float64(bounds.Dx()) * float64(options.Height) / float64(bounds.Dy()))
	}
	if options.Height == 0 {
		options.Height = int(float64(bounds.Dy()) * float64(options.Width) / float64(bounds.Dx()))
	}

	// resize the image
	newImage := imaging.Resize(img, options.Width, options.Height, imaging.Lanczos)
	return newImage, nil
}
