package plugins

import (
	"github.com/afex/hystrix-go/hystrix/metric_collector"
	"github.com/prometheus/client_golang/prometheus"
	"log"
	"strings"
)

type PrometheusCollector struct {
	circuitOpenPrefix       string
	attemptsPrefix          string
	errorsPrefix            string
	successesPrefix         string
	failuresPrefix          string
	rejectsPrefix           string
	shortCircuitsPrefix     string
	timeoutsPrefix          string
	fallbackSuccessesPrefix string
	fallbackFailuresPrefix  string
	canceledPrefix          string
	deadlinePrefix          string
	totalDurationPrefix     string
	runDurationPrefix       string
	concurrencyInUsePrefix  string
	gauges                  map[string]*prometheus.Gauge
	counters                map[string]*prometheus.Counter
	histograms              map[string]*prometheus.Histogram
}

type PrometheusCollectorClient struct {
	namespace string
}

// PrometheusCollectorConfig provides configuration for the prometheus client
type PrometheusCollectorConfig struct {
	// Namespace is the prefix that will be prepended to all metrics sent from this collector.
	Namespace string
}

func InitializePrometheusCollector(config PrometheusCollectorConfig) *PrometheusCollectorClient {
	return &PrometheusCollectorClient{
		namespace: strings.Replace(config.Namespace, "-", "_", -1),
	}
}

// NewPrometheusCollector creates a collector for a specific circuit. The
// prefix given to this circuit will be {config.Prefix}_{subsystem}_{metric}.
// Circuits with "/", ":", and "." in their names will have them replaced with "_".
func (client *PrometheusCollectorClient) NewPrometheusCollector(subsystem string) metricCollector.MetricCollector {
	subsystem = strings.Replace(subsystem, "/", "_", -1)
	subsystem = strings.Replace(subsystem, ":", "_", -1)
	subsystem = strings.Replace(subsystem, ".", "_", -1)
	c := PrometheusCollector{
		circuitOpenPrefix:       "circuitOpen",
		attemptsPrefix:          "attempts",
		errorsPrefix:            "errors",
		successesPrefix:         "successes",
		failuresPrefix:          "failures",
		rejectsPrefix:           "rejects",
		shortCircuitsPrefix:     "shortCircuits",
		timeoutsPrefix:          "timeouts",
		fallbackSuccessesPrefix: "fallbackSuccesses",
		fallbackFailuresPrefix:  "fallbackFailures",
		canceledPrefix:          "contextCanceled",
		deadlinePrefix:          "contextDeadlineExceeded",
		totalDurationPrefix:     "totalDuration",
		runDurationPrefix:       "runDuration",
		concurrencyInUsePrefix:  "concurrencyInUse",
	}
	c.gauges = make(map[string]*prometheus.Gauge)
	c.counters = make(map[string]*prometheus.Counter)
	c.histograms = make(map[string]*prometheus.Histogram)

	{
		gague := prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: client.namespace,
			Subsystem: subsystem,
			Name:      c.circuitOpenPrefix,
			Help:      "Generated hystrix metric to track " + c.circuitOpenPrefix + " for subsystem " + subsystem + ".",
		})
		c.gauges[c.circuitOpenPrefix] = &gague
		prometheus.MustRegister(gague)
	}

	{
		counters := []string{c.attemptsPrefix, c.errorsPrefix, c.successesPrefix,
			c.failuresPrefix, c.rejectsPrefix, c.shortCircuitsPrefix,
			c.timeoutsPrefix, c.fallbackSuccessesPrefix, c.fallbackFailuresPrefix,
			c.canceledPrefix, c.deadlinePrefix}
		for _, metric := range counters {
			counter := prometheus.NewCounter(prometheus.CounterOpts{
				Namespace: client.namespace,
				Subsystem: subsystem,
				Name:      metric,
				Help:      "Generated hystrix metric to track " + metric + " for subsystem " + subsystem + ".",
			})
			c.counters[metric] = &counter
			prometheus.MustRegister(counter)
		}
	}

	{
		histograms := []string{c.totalDurationPrefix, c.runDurationPrefix, c.concurrencyInUsePrefix}
		for _, metric := range histograms {
			histogram := prometheus.NewHistogram(prometheus.HistogramOpts{
				Namespace: client.namespace,
				Subsystem: subsystem,
				Name:      metric,
				Help:      "Generated hystrix metric to track " + metric + " for subsystem " + subsystem + ".",
			})
			c.histograms[metric] = &histogram
			prometheus.MustRegister(histogram)
		}
	}
	return c
}

func (c *PrometheusCollector) setGauge(metric string, value float64) {
	if gauge, ok := c.gauges[metric]; ok {
		(*gauge).Set(value)
	} else {
		log.Printf("Error: '%s' is not a valid gague", metric)
	}
}

func (c *PrometheusCollector) incrementCounterMetric(metric string, i float64) {
	if i == 0 {
		return
	}
	if counter, ok := c.counters[metric]; ok {
		(*counter).Add(i)
	} else {
		log.Printf("Error: '%s' is not a valid counter", metric)
	}
}

func (c *PrometheusCollector) updateHistogramMetric(metric string, dur float64) {
	if histogram, ok := c.histograms[metric]; ok {
		(*histogram).Observe(dur)
	} else {
		log.Printf("Error: '%s' is not a valid timer", metric)
	}
}

func (c PrometheusCollector) Update(r metricCollector.MetricResult) {
	if r.Successes > 0 {
		c.setGauge(c.circuitOpenPrefix, 0)
	} else if r.ShortCircuits > 0 {
		c.setGauge(c.circuitOpenPrefix, 1)
	}

	c.incrementCounterMetric(c.attemptsPrefix, r.Attempts)
	c.incrementCounterMetric(c.errorsPrefix, r.Errors)
	c.incrementCounterMetric(c.successesPrefix, r.Successes)
	c.incrementCounterMetric(c.failuresPrefix, r.Failures)
	c.incrementCounterMetric(c.rejectsPrefix, r.Rejects)
	c.incrementCounterMetric(c.shortCircuitsPrefix, r.ShortCircuits)
	c.incrementCounterMetric(c.timeoutsPrefix, r.Timeouts)
	c.incrementCounterMetric(c.fallbackSuccessesPrefix, r.FallbackSuccesses)
	c.incrementCounterMetric(c.fallbackFailuresPrefix, r.FallbackFailures)
	c.incrementCounterMetric(c.canceledPrefix, r.ContextCanceled)
	c.incrementCounterMetric(c.deadlinePrefix, r.ContextDeadlineExceeded)
	c.updateHistogramMetric(c.totalDurationPrefix, r.TotalDuration.Seconds())
	c.updateHistogramMetric(c.runDurationPrefix, r.RunDuration.Seconds())
	c.updateHistogramMetric(c.concurrencyInUsePrefix, r.ConcurrencyInUse)
}

// Reset is a noop operation in this collector.
func (c PrometheusCollector) Reset() {}
