package main_test

import (
	"log"
	"reflect"
	"testing"

	hourglass "github.com/devpablocristo/golang-examples/interview-code/22-hourglass"
)

func TestSumHourglass(t *testing.T) {

	//a := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	h := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	got := hourglass.SumHourglass(h)
	want := 35

	if want != got {
		t.Fatalf("ERROR: want: %d - got: %d", want, got)
	}

	log.Printf("TEST PASSED: want: %d - got: %d", want, got)

}

func TestGetHourglass(t *testing.T) {

	//a := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	a := [6][6]int{
		{1, 1, 1, 2, 2, 2},
		{7, 7, 7, 1, 2, 3},
		{9, 9, 9, 5, 6, 7},
		{3, 3, 3, 4, 4, 1},
		{0, 9, 8, 4, 1, 4},
		{5, 0, 7, 1, 4, 4},
	}

	var ah [16]hourglass.Hourglass

	ah[0].Hg = [3][3]int{
		{1, 1, 1},
		{7, 7, 7},
		{9, 9, 9},
	}

	ah[12].Hg = [3][3]int{
		{2, 2, 2},
		{1, 2, 3},
		{5, 6, 7},
	}

	ah[3].Hg = [3][3]int{
		{3, 3, 3},
		{0, 9, 8},
		{5, 0, 7},
	}

	ah[15].Hg = [3][3]int{
		{4, 4, 1},
		{4, 1, 4},
		{1, 4, 4},
	}

	got := hourglass.GetHourglass(a)
	want := ah

	if want[0] != got[0] {
		t.Fatalf("ERROR 0: want: %d - got: %d", want[0], got[0])
	} else if want[12] != got[12] {
		t.Fatalf("ERROR 12: want: %d - got: %d", want[12], got[12])
	} else if want[3] != got[3] {
		t.Fatalf("ERROR 3: want: %d - got: %d", want[3], got[3])
	} else if want[15] != got[15] {
		t.Fatalf("ERROR 3: want: %d - got: %d", want[15], got[15])
	}

	log.Printf("TEST PASSED: want: %d - got: %d", want[0], got[0])
	log.Printf("TEST PASSED: want: %d - got: %d", want[1], got[1])
	log.Printf("TEST PASSED: want: %d - got: %d", want[2], got[2])
	log.Printf("TEST PASSED: want: %d - got: %d", want[3], got[3])

}

func TestHighestHourglassSum(t *testing.T) {

	a := [6][6]int{
		{0, 0, 0, 0, 0, 0},
		{0, 8, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
	}

	b := [6][6]int{
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 65, 0, 0},
		{1, 2, 3, 0, 0, 0},
		{0, 4, 0, 0, 0, 0},
		{5, 6, 7, 0, 0, 0},
	}

	c := [6][6]int{
		{1, 1, 1, 0, 0, 0},
		{0, 1, 0, 0, 0, 0},
		{1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 23, 0},
		{0, 0, 0, 0, 0, 0},
	}

	gotA := hourglass.HighestHourglassSum(a)
	wantA := 8

	gotB := hourglass.HighestHourglassSum(b)
	wantB := 72

	gotC := hourglass.HighestHourglassSum(c)
	wantC := 24

	if wantA != gotA {
		t.Fatalf("ERROR A: want: %d - got: %d", wantA, gotA)
	} else if wantB != gotB {
		t.Fatalf("ERROR B: want: %d - got: %d", wantB, gotB)
	} else if wantC != gotC {
		t.Fatalf("ERROR C: want: %d - got: %d", wantC, gotC)
	}

	log.Printf("TEST PASSED: want: %d - got: %d", wantA, gotA)
	log.Printf("TEST PASSED: want: %d - got: %d", wantB, gotB)
	log.Printf("TEST PASSED: want: %d - got: %d", wantC, gotC)

}

func TestSumHourglass2(t *testing.T) {

	h := [][]int32{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	got := hourglass.SumHourglass2(h)
	want := int32(35)

	if want != got {
		t.Fatalf("ERROR: want: %d - got: %d", want, got)
	}

	log.Printf("2.TEST PASSED: want: %d - got: %d", want, got)

}

func TestGetHourglass2(t *testing.T) {

	a := [][]int32{
		{1, 1, 1, 2, 2, 2},
		{7, 7, 7, 1, 2, 3},
		{9, 9, 9, 5, 6, 7},
		{3, 3, 3, 4, 4, 1},
		{0, 9, 8, 4, 1, 4},
		{5, 0, 7, 1, 4, 4},
	}

	var ah [16]hourglass.Hourglass2

	ah[0].Hg = [][]int32{
		{1, 1, 1},
		{7, 7, 7},
		{9, 9, 9},
	}

	ah[12].Hg = [][]int32{
		{2, 2, 2},
		{1, 2, 3},
		{5, 6, 7},
	}

	ah[3].Hg = [][]int32{
		{3, 3, 3},
		{0, 9, 8},
		{5, 0, 7},
	}

	ah[15].Hg = [][]int32{
		{4, 4, 1},
		{4, 1, 4},
		{1, 4, 4},
	}

	got := hourglass.GetHourglass2(a)
	want := ah

	if reflect.DeepEqual(want[0], got[0]) {
		t.Fatalf("ERROR 0: want: %d - got: %d", want[0], got[0])
	} else if reflect.DeepEqual(want[12], got[12]) {
		t.Fatalf("ERROR 12: want: %d - got: %d", want[12], got[12])
	} else if reflect.DeepEqual(want[3], got[3]) {
		t.Fatalf("ERROR 3: want: %d - got: %d", want[3], got[3])
	} else if reflect.DeepEqual(want[15], got[15]) {
		t.Fatalf("ERROR 3: want: %d - got: %d", want[15], got[15])
	}

	log.Printf("TEST PASSED: want: %d - got: %d", want[0], got[0])
	log.Printf("TEST PASSED: want: %d - got: %d", want[1], got[1])
	log.Printf("TEST PASSED: want: %d - got: %d", want[2], got[2])
	log.Printf("TEST PASSED: want: %d - got: %d", want[3], got[3])

}
