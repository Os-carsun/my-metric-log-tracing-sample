package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	opsPlus = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "calculater_plus_count",
			Help: "The total number of plus",
		})
	opsSub = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "calculater_sub_count",
			Help: "The total number of sub",
		})
	uptime = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "uptime",
			Help: "HTTP service uptime.",
		}, nil,
	)
)

func recordUptime() {
	for range time.Tick(time.Second) {
		uptime.WithLabelValues().Inc()
	}
}

func UptimeMetrics() gin.HandlerFunc {
	prometheus.MustRegister(uptime)
	go recordUptime()
	return func(c *gin.Context) {
		c.Next()
	}
}

func OpsPlusMetrics() gin.HandlerFunc {
	prometheus.MustRegister(opsPlus)
	return func(c *gin.Context) {
		c.Next()
		opsPlus.Inc()
	}
}
func OpsSubMetrics() gin.HandlerFunc {
	prometheus.MustRegister(opsSub)
	return func(c *gin.Context) {
		c.Next()
		opsSub.Inc()
	}
}
