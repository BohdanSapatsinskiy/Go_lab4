package main

/*
double getPrice(int chambers, int material) {

    if (material == 1) {
        if (chambers == 1) return 2.5;
        return 3;
    }

    if (material == 2) {
        if (chambers == 1) return 0.5;
        return 1;
    }

    if (material == 3) {
        if (chambers == 1) return 1.5;
        return 2;
    }

    return 0;
}

double calculate(double width, double height, double price) {
    return width * height * price;
}

double addWindowsill(double cost, int has) {
    if (has == 1) return cost + 350;
    return cost;
}
*/
import "C"

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func getPrice(chambers, material string) float64 {

	var c int
	if chambers == "Однокамерний" {
		c = 1
	} else {
		c = 2
	}

	var m int
	switch material {
	case "Дерево":
		m = 1
	case "Метал":
		m = 2
	case "Металопластиковий":
		m = 3
	}

	return float64(C.getPrice(C.int(c), C.int(m)))
}

func calculate(width, height, price float64) float64 {
	return float64(C.calculate(
		C.double(width),
		C.double(height),
		C.double(price),
	))
}

func addWindowsill(cost float64, has bool) float64 {
	v := 0
	if has {
		v = 1
	}
	return float64(C.addWindowsill(C.double(cost), C.int(v)))
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