package pp

import (
	"image"
	"simple-image-processing/internal/imgproc"
)

func Blue(processedImage *image.RGBA) {

	pix := make([]uint8, len(processedImage.Pix))

	imgproc.MultithreadProcessCycle(processedImage.Rect, func(x, y int) {
		pos := y * processedImage.Stride + x*4 + 2
		imgproc.SetPixColor(pix, x, y, processedImage.Stride, 0, 0, processedImage.Pix[pos], 255)
	})

	processedImage.Pix = pix
}
