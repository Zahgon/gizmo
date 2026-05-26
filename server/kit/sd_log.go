package kit

import (
	"context"
	"fmt"

	"cloud.google.com/go/logging"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"google.golang.org/genproto/googleapis/api/monitoredres"
)

type sdLogger struct {
	project string
	monRes  *monitoredres.MonitoredResource
	lc      *logging.Client
	lgr     *logging.Logger
}

func newStackdriverLogger(ctx context.Context, logID, projectID, service, version string) (log.Logger, func() error, error) {
	_ = "STUB: not implemented"
	return *new(log.Logger), nil, nil
}

func (l sdLogger) Log(keyvals ...interface{}) error { _ = "STUB: not implemented"; return nil }

func (l sdLogger) getTraceID(traceCtx string) string { _ = "STUB: not implemented"; return "" }

const cloudTraceLogKey = "cloud-trace"

///////////////////////////////////////////////////
// below funcs are straight up copied out of go-kit/kit/log:
// https://github.com/go-kit/kit/blob/master/log/json_logger.go
// we needed the magic for keyvals => map[string]interface{} but we're doing the
// writing the JSON ourselves
///////////////////////////////////////////////////

func logKeyValsToMap(keyvals ...interface{}) (map[string]interface{}, level.Value, string) {
	_ = "STUB: not implemented"
	return nil, *new(level.Value), ""
}

// +1 to handle case when len is odd

func merge(dst map[string]interface{}, k, v interface{}) { _ = "STUB: not implemented"; return }

// We want json.Marshaler and encoding.TextMarshaller to take priority over
// err.Error() and v.String(). But json.Marshall (called later) does that by
// default so we force a no-op if it's one of those 2 case.

func safeString(str fmt.Stringer) (s string) { _ = "STUB: not implemented"; return "" }

func safeError(err error) (s interface{}) { _ = "STUB: not implemented"; return nil }
