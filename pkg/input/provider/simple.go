package provider

import (
	"github.com/Explorer1092/nuclei/v3/pkg/input/types"
	"github.com/Explorer1092/nuclei/v3/pkg/protocols/common/contextargs"
	iputil "github.com/projectdiscovery/utils/ip"
	stringsutil "github.com/projectdiscovery/utils/strings"
	"strings"
)

// SimpleInputProvider is a simple input provider for nuclei
// that acts like a No-Op and returns given list of urls as input
type SimpleInputProvider struct {
	Inputs []*contextargs.MetaInput
}

// NewSimpleInputProvider creates a new simple input provider
func NewSimpleInputProvider() *SimpleInputProvider {
	return &SimpleInputProvider{
		Inputs: make([]*contextargs.MetaInput, 0),
	}
}

// NewSimpleInputProviderWithUrls creates a new simple input provider with the given urls
func NewSimpleInputProviderWithUrls(urls ...string) *SimpleInputProvider {
	provider := NewSimpleInputProvider()
	for _, url := range urls {
		provider.Set(url)
	}
	return provider
}

// Count returns the total number of targets for the input provider
func (s *SimpleInputProvider) Count() int64 {
	return int64(len(s.Inputs))
}

// Iterate over all inputs in order
func (s *SimpleInputProvider) Iterate(callback func(value *contextargs.MetaInput) bool) {
	for _, input := range s.Inputs {
		if !callback(input) {
			break
		}
	}
}

// Set adds an item to the input provider
func (s *SimpleInputProvider) Set(value string) {
	if stringsutil.ContainsAny(value, ",") {
		parts := strings.Split(value, ",")
		// 处理2列的情况: CustomIP,Host
		if len(parts) == 2 {
			if iputil.IsIP(strings.TrimSpace(parts[0])) {
				metaInput := contextargs.NewMetaInput()
				metaInput.CustomIP = strings.TrimSpace(parts[0])
				metaInput.Input = strings.TrimSpace(parts[1])
				s.Inputs = append(s.Inputs, metaInput)
			}
		}
	} else {
		metaInput := contextargs.NewMetaInput()
		metaInput.Input = value
		s.Inputs = append(s.Inputs, metaInput)
	}
}

// SetWithProbe adds an item to the input provider with HTTP probing
func (s *SimpleInputProvider) SetWithProbe(value string, probe types.InputLivenessProbe) error {
	probedValue, err := probe.ProbeURL(value)
	if err != nil {
		return err
	}
	metaInput := contextargs.NewMetaInput()
	metaInput.Input = probedValue
	s.Inputs = append(s.Inputs, metaInput)
	return nil
}

// SetWithExclusions adds an item to the input provider if it doesn't match any of the exclusions
func (s *SimpleInputProvider) SetWithExclusions(value string) error {
	metaInput := contextargs.NewMetaInput()
	metaInput.Input = value
	s.Inputs = append(s.Inputs, metaInput)
	return nil
}

// InputType returns the type of input provider
func (s *SimpleInputProvider) InputType() string {
	return "SimpleInputProvider"
}

// Close the input provider and cleanup any resources
func (s *SimpleInputProvider) Close() {
	// no-op
}
