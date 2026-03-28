package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func getPrice(chambers, material string) float64 {
	switch material {
	case "Дерево":
		if chambers == "Однокамерний" {
			return 2.5
		}
		return 3
	case "Метал":
		if chambers == "Однокамерний" {
			return 0.5
		}
		return 1
	case "Металопластиковий":
		if chambers == "Однокамерний" {
			return 1.5
		}
		return 2
	}
	return 0
}

func main() {

	a := app.New()
	w := a.NewWindow("Калькулятор склопакета")
	w.Resize(fyne.NewSize(420, 300))

	widthEntry := widget.NewEntry()
	widthEntry.SetText("0")

	heightEntry := widget.NewEntry()
	heightEntry.SetText("0")

	materialSelect := widget.NewSelect(
		[]string{"Дерево", "Метал", "Металопластиковий"},
		nil,
	)
	materialSelect.SetSelected("Дерево")

	chamberSelect := widget.NewSelect(
		[]string{"Однокамерний", "Двокамерний"},
		nil,
	)
	chamberSelect.SetSelected("Однокамерний")

	windowsill := widget.NewCheck("Підвіконня (+350 грн)", nil)

	result := widget.NewLabel("0.00 грн")
	result.TextStyle = fyne.TextStyle{Bold: true}

	calc := widget.NewButton("Розрахувати", nil)

	calc.OnTapped = func() {
		width, _ := strconv.ParseFloat(widthEntry.Text, 64)
		height, _ := strconv.ParseFloat(heightEntry.Text, 64)

		price := getPrice(chamberSelect.Selected, materialSelect.Selected)

		total := width * height * price

		if windowsill.Checked {
			total += 350
		}

		result.SetText(fmt.Sprintf("%.2f грн", total))
	}

	form := widget.NewForm(
		widget.NewFormItem("Ширина (см)", widthEntry),
		widget.NewFormItem("Висота (см)", heightEntry),
		widget.NewFormItem("Матеріал", materialSelect),
		widget.NewFormItem("Тип склопакета", chamberSelect),
	)

	content := container.NewVBox(
		form,
		windowsill,
		widget.NewSeparator(),
		container.NewBorder(
			nil,
			nil,
			result,
			container.NewMax(calc),
		),
	)

	w.SetContent(container.NewPadded(content))
	w.ShowAndRun()
}