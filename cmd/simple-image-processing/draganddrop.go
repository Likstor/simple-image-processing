package main

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

func SetupDragAndDrop(w fyne.Window) {
	w.SetOnDropped(func(pos fyne.Position, uris []fyne.URI) {
		for _, uri := range uris {
			if _, ok := Images[uri.Path()]; ok {
				dialog.ShowError(fmt.Errorf("file \"%s\" is already opened", uri.Path()), w)
				continue
			}

			file, _ := os.Open(uri.Path())

			addNewImage(file, uri)
		}
	})
}
