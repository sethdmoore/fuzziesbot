package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/juju/loggo"
	"github.com/sethdmoore/fuzziesbot/config"
	"github.com/sethdmoore/fuzziesbot/logconfig"
	"github.com/sethdmoore/fuzziesbot/messages"
	//"log"
	"os"
	"time"
	//"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sethdmoore/telebot"
)

// root log level
var log loggo.Logger

func main() {
	conf, err := config.Get()
	log.Debugf(spew.Sdump(conf))
	if err != nil {
		log.Infof("Could not process config: %s\n", err)
		os.Exit(2)
	}

	// set package level var
	log = logconfig.New(conf.LogLevel)

	loggo.ConfigureLoggers(conf.LogLevel)

	bot, err := telebot.NewBot(conf.Token)
	if err != nil {
		log.Infof("Could not initialize bot!: %s\n", err)
		os.Exit(2)
	}

	bot.Messages = make(chan telebot.Message, 100)
	//pushchan := make(chan string)

	go messages.HandleMessages(bot, conf)
	//go appender(pushchan)

	log.Infof("Bot started!")
	bot.Start(1 * time.Second)
}
