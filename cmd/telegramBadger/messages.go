package main

import (
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/rogercoll/go-dictionary"
)

//Markdown format with bold prefix
func singleMessage(chatID int64, prefix, content string) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(chatID, "")
	msg.ParseMode = "markdown"
	msg.Text = "*" + prefix + "* " + content
	return msg
}

func dictionaryMessage(bot *tgbotapi.BotAPI, sleep time.Duration, chatID int64, data dictionary.Entry) error {
	eng := string(data.Key)
	cat := string(data.Value)
	_, err := bot.Send(singleMessage(chatID, "Eng:", eng))
	if err != nil {
		return nil
	}
	time.Sleep(sleep)
	_, err = bot.Send(singleMessage(chatID, "Cat:", cat))
	if err != nil {
		return err
	}
	time.Sleep(sleep)
	_, err = bot.Send(singleMessage(chatID, "*-----*", ""))
	return err
}
