package imgproc

import (
	"image"
	"image/color"
	"sync"
)

func solarizationFx(x uint8) uint8 {
	return uint8((-4./255. * float64(x) + 4.) * float64(x))
}

func Solarization(processedImage *image.RGBA) {
	width, height := processedImage.Bounds().Dx(), processedImage.Bounds().Dy()
	step := width / MAXPROCS
	if (float64(width) / float64(MAXPROCS)) > float64(step) {
		step++
	}

	wg := sync.WaitGroup{}

	wg.Add(MAXPROCS)
	for i := range MAXPROCS {
		go func() {
			start := step * i

			end := min(step*(i+1), width)

			for x := start; x < end; x++ {
				for y := range height {
					rgba := processedImage.At(x, y).(color.RGBA)

					processedImage.Set(x, y, color.RGBA{
						R: solarizationFx(rgba.R),
						G: solarizationFx(rgba.G),
						B: solarizationFx(rgba.B),
						A: rgba.A,
					})

				}
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
