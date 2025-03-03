package imgproc

import (
	"image"
	"image/color"
	"sync"
)

func GrayScale(processedImage *image.RGBA) {
	width, height := processedImage.Bounds().Dx(), processedImage.Bounds().Dy()
	step := width / MAXPROCS
	if (float64(width) / float64(MAXPROCS)) > float64(step) {
		step++
	}

	wg := sync.WaitGroup{}

	wg.Add(MAXPROCS)
	for i := 0; i < MAXPROCS; i++ {
		go func() {
			for x := step * i; x < step * (i + 1) && x < width; x++ {
				for y := 0; y < height; y++ {
					rgba := (processedImage.At(x, y)).(color.RGBA)

					I := uint8(0.3*float64(rgba.R) + 0.59*float64(rgba.G) + 0.11*float64(rgba.B))

					processedImage.Set(x, y, color.Gray{
						Y: I,
					})
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
