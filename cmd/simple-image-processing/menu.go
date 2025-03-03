package main

import (
	"errors"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
)

var (
	MenuOpenFile, MenuCloseAllFiles, MenuCloseCurrentFile, MenuSaveCurrentFileAs *fyne.MenuItem

	MenuUndo, MenuRedo *fyne.MenuItem

	MenuHistogram, MenuSettings *fyne.MenuItem
)

func SetupMenu(w fyne.Window) {
	fileMenu := createFileMenu(w)
	editMenu := createEditMenu()
	toolsMenu := createMenuTools()

	w.SetMainMenu(fyne.NewMainMenu(fileMenu, editMenu, toolsMenu))
}

func createFileMenu(w fyne.Window) *fyne.Menu {
	MenuOpenFile = fyne.NewMenuItem("Open", func() {
		fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, w)
				return
			}

			if reader == nil {
				log.Println("Cancelled")
				return
			}

			if _, ok := Images[reader.URI().Path()]; ok {
				dialog.ShowError(errors.New("this file is already opened"), w)
				return
			}

			addNewImage(reader, reader.URI())
		}, w)

		path, err := os.Getwd()
		if err != nil {
			dialog.ShowError(err, w)
			return
		}

		luri, err := storage.ListerForURI(storage.NewFileURI(path))
		if err == nil {
			fd.SetLocation(luri)
		}

		fd.SetFilter(storage.NewExtensionFileFilter([]string{".png", ".jpg", ".jpeg"}))

		fd.Show()
	})

	MenuCloseCurrentFile = fyne.NewMenuItem("Close current file", func() {
		delete(ImageByTab, CurrentImage.Tab)
		delete(Images, CurrentImage.URI.Path())
		ImagesContainer.Remove(CurrentImage.Tab)

		if ImagesContainer.Selected() == nil {
			MenuSaveCurrentFileAs.Disabled = true
			MenuCloseAllFiles.Disabled = true
			MenuCloseCurrentFile.Disabled = true

			MenuUndo.Disabled = true
			MenuRedo.Disabled = true
		}
	})
	MenuCloseCurrentFile.Disabled = true

	MenuCloseAllFiles = fyne.NewMenuItem("Close all files", func() {
		for ImagesContainer.Selected() != nil {
			ImagesContainer.Remove(CurrentImage.Tab)
		}

		MenuSaveCurrentFileAs.Disabled = true
		MenuCloseAllFiles.Disabled = true
		MenuCloseCurrentFile.Disabled = true

		MenuUndo.Disabled = true
		MenuRedo.Disabled = true
	})
	MenuCloseAllFiles.Disabled = true

	MenuSaveCurrentFileAs = fyne.NewMenuItem("Save as...", func() {
		fs := dialog.NewFileSave(func(writer fyne.URIWriteCloser, err error) {
			if err != nil {
				dialog.ShowError(err, w)
				return
			}

			if writer == nil {
				log.Println("Canceled")
				return
			}

			uri := writer.URI()

			isJpeg := true

			if _, ok := strings.CutSuffix(uri.Name(), ".png"); ok {
				isJpeg = false
			}

			if isJpeg {
				err := jpeg.Encode(writer, CurrentImage.BaseImage, nil)
				if err != nil {
					dialog.ShowError(err, w)
					return
				}
			} else {
				err := png.Encode(writer, CurrentImage.BaseImage)
				if err != nil {
					dialog.ShowError(err, w)
					return
				}
			}

			delete(Images, CurrentImage.URI.Path())

			CurrentImage.URI = writer.URI()
			CurrentImage.Tab.Text = writer.URI().Name()
			
			Images[CurrentImage.URI.Path()] = CurrentImage

			ImagesContainer.Refresh()
		}, w)

		path := filepath.Dir(CurrentImage.URI.Path())

		luri, err := storage.ListerForURI(storage.NewFileURI(path))
		if err == nil {
			fs.SetLocation(luri)
		}

		fs.SetFileName(CurrentImage.URI.Name())

		fs.Show()
	})
	MenuSaveCurrentFileAs.Disabled = true

	return fyne.NewMenu(
		"File", 
		MenuOpenFile, 
		fyne.NewMenuItemSeparator(), 
		MenuCloseCurrentFile, 
		MenuCloseAllFiles, 
		fyne.NewMenuItemSeparator(), 
		MenuSaveCurrentFileAs,
	)
}

func createEditMenu() *fyne.Menu {
	MenuUndo = fyne.NewMenuItem("Undo", func() {
		if CurrentImage.currentStep != nil {
			CurrentImage.Undo()
		}
	})
	MenuUndo.Disabled = true
	MenuUndo.Shortcut = &fyne.ShortcutUndo{}

	MenuRedo = fyne.NewMenuItem("Redo", func() {
		if CurrentImage.currentStep != nil {
			CurrentImage.Redo()
		}
	})
	MenuRedo.Disabled = true
	MenuRedo.Shortcut = &fyne.ShortcutRedo{}

	return fyne.NewMenu("Edit", MenuUndo, MenuRedo)
}

func createMenuTools() *fyne.Menu {
	var histogramWindow fyne.Window

	MenuHistogram = fyne.NewMenuItem("Color histogram", func() {
		histogramWindow = CreateHistogramWindow()

		histogramWindow.Show()
	})

	MenuSettings = fyne.NewMenuItem("Settings", func() {
		settingsWindow := CreateSettingsWindow()

		settingsWindow.Show()
	})

	return fyne.NewMenu("Tools", MenuHistogram, MenuSettings)
}

