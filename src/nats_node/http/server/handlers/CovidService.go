package handlers

import (
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
		}
	}
	err = xml.Unmarshal([]byte(state), &districts)
	if err != nil {
		logger.Logger.PrintError(err)
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

	validHeader, err := validateHeaders(ctx, []string{"Authorization"})
	validParameters, err := validateParameters(ctx, []string{"idDistrict"})

	if validHeader == true && validParameters == true {

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
	}

	err = xml.Unmarshal([]byte(state), &lpuList)
	if err != nil {
		logger.Logger.PrintError(err)
	}

	sendModelIfExist(ctx, lpuList, err)
}

var CovidLpuListHandler fasthttp.RequestHandler = func(ctx *fasthttp.RequestCtx) {
	defer CatchPanic(ctx)

	var state string
	var lpuList soapmodel.SoapCovidLpuListResponse
	var err error
	var validHeader bool

	validHeader, err = validateHeaders(ctx, []string{"Authorization"})

	if validHeader == true {

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
	}

	err = xml.Unmarshal([]byte(state), &lpuList)
	if err != nil {
		logger.Logger.PrintError(err)
	}

	sendModelIfExist(ctx, lpuList, err)
}

var CovidLpuIdByNameHandler fasthttp.RequestHandler = func(ctx *fasthttp.RequestCtx) {
	defer CatchPanic(ctx)

	var lpuId string

	if ctx.IsPost() == true {
		err = errors.New("Method POST not supported")
	} else {
		rc := new(context.RequestContext)
		rc.New(ctx)
		bytes, err := requestContextToBytesArray(rc)
		if err == nil {
			err = NatsConnection.Request("GetCovidLpuIdByName", bytes, &lpuId, 10*time.Minute)
		}
	}
	sendModelIfExist(ctx, lpuId, err)
}

var SpecialityListHandler fasthttp.RequestHandler = func(ctx *fasthttp.RequestCtx) {
	defer CatchPanic(ctx)

	var specialityList soapmodel.SoapSpecialityListResponse

	if ctx.IsPost() == true {
		err = errors.New("Method POST not supported")
	} else {
		rc := new(context.RequestContext)
		rc.New(ctx)
		bytes, err := requestContextToBytesArray(rc)
		if err == nil {
			err = NatsConnection.Request("GetCovidSpecialityList", bytes, &specialityList, 10*time.Minute)
		}
	}
	sendModelIfExist(ctx, specialityList, err)
}

var DoctorListHandler fasthttp.RequestHandler = func(ctx *fasthttp.RequestCtx) {
	defer CatchPanic(ctx)

	var doctorList soapmodel.SoapDoctorListResponse

	if ctx.IsPost() == true {
		err = errors.New("Method POST not supported")
	} else {
		rc := new(context.RequestContext)
		rc.New(ctx)
		bytes, err := requestContextToBytesArray(rc)
		if err == nil {
			err = NatsConnection.Request("GetCovidDoctorList", bytes, &doctorList, 10*time.Minute)
		}
	}
	sendModelIfExist(ctx, doctorList, err)
}

var AvailableAppointmentHandler fasthttp.RequestHandler = func(ctx *fasthttp.RequestCtx) {
	defer CatchPanic(ctx)

	var appointmentsList soapmodel.SoapAppointmentListResponse

	if ctx.IsPost() == true {
		err = errors.New("Method POST not supported")
	} else {
		rc := new(context.RequestContext)
		rc.New(ctx)
		bytes, err := requestContextToBytesArray(rc)
		if err == nil {
			err = NatsConnection.Request("GetCovidAppointmentList", bytes, &appointmentsList, 10*time.Minute)
		}
	}
	sendModelIfExist(ctx, appointmentsList, err)
}

var CheckPatientHandler fasthttp.RequestHandler = func(ctx *fasthttp.RequestCtx) {
	defer CatchPanic(ctx)

	var patientResponse soapmodel.SoapCheckPatientResponse

	if ctx.IsGet() == true {
		err = errors.New("Method GET not supported")
	} else {
		rc := new(context.RequestContext)
		rc.New(ctx)
		bytes, err := requestContextToBytesArray(rc)
		if err == nil {
			err = NatsConnection.Request("CheckPatient", bytes, &patientResponse, 10*time.Minute)
		}
	}
	sendModelIfExist(ctx, patientResponse, err)
}

var AddPatientHandler fasthttp.RequestHandler = func(ctx *fasthttp.RequestCtx) {
	defer CatchPanic(ctx)

	var patientResponse soapmodel.SoapCheckPatientResponse

	if ctx.IsGet() == true {
		err = errors.New("Method GET not supported")
	} else {
		rc := new(context.RequestContext)
		rc.New(ctx)
		bytes, err := requestContextToBytesArray(rc)
		if err == nil {
			err = NatsConnection.Request("AddPatient", bytes, &patientResponse, 10*time.Minute)
		}
	}
	sendModelIfExist(ctx, patientResponse, err)
}

