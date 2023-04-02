package config

import (
	"fmt"
	"os"
)

var (
	Token     string
	BotPrefix string

	config *configStruct
)

type configStruct struct {
	Token     string `json:"Token"`
	BotPrefix string `json:"BotPrefix"`
}

func ReadConfig() error {

	fmt.Println("Reading config ...")

	// file, err := ioutil.ReadFile("./docs/config.json")

	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return err
	// }

	// fmt.Println(string(file))

	// err = json.Unmarshal(file, &config)

	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return err
	// }

	// Token = config.Token
	// BotPrefix = config.BotPrefix
	// ---------------------------------
	// Implement method of env vars //

	Token = os.Getenv("TOKEN")
	BotPrefix = "No-Hello-Bot"

	if Token == "" || BotPrefix == "" {
		return fmt.Errorf("BOT_TOKEN and BOT_PREFIX environment variables must be set")
	}

	return nil
}
