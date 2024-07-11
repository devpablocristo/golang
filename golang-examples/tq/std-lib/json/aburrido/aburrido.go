package main

import (
	"encoding/json"
	"fmt"
)

type A struct {
	A string
}

type B struct {
	A string
}

type C struct {
	A
	B
}

type Cx struct {
	A A
	B B
}

func main() {
	var c C
	c.A.A = "aa"
	c.B.A = "ba"
	fmt.Printf("%#v\n", c) // => main.C{A:main.A{A:"aa"}, B:main.B{A:"ba"}}

	buf, _ := json.Marshal(c)
	fmt.Printf("%s\n", buf) // => {}

	x := make(map[string]interface{})
	x["A"] = c.A
	x["B"] = c.B
	buf2, _ := json.Marshal(x)
	fmt.Printf("%s\n", buf2) // => {"A":{"A":"aa"},"B":{"A":"ba"}}

	var cx Cx
	cx.A.A = "aa"
	cx.B.A = "ba"
	fmt.Printf("%#v\n", cx) // => main.C{A:main.A{A:"aa"}, B:main.B{A:"ba"}}

	bufx, _ := json.Marshal(cx)
	fmt.Printf("%s\n", bufx) // => {"A":{"A":"aa"},"B":{"A":"ba"}}
}
