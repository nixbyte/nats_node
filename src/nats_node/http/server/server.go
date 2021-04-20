package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	configs "nats_node/configs/http"
	"nats_node/utils/logger"
	"nats_node/utils/monitoring"
	"os"
	"strings"
	"time"

	"github.com/valyala/fasthttp"
)

type apistruct struct {
	router  map[string]fasthttp.RequestHandler
	Handler func(*fasthttp.RequestCtx)
}

var (
	server       *fasthttp.Server
	MetricServer *fasthttp.Server
	Api          apistruct
	Config       *configs.ServerConfig
	MetricConfig *configs.ServerConfig
	MetricApi    apistruct
)

func init() {
	fmt.Println("Init Server...")

	Config = configs.SetDefaultServerConfig()

	Api = apistruct{
		router: make(map[string]fasthttp.RequestHandler),
		Handler: func(ctx *fasthttp.RequestCtx) {

			var url string = string(ctx.Request.RequestURI())

			if strings.IndexByte(string(ctx.Request.RequestURI()), '?') != -1 {
				url = url[:strings.IndexByte(url, '?')]
			}

			handler, ok := Api.router[url]
			if !ok {
				ctx.NotFound()
				return
			}
			handler(ctx)
		},
	}

	server = &fasthttp.Server{
		Handler:      Api.Handler,
		Concurrency:  Config.Concurancy,
		ReadTimeout:  time.Duration(Config.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(Config.WriteTimeout) * time.Second,
	}

	if monitoring.Monitoring.WRITE_METRICS {

		MetricConfig = setMetricsConfig()

		MetricApi = apistruct{
			router: make(map[string]fasthttp.RequestHandler),
			Handler: func(ctx *fasthttp.RequestCtx) {

				var url string = string(ctx.Request.RequestURI())

				if strings.IndexByte(string(ctx.Request.RequestURI()), '?') != -1 {
					url = url[:strings.IndexByte(url, '?')]
				}

				handler, ok := MetricApi.router[url]
				if !ok {
					ctx.NotFound()
					return
				}
				handler(ctx)
			},
		}

		MetricServer = &fasthttp.Server{
			Handler:      MetricApi.Handler,
			Concurrency:  MetricConfig.Concurancy,
			ReadTimeout:  time.Duration(MetricConfig.ReadTimeout) * time.Second,
			WriteTimeout: time.Duration(MetricConfig.WriteTimeout) * time.Second,
		}
	}
}

func setMetricsConfig() *configs.ServerConfig {
	config := &configs.ServerConfig{
		"localhost",
		"8081",
		60,
		60,
		65535,
	}

	value, isSet := os.LookupEnv("ELK_CONFIG_PATH")

	if isSet && value != "" {

		fileName := value + "/metric_server_config.json"
		configFile, err := ioutil.ReadFile(fileName)
		logger.Logger.PrintError(err)

		json.Unmarshal(configFile, config)
	} else {
		fmt.Println("Server config environment variable METRIC_CONFIG_PATH not set")
		fmt.Println("Apply default config")
	}

	fmt.Println("Metrics Server config")
	fmt.Println(config)

	return config
}

func (api apistruct) AddHandlerToRoute(route string, handler fasthttp.RequestHandler) {
	api.router[route] = handler
}

func Start() {
	go func() {
		defaultAddress := Config.DefaultIP + ":" + Config.DefaultPort
		err := server.ListenAndServe(defaultAddress)
		if err != nil {
			logger.Logger.PrintError(err)
			os.Exit(1)
		}
	}()
	if monitoring.Monitoring.WRITE_METRICS {
		go func() {
			defaultMetricAddress := MetricConfig.DefaultIP + ":" + MetricConfig.DefaultPort
			err := MetricServer.ListenAndServe(defaultMetricAddress)
			if err != nil {
				logger.Logger.PrintError(err)
			}
		}()
	}
}
