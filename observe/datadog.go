package observe

import (
	datadog "github.com/DataDog/opencensus-go-exporter-datadog"
)

// DatadogExporterConfig provides configuration for the Datadog exporter.
// Config can be initiated with envconfig from environment
type DatadogExporterConfig struct {
	DatadogExporterEnabled        bool   `default:"false" split_words:"true"`
	DatadogExporterMetricsAddress string `split_words:"true"`
	DatadogExporterTracesAddress  string `split_words:"true"`
	DatadogExporterNamespace      string `default:"opencensus" split_words:"true"`
}

// NewDatadogExporter will return Datadog's opencensus exporter if it's enabled (through
// DATADOG_EXPORTER_ENABLED env variable). When the exporter is disabled, it will
// return nil. Exporter will send metrics and traces to Datadog's agent using
// addresses specified through DatadogExporterConfig
func NewDatadogExporter(config DatadogExporterConfig, onErr func(error)) (*datadog.Exporter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getDatadogOpts returns Datadog Options that you can pass directly
// to the OpenCensus exporter or other libraries.
func getDatadogOpts(config DatadogExporterConfig, service, version string, onErr func(err error)) datadog.Options {
	_ = "STUB: not implemented"
	return *

	// Namespace specifies the namespaces to which metric keys are appended.
	// TODO: Figure out what the namespace should be. Can be either a projectID or something else.
	new(datadog.Options)
}

// Service specifies the service name used for tracing.

// TraceAddr specifies the host[:port] address of the Datadog Trace Agent.
// It defaults to localhost:8126.

// StatsAddr specifies the host[:port] address for DogStatsD. It defaults
// to localhost:8125.

// OnError specifies a function that will be called if an error occurs during
// processing stats or metrics.

// // Tags specifies a set of global tags to attach to each metric.
// Tags []string

// GlobalTags holds a set of tags that will automatically be applied to all
// exported spans.

// // DisableCountPerBuckets specifies whether to emit count_per_bucket metrics
// DisableCountPerBuckets bool
