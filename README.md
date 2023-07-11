# Imago - stand-alone image manipulation

Imago is a simple command line to make some image manipulation.

At this time, Imago can:

- Drop shadow (with offset) on any image file. PNG alpha is managed (the shodow is applied to the alpha channel of the image, not the borders),
- resize image

Imago can get the image from file or standard input, and can save image in a new file or to standard output (so you can pipe imago commands)

# Usage

You should refer to the `help` commands for each subcommand.

## Drop shadow

```bash

Add a drop shadow to an image

Usage:
  imago dropshadow [flags] input output

Flags:
  -h, --help            help for dropshadow
  -f, --offset int      offset of the shadow (in pixels)
  -o, --opacity float   opacity of the shadow (default 0.8)
  -s, --sigma float     Blur sigma (standard deviation, strength of the blur) (default 25)
  -z, --size int        size to extend the image (in pixels)
```

## Resize

```bash
Resize an image

Usage:
  imago resize [flags] input output

Flags:
  -d, --dimensions string   dimensions to resize the image (e.g. 100x200). If one value is set to 0, the value is calculated to respect the aspect-ratio. (default "0x0")
  -h, --help                help for resize
  -p, --percent float       percent to resize the image. If set, width and height are ignored
```

# Installation

## Using "Go"

You can install the binary for your platform using `go install`:

```bash
go install -u github.com/metal3d/imago@latest
```

## Get the release

You can also download one of the binaries in the [release page](https://github.com/metal3d/Imago/releases). Place the binary in you `$PATH` and rename it to `imago`.

# Why not ImageMagik (convert command)

Of course, for any complex and very optimized operations, you **should** use ImageMagik.

Imago is not intended to replace or to be a concurrent of ImageMagik. It is designed to be:

- stand-alone (it's a static binary)
- only for very simple operations
- for tiny resources systems
- very easy arguments

For example, it is very easy to use with LaTeX `\ShellEscape` as there is no backslash to use. While ImageMagik can do very complete opertion to create a drop shadow, and "repage" the image, Imago has only one method to do it.

Imago will never do this (unless the operation is simple, and our opinion changes):

- propose an argument to script operations
- make image convertion, especially to SVG
- complex effects, animations

Imago's aim is to keep things simple and offer basic, straightforward operations efficiently.

