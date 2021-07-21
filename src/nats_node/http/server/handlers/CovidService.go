package handlers

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	model "nats_node/http/model/json"
	soapmodel "nats_node/http/model/soap"
	context "nats_node/nats/model"
	"nats_node/utils/logger"
	"time"

	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

var CovidSwaggerHandler fasthttp.RequestHandler = fasthttpadaptor.NewFastHTTPHandler(
	httpSwagger.Handler(
		httpSwagger.URL("/covid.json"),
	))

var TokenHandler fasthttp.RequestHandler = func(ctx *fasthttp.RequestCtx) {
	defer CatchPanic(ctx)

	var token model.Token

	if ctx.IsPost() == true {
		err = errors.New("method POST not supported")
	} else {

		rc := new(context.RequestContext)
		rc.New(ctx)

		bytes, err := requestContextToBytesArray(rc)
		if err != nil {
			logger.Logger.PrintError(err)
		}
		err = NatsConnection.Request("GetCovidToken", bytes, &token, 10*time.Minute)
		if err != nil {
			logger.Logger.PrintError(err)
		}

	}
	sendModelIfExist(ctx, token, err)
}

var TokenExpirationHandler fasthttp.RequestHandler = func(ctx *fasthttp.RequestCtx) {
	defer CatchPanic(ctx)

	var tokenExpiration model.TokenExpiration
	var validParameters bool
	var err error

	if ctx.IsPost() == true {
		err = errors.New("method POST not supported")
	} else {

		validParameters, err = validateParameters(ctx, []string{"token"})

		if validParameters == true {
			rc := new(context.RequestContext)
			rc.New(ctx)
			bytes, err := requestContextToBytesArray(rc)
			if err != nil {
				logger.Logger.PrintError(err)
			}
			err = NatsConnection.Request("GetCovidTokenExpiration", bytes, &tokenExpiration, 10*time.Minute)
			if err != nil {
				logger.Logger.PrintError(err)
			}
		}
	}
	sendModelIfExist(ctx, tokenExpiration, err)
}

var DistrictListHandler fasthttp.RequestHandler = func(ctx *fasthttp.RequestCtx) {
	defer CatchPanic(ctx)

	var state string
	var districts soapmodel.SoapDistrictListResponse
	var err error
	var validHeader bool

	if ctx.IsPost() == true {
		err = errors.New("method POST not supported")
	} else {

		validHeader, err = validateHeaders(ctx, []string{"Authorization"})

		if validHeader == true {

			rc := new(context.RequestContext)
			rc.New(ctx)
			bytes, err := requestContextToBytesArray(rc)
			if err != nil {
				logger.Logger.PrintError(err)
			}
			err = NatsConnection.Request("GetCovidDistrictsList", bytes, &state, 10*time.Minute)
			if err != nil {
				logger.Logger.PrintError(err)
			}
			err = xml.Unmarshal([]byte(state), &districts)
			if err != nil {
				logger.Logger.PrintError(err)
			}
		}
	}
	sendModelIfExist(ctx, districts, err)
}

var DistrictsHandler fasthttp.RequestHandler = func(ctx *fasthttp.RequestCtx) {
	defer CatchPanic(ctx)

	var districtsStrings []string
	var err error
	var validHeader bool

	if ctx.IsPost() == true {
		err = errors.New("Method POST not supported")
	} else {
		validHeader, err = validateHeaders(ctx, []string{"Authorization"})

		if validHeader == true {
			rc := new(context.RequestContext)
			rc.New(ctx)
			bytes, err := requestContextToBytesArray(rc)
			if err == nil {
				logger.Logger.PrintError(err)
			}
			err = NatsConnection.Request("GetCovidDistrictStringsList", bytes, &districtsStrings, 10*time.Minute)
			if err == nil {
				logger.Logger.PrintError(err)
			}
		}
	}
	sendModelIfExist(ctx, districtsStrings, err)
}

