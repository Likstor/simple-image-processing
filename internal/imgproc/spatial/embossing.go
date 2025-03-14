package sp

import (
	"image"
	"simple-image-processing/internal/imgproc"
	pp "simple-image-processing/internal/imgproc/point"
)

type EmbossingType int

const (
	EmbossingH1 EmbossingType = iota
	EmbossingH2
)

var (
	EmbossingKernelH1 = [][]float64{
		{0, 1, 0},
		{-1, 0, 1},
		{0, -1, 0},
	}

	EmbossingKernelH2 = [][]float64{
		{0, -1, 0},
		{1, 0, -1},
		{0, 1, 0},
	}
)

func Embossing(processedImage *image.RGBA, embossingKernelType EmbossingType) {
	var kernel imgproc.Matrix

	switch embossingKernelType {
	case EmbossingH1:
		kernel = EmbossingKernelH1
	case EmbossingH2:
		kernel = EmbossingKernelH2
	}

	pp.GrayScale(processedImage)

	convPix := make([]uint8, len(processedImage.Pix))

	imgproc.MultithreadProcessCycle(processedImage.Rect, func(x, y int) {
		c := imgproc.ConvolutionOneChannel(processedImage, kernel, x, y, imgproc.RedChannel) + 128

		imgproc.SetPixColorWithLimitFrom0To255(convPix, x, y, processedImage.Stride, c, c, c, 255)
	})

	processedImage.Pix = convPix
}