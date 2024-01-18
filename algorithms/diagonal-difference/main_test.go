package diagonaldifference_test

import (
	"testing"

	dd "github.com/devpablocristo/golang-examples/interview-code/27-diagonal-difference"
)

func TestDiagonalDifference(t *testing.T) {
	t.Run("TestDiagonalDifference", func(t *testing.T) {
		t.Log("TestDiagonalDifference")
		want := int32(52)

		got := dd.DiagonalDifference([][]int32{
			{6, 6, 7, -10, 9, -3, 8, 9, -1},
			{9, 7, -10, 6, 4, 1, 6, 1, 1},
			{-1, -2, 4, -6, 1, -4, -6, 3, 9},
			{-8, 7, 6, -1, -6, -6, 6, -7, 2},
			{-10, -4, 9, 1, -7, 8, -5, 3, -5},
			{-8, -3, -4, 2, -3, 7, -5, 1, -5},
			{-2, -7, -4, 8, 3, -1, 8, 2, 3},
			{-3, 4, 6, -7, -7, -8, -3, 9, -6},
			{-2, 0, 5, 4, 4, 4, -3, 3, 0}})

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
