package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/cmd/fyne_settings/settings"
	"fyne.io/fyne/v2/theme"
)

func CreateSettingsWindow() fyne.Window {
	settingsWindow := fyne.CurrentApp().NewWindow("Color histogram tool")
	settingsWindow.SetIcon(theme.Icon(theme.IconNameSettings))

	settingsWindow.SetContent(settings.NewSettings().LoadAppearanceScreen(settingsWindow))

	return settingsWindow
}