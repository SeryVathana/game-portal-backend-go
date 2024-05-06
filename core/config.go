package core

import (
	"fmt"
	"io"
	"os"

	"direxplay.com/game-portal-backend-go/model"
	"gopkg.in/yaml.v2"
)

var appConfig *model.Config

func LoadConfig() *model.Config {
	if appConfig == nil {
		fmt.Println("Loading configuration file...")

		f, err := os.Open("config.yaml")
		if err != nil {
			fmt.Println("Unable to load config.yaml!")
			panic(err)
		}
		buffer, err := io.ReadAll(f)
		if err != nil {
			panic(err)
		}
		err = yaml.Unmarshal(buffer, &appConfig)
		if err != nil {
			fmt.Println("Unable to parse config.yaml!")
			panic(err)
		}

		fmt.Println("Configuration loaded!")
	}
	return appConfig
}
