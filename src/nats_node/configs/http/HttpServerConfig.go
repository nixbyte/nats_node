package configs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"nats_node/utils/logger"
	"os"
	"time"

	"github.com/valyala/fasthttp"
)

type NatsNodeHttpServersConfig struct {
	HttpServerCfg    *ServerConfig `json:"response_server"`
	MetricServerCfg  *ServerConfig `json:"metric_server"`
	SwaggerServerCfg *ServerConfig `json:"swagger_server"`
}

func (natsNodeConfig *NatsNodeHttpServersConfig) String() string {
	return fmt.Sprintf("{\nresponse_server : %v\nmetric_server : %v\nswagger_server : %v\n}", natsNodeConfig.HttpServerCfg, natsNodeConfig.MetricServerCfg, natsNodeConfig.SwaggerServerCfg)
}

type ServerConfig struct {
	DefaultIP    string `json:"server_hostname"`
	DefaultPort  string `json:"server_port"`
	ReadTimeout  int    `json:"server_read_timeout"`
	WriteTimeout int    `json:"server_write_timeout"`
	Concurancy   int    `json:"server_concurancy"`
}

func (serverConfig *ServerConfig) String() string {
	return fmt.Sprintf("{\n    server_hostname : %s,\n    server_port : %s,\n    server_read_timeout : %d,\n    server_write_timeout : %d,\n    server_concurancy : %d\n  }    ", serverConfig.DefaultIP, serverConfig.DefaultPort, serverConfig.ReadTimeout, serverConfig.WriteTimeout, serverConfig.Concurancy)
}

func (config ServerConfig) InitServer(handler func(ctx *fasthttp.RequestCtx)) *fasthttp.Server {
	return &fasthttp.Server{
		Handler:      handler,
		Concurrency:  config.Concurancy,
		ReadTimeout:  time.Duration(config.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(config.WriteTimeout) * time.Second,
	}
}

func SetDefaultNatsNodeHttpServerConfig() *NatsNodeHttpServersConfig {
	responseServerConfig := &ServerConfig{
		"localhost",
		"8080",
		60,
		60,
		65535,
	}

	metricServerConfig := &ServerConfig{
		"localhost",
		"8081",
		60,
		60,
		65535,
	}

	swaggerServerConfig := &ServerConfig{
		"localhost",
		"8082",
		60,
		60,
		65535,
	}

	config := &NatsNodeHttpServersConfig{
		responseServerConfig,
		metricServerConfig,
		swaggerServerConfig,
	}

	value, isSet := os.LookupEnv("ELK_CONFIG_PATH")

	if isSet && value != "" {

		fileName := value + "/nats_http_servers_config.json"
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
		fmt.Println("Nats http Servers config environment variable ELK_CONFIG_PATH not set")
		fmt.Println("Apply default config")
	}

	fmt.Println("Nats http servers config")
	fmt.Printf("Config - %v", config)

	return config
}
