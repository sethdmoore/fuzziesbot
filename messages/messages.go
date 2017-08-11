package messages

import (
	//"errors"
	"github.com/davecgh/go-spew/spew"
	"github.com/juju/loggo"
	"github.com/sethdmoore/fuzziesbot/commands"
	"github.com/sethdmoore/fuzziesbot/config"
	"github.com/sethdmoore/fuzziesbot/util"
	"github.com/sethdmoore/telebot"
	"strconv"
)

var log loggo.Logger

func HandleReply(message telebot.Message, bot *telebot.Bot) error {
	authorized, err := util.IsAuthorized(message, bot)
	if err != nil {
		return err
	}

	if authorized {
		err := bot.SendMessage(message.ReplyTo.OriginalSender, message.Text, nil)
		if err != nil {
			return err
		}
		return nil
	} else {
		bot.SendMessage(message.Chat, "Unauthorized", nil)
		if err != nil {
			return err
		}
		return nil
	}
}

// HandleGeneral handles general messages (personal or admin group chat)
func HandleGeneral(m telebot.Message, bot *telebot.Bot, c *config.Config) error {
	var dest telebot.Chat
	if m.ContainsCommand() {
		log.Debugf("Received command!")

	} else if m.Chat.ID == c.WaitingRoomChatID {
		log.Debugf("Message does not need to be handled.")
		log.Tracef(spew.Sdump(m))
		return nil
	} else if m.IsPersonal() {
		dest.ID = c.WaitingRoomChatID
		err := bot.ForwardMessage(dest, m)
		if err != nil {
			log.Errorf("Error forwarding message: %s\n", err)
			return err
		}
	} else {
		log.Infof("Unhandled message")
		log.Debugf(spew.Sdump(m))
	}
	return nil
}

func HandleMessages(bot *telebot.Bot, c *config.Config) {
	log = loggo.GetLogger("messages")

	for message := range bot.Messages {
		log.Debugf("Message from %s at %v from %s with text %s\n", strconv.FormatInt(message.Chat.ID, 10), message.Unixtime, message.Sender.Username, message.Text)

		// if someone in the waiting room has replied to the forward
		if util.IsProxyReply(c.WaitingRoomChatID, message) {
			log.Debugf("Received eligible reply from waiting room")
			err := HandleReply(message, bot)
			if err != nil {
				log.Warningf("Error checking chat privileges: %s\n", err)
				continue
			}

		} else {
			err := HandleGeneral(message, bot, c)
			if err != nil {
				log.Warningf("")
				continue
			}
		}

	}
}
