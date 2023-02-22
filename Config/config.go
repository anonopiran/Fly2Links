package config

import (
	"encoding/json"
	"fmt"
	"os"

	"sync"

	"github.com/ilyakaznacheev/cleanenv"
	log "github.com/sirupsen/logrus"
)

var lock = &sync.Mutex{}
var settingsInstance *SettingsType
var profileInstance *[]ProfileType

func Describe() {
	var cfg SettingsType
	help, err := cleanenv.GetDescription(&cfg, nil)
	if err != nil {
		log.WithError(err).Panic("can not generate description")
	}
	log.Println(help)
}

func Config() *SettingsType {
	if settingsInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		var err error
		settingsInstance = &SettingsType{}
		if _, err_file := os.Stat(".env"); err_file == nil {
			log.Info("found .env file")
			err = cleanenv.ReadConfig(".env", settingsInstance)
		} else {
			log.Info("no .env file found")
			err = cleanenv.ReadEnv(settingsInstance)
		}
		if err != nil {
			log.WithError(err).Panic("can not initiate configuration")
		}
		log.WithField("data", fmt.Sprintf("%+v", settingsInstance)).Debug("Parsed Configuration")
	}
	return settingsInstance
}

func Profile() *[]ProfileType {
	if profileInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		profileInstance = &[]ProfileType{}
		jsonFile, _ := os.ReadFile(string(Config().ProfilePath))
		json.Unmarshal(jsonFile, profileInstance)
		log.WithField("data", fmt.Sprintf("%+v", profileInstance)).Debug("Parsed Profile")
	}
	return profileInstance
}
