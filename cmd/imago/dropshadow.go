package main

import (
	"fmt"
	"image/png"

	"github.com/metal3d/imago/operations"
)

func dropShadow(input, output string, options *operations.DropShadowOptions) error {
	Log(fmt.Sprintf("DropShadow: %v -> %v", input, output))

	// open input file / stdin
	inStream, err := inputStream(&input)
	if err != nil {
		return fmt.Errorf("failed to open image: %w", err)
	}
	defer inStream.Close()

	// decode
	im, err := decodeImageFile(inStream)
	if err != nil {
		return fmt.Errorf("failed to decode image: %w", err)
	}

	// drop shadow
	newImage, err := operations.DropShadow(im, options)
	if err != nil {
		return fmt.Errorf("failed to drop shadow: %w", err)
	}

	// encode to output file / stdout
	outStream, err := outputStream(&output)
	if err != nil {
		return fmt.Errorf("failed to open image: %w", err)
	}
	defer outStream.Close()
	err = png.Encode(outStream, newImage)
	if err != nil {
		return fmt.Errorf("failed to encode image: %w", err)
	}
	return nil
}
