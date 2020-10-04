package modules

import (
	"io"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
)

type WebAPI interface {
	RegisterInRouter(router gin.IRouter, path string)
	InitTracer() (opentracing.Tracer, io.Closer)
}

type WSAPI interface {
	Init(router *gin.Engine)
	Start()
	InitTracer() (opentracing.Tracer, io.Closer)
}
