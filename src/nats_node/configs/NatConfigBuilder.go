package configs

import (
	"github.com/nats-io/nats.go"
)

type ConnectionType int

const (
	DEFAULT ConnectionType = iota
	ENCODED
)

func (t ConnectionType) String() string {
	return [...]string{"DEFAULT", "ENCODED"}[t]
}

type NatsConnections struct {
	ConType           ConnectionType
	EncodedConnection *nats.EncodedConn
	Connection        *nats.Conn
}

var NatsConf *NatsConfig

type NatsConfigBuilder interface {
	SetMessagingConfig(msgConf Messaging)
	SetSubscriberConfig(supConf Subscriber)
	SetPublisherConfig(pubConf Publisher)
	SetConnectConfig(conConf Connect)
	SetReconnectConfig(reconConf Reconnect)
	SetSecurityConfig(secConf Security)
	Config() (*NatsConfig, error)
}

func init() {
	NatsConf = &NatsConfig{
		ConnectConf: Connect{
			ConnectionName:          "api",
			ConnectionType:          "server",
			Servers:                 []string{"localhost:4222"},
			VerboseMod:              true,
			EchoMod:                 true,
			ConnectionTimeout:       10,
			ConnectionTimeoutDigits: "sec",
			PingMod:                 false,
			PingInterval:            10,
			PingIntervalDigits:      "sec",
			PingMaxOutstanding:      5,
		},
		ReconnectConf: Reconnect{
			ReconnectAttempts:     5,
			ReconnectWait:         60,
			ReconnectWaitDigits:   "sec",
			ListenReconnectEvents: false,
		},
		MessagingConf: Messaging{
			MsgMod:          "sync",
			MsgWaitInterval: 1,
			MsgWaitDigits:   "sec",
		},
		SubscriberConf: Subscriber{
			ReceiveMsgMod:      "sync",
			UnsubscribeAfter:   0,
			ReplyToSubject:     "subject",
			WildcardSubject:    "sub*",
			Wildcard:           false,
			QueueMod:           false,
			QueueName:          "queue",
			IsJSONData:         true,
			DrainMod:           false,
			DrainTimeout:       30,
			DrainTimeoutDigits: "sec",
		},
		PublisherConf: Publisher{
			ReplyToSubject: "subject",
			IsJSONData:     true,
		},
		SecurityConf: Security{
			AuthMod:  false,
			AuthType: "login",
			CredentialConf: Credential{
				LoginConf: Login{
					Username: "testuser",
					Password: "testpassword",
				},
				TLSConf: TLS{
					Path:               "./",
					RootCAFilename:     "RootCA.pem",
					ClientCertFilename: "ClientCer.pem",
					ClientKeyFilename:  "ClientKey.pem",
				},
				Token: "testtoken",
				Jwt:   "jwtstring",
				Seed:  "./seed.txt",
			},
		},
	}
}

func GetDefaultConfig() *NatsConfig {
	return NatsConf
}
