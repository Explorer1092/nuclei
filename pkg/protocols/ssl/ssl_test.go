package ssl

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

<<<<<<< HEAD
	"github.com/Explorer1092/nuclei/v3/pkg/model"
	"github.com/Explorer1092/nuclei/v3/pkg/model/types/severity"
	"github.com/Explorer1092/nuclei/v3/pkg/output"
	"github.com/Explorer1092/nuclei/v3/pkg/protocols/common/contextargs"
	"github.com/Explorer1092/nuclei/v3/pkg/testutils"
=======
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
>>>>>>> 419f08f61ce5ca2d3f0eae9fe36dc7c44c1f532a:pkg/protocols/ssl/ssl_test.go
>>>>>>> projectdiscovery-main
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
