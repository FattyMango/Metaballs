package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	// "fyne.io/fyne/v2/container"
)

func main() {

	a := app.New()
	w := a.NewWindow("Metaballs")
	screen := NewScreen()

	screen.raster = canvas.NewRaster(screen.draw)
	screen.ExtendBaseWidget(screen)
	w.SetContent(screen)
	w.Resize(fyne.NewSize(512, 512))
	screen.Animate()
	w.ShowAndRun()
}
