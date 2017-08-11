package util

import (
	"github.com/sethdmoore/telebot"
)

func IsProxyReply(waitingRoom int64, message telebot.Message) bool {
	if message.Chat.ID == int64(waitingRoom) && message.IsReply() {
		return true
	} else {
		return false
	}
}

func IsAuthorized(message telebot.Message, bot *telebot.Bot) (bool, error) {
	if message.ReplyTo.Sender.ID != bot.Identity.ID {
		return false, nil
	}

	cm, err := bot.GetChatMember(message.Chat, message.Sender)
	if err != nil {
		return false, err
	}

	if cm.Status == "administrator" || cm.Status == "creator" {
		return true, nil
	}

	return false, nil

}
