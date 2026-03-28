package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func getPrice(country, season string) float64 {

	switch country {

	case "Болгарія":
		if season == "Літо" {
			return 100
		}
		return 150

	case "Німеччина":
		if season == "Літо" {
			return 160
		}
		return 200

	case "Польща":
		if season == "Літо" {
			return 120
		}
		return 180
	}

	return 0
}

func calculate(price float64, days int, vouchers int) float64 {
	return price * float64(days) * float64(vouchers)
}

func addGuide(total float64, days int, has bool) float64 {
	if has {
		return total + float64(days*50)
	}
	return total
}

func addLux(total float64, isLux bool) float64 {
	if isLux {
		return total * 1.2
	}
	return total
}

func main() {

	a := app.New()
	w := a.NewWindow("Калькулятор туру")
	w.Resize(fyne.NewSize(420, 320))

	countrySelect := widget.NewSelect(
		[]string{"Болгарія", "Німеччина", "Польща"},
		nil,
	)
	countrySelect.SetSelected("Болгарія")

	seasonSelect := widget.NewSelect(
		[]string{"Літо", "Зима"},
		nil,
	)
	seasonSelect.SetSelected("Літо")

	daysEntry := widget.NewEntry()
	daysEntry.SetText("7")

	vouchersEntry := widget.NewEntry()
	vouchersEntry.SetText("2")

	guideCheck := widget.NewCheck("Індивідуальний гід ($50/день)", nil)
	luxCheck := widget.NewCheck("Номер люкс (+20%)", nil)

	result := widget.NewLabel("$0")
	result.TextStyle = fyne.TextStyle{Bold: true}

	button := widget.NewButton("Розрахувати", func() {

		days, _ := strconv.Atoi(daysEntry.Text)
		vouchers, _ := strconv.Atoi(vouchersEntry.Text)

		price := getPrice(countrySelect.Selected, seasonSelect.Selected)

		total := calculate(price, days, vouchers)

		total = addGuide(total, days, guideCheck.Checked)
		total = addLux(total, luxCheck.Checked)

		result.SetText(fmt.Sprintf("$%.2f", total))
	})

	form := widget.NewForm(
		widget.NewFormItem("Країна", countrySelect),
		widget.NewFormItem("Сезон", seasonSelect),
		widget.NewFormItem("Днів", daysEntry),
		widget.NewFormItem("Путівок", vouchersEntry),
	)

	content := container.NewVBox(
		form,
		guideCheck,
		luxCheck,
		widget.NewSeparator(),
		container.NewBorder(
			nil,
			nil,
			result,
			container.NewMax(button),
		),
	)

	w.SetContent(container.NewPadded(content))
	w.ShowAndRun()
}