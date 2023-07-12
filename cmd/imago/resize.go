package main

import (
	"fmt"
	"image/png"
	"os"

	"github.com/metal3d/imago/operations"
)

func resize(input, output string, options *operations.ResizeOptions) error {
	Log(fmt.Sprintf("Resize: %v -> %v", input, output))

	inStream, err := inputStream(&input, os.Stdin)
	if err != nil {
		return fmt.Errorf("failed to open image: %w", err)
	}
	defer inStream.Close()

	outStream, err := outputStream(&output, os.Stdout)
	if err != nil {
		return fmt.Errorf("failed to open image: %w", err)
	}
	defer outStream.Close()

	im, err := decodeImageFile(inStream)
	if err != nil {
		return fmt.Errorf("failed to decode image: %w", err)
	}

	resized, err := operations.Resize(im, options)
	if err != nil {
		return fmt.Errorf("failed to resize image: %w", err)
	}

	err = png.Encode(outStream, resized)
	if err != nil {
		return fmt.Errorf("failed to encode image: %w", err)
	}
	return nil
}
