package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	"log"
	"os"
	dropshadow "shagow/shadow"
)

var (
	version = "dev"
)

func main() {

	var (
		sigma       = flag.Float64("sigma", 25, "Blur sigma (standard deviation, strength of the blur)")
		opacity     = flag.Float64("opacity", 0.8, "opacity of the shadow")
		size        = flag.Int("size", 0, "size to extend the image (in pixels)")
		offset      = flag.Int("offset", 0, "offset of the shadow (in pixels)")
		versionFlag = flag.Bool("version", false, "Print version information and quit")
	)
	flag.Usage = func() {
		fmt.Printf("Usage: %s [options] <input> <output>", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	if *versionFlag {
		fmt.Println("Version:", version)
		os.Exit(0)
	}

	if flag.NArg() != 2 {
		flag.Usage()
		os.Exit(1)
	}

	imageFile := flag.Arg(0)
	destFile := flag.Arg(1)
	log.Printf("Drop shadow: %s -> %s", imageFile, destFile)

	// decode
	orig, err := os.Open(imageFile)
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}
	defer orig.Close()
	im, _, err := image.Decode(orig)
	if err != nil {
		log.Fatalf("failed to decode image: %v", err)
	}

	// drop shadow
	options := dropshadow.NewDropShadowOptions()
	options.Size = *size
	options.Sigma = *sigma
	options.Opacity = *opacity
	options.Offset = image.Pt(*offset, *offset)
	newImage := dropshadow.DropShadow(im, options)

	// encode
	dest, err := os.Create(destFile)
	if err != nil {
		log.Fatalf("failed to create image: %v", err)
	}
	defer dest.Close()
	err = png.Encode(dest, newImage)
	if err != nil {
		log.Fatalf("failed to encode image: %v", err)
	}

}
