package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"nats_node/configs"
	"nats_node/utils/logger"
	"nats_node/utils/monitoring"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var (
	Client *HttpWorker
)

type HttpWorker struct {
	HttpClient *http.Client
	Config     *configs.ClientConfig
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

func init() {
	fmt.Println("Init Client...")

	Client = NewClient(configs.SetDefaultClientConfig())

}

//-------Private functionality

func (t RequestType) string() string {
	return [...]string{"GET", "POST", "PUT", "DELETE"}[t]
}

func prepareRequest(request *HttpRequest) (*http.Request, error) {
	if strings.Compare(request.Endpoint, "") == 0 {
		return nil, errors.New("Parameter endpoint is empty string")
	}

	if strings.Compare(request.Hostname, "") == 0 {
		request.Hostname = Client.Config.DefaultHostName
	}

	var req *http.Request
	var err error

	if strings.IndexByte(request.Endpoint, '/') != 0 {
		var sb strings.Builder
		sb.WriteRune('/')
		sb.WriteString(request.Endpoint)
		request.Endpoint = sb.String()
	}

	fmt.Println("Request Body")
	fmt.Println(request.Body)
	fmt.Println(len(request.BodyMap))

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

func sendRequest(request *HttpRequest) ([]byte, error) {
	req, err := prepareRequest(request)
	if err != nil {
		return nil, err
	}
	logger.Logger.PrintRequestDebug(req)

	resp, err := Client.HttpClient.Do(req)
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
func SendRequest(request *HttpRequest) ([]byte, error) {
	return sendRequest(request)
}

func (request HttpRequest) PrepareMetricName(prefix string) string {
	metricName := request.Endpoint

	if strings.IndexByte(request.Endpoint, '?') != -1 {
		metricName = request.Endpoint[:strings.IndexByte(request.Endpoint, '?')]
	}
	metric := strings.Split(string(metricName), "/")
	name := metric[len(metric)-1]

	return prefix + name
}

func NewRequest() *HttpRequest {
	return &HttpRequest{0, "", "", make(map[string]string), url.Values{}, &bytes.Reader{}, make(map[string]interface{})}
}

func NewClient(config *configs.ClientConfig) *HttpWorker {
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
