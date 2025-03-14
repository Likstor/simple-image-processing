package sp

import (
	"image"
	"image/color"
	"simple-image-processing/internal/imgproc"
	"slices"
)

func Median(processedImage *image.RGBA, filterSize int) {
	width, height := processedImage.Rect.Dx(), processedImage.Rect.Dy()

	convPix := make([]uint8, len(processedImage.Pix))

	dummy := make(imgproc.Matrix, filterSize)

	size := filterSize * filterSize
	
	m := size / 2

	imgproc.MultithreadProcessCycle(processedImage.Rect, func(x, y int) {
		rValues := make([]uint8, size)
		gValues := make([]uint8, size)
		bValues := make([]uint8, size)

		index := 0
		imgproc.KernelCycle(dummy, width, height, x, y, func(x, y, i, j int) {
			rgba := processedImage.At(x, y).(color.RGBA)
			
			rValues[index] = rgba.R
			gValues[index] = rgba.G
			bValues[index] = rgba.B
			index++
		})

		slices.Sort(rValues)

		slices.Sort(gValues)

		slices.Sort(bValues)


		imgproc.SetPixColor(convPix, x, y, processedImage.Stride, rValues[m], gValues[m], bValues[m], 255)
	})

	processedImage.Pix = convPix
}
