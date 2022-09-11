package main

import (
	"io/ioutil"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	c "fyne.io/fyne/v2/container"
	d "fyne.io/fyne/v2/dialog"
	s "fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Open File")
	w.Resize(fyne.NewSize(400, 400))

	btn := widget.NewButton("Open .go files", func() {
		fd := d.NewFileOpen(func(r fyne.URIReadCloser, _ error) {
			data, _ := ioutil.ReadAll(r)
			result := fyne.NewStaticResource("name", data)

			entry := widget.NewMultiLineEntry()
			entry.SetText(string(result.StaticContent))

			w := fyne.CurrentApp().NewWindow(string(result.StaticName))
			w.SetContent(c.NewScroll(entry))
			w.Resize(fyne.NewSize(400, 400))
			w.Show()
		}, w)
		fd.SetFilter(s.NewExtensionFileFilter([]string{".go"}))
		fd.Show()
	})

	w.SetContent(c.NewVBox(btn))
	w.ShowAndRun()
}
