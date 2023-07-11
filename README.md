# Imago - stand-alone image manipulation

Imago is a simple command line to make some image manipulation.

At this time, Imago can:

- Drop shadow (with offset) on any image file. PNG alpha is managed (the shodow is applied to the alpha channel of the image, not the borders),
- resize image

Imago can get the image from file or standard input, and can save image in a new file or to standard output (so you can pipe imago commands)

# Usage

# Installation

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

