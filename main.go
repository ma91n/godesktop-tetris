package main

import (
	_ "embed"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

//go:embed リス.png
var iconData []byte

func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	w.SetIcon(&fyne.StaticResource{
		StaticName:    "リス",
		StaticContent: iconData,
	})

	w.SetContent(container.NewVBox(
		widget.NewLabel("Hello Fyne!"),
		widget.NewButton("Hi!", func() {
			widget.NewLabel("Hello Fyne!").SetText("Welcome :)")
		}),
	))

	w.ShowAndRun()
}
