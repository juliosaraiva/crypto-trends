package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/google/generative-ai-go/genai"
	"github.com/juliosaraiva/crypto-trends/coinmarketcap"
	"github.com/juliosaraiva/crypto-trends/coinmarketcap/model"
	"google.golang.org/api/option"
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

	queryListing := url.Values{}
	queryHistorical := url.Values{}

	queryListing.Add("symbol", symbolList)
	queryHistorical.Add("symbol", symbolList)

	now := time.Now()
	ninetyDaysAgo := now.AddDate(0, 0, -30)
	unixTimestamp := ninetyDaysAgo.Unix()

	timestampString := strconv.FormatInt(unixTimestamp, 10)

	queryHistorical.Add("time_start", timestampString)

	listing, _ := coinmarketcap.GetCrypto(ctx, queryListing, reqHeaders)
	historical, _ := coinmarketcap.GetHistoricalPrices(ctx, queryHistorical, reqHeaders)

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

	JSONMarshal, _ := json.Marshal(cryptoList)

	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("API_KEY")))
	if err != nil {
		log.Fatal(err.Error())
	}

	defer client.Close()
	model := client.GenerativeModel("gemini-pro")
	resp, err := model.GenerateContent(ctx, genai.Text("Based on the given JSON, I want you to do some analysis to identify trends in the cryptos listed in this JSON"+string(JSONMarshal)))
	if err != nil {
		fmt.Println(err.Error())
	}

	JSONResult, _ := json.Marshal(resp)

	fmt.Println(string(JSONResult))
}