var LpuListHandler fasthttp.RequestHandler = func(ctx *fasthttp.RequestCtx) {
	defer CatchPanic(ctx)

	var state string
	var lpuList soapmodel.SoapLpuListResponse
	var err error
	var validHeader bool
	var validParameters bool

	validHeader, err = validateHeaders(ctx, []string{"Authorization"})

	if validHeader == true {
		validParameters, err = validateParameters(ctx, []string{"idDistrict"})
		if validParameters == true {
			rc := new(context.RequestContext)
			rc.New(ctx)
			bytes, err := requestContextToBytesArray(rc)
			if err != nil {
				logger.Logger.PrintError(err)
			}
			err = NatsConnection.Request("GetLpuList", bytes, &state, 10*time.Minute)
			if err != nil {
				logger.Logger.PrintError(err)
			}
			err = xml.Unmarshal([]byte(state), &lpuList)
			if err != nil {
				logger.Logger.PrintError(err)
			}
		}
	}
	sendModelIfExist(ctx, lpuList, err)
}

var CovidLpuListHandler fasthttp.RequestHandler = func(ctx *fasthttp.RequestCtx) {
	defer CatchPanic(ctx)

	var state string
	var lpuList soapmodel.SoapCovidLpuListResponse
	var err error
	var validHeader bool
	var validParameter bool

	validHeader, err = validateHeaders(ctx, []string{"Authorization"})

	if validHeader == true {
		validParameter, err = validateParameters(ctx, []string{"idDistrict"})
		if validParameter == true {

			rc := new(context.RequestContext)
			rc.New(ctx)
			bytes, err := requestContextToBytesArray(rc)
			if err != nil {
				logger.Logger.PrintError(err)
			}
			err = NatsConnection.Request("GetCovidLpuList", bytes, &state, 10*time.Minute)
			if err != nil {
				logger.Logger.PrintError(err)
			}
			err = xml.Unmarshal([]byte(state), &lpuList)
			if err != nil {
				logger.Logger.PrintError(err)
			}
		}
	}
	sendModelIfExist(ctx, lpuList, err)
}

var SpecialityListHandler fasthttp.RequestHandler = func(ctx *fasthttp.RequestCtx) {
	defer CatchPanic(ctx)

	var state string
	var specialityList soapmodel.SoapSpecialityListResponse
	var err error
	var validHeader bool
	var validParameter bool

	if ctx.IsPost() == true {
		err = errors.New("Method POST not uspported")
	} else {

		validHeader, err = validateHeaders(ctx, []string{"Authorization"})

		if validHeader == true {
			validParameter, err = validateParameters(ctx, []string{"idLpu"})
			if validParameter == true {
				rc := new(context.RequestContext)
				rc.New(ctx)
				bytes, err := requestContextToBytesArray(rc)
				if err != nil {
					logger.Logger.PrintError(err)
				}
				err = NatsConnection.Request("GetCovidSpecialityList", bytes, &state, 10*time.Minute)
				if err != nil {
					logger.Logger.PrintError(err)
				}
				err = xml.Unmarshal([]byte(state), &specialityList)
				if err != nil {
					logger.Logger.PrintError(err)
				}
			}
		}
	}

	sendModelIfExist(ctx, specialityList, err)
}

var DoctorListHandler fasthttp.RequestHandler = func(ctx *fasthttp.RequestCtx) {
	defer CatchPanic(ctx)

	var state string
	var doctorList soapmodel.SoapDoctorListResponse
	var err error
	var validHeader bool
	var validParameter bool

	if ctx.IsPost() == true {
		err = errors.New("Method POST not supported")
	} else {
		validHeader, err = validateHeaders(ctx, []string{"Authorization"})

		if validHeader == true {
			validParameter, err = validateParameters(ctx, []string{"idSpeciality", "idLpu"})
			if validParameter == true {
				rc := new(context.RequestContext)
				rc.New(ctx)
				bytes, err := requestContextToBytesArray(rc)
				if err != nil {
					logger.Logger.PrintError(err)
				}
				err = NatsConnection.Request("GetCovidDoctorList", bytes, &state, 10*time.Minute)
				if err != nil {
					logger.Logger.PrintError(err)
				}
				err = xml.Unmarshal([]byte(state), &doctorList)
				if err != nil {
					logger.Logger.PrintError(err)
				}
			}
		}
	}
	sendModelIfExist(ctx, doctorList, err)
}

