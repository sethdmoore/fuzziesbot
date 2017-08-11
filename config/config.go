package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/sethdmoore/fuzziesbot/constants"
)

type Config struct {
	// Bot token
	Token string `required:"true"`

	// This chat ID is for forwarding messages
	WaitingRoomChatID int64 `split_words:"true" required:"true"`

	// This chat ID is required for the exportChatInviteLink tgapi method
	MainRoomChatID int64 `split_words:"true" required:"true"`

	// set the log level. See the juju docs for more info
	// https://github.com/juju/loggo/blob/master/level.go#L13-L21
	LogLevel string `split_words:"true" default:"INFO"`
}

func Get() (*Config, error) {
	var c Config
	err := envconfig.Process(constants.Prefix, &c)

	if err != nil {
		return nil, err
	}

	return &c, nil
}
