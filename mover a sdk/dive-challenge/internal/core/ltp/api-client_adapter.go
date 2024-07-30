package ltp

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type ApiClient struct {
	BaseURL string
	Client  *http.Client
}

func NewExternalAPI(baseURL string) APIClientPort {
	return &ApiClient{
		BaseURL: baseURL,
		Client:  &http.Client{Timeout: 10 * time.Second},
	}
}

func (a *ApiClient) GetKrakenLTPs(ctx context.Context, pairs []string) ([]LTP, time.Time, error) {
	var lastTradedPrices []LTPdto
	var latestTime time.Time

	for _, pair := range pairs {
		url := fmt.Sprintf("%s/0/public/Ticker?pair=%s", a.BaseURL, pair)
		resp, err := a.Client.Get(url)
		if err != nil {
			return nil, time.Time{}, err
		}
		defer resp.Body.Close()

		var tickerResp KrakenResponse
		if err := json.NewDecoder(resp.Body).Decode(&tickerResp); err != nil {
			return nil, time.Time{}, err
		}

		if len(tickerResp.Error) > 0 {
			return nil, time.Time{}, fmt.Errorf("api error: %v", tickerResp.Error)
		}

		currentTime := time.Now()
		if currentTime.After(latestTime) {
			latestTime = currentTime
		}

		quote, exists := tickerResp.Result[pair]
		if !exists {
			return nil, time.Time{}, fmt.Errorf("pair %s not found in response", pair)
		}

		quote.Pair = pair
		lastTradedPrices = append(lastTradedPrices, quote)
	}

	return listToDomain(lastTradedPrices), latestTime, nil
}
