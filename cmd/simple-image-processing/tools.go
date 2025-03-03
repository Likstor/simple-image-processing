package main

import (
	"fmt"
	"image/color"
	"math"
	"simple-image-processing/internal/imgproc"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type Tool struct {
	Canvas fyne.CanvasObject
	Title  string
}

func CreateToolsContainer(w fyne.Window) fyne.CanvasObject {
	tools := make([]Tool, 0)

	tools = append(
		tools,
		CreateGrayScaleMenu(w),
		CreateAdjustBrightnessMenu(w),
		CreateNegativeMenu(w),
		CreateBinarizationMenu(w),
		CreateContrastMenu(w),
		CreateGammaConversionMenu(w),
		CreateQuantizationMenu(w),
	)

	list := widget.NewList(
		func() int {
			return len(tools)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Template tool")
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			item.(*widget.Label).SetText(tools[id].Title)
		},
	)

	selectLabel := container.NewCenter(widget.NewLabel("Select a tool from the list"))
	content := container.NewStack(selectLabel)

	list.OnSelected = func(id widget.ListItemID) {
		content.Objects = []fyne.CanvasObject{
			tools[id].Canvas,
		}
	}

	list.OnUnselected = func(id widget.ListItemID) {
		content.Objects = []fyne.CanvasObject{
			selectLabel,
		}
	}

	return container.NewHSplit(list, content)
}

func CreateNegativeMenu(w fyne.Window) Tool {
	thresholdSlider := widget.NewSlider(0, 255)
	thresholdSlider.Step = 1

	thresholdValue := widget.NewLabel(strconv.Itoa(int(thresholdSlider.Value)))
	thresholdValue.Resize(thresholdValue.MinSize())

	thresholdTitle := container.NewCenter(widget.NewLabel("Threshold"))

	thresholdContent := container.NewBorder(thresholdTitle, nil, nil, thresholdValue, thresholdSlider)

	thresholdSlider.OnChanged = func(f float64) {
		thresholdValue.SetText(strconv.Itoa(int(thresholdSlider.Value)))
	}

	btn := widget.NewButton("Accept", func() {
		if CurrentImage == nil {
			dialog.ShowError(ErrImageNotSelected, w)
			return
		}

		CurrentImage.SaveStep()

		imgproc.NegativeWithThreshold(CurrentImage.BaseImage, uint8(thresholdSlider.Value))

		CurrentImage.Refresh()
	})

	params := container.NewVBox(thresholdContent)

	paramsScroll := container.NewScroll(params)
	paramsScroll.ScrollToBottom()

	content := container.NewBorder(
		nil,
		btn,
		nil,
		nil,
		paramsScroll,
	)

	return Tool{
		Canvas: content,
		Title:  "Negative",
	}
}

func CreateGrayScaleMenu(w fyne.Window) Tool {
	btn := widget.NewButton("Accept", func() {
		if CurrentImage == nil {
			dialog.ShowError(ErrImageNotSelected, w)
			return
		}

		CurrentImage.SaveStep()

		imgproc.GrayScale(CurrentImage.BaseImage)

		CurrentImage.Refresh()
	})

	content := container.NewBorder(
		nil,
		btn,
		nil,
		nil,
	)

	return Tool{
		Canvas: content,
		Title:  "GrayScale",
	}
}

func pickColor(colorParam *color.RGBA, buttonLabel string, w fyne.Window) fyne.CanvasObject {
	rectangle := canvas.NewRectangle(colorParam)
	size := 2 * theme.IconInlineSize()
	rectangle.SetMinSize(fyne.NewSize(size, size))

	advancedColorPicker := dialog.NewColorPicker("Pick a color", "", func(c color.Color) {
		*colorParam = color.RGBA(c.(color.NRGBA))
		rectangle.FillColor = c
		rectangle.Refresh()
	}, w)
	advancedColorPicker.Advanced = true

	callPicker := widget.NewButton(buttonLabel, func() {
		advancedColorPicker.SetColor(colorParam)
		advancedColorPicker.Show()
	})

	return container.NewBorder(
		nil,
		nil,
		container.NewHBox(rectangle, callPicker),
		nil,
	)
}

func CreateBinarizationMenu(w fyne.Window) Tool {
	thresholdSlider := widget.NewSlider(0, 255)
	thresholdSlider.Step = 1

	thresholdValue := widget.NewLabel(strconv.Itoa(int(thresholdSlider.Value)))
	thresholdValue.Resize(thresholdValue.MinSize())

	thresholdTitle := container.NewCenter(widget.NewLabel("Threshold"))

	thresholdContent := container.NewBorder(thresholdTitle, nil, nil, thresholdValue, thresholdSlider)

	thresholdSlider.OnChanged = func(f float64) {
		thresholdValue.SetText(strconv.Itoa(int(thresholdSlider.Value)))
	}

	color1 := &color.RGBA{
		R: 255,
		B: 255,
		G: 255,
		A: 255,
	}
	picker1 := pickColor(color1, "Pick a first color", w)

	color2 := &color.RGBA{
		R: 0,
		B: 0,
		G: 0,
		A: 255,
	}
	picker2 := pickColor(color2, "Pick a second color", w)

	btn := widget.NewButton("Accept", func() {
		if CurrentImage == nil {
			dialog.ShowError(ErrImageNotSelected, w)
			return
		}

		CurrentImage.SaveStep()

		imgproc.ThresholdBinarization(CurrentImage.BaseImage, uint8(thresholdSlider.Value), color1, color2)

		CurrentImage.Refresh()
	})

	params := container.NewVBox(thresholdContent, picker1, picker2)

	paramsScroll := container.NewScroll(params)
	paramsScroll.ScrollToBottom()

	content := container.NewBorder(
		nil,
		btn,
		nil,
		nil,
		paramsScroll,
	)

	return Tool{
		Canvas: content,
		Title:  "Binarization",
	}
}

var (
	increaseContrast = "Increase"
	decreaseContrast = "Decrease"
)

func CreateContrastMenu(w fyne.Window) Tool {
	currentType := increaseContrast

	contrastTypeSelector := widget.NewSelect(
		[]string{increaseContrast, decreaseContrast},
		func(selected string) {
			currentType = selected
		},
	)
	contrastTypeSelector.Selected = increaseContrast

	contrastTypeLabel := widget.NewLabel("Type: ")

	contrastTypeContent := container.NewHBox(contrastTypeLabel, contrastTypeSelector)

	q1Slider := widget.NewSlider(0, 254)
	q1Slider.SetValue(0)
	q1Slider.Step = 1

	q1Value := widget.NewLabel(strconv.Itoa(int(q1Slider.Value)))
	q1Value.Resize(q1Value.MinSize())

	q1Title := container.NewCenter(widget.NewLabel("Q1"))

	q1Content := container.NewBorder(q1Title, nil, nil, q1Value, q1Slider)

	q2Slider := widget.NewSlider(1, 255)
	q2Slider.SetValue(255)
	q2Slider.Step = 1

	q2Value := widget.NewLabel(strconv.Itoa(int(q2Slider.Value)))
	q2Value.Resize(q2Value.MinSize())

	q2Title := container.NewCenter(widget.NewLabel("Q2"))

	q2Content := container.NewBorder(q2Title, nil, nil, q2Value, q2Slider)

	q1Slider.OnChanged = func(f float64) {
		q1Value.SetText(strconv.Itoa(int(q1Slider.Value)))

		if q1Slider.Value >= q2Slider.Value {
			q2Slider.SetValue(q2Slider.Value + 1)
		}
	}

	q2Slider.OnChanged = func(f float64) {
		q2Value.SetText(strconv.Itoa(int(q2Slider.Value)))

		if q2Slider.Value <= q1Slider.Value {
			q1Slider.SetValue(q1Slider.Value - 1)
		}
	}

	btn := widget.NewButton("Accept", func() {
		if CurrentImage == nil {
			dialog.ShowError(ErrImageNotSelected, w)
			return
		}

		CurrentImage.SaveStep()

		switch currentType {
		case increaseContrast:
			imgproc.IncreaseContrast(CurrentImage.BaseImage, uint8(q1Slider.Value), uint8(q2Slider.Value))
		case decreaseContrast:
			imgproc.DecreaseContrast(CurrentImage.BaseImage, uint8(q1Slider.Value), uint8(q2Slider.Value))
		}

		CurrentImage.Refresh()
	})

	params := container.NewVBox(contrastTypeContent, q1Content, q2Content)

	paramsScroll := container.NewScroll(params)
	paramsScroll.ScrollToBottom()

	content := container.NewBorder(
		nil,
		btn,
		nil,
		nil,
		paramsScroll,
	)

	return Tool{
		Canvas: content,
		Title:  "Contrast",
	}
}

func CreateAdjustBrightnessMenu(w fyne.Window) Tool {
	paramSlider := widget.NewSlider(-255, 255)
	paramSlider.Step = 1

	paramValue := widget.NewLabel(strconv.Itoa(int(paramSlider.Value)))
	paramValue.Resize(paramValue.MinSize())

	paramTitle := container.NewCenter(widget.NewLabel("Brightness parameter"))

	paramContent := container.NewBorder(paramTitle, nil, nil, paramValue, paramSlider)

	paramSlider.OnChanged = func(f float64) {
		paramValue.SetText(strconv.Itoa(int(paramSlider.Value)))
	}

	btn := widget.NewButton("Accept", func() {
		if CurrentImage == nil {
			dialog.ShowError(ErrImageNotSelected, w)
			return
		}

		CurrentImage.SaveStep()

		imgproc.AdjustBrightness(CurrentImage.BaseImage, int(paramSlider.Value))

		CurrentImage.Refresh()
	})

	params := container.NewVBox(paramContent)

	paramsScroll := container.NewScroll(params)
	paramsScroll.ScrollToBottom()

	content := container.NewBorder(
		nil,
		btn,
		nil,
		nil,
		paramsScroll,
	)

	return Tool{
		Canvas: content,
		Title:  "Brightness",
	}
}

func CreateGammaConversionMenu(w fyne.Window) Tool {
	currentGamma := 2.

	gammaSlider := widget.NewSlider(-255, 255)
	gammaSlider.Step = 1
	gammaSlider.SetValue(currentGamma)

	gammaValue := widget.NewLabel(fmt.Sprintf("Gamma: %0.f", currentGamma))

	gammaTitle := container.NewCenter(widget.NewLabel("Gamma parameter"))

	gammaContent := container.NewVBox(gammaTitle,gammaValue, gammaSlider)

	gammaSlider.OnChanged = func(f float64) {
		if gammaSlider.Value < 2 && gammaSlider.Value >= 0 {
			gammaSlider.SetValue(-2)
		} else if gammaSlider.Value > -2 && gammaSlider.Value < 0 {
			gammaSlider.SetValue(2)
		}

		if gammaSlider.Value < 0 {
			gammaValue.SetText(fmt.Sprintf("Gamma: 1/%0.f", math.Abs(gammaSlider.Value)))
		} else {
			gammaValue.SetText(fmt.Sprintf("Gamma: %0.f", gammaSlider.Value))
		}
	}

	btn := widget.NewButton("Accept", func() {
		if CurrentImage == nil {
			dialog.ShowError(ErrImageNotSelected, w)
			return
		}

		CurrentImage.SaveStep()

		if gammaSlider.Value < 0 {
			currentGamma = 1. / math.Abs(gammaSlider.Value)
		} else {
			currentGamma = gammaSlider.Value
		}

		imgproc.GammaConversion(CurrentImage.BaseImage, currentGamma)

		CurrentImage.Refresh()
	})

	params := container.NewVBox(gammaContent)

	paramsScroll := container.NewScroll(params)
	paramsScroll.ScrollToBottom()

	content := container.NewBorder(
		nil,
		btn,
		nil,
		nil,
		paramsScroll,
	)

	return Tool{
		Canvas: content,
		Title:  "Gamma conversion",
	}
}

func CreateQuantizationMenu(w fyne.Window) Tool {
	kSlider := widget.NewSlider(1, 255)
	kSlider.SetValue(1)
	kSlider.Step = 1

	kValue := widget.NewLabel(strconv.Itoa(int(kSlider.Value)))
	kValue.Resize(kValue.MinSize())

	kTitle := container.NewCenter(widget.NewLabel("Parameter K"))

	kContent := container.NewBorder(kTitle, nil, nil, kValue, kSlider)

	kSlider.OnChanged = func(f float64) {
		kValue.SetText(strconv.Itoa(int(kSlider.Value)))
	}

	btn := widget.NewButton("Accept", func() {
		if CurrentImage == nil {
			dialog.ShowError(ErrImageNotSelected, w)
			return
		}

		CurrentImage.SaveStep()

		imgproc.Quantization(CurrentImage.BaseImage, int(kSlider.Value))

		CurrentImage.Refresh()
	})

	params := container.NewVBox(kContent)

	paramsScroll := container.NewScroll(params)
	paramsScroll.ScrollToBottom()

	content := container.NewBorder(
		nil,
		btn,
		nil,
		nil,
		paramsScroll,
	)

	return Tool{
		Canvas: content,
		Title:  "Quantization",
	}
}