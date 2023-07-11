package operations

import (
	"image"
	"image/color"

	"github.com/disintegration/imaging"
)

// DropShadowOptions is a struct that contains the options for the DropShadow function.
//
// The size is added to the width and height of the original image, and the shadow
// is placed at the offset.
// The sigma is the amount of blur to apply to the shadow.
// The opacity is the opacity of the shadow, 0 is transparent, 1 is opaque.
// The offset is the offset of the shadow from the original image. Only the first
// offset is used. If no offset is provided, the shadow is placed directly behind the original image.
type DropShadowOptions struct {
	Size    int
	Sigma   float64
	Opacity float64
	Offset  image.Point
}

// NewDropShadowOptions returns a new DropShadowOptions struct with default values.
func NewDropShadowOptions() *DropShadowOptions {
	return &DropShadowOptions{
		Size:    20,
		Sigma:   10,
		Opacity: 0.8,
		Offset:  image.Pt(0, 0),
	}
}

// DropShadow takes an image, and returns a new image with a drop shadow.
func DropShadow(img image.Image, options *DropShadowOptions) (image.Image, error) {
	if options == nil {
		options = NewDropShadowOptions()
	}
	bounds := img.Bounds()

	newSize := image.Pt(
		bounds.Dx()+options.Offset.X+options.Size*2,
		bounds.Dy()+options.Offset.Y+options.Size*2,
	)

	// create a alpha mask for the shadow
	shadow := imaging.New(newSize.X, newSize.Y, color.Alpha{0})

	for x := 0; x < newSize.X; x++ {
		for y := 0; y < newSize.Y; y++ {
			_, _, _, a := img.At(x, y).RGBA()
			shadow.Set(
				x+options.Size+options.Offset.X,
				y+options.Size+options.Offset.Y,
				color.RGBA{0, 0, 0, uint8(float64(a) * options.Opacity)},
			)
		}
	}
	shadow = imaging.Blur(shadow, options.Sigma)
	shadow = imaging.Overlay(shadow, img, image.Pt(options.Size, options.Size), 1)

	return shadow, nil

}