var UpdatePhoneHandler fasthttp.RequestHandler = func(ctx *fasthttp.RequestCtx) {
	defer CatchPanic(ctx)

	var updatePhoneResponse soapmodel.SoapUpdatePhoneResponse

	if ctx.IsGet() == true {
		err = errors.New("Method GET not supported")
	} else {
		rc := new(context.RequestContext)
		rc.New(ctx)
		bytes, err := requestContextToBytesArray(rc)
		if err == nil {
			err = NatsConnection.Request("UpdatePhone", bytes, &updatePhoneResponse, 10*time.Minute)
		}
	}
	sendModelIfExist(ctx, updatePhoneResponse, err)
}

var SetAppointmentHandler fasthttp.RequestHandler = func(ctx *fasthttp.RequestCtx) {
	defer CatchPanic(ctx)

	var setAppointmentResponse soapmodel.SoapSetAppointmentResponse

	if ctx.IsGet() == true {
		err = errors.New("Method GET not supported")
	} else {
		rc := new(context.RequestContext)
		rc.New(ctx)
		bytes, err := requestContextToBytesArray(rc)
		if err == nil {
			err = NatsConnection.Request("SetAppointment", bytes, &setAppointmentResponse, 10*time.Minute)
		}
	}
	sendModelIfExist(ctx, setAppointmentResponse, err)
}

var DeleteAppointmentHandler fasthttp.RequestHandler = func(ctx *fasthttp.RequestCtx) {
	defer CatchPanic(ctx)

	var deleteAppointmentResponse soapmodel.SoapDeleteAppointmentResponse

	if ctx.IsGet() == true {
		err = errors.New("Method GET not supported")
	} else {
		rc := new(context.RequestContext)
		rc.New(ctx)
		bytes, err := requestContextToBytesArray(rc)
		if err == nil {
			err = NatsConnection.Request("DeleteAppointment", bytes, &deleteAppointmentResponse, 10*time.Minute)
		}
	}
	sendModelIfExist(ctx, deleteAppointmentResponse, err)
}

var CovidLpuNamesHandler fasthttp.RequestHandler = func(ctx *fasthttp.RequestCtx) {
	defer CatchPanic(ctx)

	var covidLpuNames model.CovidLpuNames

	if ctx.IsPost() == true {
		err = errors.New("Method POST not supported")
	} else {
		rc := new(context.RequestContext)
		rc.New(ctx)
		bytes, err := requestContextToBytesArray(rc)
		if err == nil {
			err = NatsConnection.Request("GetCovidLpuNames", bytes, &covidLpuNames, 10*time.Minute)
		}
	}
	sendModelIfExist(ctx, covidLpuNames, err)
}

var CovidLpuIdsHandler fasthttp.RequestHandler = func(ctx *fasthttp.RequestCtx) {
	defer CatchPanic(ctx)

	var covidLpuIds model.CovidLpuIds

	if ctx.IsPost() == true {
		err = errors.New("Method POST not supported")
	} else {
		rc := new(context.RequestContext)
		rc.New(ctx)
		bytes, err := requestContextToBytesArray(rc)
		if err == nil {
			err = NatsConnection.Request("GetCovidLpuIds", bytes, &covidLpuIds, 10*time.Minute)
		}
	}
	sendModelIfExist(ctx, covidLpuIds, err)
}

var CovidSpecialityNamesHandler fasthttp.RequestHandler = func(ctx *fasthttp.RequestCtx) {
	defer CatchPanic(ctx)

	var covidSpecialityNames model.CovidSpecialityNames

	if ctx.IsPost() == true {
		err = errors.New("Method POST not supported")
	} else {
		rc := new(context.RequestContext)
		rc.New(ctx)
		bytes, err := requestContextToBytesArray(rc)
		if err == nil {
			err = NatsConnection.Request("GetCovidSpecialityNames", bytes, &covidSpecialityNames, 10*time.Minute)
		}
	}
	sendModelIfExist(ctx, covidSpecialityNames, err)
}

var CovidSpecialityIdsHandler fasthttp.RequestHandler = func(ctx *fasthttp.RequestCtx) {
	defer CatchPanic(ctx)

	var covidSpecialityIds model.CovidSpecialityIds

	if ctx.IsPost() == true {
		err = errors.New("Method POST not supported")
	} else {
		rc := new(context.RequestContext)
		rc.New(ctx)
		bytes, err := requestContextToBytesArray(rc)
		if err == nil {
			err = NatsConnection.Request("GetCovidSpecialityIds", bytes, &covidSpecialityIds, 10*time.Minute)
		}
	}
	sendModelIfExist(ctx, covidSpecialityIds, err)
}

