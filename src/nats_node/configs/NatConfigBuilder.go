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

type NatsConfigBuilder interface {
	SetConnectConfig(conConf Connect)
	SetMessagingConfig(msgConf Messaging)
	SetSubscriberConfig(supConf Subscriber)
	SetPublisherConfig(pubConf Publisher)
	SetReconnectConfig(reconConf Reconnect)
	SetSecurityConfig(secConf Security)
	Config() (*NatsConfig, error)
}

func (c *NatsDefaultConfigBuilder) SetConnectConfig(conConf Connect) {
	c.ConnectConf = conConf
}

func (c *NatsDefaultConfigBuilder) SetMessagingConfig(msgConf Messaging) {
	c.MessagingConf = msgConf
}

func (c *NatsDefaultConfigBuilder) SetSubscriberConfig(subConf Subscriber) {
	c.SubscriberConf = subConf
}

func (c *NatsDefaultConfigBuilder) SetPublisherConfig(pubConf Publisher) {
	c.PublisherConf = pubConf
}

func (c *NatsDefaultConfigBuilder) SetReconnectConfig(reconConf Reconnect) {
	c.ReconnectConf = reconConf
}

func (c *NatsDefaultConfigBuilder) SetSecurityConfig(secConf Security) {
	c.SecurityConf = secConf
}

func (c *NatsDefaultConfigBuilder) Config() (*NatsConfig, error) {
	NatsConf := &NatsConfig{
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

	c.ConnectConf.parceConnectionConfig(&NatsConf.ConnectConf)

	return NatsConf, nil
}

func (b *Connect) parceConnectionConfig(c *Connect) {
	if len(b.ConnectionName) != 0 {
		c.ConnectionName = b.ConnectionName
	}

	if len(b.ConnectionType) != 0 {
		c.ConnectionType = b.ConnectionType
	}

	if len(b.ConnectionTimeoutDigits) != 0 {
		c.ConnectionTimeoutDigits = b.ConnectionTimeoutDigits
	}

	if b.ConnectionTimeout != 0 {
		c.ConnectionTimeout = b.ConnectionTimeout
	}

	if len(b.Servers) != 0 {
		if len(b.Servers[0]) != 0 {
			c.Servers = b.Servers
		}
	}

	if len(b.PingIntervalDigits) != 0 && b.PingMod == true {
		c.PingIntervalDigits = b.PingIntervalDigits
		c.PingMod = true
	} else {
		c.PingMod = false
	}

	if b.PingInterval != 0 && b.PingMod == true {
		c.PingInterval = b.PingInterval
		c.PingMod = true
	} else {
		c.PingMod = false
	}

	if b.PingMaxOutstanding != 0 && b.PingMod == true {
		c.PingMaxOutstanding = b.PingMaxOutstanding
		c.PingMod = true
	} else {
		c.PingMod = false
	}

	c.EchoMod = b.EchoMod
	c.VerboseMod = b.VerboseMod
}
