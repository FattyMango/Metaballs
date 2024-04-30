package main

import (
	"image"
	"image/color"
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"

	// "fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/StephaneBunel/bresenham"
)

// ISO Contour Value
const iso = 1

type Screen struct {
	widget.BaseWidget
	raster *canvas.Raster
	group  *Group

	resolution Resolution
}

func NewScreen(g *Group, r Resolution) *Screen {
	s := &Screen{
		group:      g,
		resolution: r,
	}
	s.raster = canvas.NewRaster(s.draw)
	return s
}

func (bw *Screen) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(bw.raster)
}

// Called when s.Refresh() is called
// Re-render the screen
func (s *Screen) draw(w, h int) image.Image {

	// fgcolor := theme.WarningColor()
	img := image.NewRGBA(image.Rect(0, 0, int(s.Size().Height), int(s.Size().Width)))

	size := float32(max(w, h))
	g := int(math.Ceil(float64(size) / float64(s.resolution)))
	//red
	color := color.RGBA{196, 77, 86,1}

	for row := 0; row < h; row += g {
		y := float32(row) / size
		y2 := float32(row+g) / size

		for col := 0; col < w; col += g {

			x := float32(col) / size
			x2 := float32(col+g) / size

			a := s.group.value(x, y)
			b := s.group.value(x2, y)
			c := s.group.value(x2, y2)
			d := s.group.value(x, y2)

			sq := Square{a, b, c, d}
			lines := MarchSquare(sq, col, row, g)
			for _, line := range lines {

				bresenham.DrawLine(img, line.x1, line.y1, line.x2, line.y2, color)
			}

		}

	}
	return img
}
