package request

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"nats_node/http/client"
	model "nats_node/http/model/json"
	soapmodel "nats_node/http/model/soap"
	context "nats_node/nats/model"
	"nats_node/utils/logger"
	"nats_node/utils/monitoring"
	"strconv"
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

func GetSpecialityList() {

	sub, err := NatsConnection.Conn.SubscribeSync("GetCovidSpecialityList")
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

			var validParameters bool
			validParameters, err = validateParameters(context, []string{"idLpu"})
			if err != nil {
				logger.Logger.PrintError(err)
			}

			var respBytes []byte

			if validParameters == true {

				pa := soapmodel.SoapSpecialityListRequest{
					IdLpu: context.QueryArgs["idLpu"],
					Guid:  GUID,
				}

				var response *soapmodel.SoapSpecialityListResponse = &soapmodel.SoapSpecialityListResponse{}
				authorization := context.Headers["Authorization"]

				if value, found := c.Get(pa.IdLpu); found {
					response = value.(*soapmodel.SoapSpecialityListResponse)
					respBytes, err = xml.Marshal(response)
					if err != nil {
						logger.Logger.PrintError(err)
					}
				} else {

					respBytes, err = client.SoapCallHandleResponse("http://r78-rc.zdrav.netrika.ru/hub25/HubService.svc", "http://tempuri.org/IHubService/GetSpesialityList", authorization, pa, response)

					if err != nil {
						logger.Logger.PrintError(err)
					} else {
						if response.Body.GetSpesialityListResponse.GetSpesialityListResult.Success == "true" {
							c.Set(pa.IdLpu, response, cache.DefaultExpiration)
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
func GetDoctorList() {

	sub, err := NatsConnection.Conn.SubscribeSync("GetCovidDoctorList")
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

			var validParameters bool
			validParameters, err = validateParameters(context, []string{"idSpeciality", "idLpu"})
			if err != nil {
				logger.Logger.PrintError(err)
			}

			var respBytes []byte

			if validParameters == true {

				pa := soapmodel.SoapDoctorListRequest{
					IdLpu:        context.QueryArgs["idLpu"],
					IdSpesiality: context.QueryArgs["idSpeciality"],
					Guid:         GUID,
				}

				var response *soapmodel.SoapDoctorListResponse = &soapmodel.SoapDoctorListResponse{}
				authorization := context.Headers["Authorization"]

				if value, found := c.Get(pa.IdLpu + pa.IdSpesiality); found {
					response = value.(*soapmodel.SoapDoctorListResponse)
					respBytes, err = xml.Marshal(response)
					if err != nil {
						logger.Logger.PrintError(err)
					}
				} else {

					respBytes, err = client.SoapCallHandleResponse("http://r78-rc.zdrav.netrika.ru/hub25/HubService.svc", "http://tempuri.org/IHubService/GetDoctorList", authorization, pa, response)

					if err != nil {
						logger.Logger.PrintError(err)
					} else {
						if response.Body.GetDoctorListResponse.GetDoctorListResult.Success == "true" {
							c.Set(pa.IdLpu+pa.IdSpesiality, response, 10*time.Minute)
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
func GetAppointmentList() {
	sub, err := NatsConnection.Conn.SubscribeSync("GetCovidAppointmentList")
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

			var validParameters bool
			validParameters, err = validateParameters(context, []string{"idDoc", "idLpu"})
			if err != nil {
				logger.Logger.PrintError(err)
			}

			var respBytes []byte

			if validParameters == true {

				pa := soapmodel.SoapAppointmentListRequest{
					IdDoc:      context.QueryArgs["idDoc"],
					IdLpu:      context.QueryArgs["idLpu"],
					VisitStart: time.Now().Format("2006-01-02T15:04:05"),
					VisitEnd:   time.Now().Format("2006-01-02T15:04:05"),
					Guid:       GUID,
				}

				var response *soapmodel.SoapAppointmentListResponse = &soapmodel.SoapAppointmentListResponse{}
				authorization := context.Headers["Authorization"]

				if value, found := c.Get(pa.IdDoc + pa.IdLpu); found {
					response = value.(*soapmodel.SoapAppointmentListResponse)
					respBytes, err = xml.Marshal(response)
					if err != nil {
						logger.Logger.PrintError(err)
					}
				} else {

					respBytes, err = client.SoapCallHandleResponse("http://r78-rc.zdrav.netrika.ru/hub25/HubService.svc", "http://tempuri.org/IHubService/GetAvaibleAppointments", authorization, pa, response)

					if err != nil {
						logger.Logger.PrintError(err)
					} else {
						if response.Body.GetAvaibleAppointmentsResponse.GetAvaibleAppointmentsResult.Success == "true" {
							c.Set(pa.IdDoc+pa.IdLpu, response, 1*time.Hour)
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
func CheckPatient() {
	sub, err := NatsConnection.Conn.SubscribeSync("CheckPatient")
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

			var response *soapmodel.SoapCheckPatientResponse = &soapmodel.SoapCheckPatientResponse{}
			var person *model.Patient = &model.Patient{
				Firstname:     "",
				Lastname:      "",
				Middlename:    "",
				Phone:         "",
				Birthday:      0,
				IdLpu:         "",
				IdPat:         "",
				IdAppointment: "",
				Snils:         "",
				DocumentN:     "",
				DocumentS:     "",
				PolisN:        "",
				PolisS:        "",
			}

			err = json.Unmarshal(context.Body, person)
			if err != nil {
				logger.Logger.PrintError(err)
			}
			var respBytes []byte
			validPatient, err := validatePatient(*person)

			if validPatient == true {

				soapPat := soapmodel.SoapPatient{
					Birthday:   time.Unix(person.Birthday, 0).UTC().Format("2006-01-02T15:04:05"),
					Name:       person.Firstname,
					SecondName: person.Lastname,
					Surname:    person.Middlename,
					Snils:      person.Snils,
					Document_N: person.DocumentN,
					Document_S: person.DocumentS,
					Polis_N:    person.PolisN,
					Polis_S:    person.PolisS,
				}

				pa := soapmodel.SoapCheckPatientRequest{
					Pat:   soapPat,
					IdLpu: person.IdLpu,
					Guid:  GUID,
				}

				authorization := context.Headers["Authorization"]

				respBytes, err = client.SoapCallHandleResponse("http://r78-rc.zdrav.netrika.ru/hub25/HubService.svc", "http://tempuri.org/IHubService/CheckPatient", authorization, pa, response)
				if err != nil {
					logger.Logger.PrintError(err)
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
func AddPatient() {
	sub, err := NatsConnection.Conn.SubscribeSync("AddPatient")
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

			var response *soapmodel.SoapAddPatientResponse = &soapmodel.SoapAddPatientResponse{}
			var person *model.Patient = &model.Patient{
				Firstname:     "",
				Lastname:      "",
				Middlename:    "",
				Phone:         "",
				Birthday:      0,
				IdLpu:         "",
				IdPat:         "",
				IdAppointment: "",
				Snils:         "",
				DocumentN:     "",
				DocumentS:     "",
				PolisN:        "",
				PolisS:        "",
			}

			err = json.Unmarshal(context.Body, person)
			if err != nil {
				logger.Logger.PrintError(err)
			}
			var respBytes []byte
			validPatient, err := validatePatient(*person)

			if validPatient == true {

				soapPat := soapmodel.SoapPatient{
					Birthday:   time.Unix(person.Birthday, 0).UTC().Format("2006-01-02T15:04:05"),
					Name:       person.Firstname,
					SecondName: person.Lastname,
					Surname:    person.Middlename,
					Snils:      person.Snils,
					Document_N: person.DocumentN,
					Document_S: person.DocumentS,
					Polis_N:    person.PolisN,
					Polis_S:    person.PolisS,
					CellPhone:  person.Phone,
					HomePhone:  person.Phone,
				}

				pa := soapmodel.SoapAddPatientRequest{
					Pat:   soapPat,
					IdLpu: person.IdLpu,
					Guid:  GUID,
				}

				authorization := context.Headers["Authorization"]

				respBytes, err = client.SoapCallHandleResponse("http://r78-rc.zdrav.netrika.ru/hub25/HubService.svc", "http://tempuri.org/IHubService/AddNewPatient", authorization, pa, response)
				if err != nil {
					logger.Logger.PrintError(err)
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

func UpdatePhone() {
	sub, err := NatsConnection.Conn.SubscribeSync("UpdatePhone")
	if err != nil {
		logger.Logger.PrintError(err)
	}

	for {
		// Wait for a message
		msg, err := sub.NextMsg(10 * time.Minute)
		if err != nil {
			logger.Logger.PrintError(err)
		} else {
			var respBytes []byte
			var response *soapmodel.SoapUpdatePhoneResponse = &soapmodel.SoapUpdatePhoneResponse{}
			var person *model.Patient = &model.Patient{
				Firstname:     "",
				Lastname:      "",
				Middlename:    "",
				Phone:         "",
				Birthday:      0,
				IdLpu:         "",
				IdPat:         "",
				IdAppointment: "",
			}

			requestBytes, err := GetBytesFromNatsBase64Msg(msg.Data)
			if err != nil {
				logger.Logger.PrintError(err)
			}

			context, err := GetRequestContextFromBytesArray(requestBytes)
			if err != nil {
				logger.Logger.PrintError(err)
			}

			err = json.Unmarshal(context.Body, person)
			if err != nil {
				logger.Logger.PrintError(err)
			}

			pa := soapmodel.SoapUpdatePhoneRequest{
				IdLpu:     person.IdLpu,
				IdPat:     person.IdPat,
				HomePhone: person.Phone,
				CellPhone: person.Phone,
				Guid:      GUID,
			}

			authorization := context.Headers["Authorization"]

			respBytes, err = client.SoapCallHandleResponse("http://r78-rc.zdrav.netrika.ru/hub25/HubService.svc", "http://tempuri.org/IHubService/UpdatePhoneByIdPat", authorization, pa, response)
			if err != nil {
				logger.Logger.PrintError(err)
			}
			err = msg.Respond(respBytes)
			if err != nil {
				logger.Logger.PrintError(err)
			}
		}
	}
	defer NatsConnection.Close()

}
func SetAppointment() {
	sub, err := NatsConnection.Conn.SubscribeSync("SetAppointment")
	if err != nil {
		logger.Logger.PrintError(err)
	}

	for {
		// Wait for a message
		msg, err := sub.NextMsg(10 * time.Minute)
		if err != nil {
			logger.Logger.PrintError(err)
		} else {
			var respBytes []byte
			var response *soapmodel.SoapSetAppointmentResponse = &soapmodel.SoapSetAppointmentResponse{}
			var person *model.Patient = &model.Patient{
				Firstname:     "",
				Lastname:      "",
				Middlename:    "",
				Phone:         "",
				Birthday:      0,
				IdLpu:         "",
				IdPat:         "",
				IdAppointment: "",
			}

			requestBytes, err := GetBytesFromNatsBase64Msg(msg.Data)
			if err != nil {
				logger.Logger.PrintError(err)
			}

			context, err := GetRequestContextFromBytesArray(requestBytes)
			if err != nil {
				logger.Logger.PrintError(err)
			}

			err = json.Unmarshal(context.Body, person)
			if err != nil {
				logger.Logger.PrintError(err)
			}

			pa := soapmodel.SoapSetAppointmentRequest{
				IdAppointment: person.IdAppointment,
				IdLpu:         person.IdLpu,
				IdPat:         person.IdPat,
				Guid:          GUID,
			}

			authorization := context.Headers["Authorization"]

			respBytes, err = client.SoapCallHandleResponse("http://r78-rc.zdrav.netrika.ru/hub25/HubService.svc", "http://tempuri.org/IHubService/SetAppointment", authorization, pa, response)
			if err != nil {
				logger.Logger.PrintError(err)
			}
			err = msg.Respond(respBytes)
			if err != nil {
				logger.Logger.PrintError(err)
			}
		}
	}
	defer NatsConnection.Close()
}
func DeleteAppointment() {
	sub, err := NatsConnection.Conn.SubscribeSync("DeleteAppointment")
	if err != nil {
		logger.Logger.PrintError(err)
	}

	for {
		// Wait for a message
		msg, err := sub.NextMsg(10 * time.Minute)
		if err != nil {
			logger.Logger.PrintError(err)
		} else {
			var respBytes []byte
			var response *soapmodel.SoapDeleteAppointmentResponse = &soapmodel.SoapDeleteAppointmentResponse{}
			var person *model.Patient = &model.Patient{
				Firstname:     "",
				Lastname:      "",
				Middlename:    "",
				Phone:         "",
				Birthday:      0,
				IdLpu:         "",
				IdPat:         "",
				IdAppointment: "",
			}

			requestBytes, err := GetBytesFromNatsBase64Msg(msg.Data)
			if err != nil {
				logger.Logger.PrintError(err)
			}

			context, err := GetRequestContextFromBytesArray(requestBytes)
			if err != nil {
				logger.Logger.PrintError(err)
			}

			err = json.Unmarshal(context.Body, person)
			if err != nil {
				logger.Logger.PrintError(err)
			}

			pa := soapmodel.SoapDeleteAppointmentRequest{
				IdAppointment: person.IdAppointment,
				IdLpu:         person.IdLpu,
				IdPat:         person.IdPat,
				Guid:          GUID,
			}

			authorization := context.Headers["Authorization"]

			respBytes, err = client.SoapCallHandleResponse("http://r78-rc.zdrav.netrika.ru/hub25/HubService.svc", "http://tempuri.org/IHubService/CreateClaimForRefusal", authorization, pa, response)
			if err != nil {
				logger.Logger.PrintError(err)
			}
			err = msg.Respond(respBytes)
			if err != nil {
				logger.Logger.PrintError(err)
			}
		}
	}
	defer NatsConnection.Close()
}

func GetCovidAppointmentCount() {
	sub, err := NatsConnection.Conn.SubscribeSync("GetCovidAppointmentCount")
	if err != nil {
		logger.Logger.PrintError(err)
	}

	for {
		// Wait for a message
		msg, err := sub.NextMsg(10 * time.Minute)
		if err != nil {
			logger.Logger.PrintError(err)
		} else {
			lpuList := &soapmodel.SoapCovidLpuListResponse{}
			districtList := &soapmodel.SoapDistrictListResponse{}
			response := &model.CovidAppointmentCountResponse{}

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

			if value, found := c.Get("tickets"); found {
				response = value.(*model.CovidAppointmentCountResponse)
			} else {
				districtList, err = getDistrictList(context)
				if err != nil {
					logger.Logger.PrintError(err)
				}

				if districtList.Body.GetDistrictListResponse.GetDistrictListResult.Success == "true" {
					districts := districtList.Body.GetDistrictListResponse.GetDistrictListResult.Districts.District
					response.Districts = []model.CovidAppointmentCount{}
					for _, district := range districts {
						response.Districts = append(response.Districts, model.CovidAppointmentCount{
							DistrictName:     district.DistrictName,
							DistrictId:       district.IdDistrict,
							AppointmentCount: 0,
						})
					}
					lpuList, err = getCovidLpuList(context)

					if err != nil {
						logger.Logger.PrintError(err)
					}

					if lpuList.Body.GetLpusResponse.GetLpusResult.Success == "true" {
						lpus := lpuList.Body.GetLpusResponse.GetLpusResult.Lpus.Lpu
						for i := range districts {
							for _, lpu := range lpus {
								if response.Districts[i].DistrictId == lpu.DistrictId {
									count, err := strconv.Atoi(lpu.CountOfAvailableCovidAppointments)
									if err == nil {
										response.Districts[i].AppointmentCount = response.Districts[i].AppointmentCount + count
									}
								}
							}
						}
						c.Set("tickets", response, 1*time.Hour)
					}
				}
			}
			var respBytes []byte
			respBytes, err = json.Marshal(response)
			if err != nil {
				logger.Logger.PrintError(err)
			}
			err = msg.Respond(respBytes)
			if err != nil {
				logger.Logger.PrintError(err)
			}
		}
	}
}

func getDistrictList(ctx *context.RequestContext) (*soapmodel.SoapDistrictListResponse, error) {

	pa := soapmodel.SoapDistrictRequest{
		Guid: GUID,
	}

	var response *soapmodel.SoapDistrictListResponse = &soapmodel.SoapDistrictListResponse{}

	err := checkHeader(ctx)
	if err != nil {
		return nil, err
	} else {

		authorization := ctx.Headers["Authorization"]

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
	}
	return response, err
}
func getCovidLpuList(ctx *context.RequestContext) (*soapmodel.SoapCovidLpuListResponse, error) {
	pa := soapmodel.SoapCovidLpuListRequest{
		Guid: GUID,
	}

	var response *soapmodel.SoapCovidLpuListResponse = &soapmodel.SoapCovidLpuListResponse{}

	err := checkHeader(ctx)
	if err != nil {
		return response, err
	}

	authorization := ctx.Headers["Authorization"]

	if value, found := c.Get("lpus"); found {
		response = value.(*soapmodel.SoapCovidLpuListResponse)
	} else {

		_, err = client.SoapCallHandleResponse("http://r78-rc.zdrav.netrika.ru/hub25/CovidLpuService.svc", "http://tempuri.org/ICovidLpuService/GetLpus", authorization, pa, response)
		if err != nil {
			logger.Logger.PrintError(err)
		} else {
			if response.Body.GetLpusResponse.GetLpusResult.Success == "true" {
				lpus := response.Body.GetLpusResponse.GetLpusResult.Lpus.Lpu

				for i, item := range lpus {
					if item.CountOfAvailableCovidAppointments == "0" {
						fmt.Println("remove")
						response.Body.GetLpusResponse.GetLpusResult.Lpus.Lpu = removeEmptyLpus(lpus, i)
					}
				}
				c.Set("lpus", response, cache.DefaultExpiration)
			}
		}
	}

	return response, err
}

func GetCovidLpuNames() {
	sub, err := NatsConnection.Conn.SubscribeSync("GetCovidLpuNames")
	if err != nil {
		logger.Logger.PrintError(err)
	}

	for {
		// Wait for a message
		msg, err := sub.NextMsg(10 * time.Minute)
		if err != nil {
			logger.Logger.PrintError(err)
		} else {
			names := &model.CovidLpuNames{
				Names: []string{},
			}

			requestBytes, err := GetBytesFromNatsBase64Msg(msg.Data)
			if err != nil {
				logger.Logger.PrintError(err)
			}

			context, err := GetRequestContextFromBytesArray(requestBytes)
			if err != nil {
				logger.Logger.PrintError(err)
			}

			response, err := getCovidLpuList(context)

			if err != nil {
				logger.Logger.PrintError(err)
			}

			_, err = validateParameters(context, []string{"districtName"})
			if err != nil {
				logger.Logger.PrintError(err)
			}

			if response.Body.GetLpusResponse.GetLpusResult.Success == "true" {
				for _, item := range response.Body.GetLpusResponse.GetLpusResult.Lpus.Lpu {
					if item.DistrictName == context.QueryArgs["districtName"] {
						names.Names = append(names.Names, item.ShortName)
					}
				}
			}
			var respBytes []byte
			respBytes, err = json.Marshal(names)
			if err != nil {
				logger.Logger.PrintError(err)
			}
			err = msg.Respond(respBytes)
			if err != nil {
				logger.Logger.PrintError(err)
			}
		}
	}
	defer NatsConnection.Close()
}
func GetCovidLpuIds() {
	sub, err := NatsConnection.Conn.SubscribeSync("GetCovidLpuIds")
	if err != nil {
		logger.Logger.PrintError(err)
	}

	for {
		// Wait for a message
		msg, err := sub.NextMsg(10 * time.Minute)
		if err != nil {
			logger.Logger.PrintError(err)
		} else {
			ids := &model.CovidLpuIds{
				Ids: []string{},
			}

			requestBytes, err := GetBytesFromNatsBase64Msg(msg.Data)
			if err != nil {
				logger.Logger.PrintError(err)
			}

			context, err := GetRequestContextFromBytesArray(requestBytes)
			if err != nil {
				logger.Logger.PrintError(err)
			}

			response, err := getCovidLpuList(context)

			if err != nil {
				logger.Logger.PrintError(err)
			}

			_, err = validateParameters(context, []string{"districtId"})
			if err != nil {
				logger.Logger.PrintError(err)
			}

			if response.Body.GetLpusResponse.GetLpusResult.Success == "true" {
				for _, item := range response.Body.GetLpusResponse.GetLpusResult.Lpus.Lpu {
					if item.DistrictId == context.QueryArgs["districtId"] {
						ids.Ids = append(ids.Ids, item.ID)
					}
				}
			}
			var respBytes []byte
			respBytes, err = json.Marshal(ids)
			if err != nil {
				logger.Logger.PrintError(err)
			}
			err = msg.Respond(respBytes)
			if err != nil {
				logger.Logger.PrintError(err)
			}
		}
	}
	defer NatsConnection.Close()
}
func GetCovidLpuIdByName()      {}
func GetCovidSpecialityNames()  {}
func GetCovidSpecialityIds()    {}
func GetCovidSpecialityId()     {}
func GetCovidDocNames()         {}
func GetCovidDocIds()           {}
func GetCovidDocId()            {}
func GetCovidAppointmentTimes() {}
func GetCovidAppointmentIds()   {}
