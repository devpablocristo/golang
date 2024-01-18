// How can you sort a slice of custom structs with the help of an example?
// We can sort slices of custom structs by using sort.Sort and sort.Stable functions. These methods sort any collection that implements sort.Interface interface that has Len(), Less() and Swap() methods as shown in the code below:

package main

import (
	"fmt"
	"sort"
)

type Interface interface {
	// Find number of elements in collection
	Len() int

	// Less method is used for identifying which elements among index i and j are lesser and is used for sorting
	Less(i, j int) bool

	// Swap method is used for swapping elements with indexes i and j
	Swap(i, j int)
}

// Consider an example of a Human Struct having name and age attributes.

type Human struct {
	name string
	age  int
}

// Also, consider we have a slice of struct Human of type AgeFactor that needs to be sorted based on age. The AgeFactor implements the methods of the sort.Interface. Then we can call sort.Sort() method on the audience as shown in the below code:

// AgeFactor implements sort.Interface that sorts the slice based on age field.
type AgeFactor []Human

func (a AgeFactor) Len() int           { return len(a) }
func (a AgeFactor) Less(i, j int) bool { return a[i].age < a[j].age }
func (a AgeFactor) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {
	audience := []Human{
		{"Alice", 35},
		{"Bob", 45},
		{"James", 25},
	}
	sort.Sort(AgeFactor(audience))
	fmt.Println(audience)
}

// This code would output:

// [{James 25} {Alice 35} {Bob 45}]
