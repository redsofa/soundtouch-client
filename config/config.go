package config

import (
	"encoding/json"
	"github.com/redsofa/soundtouch-client/logger"
	"os"
)

var ServerConf Config

type Config struct {
	Port          int
	SoundTouchUrl string
}

func ReadServerConf(directory string) {
	f, err := os.Open(directory + "config.json")
	defer f.Close()

	if err != nil {
		logger.Error.Println(err)
	}

	decoder := json.NewDecoder(f)

	err = decoder.Decode(&ServerConf)
	if err != nil {
		logger.Error.Println(err)
	}
}
