package imgproc

import (
	"runtime"
)

var (
	MAXPROCS = 0
	HISTOGRAM_HEIGHT = 720
	HISTOGRAM_WIDTH = 1280
	HISTOGRAM_PIXEL_COUNT = HISTOGRAM_HEIGHT * HISTOGRAM_WIDTH
)

func init() {
	MAXPROCS = runtime.NumCPU()
}