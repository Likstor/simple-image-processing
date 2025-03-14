package sp

import (
	"image"
	"math"
	"simple-image-processing/internal/imgproc"
	pp "simple-image-processing/internal/imgproc/point"
)

var (
	KirschKernelH1 = [][]float64{
		{-3, -3, -3},
		{-3, 0, -3},
		{5, 5, 5},
	}

	KirschKernelH2 = [][]float64{
		{-3, -3, -3},
		{5, 0, -3},
		{5, 5, -3},
	}

	KirschKernelH3 = [][]float64{
		{5, -3, -3},
		{5, 0, -3},
		{5, -3, -3},
	}

	KirschKernelH4 = [][]float64{
		{5, 5, -3},
		{5, 0, -3},
		{-3, -3, -3},
	}

	KirschKernelH5 = [][]float64{
		{5, 5, 5},
		{-3, 0, -3},
		{-3, -3, -3},
	}

	KirschKernelH6 = [][]float64{
		{-3, 5, 5},
		{-3, 0, 5},
		{-3, -3, -3},
	}

	KirschKernelH7 = [][]float64{
		{-3, -3, 5},
		{-3, 0, 5},
		{-3, -3, 5},
	}

	KirschKernelH8 = [][]float64{
		{-3, -3, -3},
		{-3, 0, 5},
		{-3, 5, 5},
	}
)

func Kirsch(processedImage *image.RGBA) {
	pp.GrayScale(processedImage)

	convPix := make([]uint8, len(processedImage.Pix))

	imgproc.MultithreadProcessCycle(processedImage.Rect, func(x, y int) {
		c1 := imgproc.ConvolutionOneChannel(processedImage, KirschKernelH1, x, y, imgproc.RedChannel)
		c2 := imgproc.ConvolutionOneChannel(processedImage, KirschKernelH2, x, y, imgproc.RedChannel)
		c3 := imgproc.ConvolutionOneChannel(processedImage, KirschKernelH3, x, y, imgproc.RedChannel)
		c4 := imgproc.ConvolutionOneChannel(processedImage, KirschKernelH4, x, y, imgproc.RedChannel)
		c5 := imgproc.ConvolutionOneChannel(processedImage, KirschKernelH5, x, y, imgproc.RedChannel)
		c6 := imgproc.ConvolutionOneChannel(processedImage, KirschKernelH6, x, y, imgproc.RedChannel)
		c7 := imgproc.ConvolutionOneChannel(processedImage, KirschKernelH7, x, y, imgproc.RedChannel)
		c8 := imgproc.ConvolutionOneChannel(processedImage, KirschKernelH8, x, y, imgproc.RedChannel)

		I := max(
			math.Abs(c1),
			math.Abs(c2),
			math.Abs(c3),
			math.Abs(c4),
			math.Abs(c5),
			math.Abs(c6),
			math.Abs(c7),
			math.Abs(c8),
		)

		imgproc.SetPixColorWithLimitFrom0To255(convPix, x, y, processedImage.Stride, I, I, I, 255)
	})

	processedImage.Pix = convPix
}
