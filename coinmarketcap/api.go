package coinmarketcap

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"sort"

	"github.com/juliosaraiva/crypto-trends/types"
)

const (
	DOMAIN   = "pro-api.coinmarketcap.com"
	PROTOCOL = "https://"
)

const (
	ListingLatestEndpoint   = "/v2/cryptocurrency/quotes/latest"
	CategoriesEndpoint      = "/v1/cryptocurrency/categories"
	TrendingLatestEndpoint  = "/v1/cryptocurrency/trending/latest"
	CategoryEndpoint        = "/v1/cryptocurrency/category"
	QuoteHistoricalEndpoint = "/v2/cryptocurrency/quotes/historical"
)

const (
	ListingLatest int = iota
	Categories
	Category
	TrendingLatest
	ListingsHistory
)

func FilterAITrends(trends *types.Data) []*types.CoinLatest {
	aiKeywords := []string{"ai-big-data", "ai"}
	aiTrends := []*types.CoinLatest{}
	for _, trend := range trends.Data {
		tags := trend.Tags
		for _, keyword := range aiKeywords {
			idx := sort.SearchStrings(tags, keyword)
			if idx < len(trend.Tags) && tags[idx] == keyword {
				aiTrends = append(aiTrends, trend)
				break
			}
		}
	}
	return aiTrends
}

func ListingCrypto() {
	var URL string
	apiKey := os.Getenv("COINMARKETCAP_API_KEY")
	endpoint := flag.Int("endpoint", 0, "Endpoint type")
	time_period := flag.String("time_period", "", "Time for the latest trending coins.")
	flag.Parse()

	switch *endpoint {
	case ListingLatest:
		URL = PROTOCOL + DOMAIN + CategoriesEndpoint
	case Categories:
		URL = PROTOCOL + DOMAIN + ListingLatestEndpoint
	case Category:
		URL = PROTOCOL + DOMAIN + CategoryEndpoint
	case TrendingLatest:
		URL = PROTOCOL + DOMAIN + TrendingLatestEndpoint
	case ListingsHistory:
		URL = PROTOCOL + DOMAIN + QuoteHistoricalEndpoint
	default:
		fmt.Println("Invalid option!")
		os.Exit(1)
	}
	client := &http.Client{}
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := url.Values{}
	if len(*time_period) > 0 && *endpoint == Categories {
		q.Add("time_period", *time_period)
	} else if *endpoint == Category {
		q.Add("id", "6051a81a66fc1b42617d6db7")
		q.Add("limit", "1000")
	} else if *endpoint == ListingLatest {
		q.Add("sort", "percent_change_7d")
		q.Add("sort_dir", "asc")
		q.Add("aux", "num_market_pairs,cmc_rank,date_added,tags,platform,max_supply,circulating_supply,total_supply,market_cap_by_total_supply,volume_24h_reported,volume_7d,volume_7d_reported,volume_30d,volume_30d_reported,is_market_cap_included_in_calc")
	}

	q.Add("start", "1")
	q.Add("limit", "5000")
	q.Add("convert", "USD")

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", apiKey)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil || resp.StatusCode != 200 {
		fmt.Println(string(respBody))
		os.Exit(1)
	}

	coin := types.Data{}
	json.Unmarshal([]byte(respBody), &coin)
	aiTokens := FilterAITrends(&coin)
	result, err := json.Marshal(aiTokens)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(string(result))
}
