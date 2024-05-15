package generators

import (
<<<<<<< HEAD:v2/pkg/protocols/common/generators/options.go
	"github.com/Explorer1092/nuclei/v2/pkg/types"
=======
	"github.com/projectdiscovery/nuclei/v3/pkg/types"
>>>>>>> 7f556f8e33080a9eb2e52a9c14fa2e16f32f62c3:pkg/protocols/common/generators/options.go
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
