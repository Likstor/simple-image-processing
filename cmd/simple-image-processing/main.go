package main

import (
	"runtime"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

var (
	ImagesContainer = &container.AppTabs{}

	Images     = make(map[string]*ImageData)
	ImageByTab = make(map[*container.TabItem]*ImageData)

	CurrentImage *ImageData

	MainWindow fyne.Window
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	imageEditor := app.NewWithID("sip")
	w := imageEditor.NewWindow("Image")

	w.SetIcon(theme.Icon(theme.IconNameColorPalette))

	w.Resize(fyne.NewSize(800, 600))
	w.SetMaster()

	ImagesContainer = container.NewAppTabs()

	MainWindow = w

	SetupDragAndDrop(w)

	SetupMenu(w)

	nav := CreateNav(w)

	content := container.NewHSplit(nav, ImagesContainer)

	ImagesContainer.OnSelected = func(tab *container.TabItem) {
		CurrentImage = ImageByTab[tab]

		MenuSaveCurrentFileAs.Disabled = false
		MenuCloseAllFiles.Disabled = false
		MenuCloseCurrentFile.Disabled = false
		
		MenuRedo.Disabled = false
		MenuUndo.Disabled = false
	}

	overlay := container.NewStack(content)

	w.SetContent(overlay)

	w.Canvas().AddShortcut(&fyne.ShortcutUndo{}, func(shortcut fyne.Shortcut) {
		CurrentImage.Undo()
	})

	w.Canvas().AddShortcut(&fyne.ShortcutRedo{}, func(shortcut fyne.Shortcut) {
		CurrentImage.Redo()
	})

	w.ShowAndRun()
}
