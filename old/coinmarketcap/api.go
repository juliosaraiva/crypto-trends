package coinmarketcap

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/juliosaraiva/crypto-trends/coinmarketcap/model"
	"github.com/juliosaraiva/crypto-trends/config"
)

var apiKey string = config.GetConfig("COINMARKETCAP_API_KEY")

const (
	BaseURL = "https://pro-api.coinmarketcap.com"
	// Endpoints
	CryptoCategoriesEndpoint    = "/v1/cryptocurrency/categories"
	CryptoCategoryEndpoint      = "/v1/cryptocurrency/category"
	ListCryptoEndpoint          = "/v1/cryptocurrency/map"
	CryptoLatestEndpoint        = "/v2/cryptocurrency/quotes/latest"
	CryptoHistoricalEndpoint    = "/v2/cryptocurrency/quotes/historical"
	CryptoOHLCVHistoricalPrices = "/v2/cryptocurrency/ohlcv/historical"
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

func GetCrypto(ctx context.Context, query url.Values, headers map[string][]string) (map[string]*model.Listing, error) {
	endpointURL := BaseURL + CryptoLatestEndpoint

	client := &http.Client{}
	req, err := http.NewRequest("GET", endpointURL, nil)
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
		fmt.Println("Inside")
		fmt.Println(string(respError))
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

func GetHistoricalPrices(ctx context.Context, query url.Values, headers map[string][]string) (map[string][]*model.CryptoHistorical, error) {
	endpointURL := BaseURL + CryptoHistoricalEndpoint
	client := &http.Client{}
	req, err := http.NewRequest("GET", endpointURL, nil)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
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
		return nil, fmt.Errorf("%s", err.Error())
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil || resp.StatusCode != 200 {
		return nil, fmt.Errorf("%s", respBody)
	}

	defer resp.Body.Close()

	h := model.CryptoHistoricalData{}
	err = json.Unmarshal(respBody, &h)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}

	return h.Data, nil
}

func ListCryptocurrencies(ctx context.Context, query url.Values, headers map[string][]string) ([]*model.Cryptocurrency, error) {
	endpointURL := BaseURL + ListCryptoEndpoint
	client := &http.Client{}
	req, err := http.NewRequest("GET", endpointURL, nil)
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

func GetOHLCVHistoricalPrices(ctx context.Context, query url.Values, reqHeaders map[string][]string) (*model.OHLCVHistorical, error) {
	endpointURL := BaseURL + CryptoOHLCVHistoricalPrices
	client := &http.Client{}
	req, err := http.NewRequest("GET", endpointURL, nil)
	if err != nil {
		return nil, err
	}

	if len(reqHeaders) == 0 {
		req.Header.Set("Accept", "application/json")
		req.Header.Set("Accept-Encoding", "'deflate,gzip'")
	}

	req.Header.Add("X-CMC_PRO_API_KEY", apiKey)

	for k, v := range reqHeaders {
		req.Header.Add(k, v[0])
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

	var ohlcvHistorical *model.OHLCVHistoricalData
	if err = json.Unmarshal(respBody, &ohlcvHistorical); err != nil {
		return nil, err
	}

	return ohlcvHistorical.Data, nil
}
