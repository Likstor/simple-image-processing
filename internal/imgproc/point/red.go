package pp

import (
	"image"
	"simple-image-processing/internal/imgproc"
)

func Red(processedImage *image.RGBA) {

	pix := make([]uint8, len(processedImage.Pix))

	imgproc.MultithreadProcessCycle(processedImage.Rect, func(x, y int) {
		pos := y * processedImage.Stride + x*4
		imgproc.SetPixColor(pix, x, y, processedImage.Stride, processedImage.Pix[pos], 0, 0, 255)
	})

	processedImage.Pix = pix
}
