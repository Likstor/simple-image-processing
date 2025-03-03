package imgproc

import (
	"image"
	"image/color"
	"log"
	"math"
	"sync"
)

func prepareQuants(k int) map[uint8]uint8 {
	quants := make(map[uint8]uint8, 256)

	quantSize := int(math.Ceil(256. / float64(k)))

	log.Println(quantSize, k)

	for i := range k {
		c := uint8(LimitFrom0To255(quantSize - 1 + quantSize*i))

		start := quantSize * i
		end := min(quantSize + quantSize*i, 256)

		for j := start; j < end; j++ {
			quants[uint8(j)] = c
		}
	}
	return quants
}

func Quantization(processedImage *image.RGBA, k int) {
	width, height := processedImage.Bounds().Dx(), processedImage.Bounds().Dy()
	step := width / MAXPROCS
	if (float64(width) / float64(MAXPROCS)) > float64(step) {
		step++
	}

	wg := sync.WaitGroup{}

	quants := prepareQuants(k)

	wg.Add(MAXPROCS)
	for i := 0; i < MAXPROCS; i++ {
		go func() {
			for x := step * i; x < step*(i+1) && x < width; x++ {
				for y := 0; y < height; y++ {
					rgba := processedImage.At(x, y).(color.RGBA)

					processedImage.Set(x, y, color.RGBA{
						R: quants[rgba.R],
						G: quants[rgba.G],
						B: quants[rgba.B],
						A: rgba.A,
					})

				}
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
