package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/sethdmoore/fuzziesbot/config"
	"log"
	"os"
	"strconv"
	"time"
	//"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sethdmoore/telebot"
)

func messages(bot *telebot.Bot, c *config.Config) {
	var dest telebot.Chat
	for message := range bot.Messages {
		log.Printf("Message from %s at %v from %s with text %s\n", strconv.FormatInt(message.Chat.ID, 10), message.Unixtime, message.Sender.Username, message.Text)
		// if an administrator has responded to a forwarded message
		if message.Chat.ID == c.WaitingRoomChatID && message.IsReply() {
			if message.ReplyTo.Sender.ID == bot.Identity.ID {
				log.Printf("HEY A REPLY to ME")
				bot.SendMessage(message.ReplyTo.OriginalSender, message.Text, nil)
				//spew.Dump(message)
			}

		} else if message.Chat.ID == c.WaitingRoomChatID {

			log.Printf("Ignoring...\n")
			continue

		} else if message.IsPersonal() {

			//err := bot.DeleteMessage(message)

			dest.ID = c.WaitingRoomChatID
			err := bot.ForwardMessage(dest, message)
			if err != nil {
				log.Printf("Error forwarding message: %s\n", err)
			}
			//appender <- fmt.Sprintf("%v:%v", message.Unixtime)
		} else {
			log.Printf("Unhandled message")
			spew.Dump(message)
		}
	}
}

func main() {
	conf, err := config.Get()
	if err != nil {
		log.Printf("Could not process config: %s\n", err)
		os.Exit(2)
	}

	fmt.Println("%+v", conf)

	bot, err := telebot.NewBot(conf.Token)
	if err != nil {
		fmt.Printf("Could not initialize bot!: %s\n", err)
		os.Exit(2)
	}

	bot.Messages = make(chan telebot.Message, 100)
	//pushchan := make(chan string)

	go messages(bot, conf)
	//go appender(pushchan)

	bot.Start(1 * time.Second)
}
