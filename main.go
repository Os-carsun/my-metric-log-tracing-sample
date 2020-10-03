package main

import "technest/tracing-log-metric/webserver"

func main() {
	router := webserver.CreateServer()
	router.Run("localhost:8889")

}
