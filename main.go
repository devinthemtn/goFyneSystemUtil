package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

func main() {
	a := app.New()
	w := a.NewWindow("Main window")

	w.SetContent(widget.NewLabel("Hello World!"))
	button := widget.NewButton("Speed test", func() {
		// =====> do some stuff <==== //
	})

	title := canvas.NewText("Get Your Useless Facts", color.White)
	title.TextStyle = fyne.TextStyle{
		Bold: true,
	}
	title.Alignment = fyne.TextAlignCenter
	title.TextSize = 24

	factText := widget.NewLabel("test123")
	hBox := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), button, layout.NewSpacer())
	vBox := container.New(layout.NewVBoxLayout(), title, hBox, widget.NewSeparator(), factText)

	w.SetContent(vBox)
	w.Show()
	a.Run()
}
