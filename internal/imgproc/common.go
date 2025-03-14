package imgproc

import (
	"image"
	"image/color"
	"math"
	"sync"
)

const (
	RedChannel = iota
	GreenChannel
	BlueChannel
)

type Matrix [][]float64

func DrawRectangle(dst *image.RGBA, x1, y1, x2, y2 int, c color.RGBA) {
	for x := x1; x < x2; x++ {
		for y := y1; y < y2; y++ {

			currentColor := dst.At(x, y).(color.RGBA)

			alpha := float64(c.A) / 255.0
			newR := uint8(float64(currentColor.R)*(1-alpha) + float64(c.R)*(alpha))
			newG := uint8(float64(currentColor.G)*(1-alpha) + float64(c.G)*(alpha))
			newB := uint8(float64(currentColor.B)*(1-alpha) + float64(c.B)*(alpha))

			dst.Set(x, y, &image.Uniform{
				color.RGBA{
					R: newR,
					G: newG,
					B: newB,
					A: 255,
				},
			},
			)
		}
	}
}

func LimitFrom0To255[T int | float64](value T) uint8 {
	if value > 255 {
		return 255
	} else if value < 0 {
		return 0
	}

	return uint8(value)
}

func MultithreadProcessCycle(rec image.Rectangle, f func(x, y int)) {
	width, height := rec.Dx(), rec.Dy()
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
					f(x, y)
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func KernelCycle(kernel Matrix, imageWidth, imageHeight, xKernel, yKernel int, f func(x, y, i, j int)) {
	p := len(kernel) / 2
	xStart := xKernel - p

	yStart := yKernel - p

	for i := range len(kernel) {
		for j := range len(kernel) {
			x := xStart + j
			y := yStart + i

			switch {
			case x < 0:
				x = 0
			case x >= imageWidth:
				x = imageWidth - 1
			}

			switch {
			case y < 0:
				y = 0
			case y >= imageHeight:
				y = imageHeight - 1
			}

			f(x, y, i, j)
		}
	}
}

func Convolution(img *image.RGBA, kernel Matrix, xKernel, yKernel int) (r float64, g float64, b float64) {
	width, height := img.Rect.Dx(), img.Rect.Dy()

	KernelCycle(kernel, width, height, xKernel, yKernel, func(x, y, i, j int) {
		rgba := img.At(x, y).(color.RGBA)

		coefficient := kernel[i][j]

		r += float64(rgba.R) * coefficient
		g += float64(rgba.G) * coefficient
		b += float64(rgba.B) * coefficient
	})

	return
}

func ConvolutionOneChannel(img *image.RGBA, kernel Matrix, xKernel, yKernel, channel int) (c float64) {
	width, height := img.Rect.Dx(), img.Rect.Dy()

	var (
		f           func()
		rgba        color.RGBA
		coefficient float64
	)

	switch channel {
	case RedChannel:
		f = func() {
			c += float64(rgba.R) * coefficient
		}
	case GreenChannel:
		f = func() {
			c += float64(rgba.G) * coefficient
		}
	case BlueChannel:
		f = func() {
			c += float64(rgba.B) * coefficient
		}
	}

	KernelCycle(kernel, width, height, xKernel, yKernel, func(x, y, i, j int) {
		rgba = img.At(x, y).(color.RGBA)

		coefficient = kernel[i][j]

		f()
	})

	return
}

// Normalization получает массив пикселей вида {r1, g1, b1, a1, r2, g2, b2, a1, ..., rn, gn, bn, an}
// и нормализует значения для каждого канала по формуле 255 * (val - min) / (max - min).
// Если полученное значение больше 255, то оно ограничивается 255.
// Если полученное значение меньше 0, то оно ограничивается 0.
func Normalization(pix []float64) []uint8 {
	minR, maxR := math.MaxFloat64, -math.MaxFloat64
	minG, maxG := math.MaxFloat64, -math.MaxFloat64
	minB, maxB := math.MaxFloat64, -math.MaxFloat64

	end := len(pix)

	for i := 0; i < end; i += 4 {
		r := pix[i]
		g := pix[i+1]
		b := pix[i+2]

		if r < minR {
			minR = r
		}
		if r > maxR {
			maxR = r
		}
		if g < minG {
			minG = g
		}
		if g > maxG {
			maxG = g
		}
		if b < minB {
			minB = b
		}
		if b > maxB {
			maxB = b
		}
	}

	f := func(val, min, max float64) uint8 {
		if max == min {
			return 0
		}

		norm := 255 * (val - min) / (max - min)

		return LimitFrom0To255(norm)
	}

	normalized := make([]uint8, 0, end)

	for i := 0; i < end; i += 4 {
		r := pix[i]
		g := pix[i+1]
		b := pix[i+2]

		normalized = append(normalized, f(r, minR, maxR), f(g, minG, maxG), f(b, minB, maxB), 255)
	}

	return normalized
}

func KernelNormalization(kernel Matrix) [][]float64 {
	sum := 0.

	for i := range len(kernel) {
		for j := range len(kernel[i]) {
			sum += kernel[i][j]
		}
	}

	for i := range len(kernel) {
		for j := range len(kernel[i]) {
			kernel[i][j] /= sum
		}
	}

	return kernel
}

func SetPixColor(pix []uint8, x, y, stride int, r, g, b, a uint8) {
	offset := y*stride + x*4
	pix[offset] = r
	pix[offset+1] = g
	pix[offset+2] = b
	pix[offset+3] = a
}

func SetPixColorWithLimitFrom0To255(pix []uint8, x, y, stride int, r, g, b, a float64) {
	SetPixColor(pix, x, y, stride, LimitFrom0To255(r), LimitFrom0To255(g), LimitFrom0To255(b), LimitFrom0To255(a))
}