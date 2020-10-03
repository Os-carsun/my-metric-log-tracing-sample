package webserver

import (
	"technest/tracing-log-metric/webserver/middleware"
	"technest/tracing-log-metric/webserver/modules/calculater"

	"github.com/gin-gonic/gin"
)

func CreateServer() *gin.Engine {
	r := gin.New()
	r.Use(middleware.Logger())
	r.Use(gin.Logger())
	r.Use(middleware.UptimeMetrics())
	cal := calculater.Module("calculater")
	cal.RegisterInRouter(r, "/cal")
	return r
}
