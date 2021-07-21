package configs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"nats_node/utils/logger"
	"os"
)

type NatsNodeHttpClientConfig struct {
	DefaultHostName       string `json:"client_hostname"`
	ClientCHUrl           string `json:"client_churl"`
	ClientCHDatabase      string `json:"client_chdatabase"`
	ClientCHTableName     string `json:"client_chtablename"`
	Guid                  string `json:"guid"`
	DealerConnectTimeout  int    `json:"connection_timeout"`
	DealerKeepAlive       int    `json:"keepalive"`
	MaxIdleConns          int    `json:"max_idle_connections"`
	MaxIdleConnsPerHost   int    `json:"max_idle_conns_per_host"`
	TLSHandshakeTimeout   int    `json:"tls_handshake_timeout"`
	ResponseHeaderTimeout int    `json:"response_header_timeout"`
	Timeout               int    `json:"timeout"`
}

func (clientConfig *NatsNodeHttpClientConfig) String() string {
	return fmt.Sprintf("{\n    client_hostname : %s,\n    client_churl : %s,\n    client_chdatabase : %s,\n    client_chtablename : %s,\n    guid : %s,\n    connection_timeout : %d,\n    keepalive : %d,\n    max_idle_connections : %d,\n   max_idle_conns_per_host : %d,\n    tls_handshake_timeout : %d,\n    response_header_timeout : %d,\n    timeout : %d\n  }\n", clientConfig.DefaultHostName, clientConfig.ClientCHUrl, clientConfig.ClientCHDatabase, clientConfig.ClientCHTableName, clientConfig.Guid, clientConfig.DealerConnectTimeout, clientConfig.DealerKeepAlive, clientConfig.MaxIdleConns, clientConfig.MaxIdleConnsPerHost, clientConfig.TLSHandshakeTimeout, clientConfig.ResponseHeaderTimeout, clientConfig.Timeout)
}

func SetDefaultClientConfig() *NatsNodeHttpClientConfig {
	config := &NatsNodeHttpClientConfig{
		"http://localhost:8080",
		"tcp://rc1b-egh7zplyyhs7s5k8.mdb.yandexcloud.net:9440?username=elk-sport&password=EgyadeyWi&database=db1",
		"db1",
		"statistic",
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

		fileName := value + "/nats_http_client_config.json"
		configFile, err := ioutil.ReadFile(fileName)
		if err != nil {
			logger.Logger.PrintWarn(err.Error())
		} else {
			err := json.Unmarshal(configFile, config)
			if err != nil {
				logger.Logger.PrintWarn(err.Error())
			}
		}
	} else {
		logger.Logger.PrintWarn("Client config environment variable not set")
		logger.Logger.PrintWarn("Using default config")
	}

	fmt.Println("Nats http client config")
	fmt.Printf("Config - %v", config)

	return config
}
