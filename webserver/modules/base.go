package modules

import "github.com/gin-gonic/gin"

type WebAPI interface {
	RegisterInRouter(router gin.IRouter, path string)
}
