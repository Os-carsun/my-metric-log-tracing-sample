package calculater

import (
	"technest/tracing-log-metric/webserver/modules"

	"github.com/gin-gonic/gin"
)

func Module(serviceName string) modules.WebAPI {
	return &module{Config: Config{}, Service: "calculater"}
}

type Config struct {
}

type module struct {
	Config
	Service string
}

func (m *module) RegisterInRouter(router gin.IRouter, path string) {
	calGroup := router.Group(path)
	{
		calGroup.GET("/plus", plus)
		calGroup.GET("/sub", sub)
	}
}