var AvailableAppointmentHandler fasthttp.RequestHandler = func(ctx *fasthttp.RequestCtx) {
	defer CatchPanic(ctx)

	var state string
	var appointmentsList soapmodel.SoapAppointmentListResponse
	var err error
	var validHeader bool
	var validParameter bool

	if ctx.IsPost() == true {
		err = errors.New("Method POST not supported")
	} else {

		validHeader, err = validateHeaders(ctx, []string{"Authorization"})

		if validHeader == true {
			validParameter, err = validateParameters(ctx, []string{"idDoc", "idLpu"})
			if validParameter == true {
				rc := new(context.RequestContext)
				rc.New(ctx)
				bytes, err := requestContextToBytesArray(rc)
				if err != nil {
					logger.Logger.PrintError(err)
				}
				err = NatsConnection.Request("GetCovidAppointmentList", bytes, &state, 10*time.Minute)
				if err != nil {
					logger.Logger.PrintError(err)
				}
				err = xml.Unmarshal([]byte(state), &appointmentsList)
				if err != nil {
					logger.Logger.PrintError(err)
				}
			}
		}
	}
	sendModelIfExist(ctx, appointmentsList, err)
}

var CheckPatientHandler fasthttp.RequestHandler = func(ctx *fasthttp.RequestCtx) {
	defer CatchPanic(ctx)

	var state string
	var person *model.Patient = &model.Patient{}
	var patientResponse soapmodel.SoapCheckPatientResponse
	var validHeader bool
	var validModel bool
	var err error

	if ctx.IsGet() == true {
		err = errors.New("Method GET not supported")
	} else {
		validHeader, err = validateHeaders(ctx, []string{"Authorization"})
		if validHeader == true {
			rc := new(context.RequestContext)
			rc.New(ctx)

			err = json.Unmarshal(rc.Body, person)
			if err != nil {
				logger.Logger.PrintError(err)
			}

			validModel, err = validatePatient(*person)

			if validModel == true {
				bytes, err := requestContextToBytesArray(rc)
				if err != nil {
					logger.Logger.PrintError(err)
				}
				err = NatsConnection.Request("CheckPatient", bytes, &state, 10*time.Minute)
				if err != nil {
					logger.Logger.PrintError(err)
				}
				err = xml.Unmarshal([]byte(state), &patientResponse)
				if err != nil {
					logger.Logger.PrintError(err)
				}
			}
		}
	}
	sendModelIfExist(ctx, patientResponse, err)
}

var AddPatientHandler fasthttp.RequestHandler = func(ctx *fasthttp.RequestCtx) {
	defer CatchPanic(ctx)

	var state string
	var person *model.Patient = &model.Patient{}
	var patientResponse soapmodel.SoapAddPatientResponse
	var validHeader bool
	var validModel bool
	var err error

	if ctx.IsGet() == true {
		err = errors.New("Method GET not supported")
	} else {
		validHeader, err = validateHeaders(ctx, []string{"Authorization"})
		if validHeader == true {
			rc := new(context.RequestContext)
			rc.New(ctx)

			err = json.Unmarshal(rc.Body, person)
			if err != nil {
				logger.Logger.PrintError(err)
			}

			validModel, err = validatePatient(*person)

			if validModel == true {
				bytes, err := requestContextToBytesArray(rc)
				if err != nil {
					logger.Logger.PrintError(err)
				}
				err = NatsConnection.Request("AddPatient", bytes, &state, 10*time.Minute)
				if err != nil {
					logger.Logger.PrintError(err)
				}
				err = xml.Unmarshal([]byte(state), &patientResponse)
				if err != nil {
					logger.Logger.PrintError(err)
				}
			}
		}
	}
	sendModelIfExist(ctx, patientResponse, err)
}

