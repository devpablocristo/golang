package main

import (
	"fmt"

	"github.com/devpablocristo/golang-examples/std-lib/packages/hello"
	"github.com/devpablocristo/golang-examples/std-lib/packages/rectangle"

	"github.com/sirupsen/logrus"

	"rsc.io/quote"
	qV2 "rsc.io/quote/v2"
	qV3 "rsc.io/quote/v3"
)

func main() {
	fmt.Println(rectangle.Area(2.5, 2.6))
	fmt.Println(hello.Hello())
	logrus.Println("Hola logrus")

	fmt.Println(quote.Glass())
	fmt.Println(qV2.GlassV2())
	fmt.Println(qV3.GlassV3())
}
