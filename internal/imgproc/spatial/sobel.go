package sp

import (
	"image"
	"math"
	"simple-image-processing/internal/imgproc"
	pp "simple-image-processing/internal/imgproc/point"
)

var (
	SobelKernelH1 = [][]float64{
		{-1, 0, 1},
		{-2, 0, 2},
		{-1, 0, 1},
	}
	SobelKernelH2 = [][]float64{
		{1, 2, 1},
		{0, 0, 0},
		{-1, -2, -1},
	}
)

// ? Нужно ли нормализовать
func Sobel(processedImage *image.RGBA) {
	convPix := make([]uint8, len(processedImage.Pix))

	pp.GrayScale(processedImage)

	imgproc.MultithreadProcessCycle(processedImage.Rect, func(x, y int) {
		c1 := imgproc.ConvolutionOneChannel(processedImage, SobelKernelH1, x, y, imgproc.GreenChannel)
		c2 := imgproc.ConvolutionOneChannel(processedImage, SobelKernelH2, x, y, imgproc.GreenChannel)

		I := imgproc.LimitFrom0To255(math.Sqrt(math.Pow(c1, 2) + math.Pow(c2, 2)))

		imgproc.SetPixColor(convPix, x, y, processedImage.Stride, I, I, I, 255)
	})

	processedImage.Pix = convPix
}

func SobelWithNormalization(processedImage *image.RGBA) {
	convPix := make([]float64, len(processedImage.Pix))

	pp.GrayScale(processedImage)

	imgproc.MultithreadProcessCycle(processedImage.Rect, func(x, y int) {
		c1 := imgproc.ConvolutionOneChannel(processedImage, SobelKernelH1, x, y, imgproc.RedChannel)
		c2 := imgproc.ConvolutionOneChannel(processedImage, SobelKernelH2, x, y, imgproc.RedChannel)

		I := math.Sqrt(math.Pow(c1, 2) + math.Pow(c2, 2))

		offset := y*processedImage.Stride + x*4
		convPix[offset] = I
		convPix[offset+1] = I
		convPix[offset+2] = I
	},
	)

	normPix := imgproc.Normalization(convPix)

	processedImage.Pix = normPix
}
