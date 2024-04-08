package coinmarketcap

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/juliosaraiva/crypto-trends/coinmarketcap/model"
	"github.com/juliosaraiva/crypto-trends/config"
)

const (
	URL = "https://pro-api.coinmarketcap.com"
)

var apiKey string = config.Config("COINMARKETCAP_API_KEY")

const (
	ListingLatestEndpoint   = "/v2/cryptocurrency/quotes/latest"
	CategoriesEndpoint      = "/v1/cryptocurrency/categories"
	TrendingLatestEndpoint  = "/v1/cryptocurrency/trending/latest"
	CategoryEndpoint        = "/v1/cryptocurrency/category"
	QuoteHistoricalEndpoint = "/v2/cryptocurrency/quotes/historical"
	CryptoCurrencyEndpoint  = "/v2/cryptocurrency/map"
)

const (
	ListingLatest int = iota
	Categories
	Category
	TrendingLatest
	ListingsHistory
	CryptoCurrency
)

func searchTags(tags []*model.ListingTags, k string) int {
	for i, v := range tags {
		if v.Slug == k {
			return i
		}
	}
	return 0
}

func FilterByTag(trends map[string][]*model.Listing, filter string) ([]*model.Listing, error) {
	keywords := strings.Split(filter, ",")
	filteredTrends := []*model.Listing{}
	for _, v := range trends {
		tags := v[0].Tags
		for _, keyword := range keywords {
			tag := searchTags(tags, keyword)
			if tag > 0 {
				filteredTrends = append(filteredTrends, v[0])
				break
			}
		}
	}
	return filteredTrends, nil
}

func GetLatest(ctx context.Context, query url.Values) ([]*model.Listing, *map[string][]*model.Listing, error) {
	ListingEndpoint := URL + ListingLatestEndpoint

	client := &http.Client{}
	req, err := http.NewRequest("GET", ListingEndpoint, nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := query

	if len(q) == 0 {
		q.Add("aux", "num_market_pairs,cmc_rank,date_added,tags,platform,max_supply,circulating_supply,total_supply,market_cap_by_total_supply,volume_24h_reported,volume_7d,volume_7d_reported,volume_30d,volume_30d_reported,is_market_cap_included_in_calc")
		q.Add("start", "1")
		q.Add("limit", "5000")
		q.Add("convert", "USD")
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Encoding", "deflare,gzip")
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

	defer resp.Body.Close()

	var tokens *model.ListingData
	categories := q.Get("categories")
	if err = json.Unmarshal(respBody, &tokens); err != nil {
		fmt.Println(err)
	}

	// parse, _ := json.Marshal(tokens)
	fmt.Println(string(respBody))

	if len(categories) > 0 {
		filteredTokens, err := FilterByTag(tokens.Data, categories)
		if err != nil {
			return nil, nil, err
		}
		return filteredTokens, nil, nil
	}

	return nil, &tokens.Data, nil
}

func GetHistory(ctx context.Context, query *url.Values) (map[string][]*model.CoinHistorical, error) {
	HistoricalEndpoint := URL + QuoteHistoricalEndpoint
	client := &http.Client{}
	req, err := http.NewRequest("GET", HistoricalEndpoint, nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Encoding", "deflate,gzip")
	req.Header.Add("X-CMC_PRO_API_KEY", apiKey)
	req.URL.RawQuery = query.Encode()

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil || resp.StatusCode != 200 {
		return nil, fmt.Errorf("%s", respBody)
	}

	defer resp.Body.Close()

	hist := model.HistoricalData{}
	err = json.Unmarshal(respBody, &hist)
	if err != nil {
		fmt.Print(err)
	}

	return hist.Data, nil
}

func ListingCrypto(ctx context.Context, query url.Values) ([]*model.Cryptocurrency, error) {
	var cryptocurrencies []*model.Cryptocurrency = []*model.Cryptocurrency{}
	CryptoEndpoint := URL + CryptoCurrencyEndpoint
	client := &http.Client{}
	req, err := http.NewRequest("GET", CryptoEndpoint, nil)
	if err != nil {
		return nil, err
	}

	if len(query) == 0 {
		query.Add("listing_status", "active")
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Encoding", "deflate,gzip")
	req.Header.Add("X-CMC_PRO_API_KEY", apiKey)
	req.URL.RawQuery = query.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	respBody, _ := io.ReadAll(resp.Body)

	if err = json.Unmarshal(respBody, cryptocurrencies); err != nil {
		return nil, err
	}

	return cryptocurrencies, nil
}
