package imgproc

import (
	"image"
	"image/color"
	"sync"
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
	width, height := processedImage.Bounds().Dx(), processedImage.Bounds().Dy()
	step := width / MAXPROCS
	if (float64(width) / float64(MAXPROCS)) > float64(step) {
		step++
	}

	preparedColors := prepareColors(borders, colors)

	wg := sync.WaitGroup{}

	wg.Add(MAXPROCS)
	for i := range MAXPROCS {
		go func() {
			for x := step * i; x < step * (i + 1) && x < width; x++ {
				for y := range height {
					rgba := processedImage.At(x, y).(color.RGBA)

					I := uint8(0.3*float64(rgba.R) + 0.59*float64(rgba.G) + 0.11*float64(rgba.B))

					processedImage.Set(x, y, *preparedColors[I])
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()
}