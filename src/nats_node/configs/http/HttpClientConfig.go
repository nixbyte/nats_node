package configs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"nats_node/utils/logger"
	"os"
)

type ClientConfig struct {
	DefaultHostName       string `json:"client_hostname"`
	Guid                  string `json:"guid"`
	DealerConnectTimeout  int    `json:"connection_timeout"`
	DealerKeepAlive       int    `json:"keepalive"`
	MaxIdleConns          int    `json:"max_idle_connections"`
	MaxIdleConnsPerHost   int    `json:"max_idle_conns_per_host"`
	TLSHandshakeTimeout   int    `json:"tls_handshake_timeout"`
	ResponseHeaderTimeout int    `json:"response_header_timeout"`
	Timeout               int    `json:"timeout"`
}

func SetDefaultClientConfig() *ClientConfig {
	config := &ClientConfig{
		"http://r78-rc.zdrav.netrika.ru",
		"C4F530D6-E6EC-4C6E-ACE7-231ADFE928CB",
		60,
		60,
		65500,
		65500,
		20,
		60,
		120,
	}

	value, isSet := os.LookupEnv("ELK_CONFIG_PATH")

	if isSet && value != "" {

		fileName := value + "/client_config.json"
		configFile, err := ioutil.ReadFile(fileName)
		if err != nil {
			logger.Logger.PrintWarn(err.Error())
		} else {
			json.Unmarshal(configFile, &config)
		}
	} else {
		logger.Logger.PrintWarn("Client config environment variable not set")
		logger.Logger.PrintWarn("Using default config")
	}

	fmt.Println("Client config")
	fmt.Println(config)
	return config
}