var CovidSpecialityIdHandler fasthttp.RequestHandler = func(ctx *fasthttp.RequestCtx) {
	defer CatchPanic(ctx)

	var covidSpecialityId string

	if ctx.IsPost() == true {
		err = errors.New("Method POST not supported")
	} else {
		rc := new(context.RequestContext)
		rc.New(ctx)
		bytes, err := requestContextToBytesArray(rc)
		if err == nil {
			err = NatsConnection.Request("GetCovidSpecialityId", bytes, &covidSpecialityId, 10*time.Minute)
		}
	}
	sendModelIfExist(ctx, covidSpecialityId, err)
}

var CovidDocNamesHandler fasthttp.RequestHandler = func(ctx *fasthttp.RequestCtx) {
	defer CatchPanic(ctx)

	var covidDocNames model.CovidDocNames

	if ctx.IsPost() == true {
		err = errors.New("Method POST not supported")
	} else {
		rc := new(context.RequestContext)
		rc.New(ctx)
		bytes, err := requestContextToBytesArray(rc)
		if err == nil {
			err = NatsConnection.Request("GetCovidDocNames", bytes, &covidDocNames, 10*time.Minute)
		}
	}
	sendModelIfExist(ctx, covidDocNames, err)
}

var CovidDocIdsHandler fasthttp.RequestHandler = func(ctx *fasthttp.RequestCtx) {
	defer CatchPanic(ctx)

	var covidDocIds model.CovidDocIds

	if ctx.IsPost() == true {
		err = errors.New("Method POST not supported")
	} else {
		rc := new(context.RequestContext)
		rc.New(ctx)
		bytes, err := requestContextToBytesArray(rc)
		if err == nil {
			err = NatsConnection.Request("GetCovidDocIds", bytes, &covidDocIds, 10*time.Minute)
		}
	}
	sendModelIfExist(ctx, covidDocIds, err)
}

var CovidDocIdHandler fasthttp.RequestHandler = func(ctx *fasthttp.RequestCtx) {
	defer CatchPanic(ctx)

	var covidDocId string

	if ctx.IsPost() == true {
		err = errors.New("Method POST not supported")
	} else {
		rc := new(context.RequestContext)
		rc.New(ctx)
		bytes, err := requestContextToBytesArray(rc)
		if err == nil {
			err = NatsConnection.Request("GetCovidDocId", bytes, &covidDocId, 10*time.Minute)
		}
	}
	sendModelIfExist(ctx, covidDocId, err)
}

var CovidAppointmentTimesHandler fasthttp.RequestHandler = func(ctx *fasthttp.RequestCtx) {
	defer CatchPanic(ctx)

	var covidAppointmentTimes model.CovidAppointmentTimes

	if ctx.IsPost() == true {
		err = errors.New("Method POST not supported")
	} else {
		rc := new(context.RequestContext)
		rc.New(ctx)
		bytes, err := requestContextToBytesArray(rc)
		if err == nil {
			err = NatsConnection.Request("GetCovidAppointmentTimes", bytes, &covidAppointmentTimes, 10*time.Minute)
		}
	}
	sendModelIfExist(ctx, covidAppointmentTimes, err)
}

var CovidAppointmentIdsHandler fasthttp.RequestHandler = func(ctx *fasthttp.RequestCtx) {
	defer CatchPanic(ctx)

	var covidAppointmentIds model.CovidAppointmentIds

	if ctx.IsPost() == true {
		err = errors.New("Method POST not supported")
	} else {
		rc := new(context.RequestContext)
		rc.New(ctx)
		bytes, err := requestContextToBytesArray(rc)
		if err == nil {
			err = NatsConnection.Request("GetCovidAppointmentIds", bytes, &covidAppointmentIds, 10*time.Minute)
		}
	}
	sendModelIfExist(ctx, covidAppointmentIds, err)
}

var CovidAppointmentCountHandler fasthttp.RequestHandler = func(ctx *fasthttp.RequestCtx) {
	defer CatchPanic(ctx)

	var covidAppointmentCount model.CovidAppointmentCountResponse

	if ctx.IsPost() == true {
		err = errors.New("Method POST not supported")
	} else {
		rc := new(context.RequestContext)
		rc.New(ctx)
		bytes, err := requestContextToBytesArray(rc)
		if err == nil {
			err = NatsConnection.Request("GetCovidAppointmentCount", bytes, &covidAppointmentCount, 10*time.Minute)
		}
	}
	sendModelIfExist(ctx, covidAppointmentCount, err)
}
