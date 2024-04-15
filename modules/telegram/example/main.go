package main

import (
	"fmt"
	"os"

	"github.com/juliosaraiva/crypto-trends/config"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

type Bot struct {
	bot   *telego.Bot
	bh    *th.BotHandler
	token string
}

func NewBot(token string) (*Bot, error) {
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
		bot:   bot,
		bh:    bh,
		token: token,
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
		_, _ = bot.SendMessage(tu.Message(
			tu.ID(update.Message.Chat.ID),
			`{"crypto": [
				{"id": 1, "name": "Bitcoin", "symbol": "BTC", "trend": "low"},
				{"id": 4948, "name": "Nervos Network", "symbol": "CKB", "trend": "low"},
				{"id": 29513, "name": "Node AI", "symbol": "GPU", "trend": "low"}
			]}`,
		))
	}, th.CommandEqual("crypto"))
}

func main() {
	if config.GetConfig("MODULE_TELEGRAM") == "true" {
		token := config.GetConfig("TOKEN_BOT_TELEGRAM")
		bot, err := NewBot(token)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		bot.Start()
	}
}
