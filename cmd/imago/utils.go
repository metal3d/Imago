package main

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

// decodeImageFile decodes an image from the given file handle.
func decodeImageFile(input *os.File) (image.Image, error) {
	im, _, err := image.Decode(input)
	if err != nil {
		return nil, fmt.Errorf("failed to decode image: %w", err)
	}
	return im, nil
}

func Log(msg string) {
	fmt.Fprintln(os.Stderr, msg)
}

// inputStream returns a file handle for the input file, or os.Stdin if the input file is nil or "-".
func inputStream(filename *string) (*os.File, error) {
	if filename == nil || *filename == "-" {
		return os.Stdin, nil
	} else {
		inStream, err := os.Open(*filename)
		if err != nil {
			return nil, fmt.Errorf("failed to open image: %w", err)
		}
		return inStream, nil
	}
}

// outputStream returns a file handle for the output file, or os.Stdout if the output file is nil or "-".
func outputStream(filename *string) (*os.File, error) {
	if filename == nil || *filename == "-" {
		return os.Stdout, nil
	} else {
		inStream, err := os.Create(*filename)
		if err != nil {
			return nil, fmt.Errorf("failed to open image: %w", err)
		}
		return inStream, nil
	}
}
