package imgproc

import (
	"image/color"
	"sync"
)

func ThresholdBinarization(processedImage Image, threshold uint8, color1, color2 color.Color) {
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

					if uint8(0.3*float64(rgba.R)+0.59*float64(rgba.G)+0.11*float64(rgba.B)) >= threshold {
						processedImage.Set(x, y, color1)
					} else {
						processedImage.Set(x, y, color2)
					}
				}
			}
			wg.Done()
		}()
	}

	wg.Wait()
}
