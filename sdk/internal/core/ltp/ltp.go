package ltp

type LTP struct {
	Pair   string
	Amount float64
}

type InMemDB map[string]*LTP
