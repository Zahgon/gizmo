package observe

import (
	"contrib.go.opencensus.io/exporter/stackdriver"
	"contrib.go.opencensus.io/exporter/stackdriver/monitoredresource"
)

// RegisterAndObserveGCP will initiate and register Stackdriver profiling and tracing and
// metrics in environments that pass the tests in the IsGCPEnabled function. All
// exporters will be registered using the information returned by the GetServiceInfo
// function. Tracing and metrics are enabled via OpenCensus exporters. See the OpenCensus
// documentation for instructions for registering additional spans and metrics.
func RegisterAndObserveGCP(onError func(error)) error { _ = "STUB: not implemented"; return nil }

// NewStackdriverExporter will return the tracing and metrics through
// the stack driver exporter, if exists in the underlying platform.
// If exporter is registered, it returns the exporter so you can register
// it and ensure to call Flush on termination.
func NewStackdriverExporter(projectID string, onErr func(error)) (*stackdriver.Exporter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getSDOpts returns Stack Driver Options that you can pass directly
// to the OpenCensus exporter or other libraries.
func getSDOpts(projectID, service, version string, mr monitoredresource.Interface, onErr func(err error)) (*stackdriver.Options, error) {
	_ = "STUB: not implemented"

	// this is so that you can export views from your local server up to SD if you wish
	return nil, nil
}
