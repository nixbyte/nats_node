package client

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	configs "nats_node/configs/http"
	"nats_node/utils/logger"
	"nats_node/utils/monitoring"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"regexp"
	"strings"
	"time"
)

var (
	Client *HttpWorker
)

type HttpWorker struct {
	HttpClient *http.Client
	Config     *configs.NatsNodeHttpClientConfig
}

type RequestType int

const (
	GET RequestType = iota
	POST
	PUT
	DELETE
)

type HttpRequest struct {
	Rt         RequestType
	Hostname   string
	Endpoint   string
	Headers    map[string]string
	Parameters url.Values
	Body       io.Reader
	BodyMap    map[string]interface{}
}

type soapRequest struct {
	XMLName  xml.Name `xml:"x:Envelope"`
	XMLNsX   string   `xml:"xmlns:x,attr"`
	XMLNsTem string   `xml:"xmlns:tem,attr"`
	XMLNsHub string   `xml:"xmlns:hub,attr"`
	Header   soapHeader
	Body     soapBody
}

type soapBody struct {
	XMLName xml.Name `xml:"x:Body"`
	Payload interface{}
}

type soapHeader struct {
	XMLName       xml.Name `xml:"x:Header"`
	Text          string   `xml:",chardata"`
	Authorization string   `xml:"x:Authorization"`
}

func init() {
	fmt.Println("Init Client...")

	Client = NewWorker(configs.SetDefaultClientConfig())

}

//-------Private functionality

func (t RequestType) string() string {
	return [...]string{"GET", "POST", "PUT", "DELETE"}[t]
}

func (worker *HttpWorker) prepareRequest(request *HttpRequest) (*http.Request, error) {
	if strings.Compare(request.Endpoint, "") == 0 {
		return nil, errors.New("Parameter endpoint is empty string")
	}

	if strings.Compare(request.Hostname, "") == 0 {
		request.Hostname = worker.Config.DefaultHostName
	}

	var req *http.Request
	var err error

	if strings.IndexByte(request.Endpoint, '/') != 0 {
		var sb strings.Builder
		sb.WriteRune('/')
		sb.WriteString(request.Endpoint)
		request.Endpoint = sb.String()
	}

	if len(request.BodyMap) > 0 {

		bodyJson, _ := json.Marshal(request.BodyMap)

		request.Body = bytes.NewBuffer(bodyJson)
	}

	switch request.Rt {

	case GET:
		req, err = http.NewRequest(GET.string(), request.Hostname+request.Endpoint, request.Body)

		req.URL.Query().Encode()
		if monitoring.Monitoring.WRITE_METRICS == true {
			metricName := request.PrepareMetricName(GET.string())
			go monitoring.HttpMetrics.AddCounterMetric(metricName, metricName+" request count")
		}

	case POST:
		req, err = http.NewRequest(POST.string(), request.Hostname+request.Endpoint, request.Body)
		if monitoring.Monitoring.WRITE_METRICS == true {
			metricName := request.PrepareMetricName(POST.string())
			go monitoring.HttpMetrics.AddCounterMetric(metricName, metricName+" request count")
		}
	default:
		return nil, errors.New("Unknown Request type")
	}
	setHeaders(req, request.Headers)

	if request.Parameters != nil {
		req.URL.RawQuery = request.Parameters.Encode()
	}

	return req, err
}

func setHeaders(req *http.Request, headers map[string]string) {
	if headers != nil {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}
}

func (worker *HttpWorker) sendRequest(request *HttpRequest) ([]byte, error) {
	req, err := worker.prepareRequest(request)
	if err != nil {
		return nil, err
	}
	logger.Logger.PrintRequestDebug(req)

	resp, err := worker.HttpClient.Do(req)
	if err != nil {
		if monitoring.Monitoring.WRITE_METRICS == true {
			metricName := request.PrepareMetricName("ERROR")
			go monitoring.HttpMetrics.AddCounterMetric(metricName, err.Error())
		}
		return nil, err
	}

	if monitoring.Monitoring.WRITE_METRICS == true {
		var metricName string
		switch resp.StatusCode {
		case 200:
			metricName = request.PrepareMetricName("STATUS_OK")
		case 201:
			metricName = request.PrepareMetricName("STATUS_CREATED")
		case 400:
			metricName = request.PrepareMetricName("STATUS_BAD_REQUEST")
		case 401:
			metricName = request.PrepareMetricName("STATUS_UNAUTORIZED")
		case 403:
			metricName = request.PrepareMetricName("STATUS_FORBIDDEN")
		case 404:
			metricName = request.PrepareMetricName("STATUS_NOT_FOUND")
		case 405:
			metricName = request.PrepareMetricName("STATUS_METHOD_NOT_ALLOWED")
		case 500:
			metricName = request.PrepareMetricName("INTERNAL_SERVER_ERROR")
		default:
			metricName = request.PrepareMetricName("UNKNOWN_STATUS_CODE")
		}
		go monitoring.HttpMetrics.AddCounterMetric(metricName, "Status code counter")
	}

	defer resp.Body.Close()
	logger.Logger.PrintResponseDebug(resp)

	return ioutil.ReadAll(resp.Body)
}

