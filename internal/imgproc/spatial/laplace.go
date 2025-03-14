package sp

import (
	"image"
	"math"
	"simple-image-processing/internal/imgproc"
	pp "simple-image-processing/internal/imgproc/point"
)

var (
	LaplaceKernel = [][]float64{
		{0, 1, 0},
		{1, -4, 1},
		{0, 1, 0},
	}
)

func Laplace(processedImage *image.RGBA) {
	convPix := make([]uint8, len(processedImage.Pix))

	pp.GrayScale(processedImage)

	imgproc.MultithreadProcessCycle(processedImage.Rect, func(x, y int) {
		c := imgproc.ConvolutionOneChannel(processedImage, LaplaceKernel, x, y, imgproc.RedChannel)

		c = math.Abs(c)

		imgproc.SetPixColorWithLimitFrom0To255(convPix, x, y, processedImage.Stride, c, c, c, 255)
	})

	processedImage.Pix = convPix
}