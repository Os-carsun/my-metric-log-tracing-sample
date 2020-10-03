package calculater

import (
	"fmt"
	"io"
	"technest/tracing-log-metric/config"
	"technest/tracing-log-metric/webserver/middleware"
	"technest/tracing-log-metric/webserver/modules"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

func Module(serviceName string) modules.WebAPI {
	return &module{
		Config: Config{
			Tracing: &jaegercfg.Configuration{
				Sampler: &jaegercfg.SamplerConfig{
					Type:  jaeger.SamplerTypeConst,
					Param: 1,
				},
				Reporter: &jaegercfg.ReporterConfig{
					LocalAgentHostPort: config.JaegerHost,
					LogSpans:           true,
				},
			},
		},
		Tracer:  nil,
		Service: "calculater",
	}
}

type Config struct {
	Tracing *jaegercfg.Configuration
}

type module struct {
	Config
	Tracer  opentracing.Tracer
	Service string
}

func (m *module) RegisterInRouter(router gin.IRouter, path string) {
	calGroup := router.Group(path)
	{
		calGroup.GET("/plus", plus, middleware.OpsPlusMetrics(), m.tracePlus())
		calGroup.GET("/sub", sub, middleware.OpsSubMetrics(), m.traceSub())
	}
}

func (m *module) InitTracer() (opentracing.Tracer, io.Closer) {
	t, c, err := m.Config.Tracing.New(m.Service)
	if err != nil {
		panic(fmt.Sprintf("Init failed: %v\n", err))
	}
	m.Tracer = t
	return t, c
}

func (m *module) tracePlus() gin.HandlerFunc {
	return func(c *gin.Context) {
		if m.Tracer == nil {
			return
		}
		c.Next()
		spanCtx, _ := m.Tracer.Extract(opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier(c.Request.Header))

		span := m.Tracer.StartSpan("PlusResuest", ext.RPCServerOption(spanCtx))
		defer span.Finish()
		span.LogFields(
			log.String("event", "plus number"),
			log.String("type", "calculater"))
	}
}

func (m *module) traceSub() gin.HandlerFunc {
	return func(c *gin.Context) {
		if m.Tracer == nil {
			return
		}
		span := opentracing.StartSpan("SubRequest")
		span.SetTag("a", c.Query("a"))
		span.SetTag("b", c.Query("b"))
		defer span.Finish()
		span.LogFields(
			log.String("event", "sub number"),
			log.String("type", "calculater"),
			log.String("a", c.Query("a")))
	}
}
