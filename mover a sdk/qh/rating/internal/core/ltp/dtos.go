package ltp

import "strconv"

type LTPdto struct {
	Pair         string   `json:"pair"` // Currency pair
	Ask          []string `json:"a"`    // Ask price
	Bid          []string `json:"b"`    // Bid price
	Last         []string `json:"c"`    // Last trade price
	Volume       []string `json:"v"`    // Volume
	Weighted     []string `json:"p"`    // Weighted price
	Transactions []int    `json:"t"`    // Number of transactions
	Lowest       []string `json:"l"`    // Lowest price
	Highest      []string `json:"h"`    // Highest price
	Opening      string   `json:"o"`    // Opening price
}

func toDomain(dto LTPdto) *LTP {
	amount, err := strconv.ParseFloat(dto.Last[0], 64)
	if err != nil {
		return nil
	}

	return &LTP{
		Pair:   dto.Pair,
		Amount: amount,
	}
}

func listToDomain(dtos []LTPdto) []LTP {
	var ltps []LTP
	for _, dto := range dtos {
		ltp := toDomain(dto)
		if ltp != nil {
			ltps = append(ltps, *ltp)
		}
	}
	return ltps
}

type KrakenResponse struct {
	Error  []string          `json:"error"`
	Result map[string]LTPdto `json:"result"`
}