var GetPatientHistoryHandler fasthttp.RequestHandler = func(ctx *fasthttp.RequestCtx) {
	defer CatchPanic(ctx)

	var state string
	var personHistory *model.PatientHistory = &model.PatientHistory{}
	var patientHistoryResponse soapmodel.SoapGetPatientHistoryResponse
	var validHeader bool
	var err error

	if ctx.IsGet() == true {
		err = errors.New("Method GET not supported")
	} else {
		validHeader, err = validateHeaders(ctx, []string{"Authorization"})
		if validHeader == true {
			rc := new(context.RequestContext)
			rc.New(ctx)

			err = json.Unmarshal(rc.Body, personHistory)
			if err != nil {
				logger.Logger.PrintError(err)
			}

			bytes, err := requestContextToBytesArray(rc)
			if err != nil {
				logger.Logger.PrintError(err)
			}
			err = NatsConnection.Request("GetPatientHistory", bytes, &state, 10*time.Minute)
			if err != nil {
				logger.Logger.PrintError(err)
			}
			err = xml.Unmarshal([]byte(state), &patientHistoryResponse)
			if err != nil {
				logger.Logger.PrintError(err)
			}
		}
	}
	sendModelIfExist(ctx, patientHistoryResponse, err)
}

var UpdatePhoneHandler fasthttp.RequestHandler = func(ctx *fasthttp.RequestCtx) {
	defer CatchPanic(ctx)

	var state string
	var person *model.Patient
	var updatePhoneResponse soapmodel.SoapUpdatePhoneResponse
	var validHeader bool

	if ctx.IsGet() == true {
		err = errors.New("Method GET not supported")
	} else {
		validHeader, err = validateHeaders(ctx, []string{"Authorization"})
		rc := new(context.RequestContext)
		rc.New(ctx)

		err := json.Unmarshal(rc.Body, person)
		if err != nil {
			logger.Logger.PrintError(err)
		}
		if validHeader == true {
			bytes, err := requestContextToBytesArray(rc)
			if err != nil {
				logger.Logger.PrintError(err)
			}
			err = NatsConnection.Request("UpdatePhone", bytes, &state, 10*time.Minute)
			if err != nil {
				logger.Logger.PrintError(err)
			}
			err = xml.Unmarshal([]byte(state), &updatePhoneResponse)
			if err != nil {
				logger.Logger.PrintError(err)
			}
		}
	}
	sendModelIfExist(ctx, updatePhoneResponse, err)
}

var SetAppointmentHandler fasthttp.RequestHandler = func(ctx *fasthttp.RequestCtx) {
	defer CatchPanic(ctx)

	var state string
	var person *model.Patient = &model.Patient{}
	var setAppointmentResponse soapmodel.SoapSetAppointmentResponse
	var validHeader bool

	if ctx.IsGet() == true {
		err = errors.New("Method GET not supported")
	} else {
		validHeader, err = validateHeaders(ctx, []string{"Authorization"})
		if validHeader == true {
			rc := new(context.RequestContext)
			rc.New(ctx)

			err := json.Unmarshal(rc.Body, person)
			if err != nil {
				logger.Logger.PrintError(err)
			}
			bytes, err := requestContextToBytesArray(rc)
			if err != nil {
				logger.Logger.PrintError(err)
			}
			err = NatsConnection.Request("SetAppointment", bytes, &state, 10*time.Minute)
			if err != nil {
				logger.Logger.PrintError(err)
			}
			err = xml.Unmarshal([]byte(state), &setAppointmentResponse)
			if err != nil {
				logger.Logger.PrintError(err)
			}
		}
	}
	sendModelIfExist(ctx, setAppointmentResponse, err)
}

