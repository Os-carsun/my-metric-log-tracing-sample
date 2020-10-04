package chat

import (
	"technest/tracing-log-metric/config"
	"technest/tracing-log-metric/webserver/modules"

	"github.com/gin-gonic/gin"
	jaegercfg "github.com/uber/jaeger-client-go/config"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
)

func Module(serviceName string) modules.WSAPI {
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
		Service: serviceName,
	}
}

type Config struct {
	Tracing *jaegercfg.Configuration
}
type module struct {
	Config
	Tracer  opentracing.Tracer
	hub     *Hub
	Service string
}

func (m *module) Init(router *gin.Engine) {
	m.hub = newHub()
	router.LoadHTMLGlob("template/*.html")
	router.GET("/start", startClient(m.hub))
	router.GET("/home", func(c *gin.Context) {
		c.HTML(200, "home.html", nil)
	})
}

func (m *module) Start() {
	m.hub.run()
}
