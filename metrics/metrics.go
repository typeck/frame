package metrics

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/typeck/frame/log"
	"net/http"
)

type Config struct {
	Namespace 		string		`toml:"namespace"`
	Name	 		string		`toml:"name"`
	Port 			int			`toml:"port"`
}


var DefaultConfig = &Config{
	Namespace: "frame",
	Name:   "frame",
	Port:      10106,
}

type CounterVecOpts struct {
	Name      string
	Help      string
	Labels    []string
}

type counterVec struct {
	*prometheus.CounterVec
}

type HistogramVecOpts struct {
	Name      string
	Help      string
	Labels    []string
	Buckets   []float64
}

type histogramVec struct {
	*prometheus.HistogramVec
}

func NewCounterVec(opts *CounterVecOpts) *counterVec{
	vec := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: DefaultConfig.Namespace,
			Subsystem: DefaultConfig.Name,
			Name:      opts.Name,
			Help:      opts.Help,
		}, opts.Labels)
	prometheus.MustRegister(vec)
	return &counterVec{
		CounterVec: vec,
	}
}

func NewHistogramVec(opts *HistogramVecOpts) *histogramVec {
	vec := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: DefaultConfig.Namespace,
			Subsystem: DefaultConfig.Name,
			Name:      opts.Name,
			Help:      opts.Help,
			Buckets:   opts.Buckets,
		}, opts.Labels)
	prometheus.MustRegister(vec)
	return &histogramVec{
		HistogramVec: vec,
	}
}


func init() {
	mux := http.NewServeMux()
	mux.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		promhttp.Handler().ServeHTTP(w, r)
	})
	go func() {
		err := http.ListenAndServe(fmt.Sprintf(":%d", DefaultConfig.Port), mux)
		if err != nil {
			log.Panic(err)
		}
	}()
	fmt.Printf("metrics listen in :%d", DefaultConfig.Port)
}