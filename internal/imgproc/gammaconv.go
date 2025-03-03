package imgproc

import (
	"image"
	"image/color"
	"math"
	"sync"
)

func GammaConversion(processedImage *image.RGBA, gamma float64) {
	width, height := processedImage.Bounds().Dx(), processedImage.Bounds().Dy()
	step := width / MAXPROCS
	if (float64(width) / float64(MAXPROCS)) > float64(step) {
		step++
	}

	wg := sync.WaitGroup{}

	wg.Add(MAXPROCS)
	for i := range MAXPROCS {
		go func() {
			for x := step * i; x < step*(i+1) && x < width; x++ {
				for y := range height {
					rgba := processedImage.At(x, y).(color.RGBA)

					processedImage.Set(x, y, color.RGBA{
						R: LimitFrom0To255(255. * math.Pow(float64(rgba.R)/255., gamma)),
						G: LimitFrom0To255(255. * math.Pow(float64(rgba.G)/255., gamma)),
						B: LimitFrom0To255(255. * math.Pow(float64(rgba.B)/255., gamma)),
						A: rgba.A,
					})
				}
			}
			wg.Done()
		}()
	}

	wg.Wait()
}
