package main

import (
	"go-ApmCommon/middleware"
	"go-ApmCommon/model"
	"go-ApmExam1/router"
	"os"

	"github.com/urfave/negroni"
)

var config model.TomlConfig

func init() {
	config.Load()
	//EXPORT APM EXVIRONMENT
	os.Setenv("ELASTIC_APM_SERVER_URL", config.ApmServerUrl())
	os.Setenv("ELASTIC_APM_SERVICE_NAME", config.Service)
}
func main() {
	n := negroni.New(negroni.HandlerFunc(middleware.NewLoggingMiddleware(config.Logpaths.Logpath)))
	n.UseHandler(router.NewRouter())
	n.Run(config.Servers["ApmExam1"].PORT)
}
