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
	query.Add("symbol", "BTC,GPU,CKB")
	query.Add("sort", "cmc_rank")
	var reqHeaders map[string][]string = map[string][]string{}
	cryptocurrencies, _ := coinmarketcap.ListCryptocurrencies(ctx, query, reqHeaders)
	var IDList string

	for _, crypto := range cryptocurrencies {
		if len(IDList) == 0 {
			IDList = strconv.Itoa(crypto.ID)
			continue
		}

		IDList = IDList + "," + strconv.Itoa(crypto.ID)
	}

	queryListing := url.Values{}
	queryHistorical := url.Values{}

	queryListing.Add("id", IDList)
	queryHistorical.Add("id", IDList)

	now := time.Now()
	FifiteenDaysAgoUnixTimestamp := now.AddDate(0, 0, -1).Unix()
	// TodayUnixTimestamp := now.Unix()

	FifiteenDaysTimestampString := strconv.FormatInt(FifiteenDaysAgoUnixTimestamp, 10)
	// TodayTimestampString := strconv.FormatInt(TodayUnixTimestamp, 10)

	queryHistorical.Add("time_start", FifiteenDaysTimestampString)
	// queryHistorical.Add("time_end", TodayTimestampString)
	queryHistorical.Add("interval", "1h")

	listing, _ := coinmarketcap.GetCrypto(ctx, queryListing, reqHeaders)
	historical, _ := coinmarketcap.GetOHLCVHistoricalPrices(ctx, queryHistorical, reqHeaders)

	sliceIDs := strings.Split(IDList, ",")

	var cryptocurrency model.Cryptocurrency = model.Cryptocurrency{}
	var cryptoList []model.Cryptocurrency

	for _, id := range sliceIDs {
		parseJSON, _ := json.Marshal(listing[id])

		json.Unmarshal([]byte(parseJSON), &cryptocurrency)
		cryptocurrency.Historical = historical.Quotes
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
	textPrompt := `
	You are a technical analyst investing in crypto, and needs to investigate a trend to the give historical period to identify the trend based on the volume or any chart pattern for the given JSON below.

	Important: You can identify the trend using RSI (Relative Strength Index) for the given historical volume, and for the chart analyze following the principles of price action.

	After your analysis of the trend, the output must be a JSON with the following:

	- An array of json with the following contents:
	  - the name of the crypto
	  - the symbol of the crypto
	  - the ID of the crypto based on the given JSON below
	  - and I want a json key called *trend* with one of the following values: high, sideway, or low based on your analysis for the given period wether the trend is high, sideway or low.

	I am going to use your output to into my application,
	so try to be more specific as possible and just give me the information I need (JSON Content).
	Here is an example of the schema for the output expected:

	{"crypto":  list[CRYPTO]}
	CRYPTO = {"id": int, "name": str, "symbol": str, "trend": str}

	All fields are required.

	Here is the JSON content you need to analyze: `
	resp, err := aiModel.GenerateContent(ctx, genai.Text(textPrompt+string(cryptoMarshal)))
	if err != nil {
		fmt.Println(err.Error())
	}

	return resp, nil
}
