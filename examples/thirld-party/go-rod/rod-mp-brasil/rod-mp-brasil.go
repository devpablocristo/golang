package main

import (
	"fmt"

	"github.com/go-rod/rod"
)

func main() {
	const usr = "TETE7774186"
	const psw = "tlhuu2/9=grc3rdl"

	browser := rod.New().MustConnect()
	browser1 := browser.MustIncognito()
	defer browser1.MustClose()
	page := browser1.MustPage("https://www.mercadopago.com.br/").MustWaitLoad()

	page.MustElement(".nav-header-guest__link--login").MustClick()

	// read the css element ".center-card__title" with the text "¡Hola! Ingresa tu teléfono, e‑mail o usuario"
	text := page.MustElement(".center-card__title").MustText()
	fmt.Println(text)

	///////
	//colombia
	//////
	// click on "ingresar" on "https://www.mercadopago.com.co/"
	// page.MustElement(".nav-header-guest__link--login").MustClick()

	// // read the css element ".center-card__title" with the text "¡Hola! Ingresa tu teléfono, e‑mail o usuario"
	// text := page.MustElement(".center-card__title").MustText()
	// fmt.Println(text)

	// page.MustElement("input").MustInput("TETE1421848").MustType(input.Enter)
	// text = page.MustElement(".email-badge__user-name").MustText()
	// fmt.Println(text)

	// text = page.MustElement(".center-card__title").MustText()
	// fmt.Println(text)

	// page.MustElement(".andes-form-control__control").MustInput("1lXXR4WWjg").MustType(input.Enter)
	// text = page.MustElement(".card-header__title").MustText()
	// fmt.Println(text)

	// page.MustElementR("span.mp-s-link__text", "Informes").MustClick()
	// text = page.MustElement(".report-card__title").MustText()
	// fmt.Println(text)

	// page.MustElement(".report-card__title").MustClick()
	// text = page.MustElement(".andes-button--loud").MustText()
	// fmt.Println(text)

	///////
	//////

	// ctx, cancel := context.WithCancel(context.Background())
	// pageWithCancel := page.Context(ctx)

	// go func() {
	// 	time.Sleep(30 * time.Second)
	// 	cancel()
	// }()

	// page.MustElement("a.nav-header-guest__link").MustClick()
	// txt := page.MustElement("div.center-card__header").MustText()
	// fmt.Println(txt)
	// page.MustElement("input").MustInput(usr).MustType(input.Enter)
	// txt = page.MustElement(".center-card__title").MustText()
	// fmt.Println(txt)
	// page.MustElement(".andes-form-control__control").MustInput(psw).MustType(input.Enter)
	// txt = page.MustElement(".card-header__title").MustText()
	// fmt.Println(txt)

	// page.MustElementR(".mp-s-link", "Relatórios e faturamento").MustClick()
	// txt = page.MustElement(".billing-title-access").MustText()
	// fmt.Println(txt)
	// pageWithCancel.MustElementR("a.card-bills-reports-mp-link", "Ir para Relatórios").MustWaitLoad().MustClick()
	// txt = pageWithCancel.MustElement("h1").MustText()
	// fmt.Println(txt)
}
