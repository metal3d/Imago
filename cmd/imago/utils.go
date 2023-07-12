package main

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

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

func inputStream(filename *string, stream *os.File) (*os.File, error) {
	if filename == nil || *filename == "-" {
		return stream, nil
	} else {
		inStream, err := os.Open(*filename)
		if err != nil {
			return nil, fmt.Errorf("failed to open image: %w", err)
		}
		return inStream, nil
	}
}

func outputStream(filename *string, stream *os.File) (*os.File, error) {
	if filename == nil || *filename == "-" {
		return stream, nil
	} else {
		inStream, err := os.Create(*filename)
		if err != nil {
			return nil, fmt.Errorf("failed to open image: %w", err)
		}
		return inStream, nil
	}
}
