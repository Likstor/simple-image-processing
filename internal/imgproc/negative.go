package imgproc

import (
	"image/color"
	"sync"
)
func Negative(processedImage Image) {
	NegativeWithThreshold(processedImage, 0)
}

func NegativeWithThreshold(processedImage Image, threshold uint8) {
	width, height := processedImage.Bounds().Dx(), processedImage.Bounds().Dy()
	step := width / MAXPROCS
	if (float64(width) / float64(MAXPROCS)) > float64(step) {
		step++
	}

	wg := sync.WaitGroup{}

	wg.Add(MAXPROCS)
	for i := range MAXPROCS {
		go func() {
			for x := step * i; x < step * (i + 1) && x < width; x++ {
				for y := 0; y < height; y++ {
					processedImage.Set(x, y, negativeColor(processedImage.At(x, y), threshold))
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func negativeColor(reversedColor color.Color, threshold uint8) color.Color {
	switch v := reversedColor.(type) {
	case color.RGBA:
		if v.R >= threshold {
			v.R = 255 - v.R
		}

		if v.G >= threshold {
			v.G = 255 - v.G
		}

		if v.B >= threshold {
			v.B = 255 - v.B
		}

		return v
	case color.Gray:
		if v.Y >= threshold {
			v.Y = 255 - v.Y
		}

		return v
	}

	panic("unknown color type")
}
