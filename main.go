package main

import (
	"fmt"
	"image/color"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	// "os"
	"os/exec"
)

func main() {
	a := app.New()
	w := a.NewWindow("Main window")

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("error getting users's home: ", err)
	}
	usrConfigDir := homeDir + "/.config/katSystemUtil"
	w.SetContent(widget.NewLabel("Hello World!"))
	factText := widget.NewLabel("")
	button2 := widget.NewButton("Speed Test", func() {
		// =====> do some stuff <==== //
		fmt.Println("I was pressed")
		factText.SetText("running speedtest")
		cmd := exec.Command("speedtest")
		// cmd.Stdout = os.Stdout
		// cmd.Stderr = os.Stderr
		output, err := cmd.Output()
		if err != nil {
			fmt.Println("Error executing command:", err)
			return
		}
		fmt.Println(string(output))
		factText.SetText(string(output))
	})
	button := widget.NewButton("Test", func() {
		// =====> do some stuff <==== //
		fmt.Println("I was pressed")
		cmd := exec.Command("ls", "-l")
		output, err := cmd.Output()
		if err != nil {
			fmt.Println("Error executing command:", err)
			return
		}
		// fmt.Println(string(output))
		factText.SetText(string(output))
	})
	button3 := widget.NewButton("Time Checkin", func() {
		// =====> do some stuff <==== //
		fmt.Println("I was pressed")
		//TODO: check if logfile exists || create it
		myconfigDir, err := dirExists(usrConfigDir)
		if err != nil {
			fmt.Println("error with dir: ", err)
		}
		if myconfigDir == false {
			fmt.Println("no dir exists,creating dir: ", usrConfigDir)

			if err := os.MkdirAll(usrConfigDir, os.ModePerm); err != nil {
				fmt.Println("error ", err)
			}
		}

		//TODO: write date stamp to file
		// factText.SetText(string(output))
	})

	title := canvas.NewText("System Utils", color.White)
	title.TextStyle = fyne.TextStyle{
		Bold: true,
	}
	title.Alignment = fyne.TextAlignCenter
	title.TextSize = 24

	hBox := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), button, layout.NewSpacer())
	hBox2 := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), button2, layout.NewSpacer())
	hBox3 := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), button3, layout.NewSpacer())
	vBox := container.New(layout.NewVBoxLayout(), title, hBox, hBox2, hBox3, widget.NewSeparator(), factText)

	w.SetContent(vBox)
	w.Show()
	a.Run()
}

func dirExists(path string) (bool, error) {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false, nil
	}
	if err != err {
		return false, err
	}
	return info.IsDir(), nil
}
