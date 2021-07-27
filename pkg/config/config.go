package config

import (
	"log"
	"sync"

	"github.com/spf13/viper"
)

var (
	mainConfig *config
	singleton  sync.Once
)

type config struct {
	Flight struct {
		Host string
		Port int32
	}

	Logger struct {
		Level string
		File  string
	}
}

func InitConfig(path string) {
	singleton.Do(func() {
		mainConfig = readConfig(path)
	})
}

func GetMainConfig() *config {
	return mainConfig
}

func readConfig(path string) *config {
	viper.AddConfigPath(path)
	viper.SetConfigName("main")

	var config config
	setDefaultValues()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}

	return &config
}

func setDefaultValues() {
	viper.SetDefault("Flight.Port", 633)
	viper.SetDefault("Logger.Level", "error")
}