var DeleteAppointmentHandler fasthttp.RequestHandler = func(ctx *fasthttp.RequestCtx) {
	defer CatchPanic(ctx)

	var state string
	var person *model.Patient = &model.Patient{}
	var deleteAppointmentResponse soapmodel.SoapDeleteAppointmentResponse
	var validHeader bool

	if ctx.IsGet() == true {
		err = errors.New("Method GET not supported")
	} else {
		validHeader, err = validateHeaders(ctx, []string{"Authorization"})
		if validHeader == true {
			rc := new(context.RequestContext)
			rc.New(ctx)

			err := json.Unmarshal(rc.Body, person)
			if err != nil {
				logger.Logger.PrintError(err)
			}
			bytes, err := requestContextToBytesArray(rc)
			if err != nil {
				logger.Logger.PrintError(err)
			}
			err = NatsConnection.Request("DeleteAppointment", bytes, &state, 10*time.Minute)
			if err != nil {
				logger.Logger.PrintError(err)
			}
			err = xml.Unmarshal([]byte(state), &deleteAppointmentResponse)
			if err != nil {
				logger.Logger.PrintError(err)
			}
		}
	}
	sendModelIfExist(ctx, deleteAppointmentResponse, err)
}

var CovidAppointmentCountHandler fasthttp.RequestHandler = func(ctx *fasthttp.RequestCtx) {
	defer CatchPanic(ctx)

	var validHeader bool
	var covidAppointmentCount model.CovidAppointmentCountResponse

	if ctx.IsPost() == true {
		err = errors.New("Method POST not supported")
	} else {

		validHeader, err = validateHeaders(ctx, []string{"Authorization"})
		if validHeader == true {

			rc := new(context.RequestContext)
			rc.New(ctx)
			bytes, err := requestContextToBytesArray(rc)
			if err == nil {
				err = NatsConnection.Request("GetCovidAppointmentCount", bytes, &covidAppointmentCount, 10*time.Minute)
			}
		}
	}
	sendModelIfExist(ctx, covidAppointmentCount, err)
}

var CovidAppointmentTimesHandler fasthttp.RequestHandler = func(ctx *fasthttp.RequestCtx) {
	defer CatchPanic(ctx)

	var validHeader bool
	var validParameter bool
	var covidAppointmentTimes model.CovidAppointmentTimes

	if ctx.IsPost() == true {
		err = errors.New("Method POST not supported")
	} else {
		validHeader, err = validateHeaders(ctx, []string{"Authorization"})
		if validHeader == true {
			validParameter, err = validateParameters(ctx, []string{"idDoc", "idLpu"})
			if validParameter == true {
				rc := new(context.RequestContext)
				rc.New(ctx)
				bytes, err := requestContextToBytesArray(rc)
				if err == nil {
					err = NatsConnection.Request("GetCovidAppointmentTimes", bytes, &covidAppointmentTimes, 10*time.Minute)
				}
			}
		}
	}
	sendModelIfExist(ctx, covidAppointmentTimes, err)
}

var CovidLpuNamesHandler fasthttp.RequestHandler = func(ctx *fasthttp.RequestCtx) {
	defer CatchPanic(ctx)

	var validHeader bool
	var covidLpuNames model.CovidLpuNames
	var validParameter bool

	if ctx.IsPost() == true {
		err = errors.New("Method POST not supported")
	} else {
		validHeader, err = validateHeaders(ctx, []string{"Authorization"})
		if validHeader == true {
			validParameter, err = validateParameters(ctx, []string{"districtName"})
			if validParameter == true {
				rc := new(context.RequestContext)
				rc.New(ctx)
				bytes, err := requestContextToBytesArray(rc)
				if err == nil {
					err = NatsConnection.Request("GetCovidLpuNames", bytes, &covidLpuNames, 10*time.Minute)
				}
			}
		}
	}
	sendModelIfExist(ctx, covidLpuNames, err)
}

