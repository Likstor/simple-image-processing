package pp

import (
	"image"
	"simple-image-processing/internal/imgproc"
)

func Green(processedImage *image.RGBA) {

	pix := make([]uint8, len(processedImage.Pix))

	imgproc.MultithreadProcessCycle(processedImage.Rect, func(x, y int) {
		pos := y * processedImage.Stride + x*4 + 1
		imgproc.SetPixColor(pix, x, y, processedImage.Stride, 0, processedImage.Pix[pos], 0, 255)
	})

	processedImage.Pix = pix
}
