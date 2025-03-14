package pp

import (
	"image"
	"image/color"
	"simple-image-processing/internal/imgproc"
)

func AdjustBrightness(processedImage *image.RGBA, param int) {
	imgproc.MultithreadProcessCycle(processedImage.Rect, func(x, y int) {
		rgba := processedImage.At(x, y).(color.RGBA)

		processedImage.Set(x, y, color.RGBA{
			R: imgproc.LimitFrom0To255(int(rgba.R) + param),
			G: imgproc.LimitFrom0To255(int(rgba.G) + param),
			B: imgproc.LimitFrom0To255(int(rgba.B) + param),
			A: rgba.A,
		})
	})
}
