package bot

import (
	"bp-telegram-bot/internal/api"
	"bp-telegram-bot/internal/bot/constants"
	"time"

	"gopkg.in/telebot.v3"
)

type Bot struct {
	tb    *telebot.Bot
	api   *api.Client
	state map[int64]constants.State

	buffer map[int64]map[string]any
}

func NewBot(token string, api *api.Client) *Bot {
	tb, err := telebot.NewBot(telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})

	//log.Println(tb, err)
	if err != nil {
		return nil
	}

	return &Bot{
		tb:     tb,
		api:    api,
		state:  make(map[int64]constants.State),
		buffer: make(map[int64]map[string]any),
	}
}

func (bot *Bot) Start() {
	bot.InitRoutes()
	bot.tb.Start()
}
