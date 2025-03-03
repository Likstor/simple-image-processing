package main

import (
	"image"
	"image/color"
	"image/draw"
	"simple-image-processing/internal/imgproc"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func CreateHistogramWindow() fyne.Window {
	lineColor := color.RGBA{
			R: 200,
			G: 200,
			B: 200,
			A: 255,
		}

	histogramWindow := fyne.CurrentApp().NewWindow("Color histogram tool")
	histogramWindow.SetIcon(theme.Icon(theme.IconNameColorPalette))

	hist := image.NewRGBA(imgproc.HISTOGRAM_RECTANGLE)

	backgroundColor := theme.Color(theme.ColorNameBackground)
	
	draw.Draw(hist, hist.Bounds(), &image.Uniform{backgroundColor}, image.Point{}, draw.Src)

	imgproc.DrawRectangle(hist, 0, hist.Rect.Dy() - 11, hist.Rect.Dx() - 1, hist.Rect.Dy() - 1, lineColor)

	canvasTempImage := canvas.NewImageFromImage(hist)
	canvasTempImage.FillMode = canvas.ImageFillStretch

	updateBtn := widget.NewButton("Update", func() {
		if CurrentImage == nil {
			dialog.ShowError(ErrImageNotSelected, histogramWindow)
			return
		}

		imgproc.UpdateColorHistogram(CurrentImage.BaseImage, hist, (color.RGBA)((backgroundColor).(color.NRGBA)), lineColor)

		canvasTempImage.Refresh()
	})

	content := container.NewBorder(
		nil, 
		updateBtn,
		nil,
		nil,
		canvasTempImage,
	)

	histogramWindow.SetContent(content)

	histogramWindow.Resize(fyne.NewSize(400, 200))

	return histogramWindow
}