package metrics

const (
	ERROR		=	"error"
	SUCCESS 	= 	"success"
	COUNT		=	"count"
	PANIC 		= 	"panic"
) 

var (
	Counter = NewCounterVec(&CounterVecOpts{
		Name:      "counter",
		Help:      "frame watch error count",
		Labels:    []string{"name", "type", "msg"},
	})
	
	Histogram = NewHistogramVec(&HistogramVecOpts{
		Name:      "histogram",
		Help:      "frame histogram",
		Labels:    []string{"name"},
		Buckets:   []float64{5, 10, 25, 50, 100, 250, 500, 1000, 2500},
	})
)
