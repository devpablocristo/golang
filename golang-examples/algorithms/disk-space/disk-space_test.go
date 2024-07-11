package main_test

import (
	"fmt"
	"log"
	"reflect"
	"testing"

	diskSpace "github.com/devpablocristo/golang-examples/interview-code/23-disk-space"
)

func TestSegment(t *testing.T) {
	x := int32(2)
	space := []int32{8, 2, 4, 6}

	maxWant := int32(3)
	minWant := []int32{2, 2, 4}

	maxGot, minGot := diskSpace.Segment(x, space)

	fmt.Println(minGot)

	if !reflect.DeepEqual(minGot, minWant) {
		t.Fatalf("Incorrect min: want: %d - got: %d", minWant, minGot)
	}

	if maxGot != maxWant {
		t.Fatalf("Incorrect max: want: %d - git: %d", maxWant, maxGot)
	}

	// if minGot != minWant {
	// 	t.Fatalf("Incorrect min: want: %d - git: %d", minWant, minGot)
	// }

}

func TestSliceToMap(t *testing.T) {

	x := 2
	space := []int{8, 2, 4, 6}

	want := map[int][]int{0: {8, 2, 4, 6}, 1: {2, 4, 6}, 2: {4, 6}, 3: {6}}
	got := diskSpace.SliceToMap(x, space)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("Incorrect min: want: %d - got: %d", want, got)
	}

}

func TestRemoveItemsFromSLice(t *testing.T) {
	s := make([]int, 5)
	s = []int{2, 3, 5, 7, 11, 13}

	want := []int{13}
	got := diskSpace.RemoveItemsFromSlice(s)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("Incorrect min: want: %d - got: %d", want, got)
	}

}

func TestChunkSlice(t *testing.T) {
	x := 2
	space := []int{8, 2, 4, 6}
	want := [][]int{{8, 2}, {4, 6}}

	got := diskSpace.ChunkSlice(space, x)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("Incorrect: want: %d - got: %d", want, got)
	}

	log.Printf("TEST PASSED!: want: %d - got: %d", want, got)
}
