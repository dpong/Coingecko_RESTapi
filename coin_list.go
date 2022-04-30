package coingeckoapi

import "net/http"

type CoinListResponse struct {
	ID     string `json:"id"`
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
}

// opt = including platform info inside or not
func (b *Client) CoinList(platform bool) ([]CoinListResponse, error) {
	type opt struct {
		Platform bool `url:"include_platform"`
	}
	input := opt{
		Platform: platform,
	}
	res, err := b.do("spot", http.MethodGet, "coins/list", input, false, false)
	if err != nil {
		return nil, err
	}
	result := []CoinListResponse{}
	err = json.Unmarshal(res, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
