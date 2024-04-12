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
	output, err := Gemini(cryptoList)
	if err != nil {
		fmt.Println(err.Error())
	}

	JSONOutput, _ := json.Marshal(output)

	fmt.Println(string(JSONOutput))
}

func Gemini(cryptoList []model.Cryptocurrency) (*genai.GenerateContentResponse, error) {
	ctx := context.TODO()
	cryptoMarshal, _ := json.Marshal(cryptoList)
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		log.Fatal(err.Error())
	}

	defer client.Close()
	aiModel := client.GenerativeModel("gemini-pro")
	resp, err := aiModel.GenerateContent(ctx, genai.Text("Based on the given JSON, I want you to do some analysis to identify trends in the cryptos listed in this JSON. Use the historical_data to identify high volumes and possible bullish trends in the bitcoin price."+string(cryptoMarshal)))
	if err != nil {
		fmt.Println(err.Error())
	}

	return resp, nil
}
