package handler

import (
	"github.com/devpablocristo/golang/sdk/rating/internal/core/ltp"
)

type responseLTP struct {
	Pair   string  `json:"pair"`
	Amount float64 `json:"amount"`
}

func ToResponseLTP(l *ltp.LTP) *responseLTP {
	return &responseLTP{
		Pair:   l.Pair,
		Amount: l.Amount,
	}
}

func ToResponseLTPList(ltps []ltp.LTP) map[string][]responseLTP {
	ltpList := make([]responseLTP, len(ltps))
	for i, ltp := range ltps {
		ltpList[i] = *ToResponseLTP(&ltp)
	}
	return map[string][]responseLTP{"ltp": ltpList}
}
