package generators

import (
<<<<<<< HEAD
	"github.com/Explorer1092/nuclei/v3/pkg/types"
=======
<<<<<<< HEAD:v2/pkg/protocols/common/generators/options.go
	"github.com/Explorer1092/nuclei/v2/pkg/types"
=======
	"github.com/projectdiscovery/nuclei/v3/pkg/types"
>>>>>>> 419f08f61ce5ca2d3f0eae9fe36dc7c44c1f532a:pkg/protocols/common/generators/options.go
>>>>>>> projectdiscovery-main
)

// BuildPayloadFromOptions returns a map with the payloads provided via CLI
func BuildPayloadFromOptions(options *types.Options) map[string]interface{} {
	m := make(map[string]interface{})
	// merge with vars
	if !options.Vars.IsEmpty() {
		m = MergeMaps(m, options.Vars.AsMap())
	}

	// merge with env vars
	if options.EnvironmentVariables {
		m = MergeMaps(EnvVars(), m)
	}
	return m
}
