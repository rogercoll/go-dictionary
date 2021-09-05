package main

import (
	"math/rand"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/rogercoll/go-dictionary"
)

func autoTest(d dictionary.Dictionary, bot *tgbotapi.BotAPI, chatID int64) {
	entries, err := d.GetAll()
	if err != nil {
		msg := tgbotapi.NewMessage(chatID, err.Error())
		bot.Send(msg)
		return
	}
	if len(entries) < 1 {
		bot.Send(singleMessage(chatID, "Error:", "No stored entries, please add a definition with /add"))
	}
	for _, entry := range entries {
		err := dictionaryMessage(bot, 2*time.Second, chatID, entry)
		if err != nil {
			msg := tgbotapi.NewMessage(chatID, err.Error())
			bot.Send(msg)
			return
		}
	}
	bot.Send(singleMessage(chatID, "---", ""))
}

func sendRandom(d dictionary.Dictionary, bot *tgbotapi.BotAPI, chatID int64) {
	entries, err := d.GetAll()
	if err != nil {
		msg := tgbotapi.NewMessage(chatID, err.Error())
		bot.Send(msg)
		return
	}
	if len(entries) < 1 {
		bot.Send(singleMessage(chatID, "Error:", "No stored entries, please add a definition with /add"))
		return
	}
	min := 0
	max := len(entries)
	r := rand.Intn(max-min) + min
	err = dictionaryMessage(bot, 2*time.Second, chatID, entries[r])
	if err != nil {
		msg := tgbotapi.NewMessage(chatID, err.Error())
		bot.Send(msg)
		return
	}
}