var CovidLpuIdsHandler fasthttp.RequestHandler = func(ctx *fasthttp.RequestCtx) {
	defer CatchPanic(ctx)

	var validHeader bool
	var validParameter bool
	var covidLpuIds model.CovidLpuIds

	if ctx.IsPost() == true {
		err = errors.New("Method POST not supported")
	} else {
		validHeader, err = validateHeaders(ctx, []string{"Authorization"})
		if validHeader == true {
			validParameter, err = validateParameters(ctx, []string{"districtId"})
			if validParameter == true {
				rc := new(context.RequestContext)
				rc.New(ctx)
				bytes, err := requestContextToBytesArray(rc)
				if err == nil {
					err = NatsConnection.Request("GetCovidLpuIds", bytes, &covidLpuIds, 10*time.Minute)
				}
			}
		}
	}
	sendModelIfExist(ctx, covidLpuIds, err)
}

var CovidLpuIdByNameHandler fasthttp.RequestHandler = func(ctx *fasthttp.RequestCtx) {
	defer CatchPanic(ctx)

	var lpuId string
	var validHeader bool
	var validParameter bool
	var err error

	if ctx.IsPost() == true {
		err = errors.New("Method POST not supported")
	} else {

		validHeader, err = validateHeaders(ctx, []string{"Authorization"})

		if validHeader == true {
			validParameter, err = validateParameters(ctx, []string{"lpuName"})
			if validParameter == true {
				rc := new(context.RequestContext)
				rc.New(ctx)
				bytes, err := requestContextToBytesArray(rc)
				if err != nil {
					logger.Logger.PrintError(err)
				}
				err = NatsConnection.Request("GetCovidLpuIdByName", bytes, &lpuId, 10*time.Minute)
				if err != nil {
					logger.Logger.PrintError(err)
				}
			}
		}
	}
	sendModelIfExist(ctx, lpuId, err)
}

var CovidSpecialityNamesHandler fasthttp.RequestHandler = func(ctx *fasthttp.RequestCtx) {
	defer CatchPanic(ctx)

	var validHeader bool
	var validParameter bool
	var covidSpecialityNames model.CovidSpecialityNames

	if ctx.IsPost() == true {
		err = errors.New("Method POST not supported")
	} else {
		validHeader, err = validateHeaders(ctx, []string{"Authorization"})
		if validHeader == true {
			validParameter, err = validateParameters(ctx, []string{"idLpu"})
			if validParameter == true {
				rc := new(context.RequestContext)
				rc.New(ctx)
				bytes, err := requestContextToBytesArray(rc)
				if err == nil {
					err = NatsConnection.Request("GetCovidSpecialityNames", bytes, &covidSpecialityNames, 10*time.Minute)
				}
			}
		}
	}
	sendModelIfExist(ctx, covidSpecialityNames, err)
}

var CovidSpecialityIdsHandler fasthttp.RequestHandler = func(ctx *fasthttp.RequestCtx) {
	defer CatchPanic(ctx)

	var validHeader bool
	var validParameter bool
	var covidSpecialityIds model.CovidSpecialityIds

	if ctx.IsPost() == true {
		err = errors.New("Method POST not supported")
	} else {
		validHeader, err = validateHeaders(ctx, []string{"Authorization"})
		if validHeader == true {
			validParameter, err = validateParameters(ctx, []string{"idLpu"})
			if validParameter == true {
				rc := new(context.RequestContext)
				rc.New(ctx)
				bytes, err := requestContextToBytesArray(rc)
				if err == nil {
					err = NatsConnection.Request("GetCovidSpecialityIds", bytes, &covidSpecialityIds, 10*time.Minute)
				}
			}
		}
	}
	sendModelIfExist(ctx, covidSpecialityIds, err)
}

var CovidSpecialityIdHandler fasthttp.RequestHandler = func(ctx *fasthttp.RequestCtx) {
	defer CatchPanic(ctx)

	var validHeader bool
	var validParameter bool
	var covidSpecialityId string

	if ctx.IsPost() == true {
		err = errors.New("Method POST not supported")
	} else {
		validHeader, err = validateHeaders(ctx, []string{"Authorization"})
		if validHeader == true {
			validParameter, err = validateParameters(ctx, []string{"idLpu", "specialityName"})
			if validParameter == true {
				rc := new(context.RequestContext)
				rc.New(ctx)
				bytes, err := requestContextToBytesArray(rc)
				if err == nil {
					err = NatsConnection.Request("GetCovidSpecialityId", bytes, &covidSpecialityId, 10*time.Minute)
				}
			}
		}
	}
	sendModelIfExist(ctx, covidSpecialityId, err)
}

