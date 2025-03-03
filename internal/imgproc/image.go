package imgproc

import (
	"image"
	"image/color"
)

type Image interface {
	Set(x, y int, color color.Color)
	Bounds() image.Rectangle
	At(x, y int) color.Color
	ColorModel() color.Model
}
