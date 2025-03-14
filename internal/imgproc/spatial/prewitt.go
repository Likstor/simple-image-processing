package sp

import (
	"image"
	"math"
	"simple-image-processing/internal/imgproc"
	pp "simple-image-processing/internal/imgproc/point"
)

var (
	PrewittKernelH1 = [][]float64{
		{1, 0, -1},
		{1, 0, -1},
		{1, 0, -1},
	}

	PrewittKernelH2 = [][]float64{
		{-1, -1, -1},
		{0, 0, 0},
		{1, 1, 1},
	}
)

func Prewitt(processedImage *image.RGBA) {
	pp.GrayScale(processedImage)

	convPix := make([]uint8, len(processedImage.Pix))

	imgproc.MultithreadProcessCycle(processedImage.Rect, func(x, y int) {
		c1 := imgproc.ConvolutionOneChannel(processedImage, PrewittKernelH1, x, y, imgproc.RedChannel)
		c2 := imgproc.ConvolutionOneChannel(processedImage, PrewittKernelH2, x, y, imgproc.RedChannel)

		I := math.Sqrt(math.Pow(c1, 2) + math.Pow(c2, 2))

		imgproc.SetPixColorWithLimitFrom0To255(convPix, x, y, processedImage.Stride, I, I, I, 255)
	})

	processedImage.Pix = convPix
}