var CovidDocNamesHandler fasthttp.RequestHandler = func(ctx *fasthttp.RequestCtx) {
	defer CatchPanic(ctx)

	var validHeader bool
	var validParameter bool
	var covidDocNames model.CovidDocNames

	if ctx.IsPost() == true {
		err = errors.New("Method POST not supported")
	} else {
		validHeader, err = validateHeaders(ctx, []string{"Authorization"})
		if validHeader == true {
			validParameter, err = validateParameters(ctx, []string{"idLpu", "idSpeciality"})
			if validParameter == true {
				rc := new(context.RequestContext)
				rc.New(ctx)
				bytes, err := requestContextToBytesArray(rc)
				if err == nil {
					err = NatsConnection.Request("GetCovidDocNames", bytes, &covidDocNames, 10*time.Minute)
				}
			}
		}
	}
	sendModelIfExist(ctx, covidDocNames, err)
}

var CovidDocIdsHandler fasthttp.RequestHandler = func(ctx *fasthttp.RequestCtx) {
	defer CatchPanic(ctx)

	var validHeader bool
	var validParameter bool
	var covidDocIds model.CovidDocIds

	if ctx.IsPost() == true {
		err = errors.New("Method POST not supported")
	} else {
		validHeader, err = validateHeaders(ctx, []string{"Authorization"})
		if validHeader == true {
			validParameter, err = validateParameters(ctx, []string{"idLpu", "idSpeciality"})
			if validParameter == true {
				rc := new(context.RequestContext)
				rc.New(ctx)
				bytes, err := requestContextToBytesArray(rc)
				if err == nil {
					err = NatsConnection.Request("GetCovidDocIds", bytes, &covidDocIds, 10*time.Minute)
				}
			}
		}
	}
	sendModelIfExist(ctx, covidDocIds, err)
}

var CovidDocIdHandler fasthttp.RequestHandler = func(ctx *fasthttp.RequestCtx) {
	defer CatchPanic(ctx)

	var validHeader bool
	var validParameter bool
	var covidDocId string

	if ctx.IsPost() == true {
		err = errors.New("Method POST not supported")
	} else {
		validHeader, err = validateHeaders(ctx, []string{"Authorization"})
		if validHeader == true {
			validParameter, err = validateParameters(ctx, []string{"idLpu", "idSpeciality", "docName"})
			if validParameter == true {
				rc := new(context.RequestContext)
				rc.New(ctx)
				bytes, err := requestContextToBytesArray(rc)
				if err == nil {
					err = NatsConnection.Request("GetCovidDocId", bytes, &covidDocId, 10*time.Minute)
				}
			}
		}
	}
	sendModelIfExist(ctx, covidDocId, err)
}

var CovidAppointmentIdsHandler fasthttp.RequestHandler = func(ctx *fasthttp.RequestCtx) {
	defer CatchPanic(ctx)

	var validHeader bool
	var validParameter bool
	var covidAppointmentIds model.CovidAppointmentIds

	if ctx.IsPost() == true {
		err = errors.New("Method POST not supported")
	} else {
		validHeader, err = validateHeaders(ctx, []string{"Authorization"})
		if validHeader == true {
			validParameter, err = validateParameters(ctx, []string{"idLpu", "idDoc"})
			if validParameter == true {
				rc := new(context.RequestContext)
				rc.New(ctx)
				bytes, err := requestContextToBytesArray(rc)
				if err == nil {
					err = NatsConnection.Request("GetCovidAppointmentIds", bytes, &covidAppointmentIds, 10*time.Minute)
				}
			}
		}
	}
	sendModelIfExist(ctx, covidAppointmentIds, err)
}
