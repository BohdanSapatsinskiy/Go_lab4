package main

/*
double getPrice(int country, int season) {

    // season: 1 літо, 2 зима
    // country: 1 Болгарія, 2 Німеччина, 3 Польща

    if (country == 1) {
        if (season == 1) return 100;
        return 150;
    }

    if (country == 2) {
        if (season == 1) return 160;
        return 200;
    }

    if (country == 3) {
        if (season == 1) return 120;
        return 180;
    }

    return 0;
}

double calculate(double price, int days, int vouchers) {
    return price * days * vouchers;
}

double addGuide(double total, int days, int hasGuide) {
    if (hasGuide == 1)
        return total + (50 * days);
    return total;
}

double addLux(double total, int isLux) {
    if (isLux == 1)
        return total * 1.2;
    return total;
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

func getPrice(country, season string) float64 {

	var c int
	switch country {
	case "Болгарія":
		c = 1
	case "Німеччина":
		c = 2
	case "Польща":
		c = 3
	}

	var s int
	if season == "Літо" {
		s = 1
	} else {
		s = 2
	}

	return float64(C.getPrice(C.int(c), C.int(s)))
}
func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
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

		total := float64(C.calculate(C.double(price), C.int(days), C.int(vouchers),))

		total = float64(C.addGuide(C.double(total), C.int(days), C.int(boolToInt(guideCheck.Checked))))
		total = float64(C.addLux(C.double(total), C.int(boolToInt(luxCheck.Checked))))

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