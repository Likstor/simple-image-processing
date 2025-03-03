package imgproc

import (
	"image"
	"image/color"
)

func DrawRectangle(img Image, x1, y1, x2, y2 int, c color.RGBA) {
	for x := x1; x < x2; x++ {
		for y := y1; y < y2; y++ {

			currentColor := img.At(x, y).(color.RGBA)

			alpha := float64(c.A) / 255.0
			newR := uint8(float64(currentColor.R)*(1-alpha) + float64(c.R)*(alpha))
			newG := uint8(float64(currentColor.G)*(1-alpha) + float64(c.G)*(alpha))
			newB := uint8(float64(currentColor.B)*(1-alpha) + float64(c.B)*(alpha))

			img.Set(x, y, &image.Uniform{
				color.RGBA{
					R: newR,
					G: newG,
					B: newB,
					A: 255,
				},
			},
			)
		}
	}
}

func LimitFrom0To255[T int | float64](value T) uint8 {
	if value > 255 {
		return 255
	} else if value < 0 {
		return 0
	}

	return uint8(value)
}