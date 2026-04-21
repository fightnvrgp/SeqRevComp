package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"

	"seqrevcomp/internal/gui"
)

func main() {
	a := app.NewWithID("com.example.seqrevcomp")
	a.Settings().SetTheme(theme.LightTheme())

	win := a.NewWindow("DNA/RNA 序列反向互补工具")
	win.Resize(fyne.NewSize(1100, 700))
	win.SetMaster()

	// Setup native menu bar (macOS top-left, Windows/Linux window menu)
	win.SetMainMenu(fyne.NewMainMenu(
		fyne.NewMenu("帮助",
			fyne.NewMenuItem("使用说明", func() {
				gui.OpenHelpWindow(a)
			}),
			fyne.NewMenuItem("关于", func() {
				gui.OpenAboutWindow(a)
			}),
		),
	))

	content := gui.BuildUI(a, win)
	win.SetContent(content)

	win.ShowAndRun()
}
