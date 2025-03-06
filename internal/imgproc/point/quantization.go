package pp

import (
	"image"
	"image/color"
	"math"
	"simple-image-processing/internal/imgproc"
)

func prepareQuants(k int) [256]uint8 {
	var quants [256]uint8

	quantSize := int(math.Ceil(256. / float64(k)))

	for i := range k {
		c := uint8(imgproc.LimitFrom0To255(quantSize - 1 + quantSize*i))

		start := quantSize * i
		end := min(quantSize+quantSize*i, 256)

		for j := start; j < end; j++ {
			quants[uint8(j)] = c
		}
	}
	return quants
}

func Quantization(processedImage *image.RGBA, k int) {
	quants := prepareQuants(k)

	imgproc.MultithreadPointProcessCycle(processedImage.Rect, func(x, y int) {
		rgba := processedImage.At(x, y).(color.RGBA)

		processedImage.Set(x, y, color.RGBA{
			R: quants[rgba.R],
			G: quants[rgba.G],
			B: quants[rgba.B],
			A: rgba.A,
		})

	})
}
