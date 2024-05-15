package networkclientpool

import (
	"github.com/Explorer1092/nuclei/v2/pkg/protocols/common/protocolstate"
	"github.com/Explorer1092/nuclei/v2/pkg/types"
	"github.com/projectdiscovery/fastdialer/fastdialer"
<<<<<<< HEAD:v2/pkg/protocols/network/networkclientpool/clientpool.go
=======
	"github.com/projectdiscovery/nuclei/v3/pkg/protocols/common/protocolstate"
	"github.com/projectdiscovery/nuclei/v3/pkg/types"
>>>>>>> 7f556f8e33080a9eb2e52a9c14fa2e16f32f62c3:pkg/protocols/network/networkclientpool/clientpool.go
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
