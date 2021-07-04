package request

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"nats_node/http/client"
	soapmodel "nats_node/http/model/soap"
	context "nats_node/nats/model"
	"nats_node/utils/logger"
	"nats_node/utils/monitoring"
	"time"

	"github.com/patrickmn/go-cache"
)

var GUID string

var c *cache.Cache

func init() {
	c = cache.New(12*time.Hour, 24*time.Hour)
	GUID = client.Client.Config.Guid
	fmt.Println("Init Covid Service")
}

func checkHeader(ctx *context.RequestContext) error {

	authorization := []byte(ctx.Headers["Authorization"])

	if len(authorization) == 0 {
		return errors.New("Authorization Header not found or empty")
	} else {
		return nil
	}
}

func checkIfExistParameter(ctx *context.RequestContext, name string) (bool, error) {
	parameter := ctx.QueryArgs[name]
	if len(parameter) == 0 {
		return false, errors.New("Query parameter " + string(parameter) + " not found")
	} else {
		return true, nil
	}
}

func validateParameters(ctx *context.RequestContext, params []string) (bool, error) {
	var buffer bytes.Buffer
	for _, param := range params {
		exist, _ := checkIfExistParameter(ctx, param)
		if exist != true {
			buffer.WriteString(" " + param + " ")
		}
	}

	if len(buffer.Bytes()) != 0 {
		return false, errors.New("Query parameters" + buffer.String() + " not found")
	} else {
		return true, nil
	}
}

func GetCovidToken() {
	sub, err := NatsConnection.Conn.SubscribeSync("GetCovidToken")
	if err != nil {
		logger.Logger.PrintError(err)
	}

	for {
		// Wait for a message
		msg, err := sub.NextMsg(10 * time.Minute)
		if err != nil {
			logger.Logger.PrintError(err)
		} else {

			request := client.NewRequest()
			request.Rt = client.GET
			request.Endpoint = "/authorization/api/token"

			response, err := client.Client.SendRequest(request)
			if err != nil {
				logger.Logger.PrintError(err)

				if monitoring.Monitoring.WRITE_METRICS == true {
					metricName := request.PrepareMetricName("UNMARSHAL_ERROR")
					go monitoring.HttpMetrics.AddCounterMetric(metricName, "Counter for JSON UnmarshalFieldError")
				}
			}

			err = msg.Respond(response)
			if err != nil {
				logger.Logger.PrintError(err)
			}
		}
	}
	defer NatsConnection.Close()
}

func CheckCovidTokenExpiration() {
	sub, err := NatsConnection.Conn.SubscribeSync("GetCovidTokenExpiration")
	if err != nil {
		logger.Logger.PrintError(err)
	}

	for {
		// Wait for a message
		msg, err := sub.NextMsg(10 * time.Minute)
		if err != nil {
			logger.Logger.PrintError(err)
		} else {

			requestBytes, err := GetBytesFromNatsBase64Msg(msg.Data)
			if err != nil {
				logger.Logger.PrintError(err)
			}

			context, err := GetRequestContextFromBytesArray(requestBytes)
			if err != nil {
				logger.Logger.PrintError(err)
			}

			request := client.NewRequest()
			request.Rt = client.GET
			request.Endpoint = "/authorization/api/session"
			request.Parameters.Add("token", context.QueryArgs["token"])

			response, err := client.Client.SendRequest(request)
			if err != nil {
				logger.Logger.PrintError(err)
				if monitoring.Monitoring.WRITE_METRICS == true {
					metricName := request.PrepareMetricName("UNMARSHAL_ERROR")
					go monitoring.HttpMetrics.AddCounterMetric(metricName, "Counter for JSON UnmarshalFieldError")
				}
			}

			err = msg.Respond(response)
			if err != nil {
				logger.Logger.PrintError(err)
			}
		}
	}
	defer NatsConnection.Close()

}

