package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
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

type CounterVec struct {
	*prometheus.CounterVec
}

type HistogramVecOpts struct {
	Name      string
	Help      string
	Labels    []string
	Buckets   []float64
}

type HistogramVec struct {
	*prometheus.HistogramVec
}

func NewCounterVec(opts *CounterVecOpts) *CounterVec{
	vec := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: DefaultConfig.Namespace,
			Subsystem: DefaultConfig.Name,
			Name:      opts.Name,
			Help:      opts.Help,
		}, opts.Labels)
	prometheus.MustRegister(vec)
	return &CounterVec{
		CounterVec: vec,
	}
}

func NewHistogramVec(opts *HistogramVecOpts) *HistogramVec {
	vec := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: DefaultConfig.Namespace,
			Subsystem: DefaultConfig.Name,
			Name:      opts.Name,
			Help:      opts.Help,
			Buckets:   opts.Buckets,
		}, opts.Labels)
	prometheus.MustRegister(vec)
	return &HistogramVec{
		HistogramVec: vec,
	}
}
