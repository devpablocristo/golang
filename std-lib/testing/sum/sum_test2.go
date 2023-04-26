package sum_test

import (
	"testing"

	sum "github.com/devpablocristo/golang-examples/std-lib/test/sum"
)

func TestSum(t *testing.T) {
	total := sum.Sum(6, 6)
	if total != 12 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 12)
	}
}

func TestSumTable(t *testing.T) {
	tables := []struct {
		x   int
		y   int
		sum int
	}{
		{10, 1, 11},
		{11, 2, 13},
		{12, 3, 15},
		{13, 4, 17},
	}

	for _, table := range tables {
		total := sum.Sum(table.x, table.y)
		if total != table.sum {
			t.Errorf("Sum incorrect (%d+%d), got: %d, expected: %d.", table.x, table.y, total, table.sum)
		}
	}
}
