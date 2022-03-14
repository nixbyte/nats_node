package main

import (
	request "nats_node/http/client/requests"
	"nats_node/utils/monitoring"
)

func init() {
	monitoring.Monitoring.StartMonitoring()
	go func() {
		<-monitoring.Monitoring.StopChan
		//@todo: will be made graceful stop here
		monitoring.Monitoring.ExitChan <- monitoring.ExitCodeInterrupted
	}()
}

func main() {
	go request.GetTest()
	go request.PostTest()
	select {}
}
