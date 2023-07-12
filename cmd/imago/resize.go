package main

import (
	"fmt"
	"image/png"

	"github.com/metal3d/imago/operations"
)

func resize(input, output string, options *operations.ResizeOptions) error {
	Log(fmt.Sprintf("Resize: %v -> %v", input, output))

	// open input file / stdin
	inStream, err := inputStream(&input)
	if err != nil {
		return fmt.Errorf("failed to open image: %w", err)
	}
	defer inStream.Close()

	im, err := decodeImageFile(inStream)
	if err != nil {
		return fmt.Errorf("failed to decode image: %w", err)
	}

	// resize
	resized, err := operations.Resize(im, options)
	if err != nil {
		return fmt.Errorf("failed to resize image: %w", err)
	}

	// encode to output file / stdout
	outStream, err := outputStream(&output)
	if err != nil {
		return fmt.Errorf("failed to open image: %w", err)
	}
	defer outStream.Close()

	err = png.Encode(outStream, resized)
	if err != nil {
		return fmt.Errorf("failed to encode image: %w", err)
	}
	return nil
}
