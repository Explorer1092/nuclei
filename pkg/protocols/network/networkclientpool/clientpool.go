package networkclientpool

import (
	"github.com/Explorer1092/nuclei/v3/pkg/protocols/common/protocolstate"
	"github.com/Explorer1092/nuclei/v3/pkg/types"
	"github.com/projectdiscovery/fastdialer/fastdialer"
<<<<<<< HEAD
=======
<<<<<<< HEAD:v2/pkg/protocols/network/networkclientpool/clientpool.go
	"github.com/Explorer1092/nuclei/v2/pkg/protocols/common/protocolstate"
	"github.com/Explorer1092/nuclei/v2/pkg/types"
=======
	"github.com/projectdiscovery/nuclei/v3/pkg/protocols/common/protocolstate"
	"github.com/projectdiscovery/nuclei/v3/pkg/types"
>>>>>>> 419f08f61ce5ca2d3f0eae9fe36dc7c44c1f532a:pkg/protocols/network/networkclientpool/clientpool.go
>>>>>>> projectdiscovery-main
)

var (
	normalClient *fastdialer.Dialer
)

// Init initializes the clientpool implementation
func Init(options *types.Options) error {
	// Don't create clients if already created in the past.
	if normalClient != nil {
		return nil
	}
	normalClient = protocolstate.Dialer
	return nil
}

// Configuration contains the custom configuration options for a client
type Configuration struct{}

// Hash returns the hash of the configuration to allow client pooling
func (c *Configuration) Hash() string {
	return ""
}

// Get creates or gets a client for the protocol based on custom configuration
func Get(options *types.Options, configuration *Configuration /*TODO review unused parameters*/) (*fastdialer.Dialer, error) {
	return normalClient, nil
}
