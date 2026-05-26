// Package observe provides functions
// that help with setting tracing/metrics
// in cloud providers, mainly GCP.
package observe // import "github.com/NYTimes/gizmo/observe"

// GoogleProjectID returns the GCP Project ID
// that can be used to instantiate various
// GCP clients such as Stack Driver.
func GoogleProjectID() string { _ = "STUB: not implemented"; return "" }

// IsGAE tells you whether your program is running
// within the App Engine platform.
func IsGAE() bool { _ = "STUB: not implemented"; return false }

// GetGAEInfo returns the service and the version of the
// GAE application.
func GetGAEInfo() (service, version string) { _ = "STUB: not implemented"; return "", "" }

// IsCloudRun tells you whether your program is running
// within the Cloud Run platform.
func IsCloudRun() bool { _ = "STUB: not implemented"; return false }

// GetCloudRunInfo returns the service and the version of the
// Cloud Run application.
func GetCloudRunInfo() (service, version, config string) {
	_ = "STUB: not implemented"
	return "", "", ""
}

// GetServiceInfo returns the GCP Project ID,
// the service name and version (GAE or through
// SERVICE_NAME/SERVICE_VERSION env vars). Note
// that SERVICE_NAME/SERVICE_VERSION are not standard but
// your application can pass them in as variables
// to be included in your trace attributes
func GetServiceInfo() (projectID, service, version string) {
	_ = "STUB: not implemented"
	return "", "", ""
}

// IsGCPEnabled returns whether the running application
// is inside GCP, has access to its products or Stackdriver integration was
// not disabled.
func IsGCPEnabled() bool { _ = "STUB: not implemented"; return false }

// SkipObserve checks if the GIZMO_SKIP_OBSERVE environment variable has been populated.
// This may be used along with local development to cut down on long startup times caused
// by the 'monitoredresource.Autodetect()' call in IsGCPEnabled().
func SkipObserve() bool { _ = "STUB: not implemented"; return false }
