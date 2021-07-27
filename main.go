package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"test.com/project_structure/config"
	"test.com/project_structure/internal/platform/logger"
	"test.com/project_structure/internal/platform/logger/zerologwrapper"
)

func main() {
	config.InitConfig(getConfigDir(os.Getenv("ENV")))
	mainConfig := config.GetMainConfig()

	var err error
	//initialize logger
	{

		logWriter := os.Stdout
		if mainConfig.Logger.File != "" {
			logWriter, err = os.OpenFile(mainConfig.Logger.File, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Fatal(err)
			}
		}
		zeroLogger := zerologwrapper.
			NewZeroLogWrapper(
				zerologwrapper.
					GetZerologDefaultLogger(logWriter, mainConfig.Logger.Level, 4),
			)
		logger.InitLogger(zeroLogger)
	}
	logger.Debug("test debug")
	logger.Debugf("%+s", "test debugf")

	logger.Info("test info")
	logger.Infof("%+s", "test infof")

	logger.Warn("test warn")
	logger.Warnf("%+s", "test warnf")

	logger.Error("test error")
	logger.Errorf("%+s", "test errorf")

	logger.Fatal("test fatal")
	logger.Fatalf("%+s", "test fatalf")
}

func getConfigDir(environment string) string {
	currDir := ""
	var err error
	if environment == "" {
		environment = "development"
		currDir, err = filepath.Abs(".")
		if err != nil {
			log.Fatal(err)
		}
		currDir = filepath.Join(currDir, "files")
	}
	return filepath.Join(currDir, fmt.Sprintf("/etc/gobiz/%s", environment))
}
