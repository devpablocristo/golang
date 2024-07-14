package main

func main() {

	items := make(map[string]int)

	items["apple"] = 10
	items["orange"] = 5
	items["pear"] = 15

	total := SumItems(items)
	println(total)
}

func SumItems(items map[string]int) int {
	var r int

	for _, v := range items {
		r += v
	}

	return r
}
