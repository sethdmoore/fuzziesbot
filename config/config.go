package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/sethdmoore/fuzziesbot/constants"
)

type Config struct {
	Token             string `required:"true"`
	WaitingRoomChatID int64  `split_words:"true" required:"true"`
}

func Get() (*Config, error) {
	var c Config
	err := envconfig.Process(constants.Prefix, &c)

	if err != nil {
		return nil, err
	}

	return &c, nil
}
