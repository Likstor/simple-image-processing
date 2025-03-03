package imgproc

import (
	"image"
	"image/color"
	"sync"
)

func AdjustBrightness(processedImage *image.RGBA, param int) {
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
				for y := range height {
					rgba := processedImage.At(x, y).(color.RGBA)
					
					processedImage.Set(x, y, color.RGBA{
						R: LimitFrom0To255(int(rgba.R) + param),
						G: LimitFrom0To255(int(rgba.G) + param),
						B: LimitFrom0To255(int(rgba.B) + param),
						A: rgba.A,
					})
				}
			}
			wg.Done()
		}()
	}

	wg.Wait()
}