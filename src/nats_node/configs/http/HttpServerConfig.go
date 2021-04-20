package configs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"nats_node/utils/logger"
	"os"
)

type ServerConfig struct {
	DefaultIP    string `json:"server_hostname"`
	DefaultPort  string `json:"server_port"`
	ReadTimeout  int    `json:"server_read_timeout"`
	WriteTimeout int    `json:"server_write_timeout"`
	Concurancy   int    `json:"server_concurancy"`
}

func SetDefaultServerConfig() *ServerConfig {
	config := &ServerConfig{
		"localhost",
		"8080",
		60,
		60,
		65535,
	}

	value, isSet := os.LookupEnv("ELK_CONFIG_PATH")

	if isSet && value != "" {

		fileName := value + "/server_config.json"
		configFile, err := ioutil.ReadFile(fileName)

		if err != nil {
			logger.Logger.PrintWarn(err.Error())
		} else {
			err := json.Unmarshal(configFile, config)
			if err != nil {
				fmt.Println(err)
			}
		}
	} else {
		fmt.Println("Server config environment variable ELK_CONFIG_PATH not set")
		fmt.Println("Apply default config")
	}

	fmt.Println("Server config")
	fmt.Println(config)

	return config
}
