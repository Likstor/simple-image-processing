package pp

import (
	"image"
	"image/color"
	"simple-image-processing/internal/imgproc"
)

func solarizationFx(x uint8) uint8 {
	return uint8((-4./255.*float64(x) + 4.) * float64(x))
}

func Solarization(processedImage *image.RGBA) {
	imgproc.MultithreadProcessCycle(processedImage.Rect, func(x, y int) {
		rgba := processedImage.At(x, y).(color.RGBA)

		processedImage.Set(x, y, color.RGBA{
			R: solarizationFx(rgba.R),
			G: solarizationFx(rgba.G),
			B: solarizationFx(rgba.B),
			A: rgba.A,
		})
	})
}
