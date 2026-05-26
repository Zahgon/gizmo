package kit

import (
	"time"
)

// Config holds info required to configure a gizmo kit.Server.
//
// This struct is loaded from the environment at Run and only made public to expose
// for documentation.
type Config struct {
	// HealthCheckPath is used by server to init the proper HealthCheckHandler.
	// If empty, this will default to '/healthz'.
	HealthCheckPath string `envconfig:"GIZMO_HEALTH_CHECK_PATH"`

	// MaxHeaderBytes can be used to override the default of 1<<20.
	MaxHeaderBytes int `envconfig:"GIZMO_MAX_HEADER_BYTES"`

	// ReadTimeout can be used to override the default http server timeout of 10s.
	// The string should be formatted like a time.Duration string.
	ReadTimeout time.Duration `envconfig:"GIZMO_READ_TIMEOUT"`

	// WriteTimeout can be used to override the default http server timeout of 10s.
	// The string should be formatted like a time.Duration string.
	WriteTimeout time.Duration `envconfig:"GIZMO_WRITE_TIMEOUT"`

	// IdleTimeout can be used to override the default http server timeout of 120s.
	// The string should be formatted like a time.Duration string.
	IdleTimeout time.Duration `envconfig:"GIZMO_IDLE_TIMEOUT"`

	// ShutdownTimeout can be used to override the default http server shutdown timeout
	// of 5m.
	ShutdownTimeout time.Duration `envconfig:"GIZMO_SHUTDOWN_TIMEOUT"`

	// GOMAXPROCS can be used to override the default GOMAXPROCS.
	GOMAXPROCS int `envconfig:"GIZMO_GOMAXPROCS"`

	// HTTPAddr is the address the server implementation will bind to.
	// The default is "" (bind to all interfaces)
	HTTPAddr string `envconfig:"HTTP_ADDR"`
	// HTTPPort is the port the server implementation will serve HTTP over.
	// The default is 8080
	HTTPPort int `envconfig:"HTTP_PORT"`
	// RPCPort is the port the server implementation will serve RPC over.
	// The default is 8081.
	RPCPort int `envconfig:"RPC_PORT"`

	// Enable pprof Profiling. Off by default.
	EnablePProf bool `envconfig:"ENABLE_PPROF"`
}

func loadConfig() Config { _ = "STUB: not implemented"; return *new(Config) }

// fall back to PORT for GAE
