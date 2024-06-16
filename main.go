package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"os/exec"
)

func main() {
	a := app.New()
	w := a.NewWindow("Main window")

	w.SetContent(widget.NewLabel("Hello World!"))
	factText := widget.NewLabel("test123")
	button := widget.NewButton("Speed test", func() {
		// =====> do some stuff <==== //
		fmt.Println("I was pressed")

		cmd := exec.Command("ls", "-l")
		output, err := cmd.Output()
		if err != nil {
			fmt.Println("Error executing command:", err)
			return
		}
		fmt.Println(string(output))
		factText.SetText(string(output))
	})

	title := canvas.NewText("System Utils", color.White)
	title.TextStyle = fyne.TextStyle{
		Bold: true,
	}
	title.Alignment = fyne.TextAlignCenter
	title.TextSize = 24

	hBox := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), button, layout.NewSpacer())
	vBox := container.New(layout.NewVBoxLayout(), title, hBox, widget.NewSeparator(), factText)

	w.SetContent(vBox)
	w.Show()
	a.Run()
}
