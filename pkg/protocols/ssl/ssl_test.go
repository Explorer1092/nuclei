package ssl

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

<<<<<<< HEAD:v2/pkg/protocols/ssl/ssl_test.go
	"github.com/Explorer1092/nuclei/v2/pkg/model"
	"github.com/Explorer1092/nuclei/v2/pkg/model/types/severity"
	"github.com/Explorer1092/nuclei/v2/pkg/output"
	"github.com/Explorer1092/nuclei/v2/pkg/protocols/common/contextargs"
	"github.com/Explorer1092/nuclei/v2/pkg/testutils"
=======
	"github.com/projectdiscovery/nuclei/v3/pkg/model"
	"github.com/projectdiscovery/nuclei/v3/pkg/model/types/severity"
	"github.com/projectdiscovery/nuclei/v3/pkg/output"
	"github.com/projectdiscovery/nuclei/v3/pkg/protocols/common/contextargs"
	"github.com/projectdiscovery/nuclei/v3/pkg/testutils"
>>>>>>> 7f556f8e33080a9eb2e52a9c14fa2e16f32f62c3:pkg/protocols/ssl/ssl_test.go
)

func TestSSLProtocol(t *testing.T) {
	options := testutils.DefaultOptions

	testutils.Init(options)
	templateID := "testing-ssl"
	request := &Request{
		Address: "{{Hostname}}",
	}
	executerOpts := testutils.NewMockExecuterOptions(options, &testutils.TemplateInfo{
		ID:   templateID,
		Info: model.Info{SeverityHolder: severity.Holder{Severity: severity.Low}, Name: "test"},
	})
	err := request.Compile(executerOpts)
	require.Nil(t, err, "could not compile ssl request")

	var gotEvent output.InternalEvent
	ctxArgs := contextargs.NewWithInput(context.Background(), "scanme.sh:443")
	err = request.ExecuteWithResults(ctxArgs, nil, nil, func(event *output.InternalWrappedEvent) {
		gotEvent = event.InternalEvent
	})
	require.Nil(t, err, "could not run ssl request")
	require.NotEmpty(t, gotEvent, "could not get event items")
}

func TestGetAddress(t *testing.T) {
	address, _ := getAddress("https://scanme.sh")
	require.Equal(t, "scanme.sh:443", address, "could not get correct address")
}
