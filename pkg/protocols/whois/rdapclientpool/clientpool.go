package rdapclientpool

import (
	"github.com/Explorer1092/nuclei/v3/pkg/types"
	"github.com/projectdiscovery/gologger"
<<<<<<< HEAD
=======
<<<<<<< HEAD:v2/pkg/protocols/whois/rdapclientpool/clientpool.go
	"github.com/Explorer1092/nuclei/v2/pkg/types"
=======
	"github.com/projectdiscovery/nuclei/v3/pkg/types"
>>>>>>> 419f08f61ce5ca2d3f0eae9fe36dc7c44c1f532a:pkg/protocols/whois/rdapclientpool/clientpool.go
>>>>>>> projectdiscovery-main
	"github.com/projectdiscovery/rdap"
)

var normalClient *rdap.Client

// Init initializes the client pool implementation
func Init(options *types.Options) error {
	// Don't create clients if already created in the past.
	if normalClient != nil {
		return nil
	}

	normalClient = &rdap.Client{}
	if options.Verbose || options.Debug || options.DebugRequests || options.DebugResponse {
		normalClient.Verbose = func(text string) {
			gologger.Debug().Msgf("rdap: %s", text)
		}
	}
	return nil
}

// Configuration contains the custom configuration options for a client - placeholder
type Configuration struct{}

// Hash returns the hash of the configuration to allow client pooling - placeholder
func (c *Configuration) Hash() string {
	return ""
}

// Get creates or gets a client for the protocol based on custom configuration
func Get(options *types.Options, configuration *Configuration) (*rdap.Client, error) {
	return normalClient, nil
}
