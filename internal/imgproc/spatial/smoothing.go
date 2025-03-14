package sp

import (
	"image"
	"simple-image-processing/internal/imgproc"
)

type SmoothingType int

const (
	SmoothingH1 SmoothingType = iota
	SmoothingH2
	SmoothingH3
)

var (
	SmoothingKernelH1 = imgproc.KernelNormalization([][]float64{
		{1, 1, 1},
		{1, 1, 1},
		{1, 1, 1},
	})

	SmoothingKernelH2 = imgproc.KernelNormalization([][]float64{
		{1, 1, 1},
		{1, 2, 1},
		{1, 1, 1},
	})

	SmoothingKernelH3 = imgproc.KernelNormalization([][]float64{
		{1, 2, 1},
		{2, 4, 2},
		{1, 2, 1},
	})
)

func Smoothing(processedImage *image.RGBA, smoothingKernelType SmoothingType) {
	var kernel imgproc.Matrix

	switch smoothingKernelType {
	case SmoothingH1:
		kernel = SmoothingKernelH1
	case SmoothingH2:
		kernel = SmoothingKernelH2
	case SmoothingH3:
		kernel = SmoothingKernelH3
	}

	convPix := make([]uint8, len(processedImage.Pix))

	imgproc.MultithreadProcessCycle(processedImage.Rect, func(x, y int) {
		r, g, b := imgproc.Convolution(processedImage, kernel, x, y)

		imgproc.SetPixColorWithLimitFrom0To255(convPix, x, y, processedImage.Stride, r, g, b, 255)
	})

	processedImage.Pix = convPix
}