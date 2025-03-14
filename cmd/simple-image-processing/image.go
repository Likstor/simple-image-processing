package main

import (
	"container/list"
	"errors"
	"fmt"
	"image"
	"image/draw"
	"io"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
)

type ImageData struct {
	URI         fyne.URI
	BaseImage   *image.RGBA
	Tab         *container.TabItem
	canvasImage *canvas.Image
	steps *list.List
	currentStep *list.Element
}

func NewImageData(imgData *ImageData) {
	imgData.canvasImage = imgData.Tab.Content.(*canvas.Image)

	imgData.steps = list.New()

	Images[imgData.URI.Path()] = imgData
	ImageByTab[imgData.Tab] = imgData

	ImagesContainer.Append(imgData.Tab)

	ImagesContainer.Select(imgData.Tab)

	imgData.currentStep = imgData.steps.PushFront(imgData.canvasImage)
}

func (img *ImageData) Refresh() {
	img.canvasImage.Refresh()
}

func (img *ImageData) SaveStep() {
	dst := image.NewRGBA(img.BaseImage.Rect)
	copy(dst.Pix, img.BaseImage.Pix)

	for img.currentStep != img.steps.Front() {
		img.steps.Remove(img.steps.Front())
	}
	
	img.currentStep.Value = dst
	
	img.currentStep = img.steps.PushFront(img.BaseImage)
}

func (img *ImageData) Undo() {
	older := img.currentStep.Next()
	if older != nil {
		olderImg := older.Value.(*image.RGBA)

		img.currentStep = older

		img.canvasImage.Image = olderImg
		img.BaseImage = olderImg
		
		img.Refresh()
	}
}

func (img *ImageData) Redo() {
	newer := img.currentStep.Prev()
	if newer != nil {
		img.currentStep = newer

		newerImg := newer.Value.(*image.RGBA)

		img.canvasImage.Image = newerImg
		img.BaseImage = newerImg

		img.Refresh()
	}
}

func (img *ImageData) GetCurrentStep() *image.RGBA {
	return img.currentStep.Value.(*image.RGBA)
}


func addNewImage(reader io.Reader, uri fyne.URI) {
	baseImage, _, err := image.Decode(reader)
	if err != nil {
		if errors.Is(err, image.ErrFormat) {
			dialog.ShowError(fmt.Errorf("unknown image type"), MainWindow)
			return
		}
	}

	processedImage := image.NewRGBA(baseImage.Bounds())
	draw.Draw(processedImage, baseImage.Bounds(), baseImage, image.Point{}, draw.Src)

	image := canvas.NewImageFromImage(processedImage)
	image.FillMode = canvas.ImageFillContain
	image.ScaleMode = canvas.ImageScaleFastest

	newImageTab := container.NewTabItem(uri.Name(), image)

	imageStruct := &ImageData{
		URI:         uri,
		BaseImage:   processedImage,
		Tab:         newImageTab,
	}

	NewImageData(imageStruct)
}
