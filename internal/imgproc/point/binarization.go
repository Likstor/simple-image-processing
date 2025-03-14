package pp

import (
	"image"
	"image/color"
	"simple-image-processing/internal/imgproc"
)

func Binarization(processedImage *image.RGBA, threshold uint8, color1, color2 color.Color) {
	imgproc.MultithreadProcessCycle(processedImage.Rect, func(x, y int) {
		rgba := processedImage.At(x, y).(color.RGBA)

		if uint8(0.3*float64(rgba.R)+0.59*float64(rgba.G)+0.11*float64(rgba.B)) >= threshold {
			processedImage.Set(x, y, color1)
		} else {
			processedImage.Set(x, y, color2)
		}
	})
}
