package math

type Table struct {
	Base   int
	Result int
}

type Random interface {
	GenerateInt() int
}

func MultiplyTablesRandom(in <-chan Table, gen Random) <-chan Table {
	out := make(chan Table)

	go func() {
		for t := range in {
			t.Result = t.Base * gen.GenerateInt()
			out <- t
		}
		close(out)
	}()

	return out
}

func MultiplyTables(in <-chan Table, num int) <-chan Table {
	out := make(chan Table)

	go func() {
		for t := range in {
			t.Result = t.Base * num
			out <- t
		}
		close(out)
	}()

	return out
}
