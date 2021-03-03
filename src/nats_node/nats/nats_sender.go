package nats_client

import (
	"nats_node/configs"
	"nats_node/utils/logger"
)

func init() {
	NCConfig := configs.GetDefaultConfig()
	NatsConnect, err := NCConfig.ParseConfig()
	if err != nil {
		logger.Logger.PrintError(err)
	}
}
