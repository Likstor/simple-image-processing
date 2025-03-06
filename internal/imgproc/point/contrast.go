package pp

import (
	"image"
	"image/color"
	"simple-image-processing/internal/imgproc"
)

func IncreaseContrast(processedImage *image.RGBA, q1, q2 uint8) {
	coefficient := 255. / float64(q2-q1)

	imgproc.MultithreadPointProcessCycle(processedImage.Rect, func(x, y int) {
		rgba := processedImage.At(x, y).(color.RGBA)

		newR := imgproc.LimitFrom0To255(float64(int(rgba.R)-int(q1)) * coefficient)
		newG := imgproc.LimitFrom0To255(float64(int(rgba.G)-int(q1)) * coefficient)
		newB := imgproc.LimitFrom0To255(float64(int(rgba.B)-int(q1)) * coefficient)

		processedImage.Set(x, y, color.RGBA{
			R: newR,
			G: newG,
			B: newB,
			A: rgba.A,
		})
	})
}

func DecreaseContrast(processedImage *image.RGBA, q1, q2 uint8) {
	coefficient := float64(q2-q1) / 255.

	imgproc.MultithreadPointProcessCycle(processedImage.Rect, func(x, y int) {
		rgba := processedImage.At(x, y).(color.RGBA)

		newR := imgproc.LimitFrom0To255(float64(q1) + float64(rgba.R)*coefficient)
		newG := imgproc.LimitFrom0To255(float64(q1) + float64(rgba.G)*coefficient)
		newB := imgproc.LimitFrom0To255(float64(q1) + float64(rgba.B)*coefficient)

		processedImage.Set(x, y, color.RGBA{
			R: newR,
			G: newG,
			B: newB,
			A: rgba.A,
		})
	})
}
