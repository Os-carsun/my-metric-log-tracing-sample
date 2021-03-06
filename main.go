package main

import (
	"net/http"
	"technest/tracing-log-metric/webserver"
	wssocket "technest/tracing-log-metric/wsserver"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	router := webserver.CreateServer()
	defer webserver.CloseTracer()
	go router.Run("localhost:8889")
	go wssocket.RunWsServer("localhost:9999")
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
