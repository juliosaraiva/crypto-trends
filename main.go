package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"

	coinmarketcap "github.com/juliosaraiva/crypto-trends/coinmarketcap/historical"
)

func main() {
	histData := coinmarketcap.CoinHistorical{}
	query := url.Values{}
	resp, err := histData.Get(&query)
	if err != nil {
		log.Fatal(err)
	}
	result, err := json.Marshal(resp)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(result))
}
