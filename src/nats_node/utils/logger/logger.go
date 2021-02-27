package logger

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"
	"testing"

	"github.com/mgutz/logxi/v1"
)

type logger struct {
	logx     log.Logger
	SHOW_LOG bool
}

var Logger *logger

func init() {
	Logger = &logger{
		SHOW_LOG: true,
		logx:     log.NewLogger(log.NewConcurrentWriter(os.Stdout), "main"),
	}
}

func (logger logger) PrintError(err error) {
	if Logger.SHOW_LOG == true {
		if err != nil {
			Logger.logx.Error(err.Error())
		}
	}
}

func (logger logger) PrintStack(stack []byte) {
	if Logger.SHOW_LOG == true {
		_ = Logger.logx.Error("StackTrace", stack)
	}
}

func (logger logger) PrintRecover(recover interface{}) {
	if Logger.SHOW_LOG == true {
		_ = Logger.logx.Error("Recover Reason", recover)
	}
}

func (logger logger) PrintWarn(message string, args ...interface{}) {
	logger.logx.Warn(message, args)
}

func (logger logger) PrintRequestDebug(request *http.Request) {
	rqDump, err := httputil.DumpRequestOut(request, true)
	logger.PrintError(err)

	fmt.Println("****REQUEST------------------>")
	fmt.Println(string(rqDump))
	fmt.Println("")
}

func (logger logger) PrintResponseDebug(response *http.Response) {
	rsDump, err := httputil.DumpResponse(response, true)
	logger.PrintError(err)

	fmt.Println("****RESPONSE------------------>")
	fmt.Println(string(rsDump))
	fmt.Println("")
}

//---------For testing

func (logger logger) IfErrorThenTested(err error, t *testing.T) {
	if err != nil {
		t.Log(t.Name() + " * tested")
	} else {
		t.Fatal(t.Name() + " * failed")
	}
}

func (logger logger) IfErrorThenFailed(err error, t *testing.T) {
	if err == nil {
		t.Log(t.Name() + " * tested")
	} else {
		t.Fatal(t.Name() + " * failed")
	}
}
