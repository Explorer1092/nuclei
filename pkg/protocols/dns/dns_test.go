package dns

import (
	"testing"

	"github.com/stretchr/testify/require"

<<<<<<< HEAD
	"github.com/Explorer1092/nuclei/v3/pkg/model"
	"github.com/Explorer1092/nuclei/v3/pkg/model/types/severity"
	"github.com/Explorer1092/nuclei/v3/pkg/testutils"
=======
<<<<<<< HEAD:v2/pkg/protocols/dns/dns_test.go
	"github.com/Explorer1092/nuclei/v2/pkg/model"
	"github.com/Explorer1092/nuclei/v2/pkg/model/types/severity"
	"github.com/Explorer1092/nuclei/v2/pkg/testutils"
=======
	"github.com/projectdiscovery/nuclei/v3/pkg/model"
	"github.com/projectdiscovery/nuclei/v3/pkg/model/types/severity"
	"github.com/projectdiscovery/nuclei/v3/pkg/testutils"
>>>>>>> 419f08f61ce5ca2d3f0eae9fe36dc7c44c1f532a:pkg/protocols/dns/dns_test.go
>>>>>>> projectdiscovery-main
)

func TestDNSCompileMake(t *testing.T) {
	options := testutils.DefaultOptions

	recursion := false
	testutils.Init(options)
	const templateID = "testing-dns"
	request := &Request{
		RequestType: DNSRequestTypeHolder{DNSRequestType: A},
		Class:       "INET",
		Retries:     5,
		ID:          templateID,
		Recursion:   &recursion,
		Name:        "{{FQDN}}",
	}
	executerOpts := testutils.NewMockExecuterOptions(options, &testutils.TemplateInfo{
		ID:   templateID,
		Info: model.Info{SeverityHolder: severity.Holder{Severity: severity.Low}, Name: "test"},
	})
	err := request.Compile(executerOpts)
	require.Nil(t, err, "could not compile dns request")

	req, err := request.Make("one.one.one.one", map[string]interface{}{"FQDN": "one.one.one.one"})
	require.Nil(t, err, "could not make dns request")
	require.Equal(t, "one.one.one.one.", req.Question[0].Name, "could not get correct dns question")
}

func TestDNSRequests(t *testing.T) {
	options := testutils.DefaultOptions

	recursion := false
	testutils.Init(options)
	const templateID = "testing-dns"

	t.Run("dns-regular", func(t *testing.T) {

		request := &Request{
			RequestType: DNSRequestTypeHolder{DNSRequestType: A},
			Class:       "INET",
			Retries:     5,
			ID:          templateID,
			Recursion:   &recursion,
			Name:        "{{FQDN}}",
		}
		executerOpts := testutils.NewMockExecuterOptions(options, &testutils.TemplateInfo{
			ID:   templateID,
			Info: model.Info{SeverityHolder: severity.Holder{Severity: severity.Low}, Name: "test"},
		})
		err := request.Compile(executerOpts)
		require.Nil(t, err, "could not compile dns request")

		reqCount := request.Requests()
		require.Equal(t, 1, reqCount, "could not get correct dns request count")
	})

	// test payload requests count is correct
	t.Run("dns-payload", func(t *testing.T) {

		request := &Request{
			RequestType: DNSRequestTypeHolder{DNSRequestType: A},
			Class:       "INET",
			Retries:     5,
			ID:          templateID,
			Recursion:   &recursion,
			Name:        "{{subdomain}}.{{FQDN}}",
			Payloads:    map[string]interface{}{"subdomain": []string{"a", "b", "c"}},
		}
		executerOpts := testutils.NewMockExecuterOptions(options, &testutils.TemplateInfo{
			ID:   templateID,
			Info: model.Info{SeverityHolder: severity.Holder{Severity: severity.Low}, Name: "test"},
		})
		err := request.Compile(executerOpts)
		require.Nil(t, err, "could not compile dns request")

		reqCount := request.Requests()
		require.Equal(t, 3, reqCount, "could not get correct dns request count")
	})
}
