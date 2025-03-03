package imgproc

import (
	"image"
	"image/color"
	"image/draw"
	"slices"
)

var HISTOGRAM_RECTANGLE = image.Rectangle{
	Max: image.Point{
		X: HISTOGRAM_WIDTH,
		Y: HISTOGRAM_HEIGHT,
	},
}

func UpdateColorHistogram(img Image, dst *image.RGBA, backgroundColor, lineColor color.RGBA) {
	count := make([]float64, 256*3)
	countRedColor := count[:256]
	countGreenColor := count[256 : 2*256]
	countBlueColor := count[2*256:]

	draw.Draw(dst, dst.Bounds(), &image.Uniform{backgroundColor}, image.Point{}, draw.Src)

	imgWidth, imgHeight := img.Bounds().Dx(), img.Bounds().Dy()

	for x := range imgWidth {
		for y := 0; y < imgHeight; y++ {
			color := img.At(x, y).(color.RGBA)

			countRedColor[color.R]++
			countGreenColor[color.G]++
			countBlueColor[color.B]++
		}
	}

	maxH := slices.Max(count)

	for i := range 256 {
		drawHistogramColumn(dst, countBlueColor[i]/maxH, i * 5, color.RGBA{B: 255, A: 40})
		drawHistogramColumn(dst, countGreenColor[i]/maxH, i * 5, color.RGBA{G: 255, A: 40})
		drawHistogramColumn(dst, countRedColor[i]/maxH, i * 5, color.RGBA{R: 255, A: 40})
	}

	DrawRectangle(dst, 0, HISTOGRAM_HEIGHT - 11, HISTOGRAM_WIDTH - 1, HISTOGRAM_HEIGHT - 1, lineColor)
}

func drawHistogramColumn(histogram Image, percent float64, pos int, c color.RGBA) {
	if percent == 0 {
		return
	}

	top := (HISTOGRAM_HEIGHT - 10) - int(float64(HISTOGRAM_HEIGHT - 20) * percent)

	topColor := color.RGBA{
		R: c.R,
		G: c.G,
		B: c.B,
		A: 255,
	}

	DrawRectangle(histogram, pos, top, pos + 4, top + 4, topColor)
	DrawRectangle(histogram, pos, top + 5, pos + 4, HISTOGRAM_HEIGHT - 11, c)
}