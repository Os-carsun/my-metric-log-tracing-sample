package wssocket

import (
	"technest/tracing-log-metric/wsserver/modules/chat"

	"github.com/gin-gonic/gin"
)

func RunWsServer(address string) {
	// Create our HTTP Router
	router := gin.New()

	// Configure HTTP Router Settings
	router.RedirectTrailingSlash = true
	router.RedirectFixedPath = false
	router.HandleMethodNotAllowed = false
	router.ForwardedByClientIP = true
	router.AppEngine = false
	router.UseRawPath = false
	router.UnescapePathValues = true

	chat := chat.Module("chat")
	chat.Init(router)
	go chat.Start()

	router.Run(address)
}
