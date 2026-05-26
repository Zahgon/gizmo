package config // import "github.com/NYTimes/gizmo/config"

// EnvAppName is used as a prefix for environment variable
// names when using the LoadXFromEnv funcs.
// It defaults to empty.
var EnvAppName = ""

// LoadJSONFile is a helper function to read a config file into whatever
// config struct you need. For example, your custom config could be composed
// of one or more of the given Config, AWS, MySQL, Oracle or MongoDB structs.
func LoadJSONFile(fileName string, cfg interface{}) { _ = "STUB: not implemented"; return }
