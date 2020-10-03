package webserver

import (
	"io"
	"technest/tracing-log-metric/webserver/middleware"
	"technest/tracing-log-metric/webserver/modules/calculater"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
)

var tracerCloser = []io.Closer{}

func CreateServer() *gin.Engine {
	r := gin.New()
	r.Use(middleware.Logger())
	r.Use(gin.Logger())
	r.Use(middleware.UptimeMetrics())

	cal := calculater.Module("calculater")
	cal.RegisterInRouter(r, "/cal")

	t, c := cal.InitTracer()
	opentracing.SetGlobalTracer(t)
	tracerCloser = append(tracerCloser, c)

	return r
}

func CloseTracer() {
	for _, c := range tracerCloser {
		c.Close()
	}
}
