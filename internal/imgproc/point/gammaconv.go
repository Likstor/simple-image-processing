package pp

import (
	"image"
	"image/color"
	"math"
	"simple-image-processing/internal/imgproc"
)

func GammaConversion(processedImage *image.RGBA, gamma float64) {
	imgproc.MultithreadProcessCycle(processedImage.Rect, func(x, y int) {
		rgba := processedImage.At(x, y).(color.RGBA)

		processedImage.Set(x, y, color.RGBA{
			R: imgproc.LimitFrom0To255(255. * math.Pow(float64(rgba.R)/255., gamma)),
			G: imgproc.LimitFrom0To255(255. * math.Pow(float64(rgba.G)/255., gamma)),
			B: imgproc.LimitFrom0To255(255. * math.Pow(float64(rgba.B)/255., gamma)),
			A: rgba.A,
		})
	})
}
