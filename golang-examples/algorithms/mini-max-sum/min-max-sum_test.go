package minimaxsum_test

import (
	"testing"

	mms "github.com/devpablocristo/golang-examples/interview-code/28-mini-max-sum"
)

func TestMinMaxSum(t *testing.T) {
	t.Run("MinMaxSum1", func(t *testing.T) {

		si := []int{1, 2, 3, 4, 5}

		minWant := 10
		maxWant := 14

		minGot, maxGot := mms.MinMaxSum(si)

		if minWant != minGot {
			t.Errorf("Min: want:%d - got: %d", minWant, minGot)
		}

		if maxWant != maxGot {
			t.Errorf("Min: want:%d - got: %d", maxWant, maxGot)
		}
	})

	t.Run("MinMaxSum2", func(t *testing.T) {

		si := []int{256741038, 623958417, 467905213, 714532089, 938071625}

		minWant := 2063136757
		maxWant := 2744467344

		minGot, maxGot := mms.MinMaxSum(si)

		if minWant != minGot {
			t.Errorf("Min: want:%d - got: %d", minWant, minGot)
		}

		if maxWant != maxGot {
			t.Errorf("Min: want:%d - got: %d", maxWant, maxGot)
		}
	})

	t.Run("MinMaxSum3", func(t *testing.T) {

		si := []int{769082435, 210437958, 673982045, 375809214, 380564127}

		t.Logf("%d", si)

		minWant := 1640793344
		maxWant := 2199437821

		minGot, maxGot := mms.MinMaxSum(si)

		if minWant != minGot {
			t.Errorf("Min: want:%d - got: %d", minWant, minGot)
		}

		if maxWant != maxGot {
			t.Errorf("Min: want:%d - got: %d", maxWant, maxGot)
		}
	})

}
