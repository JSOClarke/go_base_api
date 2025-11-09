package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var UserRegistrationCounter = promauto.NewCounter(prometheus.CounterOpts{
	Name: "app_user_registrations_total",
	Help: "Total number of user registrations",
})

var RequestErrorCounter = promauto.NewCounterVec(
	prometheus.CounterOpts{Name: "app_http_request_errors_total", Help: "Count of HTTP req errors by route and code"}, []string{"endpoint", "code"},
)
