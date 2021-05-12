package model

type MessageBody struct {
	Text        string      `xml:",chardata"`
	Message     Message     `xml:"rev:Message"`
	MessageData MessageData `xml:"rev:MessageData"`
}

type MessageData struct {
	Text    string      `xml:",chardata"`
	AppData interface{} `xml:"rev:AppData"`
}

type Message struct {
	Text         string    `xml:",chardata"`
	Sender       Sender    `xml:"rev:Sender"`
	Recipient    Recipient `xml:"rev:Recipient"`
	TypeCode     string    `xml:"rev:TypeCode"`
	Status       string    `xml:"rev:Status"`
	Date         string    `xml:"rev:Date"`
	ExchangeType string    `xml:"rev:ExchangeType"`
	TestMsg      string    `xml:"rev:TestMsg"`
}

type Sender struct {
	Text string `xml:",chardata"`
	Code string `xml:"rev:Code"`
	Name string `xml:"rev:Name"`
}

type Recipient struct {
	Text string `xml:",chardata"`
	Code string `xml:"rev:Code"`
	Name string `xml:"rev:Name"`
}
