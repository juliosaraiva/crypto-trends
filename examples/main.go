package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/juliosaraiva/crypto-trends/coinmarketcap"
	"github.com/juliosaraiva/crypto-trends/coinmarketcap/model"
)

func main() {
	ctx := context.Background()
	var query url.Values = url.Values{}
	query.Add("start", "1")
	query.Add("limit", "5")
	query.Add("sort", "cmc_rank")
	var reqHeaders map[string][]string = map[string][]string{}
	cryptocurrencies, _ := coinmarketcap.ListCryptocurrencies(ctx, query, reqHeaders)
	var symbolList string

	for _, crypto := range cryptocurrencies {
		if len(symbolList) == 0 {
			symbolList = crypto.Symbol
			continue
		}

		symbolList = symbolList + "," + crypto.Symbol
	}

	q := url.Values{}

	q.Add("symbol", symbolList)

	listing, _ := coinmarketcap.GetCrypto(ctx, q, reqHeaders)
	historical, _ := coinmarketcap.GetHistoricalPrices(ctx, q, reqHeaders)

	sliceSymbols := strings.Split(symbolList, ",")

	var cryptocurrency model.Cryptocurrency = model.Cryptocurrency{}
	var cryptoList []model.Cryptocurrency

	for _, symbol := range sliceSymbols {
		parseJSON, _ := json.Marshal(listing[symbol][0])
		json.Unmarshal([]byte(parseJSON), &cryptocurrency)
		if historical[symbol][0].Quotes != nil {
			cryptocurrency.Historical = historical[symbol][0].Quotes
		}
		cryptoList = append(cryptoList, cryptocurrency)
	}

	parseJSON, _ := json.Marshal(cryptoList)
	fmt.Println(string(parseJSON))
}
