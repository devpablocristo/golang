package main

import (
	"fmt"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
)

func main() {

	// new browser
	browser := rod.New().MustConnect()
	defer browser.MustClose()

	// load page
	page := browser.MustPage("https://www.mercadopago.com.co/").MustWaitLoad()

	page.Browser().GetCookies()

	// click on "ingresar" on "https://www.mercadopago.com.co/"
	page.MustElement(".nav-header-guest__link--login").MustClick()

	// read the css element ".center-card__title" with the text "¡Hola! Ingresa tu teléfono, e‑mail o usuario"
	text := page.MustElement(".center-card__title").MustText()
	fmt.Println(text)

	page.MustElement("input").MustInput("TETE1421848").MustType(input.Enter)
	text = page.MustElement(".email-badge__user-name").MustText()
	fmt.Println(text)

	text = page.MustElement(".center-card__title").MustText()
	fmt.Println(text)

	page.MustElement(".andes-form-control__control").MustInput("1lXXR4WWjg").MustType(input.Enter)
	text = page.MustElement(".card-header__title").MustText()
	fmt.Println(text)

	page.MustElementR("span.mp-s-link__text", "Informes").MustClick()
	text = page.MustElement(".report-card__title").MustText()
	fmt.Println(text)

	page.MustElement(".report-card__title").MustClick()
	text = page.MustElement(".andes-button--loud").MustText()
	fmt.Println(text)

	// page.MustElement(".andes-button--loud").MustClick()
	// text = page.MustElement(".button.andes-button:first-child").MustText()
	// fmt.Println(text)

}
