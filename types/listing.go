package types

type ListingLatest CoinLatest

type Data struct {
	Data []*CoinLatest `json:"data"`
}
