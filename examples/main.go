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
	var headers map[string][]string = map[string][]string{}
	cryptocurrencies, _ := coinmarketcap.GetCrypto(ctx, query, headers)
	var symbols string

	for _, c := range cryptocurrencies {
		if len(symbols) == 0 {
			symbols = c.Symbol
			continue
		}

		symbols = symbols + "," + c.Symbol
	}

	q := url.Values{}

	q.Add("symbol", symbols)

	listing, _ := coinmarketcap.GetListing(ctx, q, headers)
	historical, _ := coinmarketcap.GetHistory(ctx, q, headers)

	sliceSymbols := strings.Split(symbols, ",")

	var crypto *model.Crypto = &model.Crypto{}
	var sliceCrypto []*model.Crypto

	for _, v := range sliceSymbols {
		jsonCrypto, _ := json.Marshal(listing[v][0])
		json.Unmarshal([]byte(jsonCrypto), crypto)
		if historical[v][0].Quotes != nil {
			crypto.Historical = historical[v][0].Quotes
		}
		sliceCrypto = append(sliceCrypto, crypto)
		jsonResult, _ := json.Marshal(crypto)
		fmt.Println(string(jsonResult))
	}

	// 	query.Add("symbol", symbol)
	// 	listing, _ := coinmarketcap.GetLatest(ctx, query, headers)
	// 	historical, _ := coinmarketcap.GetHistory(ctx, query, headers)
	// }
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	// fmt.Println(string(JSONOutput))
}