//Method SendRequest sending requests to http server-------
//Paratmeter rt requestType is a type of request can be GET POST PUT DELETE
//Parameter endpoint string is a url where you send request
//Parameter headers is a map[string]string with http headers wich you can send to server. Can be nil
//Parameter parameters is a url.Values interface wich contains map with query parameters. Can be nil
func (worker *HttpWorker) SendRequest(request *HttpRequest) ([]byte, error) {
	return worker.sendRequest(request)
}

func SoapCall(ws string, action string, header string, payloadInterface interface{}) ([]byte, error) {
	v := soapRequest{
		XMLNsX:   "http://schemas.xmlsoap.org/soap/envelope/",
		XMLNsTem: "http://tempuri.org/",
		XMLNsHub: "http://schemas.datacontract.org/2004/07/HubService2",
		Header: soapHeader{
			Authorization: header,
		},
		Body: soapBody{
			Payload: payloadInterface,
		},
	}
	payload, err := xml.MarshalIndent(v, "", "  ")

	timeout := time.Duration(30 * time.Second)

	client := http.Client{
		Timeout: timeout,
	}

	req, err := http.NewRequest("POST", ws, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "text/xml, multipart/related")
	req.Header.Set("SOAPAction", action)
	req.Header.Set("Content-Type", "text/xml; charset=utf-8")

	reqd, err := httputil.DumpRequest(req, true)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(reqd))

	response, err := client.Do(req)
	if err != nil {
		logger.Logger.PrintError(err)
		return nil, err
	}

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(bodyBytes))

	defer response.Body.Close()

	return bodyBytes, nil
}

func SoapCallHandleResponse(ws string, action string, header string, payloadInterface interface{}, result interface{}) ([]byte, error) {
	body, err := SoapCall(ws, action, header, payloadInterface)
	if err != nil {
		return nil, err
	}

	err = xml.Unmarshal(body, result)
	if err != nil {
		return nil, err
	}
	if monitoring.Monitoring.WRITE_METRICS == true {
		metricName := strings.Split(action, "/")
		name := metricName[len(metricName)-1]
		go monitoring.HttpMetrics.AddCounterMetric(name, "request to Hubserver count")
	}

	return body, nil
}

func (request HttpRequest) PrepareMetricName(prefix string) string {
	metricName := request.Endpoint

	if strings.IndexByte(request.Endpoint, '?') != -1 {
		metricName = request.Endpoint[:strings.IndexByte(request.Endpoint, '?')]
	}
	metric := strings.Split(string(metricName), "/")
	name := metric[len(metric)-1]

	reg, err := regexp.Compile("[а-яА-Я]")
	if err != nil {
		log.Fatal(err)
	}

	processedString := reg.ReplaceAllString(name, "")

	return prefix + "_" + processedString
}

func NewRequest() *HttpRequest {
	return &HttpRequest{0, "", "", make(map[string]string), url.Values{}, &bytes.Reader{}, make(map[string]interface{})}
}

func SetWorker(config *configs.NatsNodeHttpClientConfig) *HttpWorker {
	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   time.Duration(config.DealerConnectTimeout) * time.Second,
			KeepAlive: time.Duration(config.DealerKeepAlive) * time.Second,
		}).DialContext,
		MaxIdleConnsPerHost:   config.MaxIdleConnsPerHost,
		MaxIdleConns:          config.MaxIdleConns,
		TLSHandshakeTimeout:   time.Duration(config.TLSHandshakeTimeout) * time.Second,
		ResponseHeaderTimeout: time.Duration(config.ResponseHeaderTimeout) * time.Second,
	}

	httpClient := &http.Client{
		Timeout:   time.Duration(config.Timeout) * time.Second,
		Transport: transport,
	}
	Client = &HttpWorker{
		Config:     config,
		HttpClient: httpClient,
	}
	return Client
}

func NewWorker(config *configs.NatsNodeHttpClientConfig) *HttpWorker {
	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   time.Duration(config.DealerConnectTimeout) * time.Second,
			KeepAlive: time.Duration(config.DealerKeepAlive) * time.Second,
		}).DialContext,
		MaxIdleConnsPerHost:   config.MaxIdleConnsPerHost,
		MaxIdleConns:          config.MaxIdleConns,
		TLSHandshakeTimeout:   time.Duration(config.TLSHandshakeTimeout) * time.Second,
		ResponseHeaderTimeout: time.Duration(config.ResponseHeaderTimeout) * time.Second,
	}

	httpClient := &http.Client{
		Timeout:   time.Duration(config.Timeout) * time.Second,
		Transport: transport,
	}
	return &HttpWorker{
		Config:     config,
		HttpClient: httpClient,
	}
}
