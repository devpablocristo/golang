package main

import (
	"fmt"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
)

func main() {
	browser := rod.New().MustConnect()
	// page := browser.MustPage("https://file-examples.com/index.php/sample-documents-download/sample-pdf-download/")
	// wait := browser.MustWaitDownload()
	// page.MustElementR("a", "DOWNLOAD SAMPLE PDF FILE").MustClick()
	// err := utils.OutputFile("t.pdf", wait())
	// if err != nil {
	// 	fmt.Println(err)
	// }

	page := browser.MustPage("https://www.mercadopago.com.co/").MustWaitLoad()
	page.MustElement(".nav-header-guest__link--login").MustClick()
	fmt.Println(page.MustElement(".center-card__title").MustText())

	page.MustElement("input").MustInput("TEST9WESVJI1").MustType(input.Enter)
	fmt.Println(page.MustElement(".email-badge__user-name").MustText())
	fmt.Println(page.MustElement(".center-card__title").MustText())

	page.MustElement(".andes-form-control__control").MustInput("vtxpr`i7!gh5qa1x").MustType(input.Enter)
	fmt.Println(page.MustElement(".card-header__title").MustText())

	page.MustNavigate("https://www.mercadopago.com.co/balance/reports/settlement/")

	page.MustSearch("span.andes-button__content").MustClick().Element("csv")

	page.MustElement("span.andes-button__content")
	list := page.MustElements("span.andes-button__content")
	fmt.Println(list.First().MustText())

	for _, val := range list {
		fmt.Println(val.MustText())
		if val.MustText() == "csv" {
			val.MustClick()
		}
	}
}
