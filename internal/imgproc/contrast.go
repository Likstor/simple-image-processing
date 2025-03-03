package imgproc

import (
	"image"
	"image/color"
	"sync"
)

func IncreaseContrast(processedImage *image.RGBA, q1, q2 uint8) {
	width, height := processedImage.Bounds().Dx(), processedImage.Bounds().Dy()
	step := width / MAXPROCS
	if (float64(width) / float64(MAXPROCS)) > float64(step) {
		step++
	}

	coefficient := 255. / float64(q2 - q1)

	wg := sync.WaitGroup{}

	wg.Add(MAXPROCS)

	for i := range MAXPROCS {
		go func() {
			for x := step * i; x < step*(i+1) && x < width; x++ {
				for y := range height {
					rgba := processedImage.At(x, y).(color.RGBA)

					newR := LimitFrom0To255(float64(int(rgba.R) - int(q1)) * coefficient)
					newG := LimitFrom0To255(float64(int(rgba.G) - int(q1)) * coefficient)
					newB := LimitFrom0To255(float64(int(rgba.B) - int(q1)) * coefficient)

					processedImage.Set(x, y, color.RGBA{
						R: newR,
						G: newG,
						B: newB,
						A: rgba.A,
					})
				}
			}
			wg.Done()
		}()
	}

	wg.Wait()
}

func DecreaseContrast(processedImage *image.RGBA, q1, q2 uint8) {
	width, height := processedImage.Bounds().Dx(), processedImage.Bounds().Dy()
	step := width / MAXPROCS
	if (float64(width) / float64(MAXPROCS)) > float64(step) {
		step++
	}

	coefficient := float64(q2 - q1) / 255.

	wg := sync.WaitGroup{}

	wg.Add(MAXPROCS)

	for i := range MAXPROCS {
		go func() {
			for x := step * i; x < step*(i+1) && x < width; x++ {
				for y := range height {
					rgba := processedImage.At(x, y).(color.RGBA)

					newR := LimitFrom0To255(float64(q1) + float64(rgba.R) * coefficient)
					newG := LimitFrom0To255(float64(q1) + float64(rgba.G) * coefficient)
					newB := LimitFrom0To255(float64(q1) + float64(rgba.B) * coefficient)

					processedImage.Set(x, y, color.RGBA{
						R: newR,
						G: newG,
						B: newB,
						A: rgba.A,
					})
				}
			}
			wg.Done()
		}()
	}

	wg.Wait()
}
