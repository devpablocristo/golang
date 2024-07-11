package main_test

import (
	"reflect"
	"sort"
	"testing"

	mergeSlice "github.com/devpablocristo/golang-examples/interview-code/26-merge-slices"
)

func TestMergeSlices(t *testing.T) {
	p1 := mergeSlice.Person{Name: "John", Age: 31}
	p2 := mergeSlice.Person{Name: "Mary", Age: 41}
	p3 := mergeSlice.Person{Name: "Tom", Age: 66}
	p4 := mergeSlice.Person{Name: "Anna", Age: 52}
	p5 := mergeSlice.Person{Name: "Paul", Age: 24}
	p6 := mergeSlice.Person{Name: "Jane", Age: 33}

	Persons1 := []mergeSlice.Person{p1, p2, p3}
	Persons2 := []mergeSlice.Person{p4, p5, p6, p1}
	checkResult := []mergeSlice.Person{p1, p2, p3, p4, p5, p6}

	result := mergeSlice.MergeSlices(Persons1, Persons2)

	sort.Slice(checkResult, func(i, j int) bool {
		return checkResult[i].Name < checkResult[j].Name
	})

	sort.Slice(result, func(i, j int) bool {
		return result[i].Name < result[j].Name
	})

	if !reflect.DeepEqual(result, checkResult) {
		t.Errorf("Merge slices: Slices are not equal")
	} else {
		t.Log("Merge slices: Slices are equal")
	}
}

func TestMergePointersSlices(t *testing.T) {

	p1 := mergeSlice.Person{Name: "John", Age: 31}
	p2 := mergeSlice.Person{Name: "Mary", Age: 41}
	p3 := mergeSlice.Person{Name: "Tom", Age: 66}
	p4 := mergeSlice.Person{Name: "Anna", Age: 52}
	p5 := mergeSlice.Person{Name: "Paul", Age: 24}
	p6 := mergeSlice.Person{Name: "Jane", Age: 33}

	Persons1 := []*mergeSlice.Person{&p1, &p2, &p3, &p4, &p1, &p2, &p3, &p5, &p6, &p4, &p5, &p6}
	Persons2 := []*mergeSlice.Person{&p4, &p5, &p6, &p1, &p2, &p3, &p1, &p2, &p3, &p5, &p6, &p4, &p5, &p6}
	checkResult := []*mergeSlice.Person{&p1, &p2, &p3, &p4, &p5, &p6}

	result := mergeSlice.MergePointersSlices(Persons1, Persons2)

	sort.Slice(checkResult, func(i, j int) bool {
		return checkResult[i].Name < checkResult[j].Name
	})

	sort.Slice(result, func(i, j int) bool {
		return result[i].Name < result[j].Name
	})

	if !reflect.DeepEqual(result, checkResult) {
		t.Errorf("Merge pointers slices: Slices are not equal")
	} else {
		t.Log("Merge pointers slices: Slices are equal")
	}
}
