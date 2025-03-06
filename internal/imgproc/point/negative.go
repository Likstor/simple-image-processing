package pp

import (
	"image"
	"image/color"
	"simple-image-processing/internal/imgproc"
)

func Negative(processedImage *image.RGBA, threshold uint8) {
	imgproc.MultithreadPointProcessCycle(processedImage.Rect, func(x, y int) {
		processedImage.Set(x, y, negativeColor(processedImage.At(x, y).(color.RGBA), threshold))
	})
}

func negativeColor(reversedColor color.RGBA, threshold uint8) color.Color {
	if threshold > 0 {
		if 0.3*float64(reversedColor.R)+0.59*float64(reversedColor.G)+0.11*float64(reversedColor.B) < float64(threshold) {
			return reversedColor
		}
	}

	reversedColor.R = 255 - reversedColor.R
	reversedColor.G = 255 - reversedColor.G
	reversedColor.B = 255 - reversedColor.B

	return reversedColor
}
