package sp

import (
	"image"
	"math"
	"simple-image-processing/internal/imgproc"
	pp "simple-image-processing/internal/imgproc/point"
)

var (
	RobertsKernelH1 = [][]float64{
		{1, 0},
		{0, -1},
	}
	RobertsKernelH2 = [][]float64{
		{0, 1},
		{-1, 0},
	}
)

func Roberts(processedImage *image.RGBA) {
	convPix := make([]uint8, len(processedImage.Pix))

	pp.GrayScale(processedImage)

	imgproc.MultithreadProcessCycle(processedImage.Rect, func(x, y int) {
		c1 := imgproc.ConvolutionOneChannel(processedImage, RobertsKernelH1, x, y, imgproc.RedChannel)
		c2 := imgproc.ConvolutionOneChannel(processedImage, RobertsKernelH2, x, y, imgproc.RedChannel)

		I := imgproc.LimitFrom0To255(math.Sqrt(math.Pow(c1, 2) + math.Pow(c2, 2)))

		imgproc.SetPixColor(convPix, x, y, processedImage.Stride, I, I, I, 255)
	})

	processedImage.Pix = convPix
}

