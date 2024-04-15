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
	"github.com/juliosaraiva/crypto-trends/config"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
	"google.golang.org/api/option"
)

type Bot struct {
	bot    *telego.Bot
	bh     *th.BotHandler
	token  string
	apiKey string // Added for Gemini API
}

func NewBot(token, apiKey string) (*Bot, error) {
	bot, err := telego.NewBot(token, telego.WithDefaultDebugLogger())
	if err != nil {
		return nil, err
	}

	updates, err := bot.UpdatesViaLongPolling(nil)
	if err != nil {
		return nil, err
	}

	bh, err := th.NewBotHandler(bot, updates)
	if err != nil {
		return nil, err
	}

	return &Bot{
		bot:    bot,
		bh:     bh,
		token:  token,
		apiKey: apiKey, // Store API key securely
	}, nil
}

// Start inicia o bot.
func (b *Bot) Start() {
	defer b.bh.Stop()
	defer b.bot.StopLongPolling()

	b.registerCommands()

	b.bh.Start()
}

func (b *Bot) registerCommands() {
	b.registerBotInfoCommand()
	b.registerCriptoFirmaCommand()
	// Add more commands here (e.g., /price, /watchlist)
}

func (b *Bot) registerBotInfoCommand() {
	b.bh.Handle(func(bot *telego.Bot, update telego.Update) {
		botUser, err := bot.GetMe()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		_, _ = bot.SendMessage(tu.Message(
			tu.ID(update.Message.Chat.ID),
			fmt.Sprintf("Bot user: %+v", botUser),
		))
	}, th.CommandEqual("bot_inf"))
}

func (b *Bot) registerCriptoFirmaCommand() {
	b.bh.Handle(func(bot *telego.Bot, update telego.Update) {
		ctx := context.Background()

		var query url.Values = url.Values{}
		query.Add("symbol", "BTC,GPU,CKB") // Replace with user-specified symbols (optional)
		query.Add("sort", "cmc_rank")
		var reqHeaders map[string][]string = map[string][]string{}

		cryptocurrencies, err := coinmarketcap.ListCryptocurrencies(ctx, query, reqHeaders)
		if err != nil {
			fmt.Println("Error fetching data:", err)
			_, _ = bot.SendMessage(tu.Message(
				tu.ID(update.Message.Chat.ID),
				"Error retrieving data. Please try again later.",
			))
			return
		}

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
		fifteenDaysAgoUnixTimestamp := now.AddDate(0, 0, -15).Unix()
		fifteenDaysTimestampString := strconv.FormatInt(fifteenDaysAgoUnixTimestamp, 10)

		queryHistorical.Add("time_start", fifteenDaysTimestampString)
		// Consider adding "time_end" for more control over historical data
		queryHistorical.Add("interval", "1h") // Adjust interval as needed

		listing, err := coinmarketcap.GetCrypto(ctx, queryListing, reqHeaders)
		if err != nil {
			fmt.Println("Error fetching listing data:", err)
			_, _ = bot.SendMessage(tu.Message(
				tu.ID(update.Message.Chat.ID),
				"Error retrieving data. Please try again later.",
			))
			return
		}

		historical, err := coinmarketcap.GetOHLCVHistoricalPrices(ctx, queryHistorical, reqHeaders)
		if err != nil {
			fmt.Println("Error fetching historical data:", err)
			_, _ = bot.SendMessage(tu.Message(
				tu.ID(update.Message.Chat.ID),
				"Error retrieving data. Please try again later.",
			))
			return
		}

		sliceIDs := strings.Split(IDList, ",")

		var cryptocurrency model.Cryptocurrency
		var cryptoList []model.Cryptocurrency

		for _, id := range sliceIDs {
			parseJSON, err := json.Marshal(listing[id])
			if err != nil {
				fmt.Println("Error marshalling listing data:", err)
				continue // Skip this cryptocurrency if marshalling fails
			}

			json.Unmarshal([]byte(parseJSON), &cryptocurrency)
			cryptocurrency.Historical = historical.Quotes
			cryptoList = append(cryptoList, cryptocurrency)
		}

		// Consider using a different trend analysis service or implementing your own logic
		output, err := Gemini(cryptoList, b.apiKey) // Use your API key securely
		if err != nil {
			fmt.Println("Error analyzing trends:", err)
			_, _ = bot.SendMessage(tu.Message(
				tu.ID(update.Message.Chat.ID),
				"Error analyzing trends. Please try again later.",
			))
			return
		}

		JSONOutput, _ := json.Marshal(output)

		_, _ = bot.SendMessage(tu.Message(
			tu.ID(update.Message.Chat.ID),
			string(JSONOutput),
		))
	}, th.CommandEqual("crypto"))
}

// Gemini is a placeholder function for trend analysis. Replace with your actual implementation.
func Gemini(cryptoList []model.Cryptocurrency, apiKey string) (*genai.GenerateContentResponse, error) {
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

func main() {
	if config.GetConfig("MODULE_TELEGRAM") == "true" {
		token := config.GetConfig("TOKEN_BOT_TELEGRAM")
		apiKey := os.Getenv("GEMINI_API_KEY") // Retrieve API key securely from environment variable

		bot, err := NewBot(token, apiKey)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		bot.Start()
	}
}
