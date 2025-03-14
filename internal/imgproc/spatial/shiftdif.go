package sp

import (
	"image"
	"simple-image-processing/internal/imgproc"
	pp "simple-image-processing/internal/imgproc/point"
)

type ShiftDifType int

const (
	ShiftDifH1 ShiftDifType = iota
	ShiftDifH2
	ShiftDifH3
)

var (
	ShiftDifKernelH1 = [][]float64{
		{0, 0, 0},
		{-1, 1, 0},
		{0, 0, 0},
	}

	ShiftDifKernelH2 = [][]float64{
		{0, -1, 0},
		{0, 1, 0},
		{0, 0, 0},
	}

	ShiftDifKernelH3 = [][]float64{
		{-1, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}
)

func ShiftDif(processedImage *image.RGBA, shiftDifKernelType ShiftDifType) {
	var kernel imgproc.Matrix

	switch shiftDifKernelType {
	case ShiftDifH1:
		kernel = ShiftDifKernelH1
	case ShiftDifH2:
		kernel = ShiftDifKernelH2
	case ShiftDifH3:
		kernel = ShiftDifKernelH3
	}

	pp.GrayScale(processedImage)

	convPix := make([]uint8, len(processedImage.Pix))

	imgproc.MultithreadProcessCycle(processedImage.Rect, func(x, y int) {
		c := imgproc.ConvolutionOneChannel(processedImage, kernel, x, y, imgproc.RedChannel)

		imgproc.SetPixColorWithLimitFrom0To255(convPix, x, y, processedImage.Stride, c, c, c, 255)
	})

	processedImage.Pix = convPix
}