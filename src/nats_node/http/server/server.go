package server

import (
	"fmt"
	configs "nats_node/configs/http"
	"nats_node/utils/logger"
	"nats_node/utils/monitoring"
	"os"

	"github.com/fasthttp/router"
	_ "github.com/swaggo/http-swagger/example/go-chi/docs"

	"github.com/valyala/fasthttp"
)

var (
	NatsHttpServersConfig *configs.NatsNodeHttpServersConfig
	ApiServer             *httpServer
	MetricServer          *httpServer
)

type httpServer struct {
	HttpConfig *configs.ServerConfig
	api        *apiRouter
	server     *fasthttp.Server
}

type apiRouter struct {
	router *router.Router
}

func newRouter() *apiRouter {
	var a *apiRouter
	a = &apiRouter{
		router: router.New(),
	}
	return a
}

func init() {
	fmt.Println("Init Servers...")

	NatsHttpServersConfig = configs.SetDefaultNatsNodeHttpServerConfig()

	Api := newRouter()

	ApiServer = &httpServer{
		NatsHttpServersConfig.HttpServerCfg,
		Api,
		NatsHttpServersConfig.HttpServerCfg.InitServer(Api.router.Handler),
	}

	if monitoring.Monitoring.WRITE_METRICS {
		MetricApi := newRouter()
		MetricServer = &httpServer{
			NatsHttpServersConfig.MetricServerCfg,
			MetricApi,
			NatsHttpServersConfig.MetricServerCfg.InitServer(MetricApi.router.Handler),
		}
	}
}

func (apiServer httpServer) GetRouter() *router.Router {
	return apiServer.api.router
}

func (httpServer *httpServer) runServer() {
	go func() {
		err := httpServer.server.ListenAndServe(httpServer.HttpConfig.DefaultIP + ":" + httpServer.HttpConfig.DefaultPort)
		if err != nil {
			logger.Logger.PrintError(err)
			os.Exit(1)
		}
	}()
}

func Start() {

	ApiServer.runServer()

	if monitoring.Monitoring.WRITE_METRICS {
		MetricServer.runServer()
	}
}
