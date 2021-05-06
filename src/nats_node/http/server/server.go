package server

import (
	"fmt"
	configs "nats_node/configs/http"
	"nats_node/utils/logger"
	"nats_node/utils/monitoring"
	"os"
	"strings"

	"github.com/fasthttp/router"
	_ "github.com/swaggo/http-swagger/example/go-chi/docs"

	"github.com/valyala/fasthttp"
)

var (
	NatsHttpServersConfig *configs.NatsNodeHttpServersConfig
	ApiServer             *httpServer
	MetricServer          *httpServer
	SwaggerServer         *httpServer
)

type httpServer struct {
	HttpConfig *configs.ServerConfig
	api        *apiStructure
	server     *fasthttp.Server
}

type apiStructure struct {
	router  *router.Router
	Handler func(*fasthttp.RequestCtx)
}

func newApi() *apiStructure {
	var a *apiStructure
	a = &apiStructure{
		router: router.New(),
		Handler: func(ctx *fasthttp.RequestCtx) {

			var url string = string(ctx.Request.RequestURI())

			if strings.IndexByte(string(ctx.Request.RequestURI()), '?') != -1 {
				url = url[:strings.IndexByte(url, '?')]
			}

			handler := a.router.Handler
			handler(ctx)
		},
	}
	return a
}

func init() {
	fmt.Println("Init Servers...")

	NatsHttpServersConfig = configs.SetDefaultNatsNodeHttpServerConfig()

	Api := newApi()
	SwaggerApi := newApi()

	ApiServer = &httpServer{
		NatsHttpServersConfig.HttpServerCfg,
		Api,
		NatsHttpServersConfig.HttpServerCfg.InitServer(Api.Handler),
	}

	SwaggerServer = &httpServer{
		NatsHttpServersConfig.SwaggerServerCfg,
		SwaggerApi,
		NatsHttpServersConfig.SwaggerServerCfg.InitServer(SwaggerApi.Handler),
	}

	if monitoring.Monitoring.WRITE_METRICS {
		MetricApi := newApi()
		MetricServer = &httpServer{
			NatsHttpServersConfig.MetricServerCfg,
			MetricApi,
			NatsHttpServersConfig.MetricServerCfg.InitServer(MetricApi.Handler),
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

	SwaggerServer.runServer()

	if monitoring.Monitoring.WRITE_METRICS {
		MetricServer.runServer()
	}
}
