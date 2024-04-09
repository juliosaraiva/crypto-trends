package coinmarketcap

import (
	"context"
	"encoding/json"
	"errors"
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

type error interface {
	Error() string
}

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
	CryptoCurrencyEndpoint  = "/v1/cryptocurrency/map"
)

func searchTags(tags []*model.Tags, k string) int {
	for i, v := range tags {
		if v.Slug == k {
			return i
		}
	}
	return 0
}

func FilterByTag(categories map[string][]*model.Listing, filter string) ([]*model.Listing, error) {
	keywords := strings.Split(filter, ",")
	filteredCategories := []*model.Listing{}
	for _, v := range categories {
		tags := v[0].Tags
		for _, keyword := range keywords {
			tag := searchTags(tags, keyword)
			if tag > 0 {
				filteredCategories = append(filteredCategories, v[0])
				break
			}
		}
	}
	return filteredCategories, nil
}

func GetListing(ctx context.Context, query url.Values, headers map[string][]string) (map[string][]*model.Listing, error) {
	ListingEndpoint := URL + ListingLatestEndpoint

	client := &http.Client{}
	req, err := http.NewRequest("GET", ListingEndpoint, nil)
	if err != nil {
		return nil, err
	}

	if len(query) == 0 {
		query.Add("aux", "num_market_pairs,cmc_rank,date_added,tags,platform,max_supply,circulating_supply,total_supply,market_cap_by_total_supply,volume_24h_reported,volume_7d,volume_7d_reported,volume_30d,volume_30d_reported,is_market_cap_included_in_calc")
		query.Add("start", "1")
		query.Add("limit", "5000")
		query.Add("convert", "USD")
	}

	if len(headers) == 0 {
		req.Header.Set("Accept", "application/json")
		req.Header.Set("Accept-Encoding", "'deflare,gzip'")
	}

	for k, v := range headers {
		req.Header.Set(k, v[0])
	}

	req.Header.Add("X-CMC_PRO_API_KEY", apiKey)
	req.URL.RawQuery = query.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	} else if resp.StatusCode >= 400 {
		respError, _ := io.ReadAll(resp.Body)
		return nil, errors.New(string(respError))
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var tokens *model.ListingData
	if err = json.Unmarshal(respBody, &tokens); err != nil {
		return nil, err
	}

	return tokens.Data, nil
}

func GetHistory(ctx context.Context, query url.Values, headers map[string][]string) (map[string][]*model.CoinHistorical, error) {
	HistoricalEndpoint := URL + QuoteHistoricalEndpoint
	client := &http.Client{}
	req, err := http.NewRequest("GET", HistoricalEndpoint, nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	if len(headers) == 0 {
		req.Header.Set("Accept", "application/json")
		req.Header.Set("Accept-Encoding", "'deflate,gzip'")
	}

	req.Header.Add("X-CMC_PRO_API_KEY", apiKey)

	for k, v := range headers {
		req.Header.Set(k, v[0])
	}

	req.URL.RawQuery = query.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
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

func GetCrypto(ctx context.Context, query url.Values, headers map[string][]string) ([]*model.Cryptocurrency, error) {
	CryptoEndpoint := URL + CryptoCurrencyEndpoint
	client := &http.Client{}
	req, err := http.NewRequest("GET", CryptoEndpoint, nil)
	if err != nil {
		return nil, err
	}

	if len(query) == 0 {
		query.Add("listing_status", "active")
	}

	if len(headers) == 0 {
		req.Header.Set("Accept", "application/json")
		req.Header.Add("Accept-Encoding", "'deflate,gzip'")
	}

	req.Header.Add("X-CMC_PRO_API_KEY", apiKey)

	for k, v := range headers {
		req.Header.Set(k, v[0])
	}

	req.URL.RawQuery = query.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var cryptocurrencies *model.CryptocurrencyData
	if err = json.Unmarshal(respBody, &cryptocurrencies); err != nil {
		return nil, err
	}

	return cryptocurrencies.Data, nil
}
