package request

import (
	"encoding/json"
	"nats_node/http/client"
	"nats_node/http/client/model"
	"nats_node/utils/logger"
	"nats_node/utils/monitoring"
)

func GetToken() (*model.Token, error) {

	request := client.NewRequest()

	request.Rt = client.GET
	request.Endpoint = "/authorization/api/token"

	response, err := client.SendRequest(request)

	if err != nil {
		logger.Logger.PrintError(err)
		return nil, err
	}

	var token model.Token

	err = json.Unmarshal(response, &token)

	if err != nil {
		if monitoring.Monitoring.WRITE_METRICS == true {
			metricName := request.PrepareMetricName("UNMARSHAL_ERROR")
			go monitoring.HttpMetrics.AddCounterMetric(metricName, "Counter for JSON UnmarshalFieldError")
		}
		return nil, err
	}
	return &token, err
}
