package sp

import (
	"image"
	"simple-image-processing/internal/imgproc"
)

type SharpnessType int

const (
	SharpnessH1 SharpnessType = iota
	SharpnessH2
	SharpnessH3
)

var (
	SharpnessKernelH1 = [][]float64{
		{-1, -1, -1},
		{-1, 9, -1},
		{-1, -1, -1},
	}

	SharpnessKernelH2 = [][]float64{
		{0, -1, 0},
		{-1, 5, -1},
		{0, -1, 0},
	}

	SharpnessKernelH3 = [][]float64{
		{1, -2, 1},
		{-2, 5, -2},
		{1, -2, 1},
	}
)

func Sharpness(processedImage *image.RGBA, sharpnessKernelType SharpnessType) {
	var kernel imgproc.Matrix

	switch sharpnessKernelType {
	case SharpnessH1:
		kernel = SharpnessKernelH1
	case SharpnessH2:
		kernel = SharpnessKernelH2
	case SharpnessH3:
		kernel = SharpnessKernelH3
	}

	convPix := make([]uint8, len(processedImage.Pix))

	imgproc.MultithreadProcessCycle(processedImage.Rect, func(x, y int) {
		r, g, b := imgproc.Convolution(processedImage, kernel, x, y)

		imgproc.SetPixColorWithLimitFrom0To255(convPix, x, y, processedImage.Stride, r, g, b, 255)
	})

	processedImage.Pix = convPix
}