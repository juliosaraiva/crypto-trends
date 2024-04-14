package main

import (
	"fmt"
	"os"

	"github.com/juliosaraiva/crypto-trends/config"
	"github.com/mymmrac/telego"
)

func main() {
	// Get Bot token from environment variables
	var botToken string = config.Config("TOKEN_BOT_TELEGRAM")

	// Create bot and enable debugging info
	// Note: Please keep in mind that default logger may expose sensitive information,
	// use in development only
	// (more on configuration in examples/configuration/main.go)
	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Call method getMe (https://core.telegram.org/bots/api#getme)
	botUser, err := bot.GetMe()
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Print Bot information
	fmt.Printf("Bot user: %+v\n", botUser)
}
