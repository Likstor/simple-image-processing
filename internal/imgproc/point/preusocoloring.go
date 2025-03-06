package pp

import (
	"image"
	"image/color"
	"simple-image-processing/internal/imgproc"
)

func prepareColors(borders []uint8, colors []*color.RGBA) [256]*color.RGBA {
	start := 0

	var res [256]*color.RGBA

	for k, b := range borders {
		for i := start; i <= int(b); i++ {
			res[i] = colors[k]
		}

		start = int(b + 1)
	}

	return res
}

func PseudoColoring(processedImage *image.RGBA, borders []uint8, colors []*color.RGBA) {
	preparedColors := prepareColors(borders, colors)

	imgproc.MultithreadPointProcessCycle(processedImage.Rect, func(x, y int) {
		rgba := processedImage.At(x, y).(color.RGBA)

		I := uint8(0.3*float64(rgba.R) + 0.59*float64(rgba.G) + 0.11*float64(rgba.B))

		processedImage.Set(x, y, *preparedColors[I])
	})
}
