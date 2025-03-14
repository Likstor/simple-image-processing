package sp

import (
	"image"
	"simple-image-processing/internal/imgproc"
)

func pascalRow(n int) []float64 {
	res := make([]float64, n)
	elem := 1.
	res[0] = 1.
	res[n-1] = 1.

	end := float64(n)/2. + 1

	for k := 1.; k < end; k++ {
		elem *= (float64(n-1) + 1 - k) / k

		pos := int(k)

		res[pos] = elem
		res[n-pos-1] = elem
	}

	return res
}

func GaussKernelByPascalRow(row []float64) imgproc.Matrix {
	res := make(imgproc.Matrix, 0, len(row))

	sum := 0.

	end := len(row)/2 + 1

	for range row {
		res = append(res, make([]float64, len(row)))
	}

	n := len(row) - 1

	setVal := func(i, j int, val float64) {
		res[i][j] = val
		res[n-i][j] = val
		res[i][n-j] = val
		res[n-i][n-j] = val
	}

	for i := range end {
		for j := range end {
			val := float64(row[i] * row[j])

			setVal(i, j, val)

			sum += val

			var t1, t2 bool

			if i != n-i {
				t1 = true
				sum += val
			}

			if j != n-j {
				t2 = true
				sum += val
			}

			if t1 && t2 {
				sum += val
			}
		}
	}

	for i := range end {
		for j := range end {
			val := res[i][j] / sum

			setVal(i, j, val)
		}
	}

	return res
}

func GaussBlur(processedImage *image.RGBA, n int) {
	kernel := GaussKernelByPascalRow(pascalRow(n))

	convPix := make([]uint8, len(processedImage.Pix))

	imgproc.MultithreadProcessCycle(processedImage.Rect, func(x, y int) {
		r, g, b := imgproc.Convolution(processedImage, kernel, x, y)

		imgproc.SetPixColorWithLimitFrom0To255(convPix, x, y, processedImage.Stride, r, g, b, 255)
	})

	processedImage.Pix = convPix
}
