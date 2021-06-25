package main

import (
	"log"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/rogercoll/go-dictionary"
)

var (
	max = 100
)

func autoTest(d dictionary.Dictionary, bot *tgbotapi.BotAPI, userID int, chatID int64) {

	result, err := d.GetAll()
	if err != nil {
		msg := tgbotapi.NewMessage(chatID, err.Error())
		bot.Send(msg)
		return
	}
	msg := tgbotapi.NewMessage(chatID, "")
	msg.ParseMode = "markdown"
	for _, entry := range result {

		eng := string(entry.Key)
		cat := string(entry.Value)
		msg.Text = "*Eng:* " + eng
		//msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
		time.Sleep(time.Second * 2)
		msg.Text = "*Cat:* " + cat
		bot.Send(msg)
		time.Sleep(time.Second * 1)
		msg.Text = "*-----*"
		time.Sleep(time.Second * 1)
		bot.Send(msg)
	}
	msg.Text = "EOF"
	bot.Send(msg)
}

func RunBot(d *dictionary.Dictionary, token string) error {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return err
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	myCommands := []tgbotapi.BotCommand{
		{Command: "add", Description: "Usage: add hello hola"},
		{Command: "del", Description: "Usage: del hello"},
		{Command: "help", Description: "Display available commands"},
		{Command: "start", Description: "Start command"},
		{Command: "autotest", Description: "Start displaying translation every 2 seconds"},
	}
	err = bot.SetMyCommands(myCommands)
	if err != nil {
		return err
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}
		if update.Message.IsCommand() {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			msg.ParseMode = "markdown"
			switch update.Message.Command() {
			case "add", "Add":
				userMess := strings.Split(update.Message.Text, " ")
				if len(userMess) == 3 {
					err := *d.Insert(userMess[1], userMess[2])
					msg.Text = "*Translation added correctly*"
					if err != nil {
						msg.Text = "Error: " + err.Error()
					}
				} else {
					msg.Text = "Wrong format! Usage: /add hello hola"
				}
			case "del", "Del":
				userMess := strings.Split(update.Message.Text, " ")
				if len(userMess) == 2 {
					err := DeleteEntry(update.Message.From.ID, userMess[1])
					msg.Text = "*Entry deleted correctly*"
					if err != nil {
						msg.Text = "Error: " + err.Error()
					}
				} else {
					msg.Text = "Wrong format! Usage: /del hello"
				}
			case "status":
				msg.Text = "I'm ok."
			case "start":
				msg.Text = "Hello *BITCH*! Welcome to your personal and simple dictonary! Type */help* command to start rolling :)"
			case "help":
				for _, command := range myCommands {
					msg.Text += "*Command* " + command.Command + " *-* " + command.Description + "\n"
				}
			case "autotest":
				autoTest(bot, update.Message.From.ID, update.Message.Chat.ID)
			default:
				msg.Text = "I don't know that command"
			}
			bot.Send(msg)
		} else {
			eng, cat, err := GetRandom(update.Message.From.ID)
			if err != nil {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, err.Error())
				bot.Send(msg)
				continue
			}
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			msg.ParseMode = "markdown"
			msg.Text = "*Eng:* " + eng
			//msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
			time.Sleep(time.Second * 2)
			msg.Text = "*Cat:* " + cat
			bot.Send(msg)
		}
	}
	return nil
}
