package count_test

import (
	"testing"

	cnt "github.com/devpablocristo/golang-examples/std-lib/concurrency/01-goroutines/05-add/count"
)

func BenchmarkAdd(b *testing.B) {
	numbers := cnt.GenerateNumbers(1e7)
	for i := 0; i < b.N; i++ {
		cnt.Add(numbers)
	}
}

func BenchmarkAddConcurrent(b *testing.B) {
	numbers := cnt.GenerateNumbers(1e7)
	for i := 0; i < b.N; i++ {
		cnt.AddConcurrent(numbers)
	}
}
