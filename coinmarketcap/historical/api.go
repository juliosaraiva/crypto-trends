package coinmarketcap

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

type HistoricalQueryParam struct {
	ID          string
	Symbol      string
	TimeStart   string
	TimeEnd     string
	Count       int
	Interval    string
	Aux         string
	SkipInvalid bool
}

type HistoricalData struct {
	Data *Crypto `json:"data"`
}

type CoinHistorical struct {
	ID       int                 `json:"id,omitempty"`
	Name     string              `json:"name"`
	Symbol   string              `json:"symbol"`
	IsActive int                 `json:"is_active"`
	IsFiat   int                 `json:"is_fiat"`
	Quotes   []*HistoricalQuotes `json:"quotes"`
}

type Crypto struct {
	CORE []*CoinHistorical `json:"core"`
}

type HistoricalQuotes struct {
	Timestamp string `json:"timestamp"`
	Quote     *USD   `json:"quote"`
}

type USD struct {
	USD *USDQuoteHistorical `json:"USD"`
}

type USDQuoteHistorical struct {
	Price             float64 `json:"price"`
	Volume24h         float64 `json:"volume_24h"`
	MarketCap         float64 `json:"market_cap"`
	CirculatingSupply float64 `json:"circulating_supply"`
	TotalSupply       float64 `json:"total_supply"`
	Timestamp         string  `json:"timestamp"`
}

func (h *CoinHistorical) Get(query *url.Values) (*Crypto, error) {
	URL := "https://pro-api.coinmarketcap.com/v2/cryptocurrency/quotes/historical"
	apiKey := os.Getenv("COINMARKETCAP_API_KEY")
	client := &http.Client{}
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := query
	// q.Add("time_period", "")
	q.Add("symbol", "CORE")
	// q.Add("aux", "num_market_pairs,cmc_rank,date_added,tags,platform,max_supply,circulating_supply,total_supply,market_cap_by_total_supply,volume_24h_reported,volume_7d,volume_7d_reported,volume_30d,volume_30d_reported,is_market_cap_included_in_calc")
	// q.Add("start", "1")
	// q.Add("limit", "5000")
	// q.Add("convert", "USD")

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
		return nil, fmt.Errorf("%s", respBody)
	}

	defer resp.Body.Close()

	hist := HistoricalData{}
	err = json.Unmarshal(respBody, &hist)
	if err != nil {
		fmt.Print(err)
	}

	return hist.Data, nil
}
