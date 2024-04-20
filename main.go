package main

import (
	"image"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"

	// "fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/StephaneBunel/bresenham"
)
type BallWidget struct {
	widget.BaseWidget
	raster *canvas.Raster

}
func (bw *BallWidget) draw(w, h int) image.Image {
	fgcolor := theme.ForegroundColor()
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	bresenham.DrawLine(img, 10, 10, 20, 20, fgcolor)
	return img

}
func (bw *BallWidget) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(bw.raster)
}
func main() {

	a := app.New()
	w := a.NewWindow("Metaballs")
	ball := &BallWidget{}

	ball.raster = canvas.NewRaster(ball.draw)
	ball.ExtendBaseWidget(ball)
	w.SetContent(ball)
	w.Resize(fyne.NewSize(512, 512))

	w.ShowAndRun()
}