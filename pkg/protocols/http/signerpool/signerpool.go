package signerpool

import (
	"fmt"
	"strings"
	"sync"

<<<<<<< HEAD
	"github.com/Explorer1092/nuclei/v3/pkg/protocols/http/signer"

	"github.com/Explorer1092/nuclei/v3/pkg/types"
=======
<<<<<<< HEAD:v2/pkg/protocols/http/signerpool/signerpool.go
	"github.com/Explorer1092/nuclei/v2/pkg/protocols/http/signer"

	"github.com/Explorer1092/nuclei/v2/pkg/types"
=======
	"github.com/projectdiscovery/nuclei/v3/pkg/protocols/http/signer"

	"github.com/projectdiscovery/nuclei/v3/pkg/types"
>>>>>>> 419f08f61ce5ca2d3f0eae9fe36dc7c44c1f532a:pkg/protocols/http/signerpool/signerpool.go
>>>>>>> projectdiscovery-main
)

var (
	poolMutex  *sync.RWMutex
	clientPool map[string]signer.Signer
)

// Init initializes the clientpool implementation
func Init(options *types.Options) error {
	poolMutex = &sync.RWMutex{}
	clientPool = make(map[string]signer.Signer)
	return nil
}

// Configuration contains the custom configuration options for a client
type Configuration struct {
	SignerArgs signer.SignerArgs
}

// Hash returns the hash of the configuration to allow client pooling
func (c *Configuration) Hash() string {
	builder := &strings.Builder{}
	builder.WriteString(fmt.Sprintf("%v", c.SignerArgs))
	hash := builder.String()
	return hash
}

// Get creates or gets a client for the protocol based on custom configuration
func Get(options *types.Options, configuration *Configuration) (signer.Signer, error) {
	hash := configuration.Hash()
	poolMutex.RLock()
	if client, ok := clientPool[hash]; ok {
		poolMutex.RUnlock()
		return client, nil
	}
	poolMutex.RUnlock()

	client, err := signer.NewSigner(configuration.SignerArgs)
	if err != nil {
		return nil, err
	}

	poolMutex.Lock()
	clientPool[hash] = client
	poolMutex.Unlock()
	return client, nil
}