func GetDistrictList() {

	sub, err := NatsConnection.Conn.SubscribeSync("GetCovidDistrictsList")
	if err != nil {
		logger.Logger.PrintError(err)
	}

	for {
		// Wait for a message
		msg, err := sub.NextMsg(10 * time.Minute)
		if err != nil {
			logger.Logger.PrintError(err)
		} else {

			requestBytes, err := GetBytesFromNatsBase64Msg(msg.Data)
			if err != nil {
				logger.Logger.PrintError(err)
			}

			context, err := GetRequestContextFromBytesArray(requestBytes)
			if err != nil {
				logger.Logger.PrintError(err)
			}

			pa := soapmodel.SoapDistrictRequest{
				Guid: GUID,
			}

			var respBytes []byte
			var response *soapmodel.SoapDistrictListResponse = &soapmodel.SoapDistrictListResponse{}
			authorization := context.Headers["Authorization"]

			if value, found := c.Get("districts"); found {
				response = value.(*soapmodel.SoapDistrictListResponse)
				respBytes, err = xml.Marshal(response)
				if err != nil {
					logger.Logger.PrintError(err)
				}
			} else {

				respBytes, err = client.SoapCallHandleResponse("http://r78-rc.zdrav.netrika.ru/hub25/HubService.svc", "http://tempuri.org/IHubService/GetDistrictList", authorization, pa, response)
				if err != nil {
					logger.Logger.PrintError(err)
				} else {
					if response.Body.GetDistrictListResponse.GetDistrictListResult.Success == "true" {
						c.Set("districts", response, cache.NoExpiration)
					}
				}
			}

			err = msg.Respond(respBytes)
			if err != nil {
				logger.Logger.PrintError(err)
			}
		}
	}
	defer NatsConnection.Close()
}

func GetDistricts() {
	sub, err := NatsConnection.Conn.SubscribeSync("GetCovidDistrictStringsList")
	if err != nil {
		logger.Logger.PrintError(err)
	}

	for {
		// Wait for a message
		msg, err := sub.NextMsg(10 * time.Minute)
		if err != nil {
			logger.Logger.PrintError(err)
		} else {

			requestBytes, err := GetBytesFromNatsBase64Msg(msg.Data)
			if err != nil {
				logger.Logger.PrintError(err)
			}

			context, err := GetRequestContextFromBytesArray(requestBytes)
			if err != nil {
				logger.Logger.PrintError(err)
			}

			pa := soapmodel.SoapDistrictRequest{
				Guid: GUID,
			}

			var response *soapmodel.SoapDistrictListResponse = &soapmodel.SoapDistrictListResponse{}
			authorization := context.Headers["Authorization"]

			if value, found := c.Get("districts"); found {
				response = value.(*soapmodel.SoapDistrictListResponse)
			} else {

				_, err = client.SoapCallHandleResponse("http://r78-rc.zdrav.netrika.ru/hub25/HubService.svc", "http://tempuri.org/IHubService/GetDistrictList", authorization, pa, response)
				if err != nil {
					logger.Logger.PrintError(err)
				} else {
					if response.Body.GetDistrictListResponse.GetDistrictListResult.Success == "true" {
						c.Set("districts", response, cache.NoExpiration)
					}
				}
			}

			districts := response.Body.GetDistrictListResponse.GetDistrictListResult.Districts.District
			districtStrings := []string{}
			for _, district := range districts {
				districtStrings = append(districtStrings, district.DistrictName)
			}

			bytes, err := json.Marshal(districtStrings)

			err = msg.Respond(bytes)
			if err != nil {
				logger.Logger.PrintError(err)
			}
		}
	}
	defer NatsConnection.Close()
}

