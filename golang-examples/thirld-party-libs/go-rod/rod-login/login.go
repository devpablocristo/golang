package main

import (
	"fmt"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
)

func main() {
	//leetcode()

	mparg()
}

func mparg() {
	const username = "TESTDSYNLVWN"
	const password = "0KpyFOusH1"

	browser := rod.New().MustConnect()
	defer browser.MustClose()

	page := browser.MustPage("https://www.mercadopago.com.ar/").MustWaitLoad()

	page.MustElement(".nav-header-guest__link").MustClick()
	i, _ := page.Info()
	fmt.Println(i.URL)

	text := page.MustElement(".andes-typography.andes-typography--type-title").MustText()
	fmt.Println(text)

	page.MustElement("input").MustInput(username).MustType(input.Enter)
	text = page.MustElement(".andes-form-control--textfield").MustText()
	fmt.Println(text)

	text = page.MustElement(".center-card__title").MustText()
	fmt.Println(text)

	page.MustElement(".andes-form-control--textfield").MustInput(password).MustType(input.Enter)
	text = page.MustElement(".card-header__title").MustText()
	fmt.Println(text)

	page.MustElementR("span.mp-s-link__text", "Informes").MustClick()
	text = page.MustElement(".report-card__title").MustText()
	fmt.Println(text)

	page.MustElement(".report-card__title").MustClick()
	text = page.MustElement(".andes-button--loud").MustText()
	fmt.Println(text)
}

func leetcode() {
	const username = ""
	const password = ""

	browser := rod.New().MustConnect()

	page := browser.MustPage("https://leetcode.com/accounts/login/")

	page.MustElement("#id_login").MustInput(username)
	page.MustElement("#id_password").MustInput(password).MustType(input.Enter)

	// It will keep retrying until one selector has found a match
	elm := page.Race().Element(".nav-user-icon-base").MustHandle(func(e *rod.Element) {
		// print the username after successful login
		fmt.Println(*e.MustAttribute("title"))
	}).Element("[data-cy=sign-in-error]").MustDo()

	if elm.MustMatches("[data-cy=sign-in-error]") {
		// when wrong username or password
		panic(elm.MustText())
	}
}
