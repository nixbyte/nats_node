package configs

import "time"

type NatsConfig struct {
	MessagingConf  Messaging  `json:"messaging"`
	SubscriberConf Subscriber `json:"subscriber"`
	PublisherConf  Publisher  `json:"publisher"`
	ConnectConf    Connect    `json:"connect"`
	ReconnectConf  Reconnect  `json:"reconnect"`
	SecurityConf   Security   `json:"security"`
}

type NatsDefaultConfigBuilder struct {
	MessagingConf  Messaging  `json:"messaging"`
	SubscriberConf Subscriber `json:"subscriber"`
	PublisherConf  Publisher  `json:"publisher"`
	ConnectConf    Connect    `json:"connect"`
	ReconnectConf  Reconnect  `json:"reconnect"`
	SecurityConf   Security   `json:"security"`
}

type Messaging struct {
	MsgMod          string `json:"msgMod"`
	MsgWaitInterval int    `json:"msgWaitInterval"`
	MsgWaitDigits   string `json:"msgWaitDigits"`
}

type Subscriber struct {
	ReceiveMsgMod      string `json:"receiveMsgMod"`
	UnsubscribeAfter   int    `json:"unsubscribeAfter"`
	ReplyToSubject     string `json:"replyToSubject"`
	WildcardSubject    string `json:"wildcardSubject"`
	Wildcard           bool   `json:"wildcard"`
	QueueMod           bool   `json:"queueMode"`
	QueueName          string `json:"queueName"`
	IsJSONData         bool   `json:"isJsonData"`
	DrainMod           bool   `json:"drainMod"`
	DrainTimeout       int    `json:"drainTimeout"`
	DrainTimeoutDigits string `json:"drainTimeoutDigits"`
}

type Publisher struct {
	ReplyToSubject string `json:"replyToSubject"`
	IsJSONData     bool   `json:"isJsonData"`
}

type Connect struct {
	ConnectionName          string   `json:"connectionName"`
	ConnectionType          string   `json:"connectionType"`
	Servers                 []string `json:"servers"`
	VerboseMod              bool     `json:"verboseMod"`
	EchoMod                 bool     `json:"echoMod"`
	ConnectionTimeout       int      `json:"connectionTimeout"`
	ConnectionTimeoutDigits string   `json:"connectionTimeoutDigits"`
	PingMod                 bool     `json:"pingMod"`
	PingInterval            int      `json:"pingInterval"`
	PingIntervalDigits      string   `json:"pingIntervalDigits"`
	PingMaxOutstanding      int      `json:"pingMaxOutstanding"`
}

type Reconnect struct {
	ReconnectAttempts     int    `json:"reconnectAttempts"`
	ReconnectWait         int    `json:"reconnectWait"`
	ReconnectWaitDigits   string `json:"reconnectWaitDigits"`
	ListenReconnectEvents bool   `json:"listenReconnectEvents"`
}

type Security struct {
	AuthMod        bool       `json:"authMod"`
	AuthType       string     `json:"authType"`
	CredentialConf Credential `json:"credential"`
}

type Credential struct {
	LoginConf Login  `json:"login"`
	TLSConf   TLS    `json:"tls"`
	Token     string `json:"token"`
	Jwt       string `json:"jwt"`
	Seed      string `json:"seed"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type TLS struct {
	Path               string `json:"path"`
	RootCAFilename     string `json:"rootCAFilename"`
	ClientCertFilename string `json:"ClientCertFilename"`
	ClientKeyFilename  string `json:"ClientKeyFilename"`
}

func GetDigitsFromString(value string) time.Duration {
	switch digit := value; digit {
	case "hour":
		return time.Hour
	case "min":
		return time.Minute
	case "sec":
		return time.Second
	case "mil":
		return time.Millisecond
	default:
		return time.Second
	}
}
