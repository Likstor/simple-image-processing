package pp

import (
	"image"
	"image/color"
	"simple-image-processing/internal/imgproc"
)

func GrayScale(processedImage *image.RGBA) {
	imgproc.MultithreadPointProcessCycle(processedImage.Rect, func(x, y int) {
		rgba := (processedImage.At(x, y)).(color.RGBA)

		I := uint8(0.3*float64(rgba.R) + 0.59*float64(rgba.G) + 0.11*float64(rgba.B))

		processedImage.Set(x, y, color.Gray{
			Y: I,
		})
	})
}
