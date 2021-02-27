package monitoring

import (
	"fmt"
	"nats_node/utils/logger"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/valyala/fasthttp"
)

type monitoring struct {
	DEBUG_HTTP    bool
	WRITE_METRICS bool
	signalChan    chan os.Signal
	ExitChan      chan int
	StopChan      chan struct{}
}

const (
	ExitCodeOk int = iota
	ExitCodeInterrupted
)

var Monitoring *monitoring

func init() {
	fmt.Println("Init Monitoring...")
	Monitoring = &monitoring{
		DEBUG_HTTP:    false,
		WRITE_METRICS: true,
		signalChan:    make(chan os.Signal),
		ExitChan:      make(chan int),
		StopChan:      make(chan struct{}),
	}
}

func (monitor monitoring) StartMonitoring() {
	signal.Notify(Monitoring.signalChan, os.Interrupt)
	signal.Notify(Monitoring.signalChan, syscall.SIGTERM)

	go processSignal()
	go processExit()
}

func processSignal() {
	s := <-Monitoring.signalChan

	logger.Logger.PrintWarn("signal found", "signal", s.String())

	close(Monitoring.StopChan)
}

func processExit() {
	code := <-Monitoring.ExitChan

	logger.Logger.PrintWarn("program is exiting", "code", code)
	os.Exit(code)
}

func SetMonitoringCounter(ctx *fasthttp.RequestCtx, prefix string, message string) {
	if Monitoring.WRITE_METRICS == true {
		metricPath := string(ctx.Path())
		metricName := strings.Split(string(metricPath), "/")
		name := metricName[len(metricName)-1]
		go HttpMetrics.AddCounterMetric(name+"_"+prefix, name+" "+message)
	}
}
