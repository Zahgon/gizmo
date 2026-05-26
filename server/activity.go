package server

// ActivityMonitor can be used to count and share the number of active requests.
type ActivityMonitor struct {
	// counter for # of active requests
	reqCount uint32
}

// NewActivityMonitor will return a new ActivityMonitor instance.
func NewActivityMonitor() *ActivityMonitor { _ = "STUB: not implemented"; return nil }

// Active returns true if there are requests currently in flight.
func (a *ActivityMonitor) Active() bool { _ = "STUB: not implemented"; return false }

// CountRequest will increment the request count and signal
// the activity monitor to stay active. Call this in your server
// when you receive a request.
func (a *ActivityMonitor) CountRequest() { _ = "STUB: not implemented"; return }

// UncountRequest will decrement the active request count. Best practice is to `defer`
// this function call directly after calling CountRequest().
func (a *ActivityMonitor) UncountRequest() { _ = "STUB: not implemented"; return }

// NumActiveRequests returns the number of in-flight requests currently
// running on this server.
func (a *ActivityMonitor) NumActiveRequests() uint32 { _ = "STUB: not implemented"; return 0 }
