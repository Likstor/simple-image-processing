package pp

import (
	"image"
	"image/color"
	"simple-image-processing/internal/imgproc"
)

func Sepia(processedImage *image.RGBA) {
	imgproc.MultithreadProcessCycle(processedImage.Rect, func(x, y int) {
		rgba := (processedImage.At(x, y)).(color.RGBA)

		newR := imgproc.LimitFrom0To255((float64(rgba.R) * .393) + (float64(rgba.G) * .769) + (float64(rgba.B) * .189))
		newG := imgproc.LimitFrom0To255((float64(rgba.R) * .349) + (float64(rgba.G) * .686) + (float64(rgba.B) * .168))
		newB := imgproc.LimitFrom0To255((float64(rgba.R) * .272) + (float64(rgba.G) * .534) + (float64(rgba.B) * .131))
		
		processedImage.Set(x, y, color.RGBA{
			R: newR,
			G: newG,
			B: newB,
			A: rgba.A,
		})
	})
}