func GetLpuList() {

	sub, err := NatsConnection.Conn.SubscribeSync("GetLpuList")
	if err != nil {
		logger.Logger.PrintError(err)
	}

	for {
		// Wait for a message
		msg, err := sub.NextMsg(10 * time.Minute)
		if err != nil {
			logger.Logger.PrintError(err)
		} else {

			requestBytes, err := GetBytesFromNatsBase64Msg(msg.Data)
			if err != nil {
				logger.Logger.PrintError(err)
			}

			context, err := GetRequestContextFromBytesArray(requestBytes)
			if err != nil {
				logger.Logger.PrintError(err)
			}

			err = checkHeader(context)
			if err != nil {
				logger.Logger.PrintError(err)
			}

			valid, err := validateParameters(context, []string{"idDistrict"})
			if err != nil {
				logger.Logger.PrintError(err)
			}
			var respBytes []byte

			if valid == true {
				pa := soapmodel.SoapLpuListRequest{
					IdDistrict: context.QueryArgs["idDistrict"],
					Guid:       GUID,
				}

				var response *soapmodel.SoapLpuListResponse = &soapmodel.SoapLpuListResponse{}
				authorization := context.Headers["Authorization"]

				if value, found := c.Get(pa.IdDistrict); found {
					response = value.(*soapmodel.SoapLpuListResponse)
					respBytes, err = xml.Marshal(response)
					if err != nil {
						logger.Logger.PrintError(err)
					}
				} else {

					respBytes, err = client.SoapCallHandleResponse("http://r78-rc.zdrav.netrika.ru/hub25/HubService.svc", "http://tempuri.org/IHubService/GetLPUList", authorization, pa, response)
					if err != nil {
						logger.Logger.PrintError(err)
					} else {
						if response.Body.GetLPUListResponse.GetLPUListResult.Success == "true" {
							c.Set(pa.IdDistrict, response, cache.DefaultExpiration)
						}
					}
				}
			}
			err = msg.Respond(respBytes)
			if err != nil {
				logger.Logger.PrintError(err)
			}
		}
	}
	defer NatsConnection.Close()

}
func GetCovidLpuList() {

	sub, err := NatsConnection.Conn.SubscribeSync("GetCovidLpuList")
	if err != nil {
		logger.Logger.PrintError(err)
	}

	for {
		// Wait for a message
		msg, err := sub.NextMsg(10 * time.Minute)
		if err != nil {
			logger.Logger.PrintError(err)
		} else {

			requestBytes, err := GetBytesFromNatsBase64Msg(msg.Data)
			if err != nil {
				logger.Logger.PrintError(err)
			}

			context, err := GetRequestContextFromBytesArray(requestBytes)
			if err != nil {
				logger.Logger.PrintError(err)
			}

			err = checkHeader(context)
			if err != nil {
				logger.Logger.PrintError(err)
			}

			var respBytes []byte

			pa := soapmodel.SoapCovidLpuListRequest{
				Guid: GUID,
			}

			var response *soapmodel.SoapCovidLpuListResponse = &soapmodel.SoapCovidLpuListResponse{}
			authorization := context.Headers["Authorization"]

			if value, found := c.Get("lpus"); found {
				response = value.(*soapmodel.SoapCovidLpuListResponse)
				respBytes, err = xml.Marshal(response)
				if err != nil {
					logger.Logger.PrintError(err)
				}
			} else {

				respBytes, err = client.SoapCallHandleResponse("http://r78-rc.zdrav.netrika.ru/hub25/CovidLpuService.svc", "http://tempuri.org/ICovidLpuService/GetLpus", authorization, pa, response)
				if err != nil {
					logger.Logger.PrintError(err)
				} else {
					if response.Body.GetLpusResponse.GetLpusResult.Success == "true" {
						lpus := response.Body.GetLpusResponse.GetLpusResult.Lpus.Lpu

						for i, item := range lpus {
							if item.CountOfAvailableCovidAppointments == "0" {
								response.Body.GetLpusResponse.GetLpusResult.Lpus.Lpu = removeEmptyLpus(lpus, i)
							}
						}
						c.Set("lpus", response, cache.DefaultExpiration)
					}
				}
			}
			err = msg.Respond(respBytes)
			if err != nil {
				logger.Logger.PrintError(err)
			}
		}
	}
	defer NatsConnection.Close()

}

func removeEmptyLpus(s []soapmodel.Lpu, i int) []soapmodel.Lpu {
	return append(s[:i], s[i+1:]...)
}

func GetSpecialityList()        {}
func GetDoctorList()            {}
func GetAppointmentList()       {}
func CheckPatient()             {}
func AddPatient()               {}
func UpdatePhone()              {}
func SetAppointment()           {}
func DeleteAppointment()        {}
func GetCovidAppointmentCount() {}

func GetCovidLpuNames()         {}
func GetCovidLpuIds()           {}
func GetCovidLpuIdByName()      {}
func GetCovidSpecialityNames()  {}
func GetCovidSpecialityIds()    {}
func GetCovidSpecialityId()     {}
func GetCovidDocNames()         {}
func GetCovidDocIds()           {}
func GetCovidDocId()            {}
func GetCovidAppointmentTimes() {}
func GetCovidAppointmentIds()   {}
