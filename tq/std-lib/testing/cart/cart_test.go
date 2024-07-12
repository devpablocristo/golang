package main

import "testing"

func TestSumItems(t *testing.T) {

	items := make(map[string]int)

	items["apple"] = 1
	items["orange"] = 2
	items["pear"] = 3

	total := SumItems(items)

	if total != 6 {
		t.Errorf("Sum incorrect %d + %d + % d, got: %d, expected: %d.", items["apple"], items["orange"], items["pear"], total, 6)
	}
}
