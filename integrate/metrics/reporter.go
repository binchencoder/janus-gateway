package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"

	tm "binchencoder.com/letsgo/time"
)

var (
	// Create a histogram for record response latency (milliseconds) of ease-gateway.
	// And it will generate additional metric, for example:
	// gateway_http_response_ms_count, it is the total number of request.
	gatewayHandledHistogram = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "gateway",
			Subsystem: "http",
			Name:      "response_ms",
			Help:      "Gateway latency distributions.",
			Buckets:   prometheus.DefBuckets,
		},
		[]string{"client", "service_name", "url", "http_method", "code"},
	)

	// Create a counter for record total system errors of ease-gateway.
	gatewayErrCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "gateway",
			Subsystem: "http",
			Name:      "sys_err",
			Help:      "Gateway system error count.",
		},
		[]string{"tag"},
	)
)

func init() {
	// Register the histogram with Prometheus's default registry.
	prometheus.MustRegister(gatewayHandledHistogram)
	// Register the counter with Prometheus's default registry.
	prometheus.MustRegister(gatewayErrCounter)
}

// ReporterParam contains prometheus label value and other extra attribute.
type ReporterParam struct {
	// Request url whick config in proto file.
	Url string
	// Request success value is "0". if Error occurred,
	// Value is errcode.
	Code string
	// Request start time for calculating latency.
	StartTime   time.Time
	ServiceName string
	// Http method like PUT/POST/DELETE and so on..
	HttpMethod string
	// Request client.
	Client string
}

// RequestComplete may be invoked in any method which request is ended.
func (reporter *ReporterParam) RequestComplete() float64 {
	// convert ns to ms.
	ms := tm.MillisecondSince(reporter.StartTime)
	gatewayHandledHistogram.WithLabelValues(reporter.Client, reporter.ServiceName, reporter.Url, reporter.HttpMethod, reporter.Code).Observe(ms)
	return ms
}

// ErrCount may be invoked in any method which error has happened.
func ErrCount(tag string) {
	gatewayErrCounter.WithLabelValues(tag).Inc()
}
