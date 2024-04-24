package main

import (
	"image/color"
)



func NewRandomColor () color.Color {
	var r, g, b uint8
	r = uint8(randRangeInt(0, 255))
	g = uint8(randRangeInt(0, 255))
	b = uint8(randRangeInt(0, 255))

	return color.RGBA{r, g, b, 255}
}

