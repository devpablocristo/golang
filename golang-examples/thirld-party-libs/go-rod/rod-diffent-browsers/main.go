package main

import (
	"fmt"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
	"github.com/go-rod/rod/lib/launcher"
)

func main() {

	chrome := launcher.New().Bin("/usr/bin/google-chrome").MustLaunch()
	fmt.Println(chrome)
	browser := rod.New().ControlURL(chrome).MustConnect().MustPage("https://example.com")

	// firefox := launcher.New().Bin("/usr/bin/firefox").MustLaunch()
	// fmt.Println(firefox)
	// browser2 := rod.New().ControlURL(firefox).MustConnect().MustPage("https://example.com")

	// page2 := browser2.Browser().MustPage("https://github.com")
	// page2.MustElement("input").MustInput("git").MustType(input.Enter)

	//browser.Browser()
	//browser := rod.New().MustConnect()

	// Even you forget to close, rod will close it after main process ends.
	// defer browser.MustClose()

	// // Create a new page
	page := browser.Browser().MustPage("https://github.com")

	// // We use css selector to get the search input element and input "git"
	page.MustElement("input").MustInput("git").MustType(input.Enter)

	// Wait until css selector get the element then get the text content of it.
	text := page.MustElement(".codesearch-results p").MustText()

	fmt.Println(text)

	// Get all input elements. Rod supports query elements by css selector, xpath, and regex.
	// For more detailed usage, check the query_test.go file.
	fmt.Println("Found", len(page.MustElements("input")), "input elements")

	// Eval js on the page
	page.MustEval(`() => console.log("hello world")`)

	// Pass parameters as json objects to the js function. This MustEval will result 3
	fmt.Println("1 + 2 =", page.MustEval(`(a, b) => a + b`, 1, 2).Int())

	// When eval on an element, "this" in the js is the current DOM element.
	fmt.Println(page.MustElement("title").MustEval(`() => this.innerText`).String())

	// Output:
	// Git is the most widely used version control system.
	// Found 5 input elements
	// 1 + 2 = 3
	// Search · git · GitHub

	// h, _ := page.HTML()
	// fmt.Println(h)

	i, _ := page.Info()

	//ss, _:= page.Screenshot()

	page.MustWaitLoad().MustScreenshot("a.png")

	fmt.Println(i.URL)
}
