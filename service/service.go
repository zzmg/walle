package service

import (
	"fmt"
	"github.com/labstack/echo"
	"cradle/walle/service/middleware"
	"cradle/walle/service/api"
	"cradle/walle/common"
)

var app = echo.New()

func init() {
	app.Use(middleware.RequestCORS())

	MountAPIModule(app)
}

// RunServer starts a server
func RunServer() {
	if common.GlobalConf.CertPem != "" && common.GlobalConf.KeyPem != "" {
		app.StartTLS(common.GlobalConf.Bind, common.GlobalConf.CertPem, common.GlobalConf.KeyPem)
	} else {
		fmt.Println(app.Start(common.GlobalConf.Bind))
	}
}

func MountAPIModule(e *echo.Echo) {
	apiv1 := e.Group("/v1")
	api.MountAPI(apiv1)
	//swagger.MountSwaggerAPI(e.Group(""))
}
